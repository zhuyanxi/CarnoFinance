import sqlite3
import pandas as pd
import akshare as ak
from datetime import datetime, timedelta
import os

# 配置日志或打印格式
def log(message: str):
    print(f"[{datetime.now().strftime('%Y-%m-%d %H:%M:%S')}] {message}")

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
    下载指定起止日期的 ETF 日线数据并存入 SQLite 数据库
    
    :param start_date: 开始日期，格式为 'YYYYMMDD'
    :param end_date: 结束日期，格式为 'YYYYMMDD'
    :param db_path: SQLite 数据库文件路径
    """
    etfs = ["SZ159915", "SH515100", "SH513100", "SH518880"]
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    total_records_inserted = 0
    
    for full_code in etfs:
        clean_code = "".join(filter(str.isdigit, full_code))
        log(f"Fetching DAILY data for {full_code} from {start_date} to {end_date}...")
        
        try:
            df = ak.fund_etf_hist_em(
                symbol=clean_code,
                period="daily",
                start_date=start_date,
                end_date=end_date,
                adjust=""
            )
            
            if df.empty:
                log(f"No daily data returned for {full_code}.")
                continue
                
            # 规范化日期与字段
            df['trade_date'] = pd.to_datetime(df['日期']).dt.strftime('%Y%m%d')
            df['ts_code'] = full_code
            df = df.rename(columns={
                '开盘': 'open',
                '最高': 'high',
                '最低': 'low',
                '收盘': 'close'
            })
            
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
    period: str = "1", 
    db_path: str = "etf_data.db"
):
    """
    下载指定起止日期及时间段的 ETF 分钟级数据并存入 SQLite 数据库
    
    :param start_date: 开始日期，格式为 'YYYYMMDD' (内部自动转换为东财需要的时间格式)
    :param end_date: 结束日期，格式为 'YYYYMMDD'
    :param period: 级别。可选 '1' (1分钟), '5' (5分钟), '15' (15分钟), '30' (30分钟), '60' (60分钟)
    :param db_path: SQLite 数据库文件路径
    """
    etfs = ["SZ159915", "SH515100", "SH513100", "SH518880"]
    
    # 转换日期格式 YYYYMMDD 为 YYYY-MM-DD 09:30:00 等具体时间，使 API 更加精确
    # 东方财富分钟接口的时间参数需要 "YYYY-MM-DD HH:MM:SS" 格式
    formatted_start = f"{start_date[:4]}-{start_date[4:6]}-{start_date[6:]} 09:30:00"
    formatted_end = f"{end_date[:4]}-{end_date[4:6]}-{end_date[6:]} 15:00:00"
    
    if period == "1":
        log("Note: AKShare 1-minute ETF data only returns the last 5 trading days due to data source limits.")
    
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    total_records_inserted = 0
    
    for full_code in etfs:
        clean_code = "".join(filter(str.isdigit, full_code))
        log(f"Fetching MINUTE (period={period}) data for {full_code} from {formatted_start} to {formatted_end}...")
        
        try:
            # 接口: fund_etf_hist_min_em
            df = ak.fund_etf_hist_min_em(
                symbol=clean_code,
                period=period,
                adjust="",
                start_date=formatted_start,
                end_date=formatted_end
            )
            
            if df.empty:
                log(f"No minute data returned for {full_code}.")
                continue
            
            # 东财分钟级接口返回的列名：
            # 时间, 开盘, 收盘, 最高, 最低, 成交量, 成交额, 最新价
            df['trade_time'] = pd.to_datetime(df['时间']).dt.strftime('%Y-%m-%d %H:%M:%S')
            df['ts_code'] = full_code
            
            df = df.rename(columns={
                '开盘': 'open',
                '最高': 'high',
                '最低': 'low',
                '收盘': 'close'
            })
            
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
    db_file = "finance.db"
    
    # 1. 初始化数据库及表结构
    init_db(db_file)
    
    # 2. 设定起止日期参数
    today = datetime.now()
    end = today.strftime('%Y%m%d')
    start = (today - timedelta(days=30)).strftime('%Y%m%d')
    
    log(f"Starting pipeline. Target database: {db_file}")
    
    # 下载日线数据
    fetch_and_save_etf_data(start_date=start, end_date=end, db_path=db_file)
    
    # 下载 5 分钟线数据（5分钟支持更久的数据跨度；若需要1分钟数据，可以将 period 改为 "1"）
    fetch_and_save_etf_min_data(start_date=start, end_date=end, period="5", db_path=db_file)
    
    # 验证读取分钟数据
    if os.path.exists(db_file):
        conn = sqlite3.connect(db_file)
        test_daily = pd.read_sql("SELECT * FROM etf_daily_prices LIMIT 5", conn)
        test_min = pd.read_sql("SELECT * FROM etf_minute_prices LIMIT 5", conn)
        
        print("\n--- 数据库中 [日线] 前 5 条数据样例 ---")
        print(test_daily)
        
        print("\n--- 数据库中 [分钟线] 前 5 条数据样例 ---")
        print(test_min)
        
        conn.close()