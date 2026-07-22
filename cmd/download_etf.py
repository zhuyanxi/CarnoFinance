import sqlite3
import pandas as pd
import yfinance as yf
from datetime import datetime, timedelta
import os

etfs = ["SZ159915", "SH515100", "SH513100", "SH518880", "SH510050"]

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
    用 yfinance 下载 ETF 日线数据并存入 SQLite 数据库

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
            ticker = yf.Ticker(yf_ticker)
            df = ticker.history(start=start_date, end=end_date, interval="1d")

            if df.empty:
                log(f"No daily data returned for {full_code} ({yf_ticker}).")
                continue

            # yfinance 返回的 index 是带时区的 DatetimeIndex，转为 date 字符串
            df = df.reset_index()
            df['trade_date'] = pd.to_datetime(df['Date']).dt.strftime('%Y%m%d')
            df['ts_code'] = full_code

            # yfinance 列名: Open, High, Low, Close
            df = df.rename(columns={
                'Open': 'open',
                'High': 'high',
                'Low': 'low',
                'Close': 'close'
            })

            # ETF 价格保留小数点后 3 位
            for col in ['open', 'high', 'low', 'close']:
                df[col] = df[col].round(3)

            final_df = df[['ts_code', 'trade_date', 'open', 'high', 'low', 'close']]
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
    用 yfinance 下载 ETF 分钟级数据并存入 SQLite 数据库

    注意：yfinance 分钟级数据有严格的时间窗口限制：
      - 1m: 最近 7 天
      - 5m: 最近 60 天
      超出范围会返回空数据。

    :param start_date: 开始日期，格式为 'YYYY-MM-DD'
    :param end_date: 结束日期，格式为 'YYYY-MM-DD'
    :param period: 级别。可选 '1m', '5m', '15m', '30m', '60m'
    :param db_path: SQLite 数据库文件路径
    """
    # etfs = ["SZ159915", "SH515100", "SH513100", "SH518880"]

    # yfinance interval 参数映射
    interval_map = {
        "1": "1m",
        "5": "5m",
        "15": "15m",
        "30": "30m",
        "60": "60m",
    }
    yf_interval = interval_map.get(period, f"{period}m")

    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    total_records_inserted = 0

    for full_code in etfs:
        yf_ticker = _to_yf_ticker(full_code)
        log(f"Fetching MINUTE (period={yf_interval}) data for {full_code} ({yf_ticker}) from {start_date} to {end_date}...")

        try:
            ticker = yf.Ticker(yf_ticker)
            df = ticker.history(start=start_date, end=end_date, interval=yf_interval)

            if df.empty:
                log(f"No minute data returned for {full_code} ({yf_ticker}).")
                continue

            df = df.reset_index()
            df['trade_time'] = pd.to_datetime(df['Datetime']).dt.strftime('%Y-%m-%d %H:%M:%S')
            df['ts_code'] = full_code

            df = df.rename(columns={
                'Open': 'open',
                'High': 'high',
                'Low': 'low',
                'Close': 'close'
            })

            # ETF 价格保留小数点后 3 位
            for col in ['open', 'high', 'low', 'close']:
                df[col] = df[col].round(3)

            final_df = df[['ts_code', 'trade_time', 'open', 'high', 'low', 'close']]
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
    start = (today - timedelta(days=45)).strftime('%Y-%m-%d')

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
