# JoinQuant Short Iron Condor — 报告到策略适配说明

目标: 将 `deep-research-report.md` 的规则固化为 JoinQuant 可运行策略模板。

文件:
- [joinquant_short_iron_condor.py](joinquant_short_iron_condor.py)

要做的事（短句，caveman style）:
- 实现 `get_option_chain`：用 JoinQuant jqdata 查询期权链，返回每个合约 `code,strike,right,expiry,delta,bid,ask`。
- 实现 `place_option_order`：用 `order_option` 或平台对应下单接口（买/卖）。
- 如平台以合约 multiplier 计价，调整 `estimate_position_risk` 中的 `multiplier`。
- 根据券商保证金和合约乘数，调整 `SINGLE_POSITION_RISK` 与 `MAX_ACCOUNT_RISK`。
- 若 JoinQuant 不支持期权回测，需在本地或支持期权的环境运行。

调试步骤:
1. 在 JoinQuant 策略编辑器上传 `joinquant_short_iron_condor.py`。
2. 实现两个核心函数：`get_option_chain`、`place_option_order`。参考平台 API 文档。
3. 运行回测，观察日志：entry/roll/close 事件。
4. 调参：`TARGET_SHORT_DELTA`、`WING_SPREADS`、`PROFIT_TAKE_RATIO`、止损比例、DTE 界值。

注意:
- 模板简化了成交、滑点、手续费、合约乘数、行权与自动行权处理。上线前逐项补足。
- 这是策略框架，不构成投资建议。
