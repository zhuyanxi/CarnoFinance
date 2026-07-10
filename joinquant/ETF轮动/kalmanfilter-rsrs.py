# 克隆自聚宽文章：https://www.joinquant.com/post/42673
# 标题：【回顾3】ETF策略之核心资产轮动
# 作者：wywy1995

import numpy as np
import pandas as pd

#初始化函数 
def initialize(context):
    # 设定基准
    set_benchmark('399300.XSHE')
    # 用真实价格交易
    set_option('use_real_price', True)
    # 打开防未来函数
    set_option("avoid_future_data", True)
    # 设置滑点 https://www.joinquant.com/view/community/detail/a31a822d1cfa7e83b1dda228d4562a70
    # set_slippage(FixedSlippage(0.002))
    set_slippage(FixedSlippage(0))
    # 设置交易成本
    set_order_cost(OrderCost(open_tax=0, close_tax=0, open_commission=0.00015, close_commission=0.00015, close_today_commission=0, min_commission=5), type='fund')
    # 过滤一定级别的日志
    log.set_level('system', 'error')
    
    # 参数
    g.etf_pool = [
        # '510050.XSHG', # 50ETF
        '518880.XSHG', #黄金ETF（大宗商品）
        '513100.XSHG', #纳指100（海外资产）
        '159915.XSHE', #创业板100（成长股，科技股，题材性，中小盘
        '510880.XSHG', #红利ETF（价值股，蓝筹股，防御性，中大盘
        # '511010.XSHG', # 避险资产：10年国债ETF
        # '515100.XSHG', #低波红利100ETF
        # '588000.XSHG', #科创50ETF
        # '512890.XSHG',
        # '159985.XSHE', #豆粕ETF
        # '159919.XSHE', #沪深300ETF
        # '512100.XSHG', #中证1000ETF
        # '162411.XSHE', #华宝油气
        # '159980.XSHE', #有色期货
        # '159981.XSHE', #能源化工期货
    ]
    g.m_days = 25 #动量参考天数
    # 卡尔曼滤波超参数
    g.q_noise = 1e-4  # 过程噪声协方差（值越大，算法越相信新数据，对趋势反应更快）
    g.r_noise = 1e-2  # 观测噪声方差（值越大，越不相信单一的价格波动，滤波曲线越平滑）
    run_daily(trade, '09:30') #每天运行确保即时捕捉动量变化
    # run_daily(trade, '13:10')

# 使用卡尔曼滤波估计动态对数价格的斜率与不确定性
def kalman_trend_filter(y_series, q_noise=1e-4, r_noise=1e-2):
    """
    状态向量 theta = [slope, intercept]^T
    观测方程: y_t = [t, 1] * theta + v_t, v_t ~ N(0, r_noise)
    状态转移方程: theta_t = theta_{t-1} + w_t, w_t ~ N(0, Q)
    """
    n = len(y_series)
    x = np.arange(n, dtype=float)
    y = y_series.values
    
    # 1. 采用 OLS 的拟合结果作为卡尔曼滤波的初始状态
    slope_init, intercept_init = np.polyfit(x, y, 1)
    theta = np.array([[slope_init], [intercept_init]])
    
    # 初始协方差矩阵 P (初始不确定性设得稍大，让滤波器快速收敛)
    P = np.eye(2) * 1.0
    
    # 噪声参数
    Q = np.eye(2) * q_noise
    R = r_noise
    
    # 迭代滤波
    for t in range(n):
        # 观测矩阵 H_t = [[t, 1]]
        H = np.array([[x[t], 1.0]])
        y_obs = y[t]
        
        # 预测步 (Time Update)
        # theta_pred = theta
        P_pred = P + Q
        
        # 更新步 (Measurement Update)
        S = H @ P_pred @ H.T + R  # 创新协方差
        K = (P_pred @ H.T) / S[0, 0]  # 卡尔曼增益
        
        # 更新状态估计与协方差
        residual = y_obs - H @ theta
        theta = theta + K * residual
        P = (np.eye(2) - K @ H) @ P_pred
        
    # 返回最终提炼的动态斜率, 以及斜率估计的方差（代表不确定性）
    final_slope = theta[0, 0]
    slope_variance = P[0, 0]
    
    return final_slope, slope_variance

# 基于卡尔曼滤波动量和不确定性惩罚打分的轮动因子
def get_rank(etf_pool):
    score_list = []
    slope_list = []
    var_list = []
    
    # 第一步：计算各资产的卡尔曼斜率和估计方差
    for etf in etf_pool:
        df = attribute_history(etf, g.m_days, '1d', ['close'])
        y = np.log(df.close)
        
        # 运行卡尔曼滤波器
        slope, var = kalman_trend_filter(y, q_noise=g.q_noise, r_noise=g.r_noise)
        slope_list.append(slope)
        var_list.append(var)
        
    # 第二步：对估计方差进行归一化，以便作为风险惩罚项
    max_var = max(var_list) if max(var_list) > 0 else 1.0
    min_var = min(var_list)
    var_range = (max_var - min_var) if (max_var - min_var) > 0 else 1.0
    
    for i, etf in enumerate(etf_pool):
        slope = slope_list[i]
        var = var_list[i]
        
        # 将卡尔曼斜率转化为年化收益率
        annualized_returns = math.pow(math.exp(slope), 250) - 1
        
        # 归一化估计方差，代表估计斜率的不确定性得分 (0 ~ 1)
        # 不确定性越大，惩罚力度越强
        uncertainty_penalty = (var - min_var) / var_range
        confidence_score = 1.0 - uncertainty_penalty
        
        # 综合得分 = 年化收益率 * 置信度评分
        score = annualized_returns * confidence_score
        score_list.append(score)
        
    df = pd.DataFrame(index=etf_pool, data={'score': score_list})
    df = df.sort_values(by='score', ascending=False)
    print(df)
    rank_list = list(df.index)
    return rank_list

# 交易
def trade(context):
    target_num = 1
    target_list = get_rank(g.etf_pool)[:target_num]
    # 卖出    
    hold_list = list(context.portfolio.positions)
    for etf in hold_list:
        if etf not in target_list:
            order_target_value(etf, 0)
            print('卖出' + str(etf))
        else:
            print('继续持有' + str(etf))
    # 买入
    hold_list = list(context.portfolio.positions)
    if len(hold_list) < target_num:
        value = context.portfolio.available_cash / (target_num - len(hold_list))
        for etf in target_list:
            if context.portfolio.positions[etf].total_amount == 0:
                order_target_value(etf, value)
                print('买入' + str(etf))
