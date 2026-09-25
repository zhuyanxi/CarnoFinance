import json
import os
import sqlite3
import subprocess
import time
from datetime import datetime, timedelta
from urllib.parse import urlencode

import pandas as pd
import yfinance as yf

etfs = ["SZ159915", "SH515100", "SH513100", "SH518880", "SH510050"]
TENCENT_DAILY_KLINE_URL = "https://web.ifzq.gtimg.cn/appstock/app/fqkline/get"
TENCENT_MINUTE_KLINE_URL = "https://ifzq.gtimg.cn/appstock/app/kline/mkline"
TENCENT_MAX_RETRIES = 3
TENCENT_TIMEOUT = 15
PRICE_COLUMNS = ["open", "high", "low", "close"]

# 配置日志或打印格式
def log(message: str):
    print(f"[{datetime.now().strftime('%Y-%m-%d %H:%M:%S')}] {message}")


def _to_yf_ticker(full_code: str) -> str:
    """
    将 "SZ159915" / "SH515100" 格式转为 yfinance 格式 "159915.SZ" / "515100.SS"
    """
    prefix = full_code[:2].upper()
    digits = "".join(filter(str.isdigit, full_code))
    if prefix == "SZ":
        return f"{digits}.SZ"
    elif prefix == "SH":
        return f"{digits}.SS"
    else:
        raise ValueError(f"Unknown exchange prefix: {prefix}")


def _to_tencent_symbol(full_code: str) -> str:
    """将 "SZ159915" / "SH515100" 转为腾讯行情的代码格式。"""
    prefix = full_code[:2].upper()
    digits = "".join(filter(str.isdigit, full_code))
    market = {"SH": "sh", "SZ": "sz"}.get(prefix)
    if market is None or not digits:
        raise ValueError(f"Unknown exchange prefix or code: {full_code}")
    return f"{market}{digits}"


def _normalize_yfinance_data(
    df: pd.DataFrame,
    full_code: str,
    source_time_column: str,
    target_time_column: str,
) -> pd.DataFrame:
    df = df.reset_index()
    df = df.rename(columns={
        "Open": "open",
        "High": "high",
        "Low": "low",
        "Close": "close",
    })

    required_columns = [source_time_column, *PRICE_COLUMNS]
    missing_columns = [column for column in required_columns if column not in df.columns]
    if missing_columns:
        raise ValueError(f"missing yfinance columns: {missing_columns}")
    if df[PRICE_COLUMNS].isna().any().any():
        raise ValueError("yfinance OHLC contains missing values")

    df[target_time_column] = pd.to_datetime(df[source_time_column]).dt.strftime(
        "%Y%m%d" if target_time_column == "trade_date" else "%Y-%m-%d %H:%M:%S"
    )
    df["ts_code"] = full_code

    for column in PRICE_COLUMNS:
        df[column] = df[column].round(3)

    return df[["ts_code", target_time_column, *PRICE_COLUMNS]]


def _fetch_tencent_json(url: str):
    result = subprocess.run(
        [
            "curl",
            "--fail",
            "--silent",
            "--show-error",
            "--max-time",
            str(TENCENT_TIMEOUT),
            url,
        ],
        check=True,
        capture_output=True,
        text=True,
    )
    payload = json.loads(result.stdout)
    if payload.get("code") != 0:
        raise ValueError(f"Tencent response error: {payload.get('msg', 'unknown error')}")
    return payload.get("data") or {}


def _normalize_tencent_data(
    klines,
    full_code: str,
    target_time_column: str,
    timestamp_format: str | None = None,
) -> pd.DataFrame:
    output_time_format = "%Y%m%d" if target_time_column == "trade_date" else "%Y-%m-%d %H:%M:%S"
    rows = []

    for kline in klines:
        parts = kline if isinstance(kline, list) else kline.split(",")
        if len(parts) < 5:
            raise ValueError(f"unexpected Tencent kline format: {kline}")

        timestamp = pd.to_datetime(parts[0], format=timestamp_format, errors="raise")
        rows.append({
            "ts_code": full_code,
            target_time_column: timestamp.strftime(output_time_format),
            "open": float(parts[1]),
            "high": float(parts[3]),
            "low": float(parts[4]),
            "close": float(parts[2]),
        })

    if not rows:
        raise ValueError("empty Tencent kline response")

    df = pd.DataFrame(rows, columns=["ts_code", target_time_column, *PRICE_COLUMNS])
    for column in PRICE_COLUMNS:
        df[column] = df[column].round(3)
    return df


def _fetch_tencent_data(
    full_code: str,
    start_date: str,
    end_date: str,
    period: str,
    target_time_column: str,
) -> pd.DataFrame:
    symbol = _to_tencent_symbol(full_code)
    if target_time_column == "trade_date":
        params = f"{symbol},day,{start_date},{end_date},1000,qfq"
        url = f"{TENCENT_DAILY_KLINE_URL}?{urlencode({'param': params}, safe=',')}"
        kline_key = "qfqday"
        timestamp_format = None
    else:
        params = f"{symbol},m{period},,1000"
        url = f"{TENCENT_MINUTE_KLINE_URL}?{urlencode({'param': params}, safe=',')}"
        kline_key = f"m{period}"
        timestamp_format = "%Y%m%d%H%M"

    last_error = None
    for attempt in range(1, TENCENT_MAX_RETRIES + 1):
        try:
            symbol_data = (_fetch_tencent_json(url).get(symbol) or {})
            klines = symbol_data.get(kline_key) or []
            if target_time_column == "trade_date":
                klines = klines or symbol_data.get("day") or []
            if not klines:
                raise ValueError(f"empty Tencent {kline_key} response")

            df = _normalize_tencent_data(
                klines,
                full_code,
                target_time_column,
                timestamp_format,
            )
            if target_time_column != "trade_date":
                timestamps = pd.to_datetime(df[target_time_column])
                start_timestamp = pd.to_datetime(start_date)
                end_timestamp = pd.to_datetime(end_date)
                df = df.loc[
                    (timestamps >= start_timestamp) & (timestamps < end_timestamp)
                ].copy()
                if df.empty:
                    raise ValueError(
                        f"Tencent minute data does not cover {start_date} to {end_date}"
                    )
            return df
        except Exception as exc:
            last_error = exc
            if attempt < TENCENT_MAX_RETRIES:
                time.sleep(attempt)

    raise RuntimeError(
        f"Tencent request failed for {full_code} after {TENCENT_MAX_RETRIES} attempts"
    ) from last_error


def _fetch_market_data(
    full_code: str,
    start_date: str,
    end_date: str,
    yf_interval: str,
    fallback_period: str,
    source_time_column: str,
    target_time_column: str,
) -> pd.DataFrame:
    yf_ticker = _to_yf_ticker(full_code)
    fallback_reason = "unknown yfinance failure"

    try:
        yf_df = yf.Ticker(yf_ticker).history(
            start=start_date,
            end=end_date,
            interval=yf_interval,
        )
        if yf_df is None:
            fallback_reason = "returned None"
        elif not isinstance(yf_df, pd.DataFrame):
            fallback_reason = f"returned {type(yf_df).__name__}"
        elif yf_df.empty:
            fallback_reason = "returned an empty DataFrame"
        else:
            return _normalize_yfinance_data(
                yf_df,
                full_code,
                source_time_column,
                target_time_column,
            )
    except Exception as exc:
        fallback_reason = str(exc)

    log(
        f"yfinance unavailable for {full_code} ({fallback_reason}); "
        "trying Tencent fallback."
    )
    df = _fetch_tencent_data(
        full_code,
        start_date,
        end_date,
        fallback_period,
        target_time_column,
    )
    log(f"Using Tencent fallback for {full_code}: {len(df)} records.")
    return df


def init_db(db_path: str = "etf_data.db"):
    """
    初始化 SQLite 数据库，并创建日线和分钟级别的数据表
    """
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()

    # 1. 创建日线表结构，复合主键 (ts_code, trade_date)
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS etf_daily_prices (
            ts_code TEXT NOT NULL,
            trade_date TEXT NOT NULL,
            open REAL,
            high REAL,
            low REAL,
            close REAL,
            PRIMARY KEY (ts_code, trade_date)
        )
    """)

    # 2. 创建分钟级别表结构，复合主键 (ts_code, trade_time)
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS etf_minute_prices (
            ts_code TEXT NOT NULL,
            trade_time TEXT NOT NULL, -- 格式：YYYY-MM-DD HH:MM:SS
            open REAL,
            high REAL,
            low REAL,
            close REAL,
            PRIMARY KEY (ts_code, trade_time)
        )
    """)

    conn.commit()
    conn.close()
    log(f"Database initialized. Tables 'etf_daily_prices' and 'etf_minute_prices' are ready at '{db_path}'.")


def fetch_and_save_etf_data(
    start_date: str,
    end_date: str,
    db_path: str = "etf_data.db"
):
    """
    优先用 yfinance 下载 ETF 日线数据；无数据或请求失败时使用 Tencent fallback。
    下载的数据存入 SQLite 数据库。

    :param start_date: 开始日期，格式为 'YYYY-MM-DD'
    :param end_date: 结束日期，格式为 'YYYY-MM-DD'
    :param db_path: SQLite 数据库文件路径
    """
    # etfs = ["SZ159915", "SH515100", "SH513100", "SH518880"]
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    total_records_inserted = 0

    for full_code in etfs:
        yf_ticker = _to_yf_ticker(full_code)
        log(f"Fetching DAILY data for {full_code} ({yf_ticker}) from {start_date} to {end_date}...")

        try:
            final_df = _fetch_market_data(
                full_code,
                start_date,
                end_date,
                "1d",
                "day",
                "Date",
                "trade_date",
            )
            records = final_df.to_records(index=False).tolist()

            cursor.executemany("""
                INSERT OR REPLACE INTO etf_daily_prices (ts_code, trade_date, open, high, low, close)
                VALUES (?, ?, ?, ?, ?, ?)
            """, records)

            conn.commit()
            log(f"Successfully saved {len(records)} daily records for {full_code}.")
            total_records_inserted += len(records)

        except Exception as e:
            log(f"Error fetching/saving daily data for {full_code}: {str(e)}")
            conn.rollback()

    conn.close()
    log(f"Daily data pipeline complete. Total {total_records_inserted} records processed.")


def fetch_and_save_etf_min_data(
    start_date: str,
    end_date: str,
    period: str = "5",
    db_path: str = "etf_data.db"
):
    """
    优先用 yfinance 下载 ETF 分钟级数据；无数据或请求失败时使用 Tencent fallback。
    下载的数据存入 SQLite 数据库。

    注意：yfinance 分钟级数据有严格的时间窗口限制：
      - 1m: 最近 7 天
      - 5m: 最近 60 天
      超出范围会返回空数据。
        Tencent fallback 返回最近一段分钟数据，只有请求日期落在其返回范围内时才能补齐数据。

    :param start_date: 开始日期，格式为 'YYYY-MM-DD'
    :param end_date: 结束日期，格式为 'YYYY-MM-DD'
    :param period: 级别。可选 '1m', '5m', '15m', '30m', '60m'
    :param db_path: SQLite 数据库文件路径
    """
    # etfs = ["SZ159915", "SH515100", "SH513100", "SH518880"]

    # yfinance interval 参数映射
    period_key = str(period).lower()
    period_key = period_key.removesuffix("m")

    interval_map = {
        "1": "1m",
        "5": "5m",
        "15": "15m",
        "30": "30m",
        "60": "60m",
    }
    if period_key not in interval_map:
        raise ValueError(f"Unsupported minute period: {period}")
    yf_interval = interval_map[period_key]

    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    total_records_inserted = 0

    for full_code in etfs:
        yf_ticker = _to_yf_ticker(full_code)
        log(f"Fetching MINUTE (period={yf_interval}) data for {full_code} ({yf_ticker}) from {start_date} to {end_date}...")

        try:
            final_df = _fetch_market_data(
                full_code,
                start_date,
                end_date,
                yf_interval,
                period_key,
                "Datetime",
                "trade_time",
            )
            records = final_df.to_records(index=False).tolist()

            cursor.executemany("""
                INSERT OR REPLACE INTO etf_minute_prices (ts_code, trade_time, open, high, low, close)
                VALUES (?, ?, ?, ?, ?, ?)
            """, records)

            conn.commit()
            log(f"Successfully saved {len(records)} minute records for {full_code}.")
            total_records_inserted += len(records)

        except Exception as e:
            log(f"Error fetching/saving minute data for {full_code}: {str(e)}")
            conn.rollback()

    conn.close()
    log(f"Minute data pipeline complete. Total {total_records_inserted} records processed.")


# 示例运行逻辑
if __name__ == "__main__":
    db_file = "cmd/finance.db"

    # 1. 初始化数据库及表结构
    init_db(db_file)

    # 2. 设定起止日期参数（yfinance 使用 YYYY-MM-DD 格式）
    today = datetime.now()
    # end = today.strftime('%Y-%m-%d')
    end = (today + timedelta(days=1)).strftime('%Y-%m-%d')
    start = (today - timedelta(days=2)).strftime('%Y-%m-%d')

    log(f"Starting pipeline. Target database: {db_file}, start: {start}, end: {end}")

    # 下载日线数据
    fetch_and_save_etf_data(start_date=start, end_date=end, db_path=db_file)

    # 下载 5 分钟线数据
    # 注意：yfinance 5m 数据最多覆盖最近 60 天，请求范围超出会被 Yahoo 自动裁剪
    fetch_and_save_etf_min_data(start_date=start, end_date=end, period="5", db_path=db_file)

    # 验证读取分钟数据
    if os.path.exists(db_file):
        conn = sqlite3.connect(db_file)
        test_daily = pd.read_sql("SELECT * FROM etf_daily_prices ORDER BY trade_date DESC, ts_code DESC LIMIT 5", conn)
        test_min = pd.read_sql("SELECT * FROM etf_minute_prices ORDER BY trade_time DESC, ts_code DESC LIMIT 5", conn)

        print("\n--- 数据库中 [日线] 最后 5 条数据样例 ---")
        print(test_daily)

        print("\n--- 数据库中 [分钟线] 最后 5 条数据样例 ---")
        print(test_min)

        conn.close()
