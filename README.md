# CarnoFinance

ETF 数据采集与量化分析工具集，通过 GitHub Actions 自动执行。

## GitHub Actions：ETF 卡尔曼滤波日频评分

**工作流文件：** `.github/workflows/etf_kalman_daily.yml`

### 触发方式

| 触发 | 说明 |
|---|---|
| 定时执行 | 每天 **北京时间 18:00**（UTC 10:00）自动运行 |
| 手动执行 | 在 Actions 页面点击 `Run workflow` 手动触发 |

### 执行流程

```
download_etf.py  →  get_kalmanfilter_score.py  →  提交 HTML →  镜像推送
```

1. **`download_etf.py`** — 从 Yahoo Finance 拉取 5 只 ETF（SZ159915、SH510050、SH513100、SH518880、SH515100）的日线与 5 分钟线数据，存入 SQLite 数据库 `finance.db`。

2. **`get_kalmanfilter_score.py`** — 从 `finance.db` 读取最近 25 天收盘价，用卡尔曼滤波计算各 ETF 的趋势斜率与方差，按 `年化收益率 × 置信度` 公式排名，生成：
   - 控制台输出 DataFrame 排名表
   - 静态 HTML 报告 `data/etf_kalman_score.html`

3. **提交 HTML 报告** — 将 `data/etf_kalman_score.html` 推回本仓库 `main` 分支。

4. **镜像到 zhuyanxi.github.io** — 将 HTML 报告复制为 `etfscore/index.html`，推送到 `zhuyanxi/zhuyanxi.github.io`，可通过 `https://zhuyanxi.github.io/etfscore/` 直接访问。

### 所需 Secrets

在仓库 `Settings → Secrets and variables → Actions` 中配置：

| Secret | 说明 |
|---|---|
| `GH_PAT` | GitHub Personal Access Token，需对 `zhuyanxi/zhuyanxi.github.io` 有 Contents 写权限。仅镜像推送步骤需要，不配置则跳过镜像。 |

### 创建 PAT 步骤

1. GitHub 头像 → **Settings** → **Developer settings** → **Personal access tokens** → **Fine-grained tokens**
2. 点击 **Generate new token**
3. Repository access：选择 **Only select repositories** → 选 `zhuyanxi/zhuyanxi.github.io`
4. Permissions：**Contents** → **Read and write**
5. 生成后复制 token，粘贴到本仓库的 Secrets 中，命名为 `GH_PAT`
