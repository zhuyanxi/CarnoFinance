"""
交割单统计分析脚本
用法: python cmd/trade_stats.py data/kalmanfilter-rsrs.csv
"""
import argparse
import sys
from pathlib import Path

import pandas as pd


def parse_symbol_label(raw: str) -> str:
    """从 '50ETF(510050.XSHG)' 提取标的名称和代码"""
    return raw.strip()


def compute_stats(df: pd.DataFrame, label: str = "") -> dict:
    """
    按已过滤的卖单 DataFrame 计算统计指标。
    df 必须只包含 交易类型 == '卖' 的行。
    """
    pnl = df["平仓盈亏"]  # Series of float

    wins = pnl[pnl > 0]
    losses = pnl[pnl < 0]
    flats = pnl[pnl == 0]

    total = len(pnl)
    win_count = len(wins)
    loss_count = len(losses)
    flat_count = len(flats)

    total_pnl = pnl.sum()
    total_fee = df["手续费"].sum()

    avg_win = wins.mean() if win_count > 0 else 0.0
    avg_loss = losses.mean() if loss_count > 0 else 0.0
    max_win = wins.max() if win_count > 0 else 0.0
    max_loss = losses.min() if loss_count > 0 else 0.0

    win_rate = win_count / total * 100 if total > 0 else 0.0

    # 盈亏比 = 平均盈利 / |平均亏损|
    if avg_loss != 0:
        profit_loss_ratio = avg_win / abs(avg_loss)
    else:
        profit_loss_ratio = float("inf") if avg_win > 0 else 0.0

    # 总盈利 / 总亏损
    total_win_amount = wins.sum()
    total_loss_amount = abs(losses.sum())
    profit_factor = total_win_amount / total_loss_amount if total_loss_amount > 0 else float("inf")

    return {
        "标的": label,
        "交易次数": total,
        "盈利次数": win_count,
        "亏损次数": loss_count,
        "持平次数": flat_count,
        "胜率(%)": round(win_rate, 2),
        "总盈亏": round(total_pnl, 2),
        "总手续费": round(total_fee, 2),
        "净盈亏": round(total_pnl - total_fee, 2),
        "平均盈亏": round(total_pnl / total, 2) if total > 0 else 0.0,
        "平均盈利": round(avg_win, 2),
        "平均亏损": round(avg_loss, 2),
        "盈亏比": round(profit_loss_ratio, 2) if profit_loss_ratio != float("inf") else "∞",
        "最大盈利": round(max_win, 2),
        "最大亏损": round(max_loss, 2),
    }


def format_amount(val) -> str:
    """金额格式化，带颜色 class"""
    if isinstance(val, str):
        return val
    sign = "positive" if val >= 0 else "negative"
    return f"<span class='{sign}'>{val:+,.2f}</span>"


def format_ratio(val) -> str:
    """盈亏比格式化"""
    if isinstance(val, str):
        return val
    return f"{val:.2f}"


def generate_html(results: list[dict], total_stats: dict, source_name: str) -> str:
    """生成完整 HTML 页面"""

    # 概览卡片数据
    net_pnl = total_stats["净盈亏"]
    total_trades = total_stats["交易次数"]
    win_rate = total_stats["胜率(%)"]
    profit_ratio = total_stats["盈亏比"]

    # 标的明细行
    detail_rows = ""
    for r in results:
        detail_rows += f"""
        <tr>
            <td class="sym">{r['标的']}</td>
            <td class="num">{r['交易次数']}</td>
            <td class="num win">{r['盈利次数']}</td>
            <td class="num loss">{r['亏损次数']}</td>
            <td class="num">{r['持平次数']}</td>
            <td class="num">{r['胜率(%)']}%</td>
            <td class="num">{format_amount(r['总盈亏'])}</td>
            <td class="num">{r['总手续费']:,.2f}</td>
            <td class="num">{format_amount(r['净盈亏'])}</td>
            <td class="num">{format_amount(r['平均盈利'])}</td>
            <td class="num">{format_amount(r['平均亏损'])}</td>
            <td class="num">{format_ratio(r['盈亏比'])}</td>
            <td class="num">{format_amount(r['最大盈利'])}</td>
            <td class="num">{format_amount(r['最大亏损'])}</td>
        </tr>"""

    # 总计行
    total_row = f"""
    <tr class="total-row">
        <td class="sym">📊 总计</td>
        <td class="num">{total_stats['交易次数']}</td>
        <td class="num win">{total_stats['盈利次数']}</td>
        <td class="num loss">{total_stats['亏损次数']}</td>
        <td class="num">{total_stats['持平次数']}</td>
        <td class="num">{total_stats['胜率(%)']}%</td>
        <td class="num">{format_amount(total_stats['总盈亏'])}</td>
        <td class="num">{total_stats['总手续费']:,.2f}</td>
        <td class="num">{format_amount(total_stats['净盈亏'])}</td>
        <td class="num">{format_amount(total_stats['平均盈利'])}</td>
        <td class="num">{format_amount(total_stats['平均亏损'])}</td>
        <td class="num">{format_ratio(total_stats['盈亏比'])}</td>
        <td class="num">{format_amount(total_stats['最大盈利'])}</td>
        <td class="num">{format_amount(total_stats['最大亏损'])}</td>
    </tr>"""

    pnl_class = "positive" if net_pnl >= 0 else "negative"

    return f"""<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>交割单统计分析 — {source_name}</title>
<style>
* {{ box-sizing: border-box; margin: 0; padding: 0; }}
body {{
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", sans-serif;
    background: #f5f6fa; color: #2d3436; padding: 24px;
}}
.container {{ max-width: 1400px; margin: 0 auto; }}
h1 {{ text-align: center; font-size: 24px; margin-bottom: 8px; color: #2d3436; }}
.subtitle {{ text-align: center; font-size: 13px; color: #636e72; margin-bottom: 24px; }}

/* 概览卡片 */
.cards {{ display: grid; grid-template-columns: repeat(auto-fit, minmax(180px, 1fr)); gap: 16px; margin-bottom: 24px; }}
.card {{
    background: #fff; border-radius: 10px; padding: 20px; text-align: center;
    box-shadow: 0 1px 4px rgba(0,0,0,0.06);
}}
.card .label {{ font-size: 12px; color: #636e72; margin-bottom: 6px; }}
.card .value {{ font-size: 28px; font-weight: 700; }}
.card .value.positive {{ color: #00b894; }}
.card .value.negative {{ color: #d63031; }}

/* 表格 */
.table-wrap {{
    background: #fff; border-radius: 10px; overflow-x: auto;
    box-shadow: 0 1px 4px rgba(0,0,0,0.06);
}}
table {{ width: 100%; border-collapse: collapse; font-size: 13px; white-space: nowrap; }}
thead th {{
    background: #2d3436; color: #fff; padding: 12px 10px; text-align: right;
    position: sticky; top: 0;
}}
thead th:first-child {{ text-align: left; }}
td {{ padding: 10px; text-align: right; border-bottom: 1px solid #eee; }}
td.sym {{ text-align: left; font-weight: 500; }}
td.win {{ color: #00b894; }}
td.loss {{ color: #d63031; }}
tr.total-row td {{
    font-weight: 700; border-top: 2px solid #2d3436;
    background: #f8f9fa;
}}
tr:hover {{ background: #f8f9fa; }}
tr.total-row:hover {{ background: #e9ecef; }}

/* 正负颜色 */
.positive {{ color: #00b894; font-weight: 500; }}
.negative {{ color: #d63031; font-weight: 500; }}

.footer {{ text-align: center; font-size: 11px; color: #b2bec3; margin-top: 24px; }}
</style>
</head>
<body>
<div class="container">

<h1>📈 交割单统计分析</h1>
<p class="subtitle">数据来源: {source_name}</p>

<div class="cards">
    <div class="card">
        <div class="label">总交易回合</div>
        <div class="value">{total_trades}</div>
    </div>
    <div class="card">
        <div class="label">胜率</div>
        <div class="value">{win_rate}%</div>
    </div>
    <div class="card">
        <div class="label">净盈亏</div>
        <div class="value {pnl_class}">{net_pnl:+,.2f}</div>
    </div>
    <div class="card">
        <div class="label">盈亏比</div>
        <div class="value">{format_ratio(profit_ratio)}</div>
    </div>
</div>

<div class="table-wrap">
<table>
<thead>
<tr>
    <th>标的</th>
    <th>交易次数</th>
    <th>盈利次数</th>
    <th>亏损次数</th>
    <th>持平次数</th>
    <th>胜率(%)</th>
    <th>总盈亏</th>
    <th>总手续费</th>
    <th>净盈亏</th>
    <th>平均盈利</th>
    <th>平均亏损</th>
    <th>盈亏比</th>
    <th>最大盈利</th>
    <th>最大亏损</th>
</tr>
</thead>
<tbody>
{detail_rows}
{total_row}
</tbody>
</table>
</div>

<p class="footer">Generated by trade_stats.py</p>
</div>
</body>
</html>"""


def main():
    parser = argparse.ArgumentParser(description="交割单统计分析 — 输出 HTML 静态页面")
    parser.add_argument("csv", type=str, help="交割单 CSV 文件路径")
    parser.add_argument("-o", "--output", type=str, default=None,
                        help="输出 HTML 文件路径（默认: <csv文件名>.html）")
    args = parser.parse_args()

    csv_path = Path(args.csv)
    if not csv_path.exists():
        print(f"文件不存在: {csv_path}")
        sys.exit(1)

    output_path = Path(args.output) if args.output else csv_path.with_suffix(".html")

    df = pd.read_csv(csv_path)

    # 只取卖单 — 每笔卖单的平仓盈亏 = 一次完整交易盈亏
    sells = df[df["交易类型"] == "卖"].copy()
    sells["平仓盈亏"] = pd.to_numeric(sells["平仓盈亏"], errors="coerce")
    sells["手续费"] = pd.to_numeric(sells["手续费"], errors="coerce")

    if sells.empty:
        print("没有卖单记录，无法统计。")
        sys.exit(0)

    # 按标的名称分组统计
    symbols = sells["标的"].unique()
    results = []
    for sym in symbols:
        sub = sells[sells["标的"] == sym]
        results.append(compute_stats(sub, label=sym))

    # 按交易次数降序排列
    results.sort(key=lambda r: r["交易次数"], reverse=True)

    # 总计行
    total_stats = compute_stats(sells, label="【总计】")

    html = generate_html(results, total_stats, csv_path.name)

    output_path.write_text(html, encoding="utf-8")
    print(f"HTML 已生成: {output_path}")


if __name__ == "__main__":
    main()
