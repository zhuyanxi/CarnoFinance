import argparse
import math
import os
import sqlite3
from datetime import datetime

import numpy as np
import pandas as pd

# etfs = ["SZ159915", "SH515100", "SH513100", "SH518880"]
# -----------------------------
# 参数（与 JoinQuant 保持一致）
# -----------------------------
# ETF_POOL = ["SZ159915", "SH515100", "SH513100", "SH518880"]
ETF_POOL = ["SZ159915", "SH510050", "SH513100", "SH518880"]

M_DAYS = 25

Q_NOISE = 1e-4
R_NOISE = 1e-2


# -----------------------------
# 卡尔曼滤波（完全保持原逻辑）
# -----------------------------
def kalman_trend_filter(y_series, q_noise=1e-4, r_noise=1e-2):
    """
    状态向量 theta = [slope, intercept]^T
    """

    n = len(y_series)

    x = np.arange(n, dtype=float)
    y = y_series.values

    slope_init, intercept_init = np.polyfit(x, y, 1)

    theta = np.array([[slope_init], [intercept_init]])

    P = np.eye(2) * 1.0

    Q = np.eye(2) * q_noise
    R = r_noise

    for t in range(n):

        H = np.array([[x[t], 1.0]])
        y_obs = y[t]

        P_pred = P + Q

        S = H @ P_pred @ H.T + R
        K = (P_pred @ H.T) / S[0, 0]

        residual = y_obs - H @ theta

        theta = theta + K * residual

        P = (np.eye(2) - K @ H) @ P_pred

    final_slope = theta[0, 0]
    slope_variance = P[0, 0]

    return final_slope, slope_variance


# -----------------------------
# 从SQLite读取最近N天收盘价
# -----------------------------
def load_history(conn, symbol, days, end_date=None):
    """
    假设表结构：

    etf_daily
    ----------
    code
    trade_date
    close

    end_date: str or None, 格式 YYYYMMDD。None = 取最新数据。
    """

    if end_date is None:
        sql = """
        SELECT trade_date, close
        FROM (
            SELECT trade_date, close
            FROM etf_daily_prices
            WHERE ts_code = ?
            ORDER BY trade_date DESC
            LIMIT ?
        )
        ORDER BY trade_date ASC
        """
        df = pd.read_sql_query(sql, conn, params=(symbol, days))
    else:
        sql = """
        SELECT trade_date, close
        FROM (
            SELECT trade_date, close
            FROM etf_daily_prices
            WHERE ts_code = ? AND trade_date <= ?
            ORDER BY trade_date DESC
            LIMIT ?
        )
        ORDER BY trade_date ASC
        """
        df = pd.read_sql_query(sql, conn, params=(symbol, end_date, days))

    if len(df) < days:
        raise RuntimeError(f"{symbol} 数据不足 {days} 天")

    return df


# -----------------------------
# 计算Score（完全保持JoinQuant逻辑）
# -----------------------------
def get_rank(conn, end_date=None):

    score_list = []
    slope_list = []
    var_list = []

    # 第一步：计算Slope和Variance
    for etf in ETF_POOL:

        df = load_history(conn, etf, M_DAYS, end_date=end_date)

        y = np.log(df["close"])

        slope, var = kalman_trend_filter(
            y,
            q_noise=Q_NOISE,
            r_noise=R_NOISE,
        )

        slope_list.append(slope)
        var_list.append(var)

    # 第二步：Variance归一化
    max_var = max(var_list) if max(var_list) > 0 else 1.0
    min_var = min(var_list)

    var_range = max_var - min_var
    if var_range == 0:
        var_range = 1.0

    # 第三步：Score
    for i, etf in enumerate(ETF_POOL):

        slope = slope_list[i]
        var = var_list[i]

        annualized_returns = math.pow(math.exp(slope), 250) - 1

        uncertainty_penalty = (var - min_var) / var_range

        confidence_score = 1.0 - uncertainty_penalty

        score = annualized_returns * confidence_score

        score_list.append(score)

    result = pd.DataFrame({
        "ETF": ETF_POOL,
        "Slope": slope_list,
        "Variance": var_list,
        "Score": score_list,
    })

    result = result.sort_values(
        by="Score",
        ascending=False,
    ).reset_index(drop=True)

    return result


# -----------------------------
# HTML 导出
# -----------------------------
def export_html(result: pd.DataFrame, end_date: str | None, out_dir: str = "data") -> str:
    """生成静态 HTML 报告，返回输出文件路径。"""

    os.makedirs(out_dir, exist_ok=True)
    out_path = os.path.join(out_dir, "etf_kalman_score.html")

    now_str = datetime.now().strftime("%Y-%m-%d %H:%M")
    title = f"ETF 卡尔曼滤波评分 — {end_date or '最新'} ({now_str})"

    # 给 Score 加颜色：正绿负红
    def score_td(score: float) -> str:
        color = "#16a34a" if score >= 0 else "#dc2626"
        pct = f"{score * 100:.2f}%"
        return f'<td style="color:{color};font-weight:600">{pct}</td>'

    rows_html = ""
    for _, row in result.iterrows():
        slope = row["Slope"]
        slope_bp = f"{slope * 10000:.1f} bp"
        rows_html += f"""<tr>
            <td>{row['ETF']}</td>
            <td>{slope:.6f}</td>
            <td>{slope_bp}</td>
            <td>{row['Variance']:.6e}</td>
            {score_td(row['Score'])}
        </tr>"""

    html = f"""<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>{title}</title>
<style>
  * {{ margin: 0; padding: 0; box-sizing: border-box; }}
  body {{
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
    background: #f5f5f5; color: #1e1e1e; padding: 24px;
  }}
  .card {{
    max-width: 720px; margin: 0 auto; background: #fff;
    border-radius: 12px; box-shadow: 0 2px 8px rgba(0,0,0,.08);
    padding: 32px;
  }}
  h1 {{ font-size: 20px; margin-bottom: 8px; }}
  .sub {{ color: #888; font-size: 13px; margin-bottom: 24px; }}
  table {{ width: 100%; border-collapse: collapse; }}
  th, td {{ padding: 10px 12px; text-align: left; border-bottom: 1px solid #eee; }}
  th {{ font-size: 12px; color: #888; text-transform: uppercase; }}
  td {{ font-size: 14px; font-variant-numeric: tabular-nums; }}
  tr:last-child td {{ border-bottom: none; }}
  .footer {{ margin-top: 24px; font-size: 12px; color: #aaa; text-align: center; }}
  .footer a {{ color: #888; }}
</style>
</head>
<body>
<div class="card">
  <h1>📊 ETF 卡尔曼滤波评分</h1>
  <p class="sub">截止日期: {end_date or '最新'} &nbsp;|&nbsp; 生成时间: {now_str}
     &nbsp;|&nbsp; 参数: M={M_DAYS} Q={Q_NOISE} R={R_NOISE}</p>
  <table>
    <thead>
      <tr><th>ETF</th><th>Slope</th><th>年化斜率</th><th>Variance</th><th>Score</th></tr>
    </thead>
    <tbody>
      {rows_html}
    </tbody>
  </table>
  <p class="footer">
    ⚠️ 历史数据回测，不构成投资建议 &nbsp;|&nbsp; 数据源: Yahoo Finance &nbsp;|&nbsp; 算法: Kalman Filter
  </p>
</div>
</body>
</html>"""

    with open(out_path, "w", encoding="utf-8") as f:
        f.write(html)

    print(f"\nHTML 报告已保存至: {out_path}")
    return out_path


# -----------------------------
# 主程序
# -----------------------------
if __name__ == "__main__":

    parser = argparse.ArgumentParser(
        description="卡尔曼滤波 ETF 排名。取 end-date 往前 M_DAYS 天数据计算。"
    )
    parser.add_argument(
        "--end-date",
        type=str,
        default=None,
        help="截止日期，格式 YYYYMMDD。不传则取最新数据。",
    )
    args = parser.parse_args()

    end_date = args.end_date
    if end_date is not None:
        try:
            datetime.strptime(end_date, "%Y%m%d")
        except ValueError:
            parser.error(f"end-date 格式错误: '{end_date}'，应为 YYYYMMDD")

    conn = sqlite3.connect("finance.db")

    result = get_rank(conn, end_date=end_date)

    print(f"end_date = {end_date or '最新'}")
    print(result)

    export_html(result, end_date)

    conn.close()