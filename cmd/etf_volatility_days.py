"""
统计 ETF 每月第四个星期三收盘价起，波动达到 ±10% 所需的天数分布。

以每月第四个星期三收盘价为基准，向后逐日扫描，记录收盘价首次偏离基准
超过 ±10% 所需的自然天数。若数据范围内未触及则标记为"未触及"。

用法:
    python etf_volatility_days.py SZ159915
    python etf_volatility_days.py SZ159915 --threshold 0.15 --db finance.db
    python etf_volatility_days.py SZ159915 --plot
"""

import sqlite3
import sys
import os
from datetime import date, timedelta
from calendar import monthrange

import pandas as pd
import numpy as np


def get_db_path(db_arg: str) -> str:
    if db_arg:
        return db_arg
    script_dir = os.path.dirname(os.path.abspath(__file__))
    return os.path.join(script_dir, "finance.db")


def load_etf_data(db_path: str, ts_code: str) -> pd.DataFrame:
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
    days_until_wed = (2 - first_day.weekday()) % 7
    first_wed = first_day + timedelta(days=days_until_wed)
    return first_wed + timedelta(weeks=3)


def find_nearest_trading_day(target: date, available_dates: pd.DatetimeIndex) -> date:
    """找到不超过 target 的最近交易日"""
    available = available_dates[available_dates <= pd.Timestamp(target)]
    if available.empty:
        return None
    return available[-1].date()


def compute_days_to_threshold(
    df: pd.DataFrame, ts_code: str, threshold: float
) -> pd.DataFrame:
    """
    对每月第四个星期三，计算收盘价首次突破 ±threshold 所需的自然天数。

    返回 DataFrame，列:
      month, start_date, start_close, hit_date, hit_close, hit_pct, days
    若未触及，hit_* 列为空，days 为 NaN。
    """
    available_dates = df.index
    closes = df["close"]

    start_date = available_dates[0].date()
    end_date = available_dates[-1].date()

    # 构建所有月份的第四个星期三（作为起点）
    start_points = []
    for y in range(start_date.year, end_date.year + 1):
        for m in range(1, 13):
            fourth_wed = find_fourth_wednesday(y, m)
            if start_date <= fourth_wed <= end_date:
                trading_day = find_nearest_trading_day(fourth_wed, available_dates)
                if trading_day is not None:
                    start_points.append(((y, m), trading_day))

    if not start_points:
        print("[错误] 数据不足，无法找到任何第四个星期三。")
        sys.exit(1)

    results = []
    for (ym, start_d) in start_points:
        start_close = closes.loc[pd.Timestamp(start_d)]

        # 在 start_d 之后的所有交易日中，找首次突破 threshold 的日期
        future = available_dates[available_dates > pd.Timestamp(start_d)]
        hit_d = None
        hit_close = None
        hit_pct = None

        for fd in future:
            fc = closes.loc[fd]
            pct = (fc - start_close) / start_close
            if abs(pct) >= threshold:
                hit_d = fd.date()
                hit_close = fc
                hit_pct = pct
                break

        results.append(
            {
                "month": f"{ym[0]}-{ym[1]:02d}",
                "start_date": start_d.strftime("%Y-%m-%d"),
                "start_close": round(start_close, 3),
                "hit_date": hit_d.strftime("%Y-%m-%d") if hit_d else None,
                "hit_close": round(hit_close, 3) if hit_close is not None else None,
                "hit_pct": round(hit_pct * 100, 2) if hit_pct is not None else None,
                "days": (hit_d - start_d).days if hit_d else None,
            }
        )

    return pd.DataFrame(results)


def print_distribution(summary: pd.DataFrame, ts_code: str, threshold: float):
    """打印天数分布统计"""
    hit = summary.dropna(subset=["days"])
    missed = summary[summary["days"].isna()]
    days = hit["days"].astype(int)
    pct_threshold = threshold * 100

    print(f"\n{'=' * 65}")
    print(f"  {ts_code} 每月第四个星期三起，收盘价波动达 ±{pct_threshold:.0f}%")
    print(f"  所需自然天数分布")
    print(f"{'=' * 65}")
    print(f"  总起始月份:   {len(summary)}")
    print(f"  触及 ±{pct_threshold:.0f}%:  {len(hit)} / {len(summary)}  ({len(hit)/len(summary)*100:.1f}%)")
    print(f"  未触及:       {len(missed)} / {len(summary)}  ({len(missed)/len(summary)*100:.1f}%)")

    if len(hit) == 0:
        print("\n  无触及记录，无法计算分布。")
        print(f"{'=' * 65}\n")
        return

    print(f"\n  --- 天数分布 (触及样本 {len(hit)} 个) ---")
    print(f"  均值:         {days.mean():.1f} 天")
    print(f"  中位数:       {days.median():.0f} 天")
    print(f"  标准差:       {days.std():.1f} 天")
    print(f"  最小:         {days.min()} 天")
    print(f"  最大:         {days.max()} 天")
    print(f"  25% 分位:     {days.quantile(0.25):.0f} 天")
    print(f"  75% 分位:     {days.quantile(0.75):.0f} 天")

    # 天数区间分布
    bin_edges = [0, 7, 14, 21, 30, 60, 90, 180, 365, np.inf]
    bin_labels = ["0-7天", "7-14天", "14-21天", "21-30天", "30-60天",
                   "60-90天", "90-180天", "180-365天", ">365天"]
    print(f"\n  {'区间':<14s} {'次数':>6s}  {'占比':>8s}")
    print(f"  {'-' * 28}")
    for i in range(len(bin_edges) - 1):
        count = ((days > bin_edges[i]) & (days <= bin_edges[i + 1])).sum()
        if count > 0:
            print(f"  {bin_labels[i]:<14s} {count:>6d}  {count / len(days) * 100:>7.1f}%")

    # 涨跌方向
    up = (hit["hit_pct"] > 0).sum()
    down = (hit["hit_pct"] < 0).sum()
    print(f"\n  --- 触及方向 ---")
    print(f"  向上触及 (+{pct_threshold:.0f}%):  {up} / {len(hit)}  ({up/len(hit)*100:.1f}%)")
    print(f"  向下触及 (-{pct_threshold:.0f}%):  {down} / {len(hit)}  ({down/len(hit)*100:.1f}%)")

    # 未触及月份列表
    if len(missed) > 0:
        print(f"\n  --- 未触及的起始月份 ({len(missed)} 个) ---")
        print(f"  {', '.join(missed['month'].tolist())}")

    print(f"{'=' * 65}\n")

    # 打印触及明细表
    print("[触及明细]")
    display_cols = ["month", "start_date", "start_close", "hit_date",
                    "hit_close", "hit_pct", "days"]
    print(hit[display_cols].to_string(index=False))

    if len(missed) > 0:
        print(f"\n[未触及月份: {len(missed)} 个]")


def plot_distribution(summary: pd.DataFrame, ts_code: str, threshold: float):
    """绘制天数分布直方图"""
    try:
        import matplotlib.pyplot as plt
        import matplotlib
        matplotlib.use("TkAgg")
    except ImportError:
        print("[提示] matplotlib 未安装，跳过绘图。pip install matplotlib")
        return

    hit = summary.dropna(subset=["days"])
    if hit.empty:
        print("[提示] 无触及数据，跳过绘图。")
        return

    days = hit["days"].astype(int)
    pct_threshold = threshold * 100

    fig, axes = plt.subplots(1, 2, figsize=(14, 5))
    fig.suptitle(
        f"{ts_code} 第四个星期三起波动达 ±{pct_threshold:.0f}% 所需天数分布",
        fontsize=13,
        fontweight="bold",
    )

    # 直方图
    ax1 = axes[0]
    bins = [0, 7, 14, 21, 30, 45, 60, 90, 120, 180, 270, 365, 730]
    ax1.hist(days, bins=bins, edgecolor="white", color="steelblue", alpha=0.85)
    ax1.axvline(days.median(), color="red", linestyle="--", linewidth=1.5,
                label=f"中位数 {days.median():.0f} 天")
    ax1.set_xlabel("自然天数")
    ax1.set_ylabel("频次")
    ax1.set_title("直方图")
    ax1.legend()

    # 触及方向饼图
    ax2 = axes[1]
    up = (hit["hit_pct"] > 0).sum()
    down = (hit["hit_pct"] < 0).sum()
    ax2.pie(
        [up, down],
        labels=[f"上涨 +{pct_threshold:.0f}%\n({up})", f"下跌 -{pct_threshold:.0f}%\n({down})"],
        colors=["#c44e52", "#4c72b0"],
        autopct="%1.1f%%",
        startangle=90,
    )
    ax2.set_title("触及方向分布")

    plt.tight_layout()
    plt.show()


def main():
    if len(sys.argv) < 2:
        print("用法: python etf_volatility_days.py <ETF代码> [--threshold <阈值>] [--db <数据库路径>] [--plot]")
        print("示例: python etf_volatility_days.py SZ159915")
        print("      python etf_volatility_days.py SZ159915 --threshold 0.15 --plot")
        sys.exit(1)

    ts_code = sys.argv[1].upper()
    threshold = 0.10  # 默认 ±10%
    db_arg = None
    plot_flag = False

    args = sys.argv[2:]
    i = 0
    while i < len(args):
        if args[i] == "--threshold" and i + 1 < len(args):
            threshold = float(args[i + 1])
            i += 2
        elif args[i] == "--db" and i + 1 < len(args):
            db_arg = args[i + 1]
            i += 2
        elif args[i] == "--plot":
            plot_flag = True
            i += 1
        else:
            i += 1

    if threshold <= 0 or threshold >= 1:
        print("[错误] threshold 应在 (0, 1) 区间，例如 0.10 表示 ±10%")
        sys.exit(1)

    db_path = get_db_path(db_arg)
    if not os.path.exists(db_path):
        print(f"[错误] 数据库文件不存在: {db_path}")
        sys.exit(1)

    print(f"[信息] ETF: {ts_code}, 阈值: ±{threshold*100:.0f}%, 数据库: {db_path}")

    df = load_etf_data(db_path, ts_code)
    print(f"[信息] 加载 {len(df)} 条日线记录，"
          f"日期范围: {df.index[0].strftime('%Y-%m-%d')} ~ {df.index[-1].strftime('%Y-%m-%d')}")

    summary = compute_days_to_threshold(df, ts_code, threshold)
    print_distribution(summary, ts_code, threshold)

    if plot_flag:
        plot_distribution(summary, ts_code, threshold)


if __name__ == "__main__":
    main()
