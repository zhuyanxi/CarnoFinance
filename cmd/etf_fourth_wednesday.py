"""
统计 ETF 收盘价从每月第四个星期三到次月第四个星期三的涨跌幅分布。

用法:
    python etf_fourth_wednesday.py SZ159915
    python etf_fourth_wednesday.py SZ159915 --db finance.db
    python etf_fourth_wednesday.py SZ159915 --plot
"""

import sqlite3
import sys
import os
from datetime import date, timedelta
from calendar import monthrange
from collections import OrderedDict

import pandas as pd
import numpy as np


def get_db_path(db_arg: str) -> str:
    """解析数据库路径，默认为脚本同目录下的 finance.db"""
    if db_arg:
        return db_arg
    script_dir = os.path.dirname(os.path.abspath(__file__))
    return os.path.join(script_dir, "finance.db")


def load_etf_data(db_path: str, ts_code: str) -> pd.DataFrame:
    """从 SQLite 加载指定 ETF 的日线数据，按交易日排序"""
    conn = sqlite3.connect(db_path)
    df = pd.read_sql(
        "SELECT ts_code, trade_date, close FROM etf_daily_prices "
        "WHERE ts_code = ? ORDER BY trade_date ASC",
        conn,
        params=(ts_code,),
    )
    conn.close()

    if df.empty:
        print(f"[错误] 未找到 {ts_code} 的数据。请检查 ETF 代码和数据库。")
        sys.exit(1)

    df["trade_date"] = pd.to_datetime(df["trade_date"], format="%Y%m%d")
    df = df.set_index("trade_date").sort_index()
    return df


def find_fourth_wednesday(year: int, month: int) -> date:
    """返回指定年月的第四个星期三"""
    first_day = date(year, month, 1)
    # 星期三 weekday() = 2
    days_until_wed = (2 - first_day.weekday()) % 7
    first_wed = first_day + timedelta(days=days_until_wed)
    return first_wed + timedelta(weeks=3)


def find_nearest_trading_day(target: date, available_dates: pd.DatetimeIndex) -> date:
    """
    找到不超过 target 的最近交易日。
    若 target 当天有交易，返回 target；否则向前找最近的交易日。
    """
    available = available_dates[available_dates <= pd.Timestamp(target)]
    if available.empty:
        return None
    return available[-1].date()


def compute_monthly_wednesday_returns(
    df: pd.DataFrame, ts_code: str
) -> pd.DataFrame:
    """
    计算每月第四个星期三到次月第四个星期三的涨跌幅。

    返回 DataFrame，列: from_date, to_date, from_close, to_close, pct_change
    """
    available_dates = df.index
    closes = df["close"]

    # 确定数据覆盖的年月范围
    start_date = available_dates[0].date()
    end_date = available_dates[-1].date()

    # 生成所有月份的第四个星期三
    all_wednesdays = OrderedDict()
    for y in range(start_date.year, end_date.year + 1):
        for m in range(1, 13):
            fourth_wed = find_fourth_wednesday(y, m)
            if start_date <= fourth_wed <= end_date:
                trading_day = find_nearest_trading_day(fourth_wed, available_dates)
                if trading_day is not None:
                    all_wednesdays[(y, m)] = trading_day

    if len(all_wednesdays) < 2:
        print("[错误] 数据不足，至少需要两个月的第四个星期三数据。")
        sys.exit(1)

    # 计算相邻两个第四个星期三之间的涨跌幅
    months = list(all_wednesdays.keys())
    results = []
    for i in range(len(months) - 1):
        ym_from = months[i]
        ym_to = months[i + 1]
        date_from = all_wednesdays[ym_from]
        date_to = all_wednesdays[ym_to]

        close_from = closes.loc[pd.Timestamp(date_from)]
        close_to = closes.loc[pd.Timestamp(date_to)]
        pct = (close_to - close_from) / close_from * 100

        results.append(
            {
                "from_month": f"{ym_from[0]}-{ym_from[1]:02d}",
                "to_month": f"{ym_to[0]}-{ym_to[1]:02d}",
                "from_date": date_from.strftime("%Y-%m-%d"),
                "to_date": date_to.strftime("%Y-%m-%d"),
                "from_close": round(close_from, 3),
                "to_close": round(close_to, 3),
                "pct_change": round(pct, 2),
            }
        )

    return pd.DataFrame(results)


def print_distribution(summary: pd.DataFrame, ts_code: str):
    """打印涨跌幅分布统计"""
    pct = summary["pct_change"]

    print(f"\n{'=' * 60}")
    print(f"  {ts_code} 每月第四个星期三 → 次月第四个星期三 涨跌幅分布")
    print(f"{'=' * 60}")
    print(f"  样本数:       {len(pct)}")
    print(f"  均值:         {pct.mean():.2f}%")
    print(f"  中位数:       {pct.median():.2f}%")
    print(f"  标准差:       {pct.std():.2f}%")
    print(f"  最小值:       {pct.min():.2f}%")
    print(f"  最大值:       {pct.max():.2f}%")
    print(f"  25% 分位:     {pct.quantile(0.25):.2f}%")
    print(f"  75% 分位:     {pct.quantile(0.75):.2f}%")
    print(f"  上涨月份:     {(pct > 0).sum()} / {len(pct)}  ({(pct > 0).mean() * 100:.1f}%)")
    print(f"  下跌月份:     {(pct < 0).sum()} / {len(pct)}  ({(pct < 0).mean() * 100:.1f}%)")

    # 涨跌幅区间分布
    bins = [-np.inf, -10, -5, -3, -1, 0, 1, 3, 5, 10, np.inf]
    labels = ["<-10%", "-10~-5%", "-5~-3%", "-3~-1%", "-1~0%", "0~1%", "1~3%", "3~5%", "5~10%", ">10%"]
    print(f"\n  {'区间':<12s} {'次数':>6s}  {'占比':>8s}")
    print(f"  {'-' * 26}")
    for i in range(len(bins) - 1):
        count = ((pct > bins[i]) & (pct <= bins[i + 1])).sum()
        if count > 0:
            print(f"  {labels[i]:<12s} {count:>6d}  {count / len(pct) * 100:>7.1f}%")

    print(f"{'=' * 60}\n")

    # 打印明细表
    print(summary.to_string(index=False))


def plot_distribution(summary: pd.DataFrame, ts_code: str):
    """绘制涨跌幅分布直方图"""
    try:
        import matplotlib.pyplot as plt
        import matplotlib
        matplotlib.use("TkAgg")
    except ImportError:
        print("[提示] matplotlib 未安装，跳过绘图。pip install matplotlib")
        return

    pct = summary["pct_change"]

    fig, axes = plt.subplots(1, 2, figsize=(14, 5))
    fig.suptitle(
        f"{ts_code} 每月第四个星期三 → 次月第四个星期三 涨跌幅分布",
        fontsize=13,
        fontweight="bold",
    )

    # 直方图
    ax1 = axes[0]
    ax1.hist(pct, bins=20, edgecolor="white", color="steelblue", alpha=0.85)
    ax1.axvline(pct.mean(), color="red", linestyle="--", linewidth=1.5, label=f"均值 {pct.mean():.2f}%")
    ax1.axvline(0, color="gray", linestyle="-", linewidth=1)
    ax1.set_xlabel("涨跌幅 (%)")
    ax1.set_ylabel("频次")
    ax1.set_title("直方图")
    ax1.legend()

    # 累计收益曲线
    ax2 = axes[1]
    cumulative = (1 + pct / 100).cumprod()
    ax2.plot(
        range(len(cumulative)),
        cumulative.values,
        marker="o",
        linewidth=1.5,
        color="steelblue",
    )
    ax2.axhline(1, color="gray", linestyle="--", linewidth=1)
    ax2.set_xlabel("周期序号")
    ax2.set_ylabel("累计净值")
    ax2.set_title(f"累计净值 (初始=1, 终值={cumulative.iloc[-1]:.4f})")

    plt.tight_layout()
    plt.show()


def main():
    if len(sys.argv) < 2:
        print("用法: python etf_fourth_wednesday.py <ETF代码> [--db <数据库路径>] [--plot]")
        print("示例: python etf_fourth_wednesday.py SZ159915")
        print("      python etf_fourth_wednesday.py SZ159915 --db finance.db --plot")
        sys.exit(1)

    ts_code = sys.argv[1].upper()

    db_arg = None
    plot_flag = False
    args = sys.argv[2:]
    i = 0
    while i < len(args):
        if args[i] == "--db" and i + 1 < len(args):
            db_arg = args[i + 1]
            i += 2
        elif args[i] == "--plot":
            plot_flag = True
            i += 1
        else:
            i += 1

    db_path = get_db_path(db_arg)
    if not os.path.exists(db_path):
        print(f"[错误] 数据库文件不存在: {db_path}")
        sys.exit(1)

    print(f"[信息] ETF: {ts_code}, 数据库: {db_path}")

    df = load_etf_data(db_path, ts_code)
    print(f"[信息] 加载 {len(df)} 条日线记录，"
          f"日期范围: {df.index[0].strftime('%Y-%m-%d')} ~ {df.index[-1].strftime('%Y-%m-%d')}")

    summary = compute_monthly_wednesday_returns(df, ts_code)
    print_distribution(summary, ts_code)

    if plot_flag:
        plot_distribution(summary, ts_code)


if __name__ == "__main__":
    main()
