#!/usr/bin/env python3
"""统计交割单中持有 50ETF(510050.XSHG) 的天数。

从 data/kalmanfilter-rsrs.csv 读取交易记录，
按 FIFO 匹配买入和卖出，计算每次持有天数并汇总。
"""

import csv
from datetime import date
from pathlib import Path


CSV_PATH = Path(__file__).resolve().parent.parent / "data" / "kalmanfilter-rsrs.csv"
TARGET = "50ETF(510050.XSHG)"


def load_trades(path: Path, symbol: str) -> list[dict]:
    """加载指定标的的全部交易记录，按委托时间排序。"""
    trades = []
    with open(path, newline="", encoding="utf-8") as f:
        reader = csv.DictReader(f)
        for row in reader:
            if row["标的"] == symbol and row["状态"] == "全部成交":
                trades.append(
                    {
                        "date": date.fromisoformat(row["日期"]),
                        "type": row["交易类型"],  # 买/卖
                        "qty": int(row["成交数量"].replace("股", "").replace(",", "")),
                    }
                )
    # 按日期排序
    trades.sort(key=lambda t: t["date"])
    return trades


def compute_hold_days(trades: list[dict]) -> list[dict]:
    """FIFO 匹配买卖，返回每次持有明细。"""
    buys: list[tuple[date, int]] = []  # [(买入日期, 剩余股数), ...]
    periods: list[dict] = []

    for t in trades:
        if t["type"] == "买":
            buys.append((t["date"], t["qty"]))
        elif t["type"] == "卖":
            # 卖出的股数（卖出记录中 qty 为负数，取绝对值）
            sell_qty = abs(t["qty"])
            sell_date = t["date"]
            remaining = sell_qty
            while remaining > 0 and buys:
                buy_date, buy_qty = buys.pop(0)
                matched = min(buy_qty, remaining)
                days = (sell_date - buy_date).days
                periods.append(
                    {
                        "buy_date": buy_date,
                        "sell_date": sell_date,
                        "shares": matched,
                        "days": days,
                    }
                )
                remaining -= matched
                if buy_qty > matched:
                    # 部分成交，剩余买回队列头
                    buys.insert(0, (buy_date, buy_qty - matched))

    return periods


def main():
    trades = load_trades(CSV_PATH, TARGET)
    periods = compute_hold_days(trades)

    print(f"50ETF(510050.XSHG) 持有分析")
    print(f"{'='*60}")
    total_days = 0
    total_shares_days = 0
    for i, p in enumerate(periods, 1):
        print(
            f"  {i:2d}. {p['buy_date']} → {p['sell_date']}  "
            f"持有 {p['days']:3d} 天,  {p['shares']:,} 股"
        )
        total_days += p["days"]
        total_shares_days += p["days"] * p["shares"]

    print(f"{'='*60}")
    print(f"  交易次数 (一买一卖):     {len(periods)}")
    print(f"  持有天数 (历次合计):     {total_days} 天")

    if periods:
        first_buy = min(p["buy_date"] for p in periods)
        last_sell = max(p["sell_date"] for p in periods)
        print(f"  首次买入:                {first_buy}")
        print(f"  末次卖出:                {last_sell}")
        print(f"  跨越总日历日:            {(last_sell - first_buy).days} 天")

        # 平均每次持有天数
        avg = total_days / len(periods)
        print(f"  平均每次持有天数:         {avg:.1f} 天")

        # 按股数加权平均
        weighted_avg = total_shares_days / sum(p["shares"] for p in periods)
        print(f"  股数加权平均持有天数:     {weighted_avg:.1f} 天")


if __name__ == "__main__":
    main()
