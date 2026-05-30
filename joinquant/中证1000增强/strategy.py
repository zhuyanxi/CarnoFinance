# 导入函数库
from jqdata import *
from functools import reduce
import numpy as np

# 初始化函数，设定基准等等
def initialize(context):
    # set_benchmark('399101.XSHE')
    set_benchmark('000852.XSHG')
    # 开启避免未来函数
    set_option("avoid_future_data", True)
    # 开启动态复权模式(真实价格)
    set_option('use_real_price', True)
    # 强制撮合，仅支持限价单。使用限价单进行委托时将不对委托价格和成交数量进行任何检查而直接成交
    # set_option("match_by_signal", True)
    # 设置滑点
    # set_slippage(FixedSlippage(0))
    set_slippage(FixedSlippage(0.05))
    # 输出内容到日志 log.info()
    log.info('初始函数开始运行且全局只运行一次')
    # 过滤掉order系列API产生的比error级别低的log
    log.set_level('order', 'error')

    ### 股票相关设定 ###
    # 股票类每笔交易时的手续费是：买入时佣金万分之三，卖出时佣金万分之三加千分之一印花税, 每笔交易佣金最低扣5块钱
    set_order_cost(OrderCost(close_tax=0.001, open_commission=0.0003, close_commission=0.0003, min_commission=5), type='stock')

    g.i=0
    g.maxHoldCount = 5
    g.check_out_lists = []
    
    run_weekly(before_market_open, 1, time='9:00', reference_security='000852.XSHG')
    run_weekly(doInvestSmallMarketCap, 1, time='09:30', reference_security='000852.XSHG')
    
    
def doInvestSmallMarketCap(context):    
    codes = g.check_out_lists
    
    log.info("每z第一个交易日 - 卖出不在筛选出的列表里的股票")
    for hold in context.portfolio.long_positions:
        if hold not in codes:
            print("sell stock: {}".format(hold))
            order_target(hold, 0)
        else:
            log.info("已持有[%s]" % (hold))
        # print("sell stock: {}".format(hold))
        # order_target(hold, 0)
    log.info("当前剩余现金 %s" % (context.portfolio.available_cash))
    
    curHolds = context.portfolio.long_positions
    log.info("当前持仓 %s" % (curHolds))
    if len(codes)>len(curHolds):
        pieceCash = context.portfolio.available_cash/(len(codes)-len(curHolds))
        log.info("当前每份现金 %s" % (pieceCash))
        log.info("需要持有的股票代码 %s" % codes)
        for code in codes:
            if code in curHolds:
                continue
            order_value(code, pieceCash)


## 开盘前运行函数
def before_market_open(context):
    operationDay = context.previous_date
    
    stocksContainer = get_index_stocks('000852.XSHG', operationDay) #中证1000
    stocksContainer=[s for s in stocksContainer if not s.startswith(('688', '689'))]
    log.info("股票池剔除科创板: %s" % (len(stocksContainer)))
    if len(stocksContainer)==0:
        return
    
    # 筛选条件：
    rets = get_fundamentals(query(
        valuation.day,
        valuation.code,
        valuation.market_cap
        ).filter(
            income.np_parent_company_owners>1000000, #最近一个季度的归母净利润 10000000
            valuation.code.in_(stocksContainer)
        ).order_by(
            valuation.market_cap.asc()
        ).limit(g.maxHoldCount*30))
    codes = rets['code'].tolist()    
    codes = filter_new_stock(context,codes)
    codes = filter_st_stock(codes)
    codes = filter_paused_stock(codes)
    codes = filter_limitup_stock(context,codes)
    codes = filter_limitdown_stock(context,codes)
    log.info("可投资股票池数量: %s" % (len(codes)))
    codes=codes[:g.maxHoldCount]
    g.check_out_lists = codes
    

#2-2 过滤ST及其他具有退市标签的股票
def filter_st_stock(stock_list):
    current_data = get_current_data()
    return [stock for stock in stock_list
            if not current_data[stock].is_st
            and 'ST' not in current_data[stock].name
            and '*' not in current_data[stock].name
            and '退' not in current_data[stock].name]

#2-6 过滤次新股
def filter_new_stock(context,stock_list):
    yesterday = context.previous_date
    return [stock for stock in stock_list if not yesterday - get_security_info(stock).start_date < datetime.timedelta(days=375)]

# 过滤科创板
def filter_kcb(context,stock_list):
    return [stock for stock in stock_list if not (stock.startswith("688") and stock.startswith("689"))]


#2-1 过滤停牌股票
def filter_paused_stock(stock_list):
    current_data = get_current_data()
    return [stock for stock in stock_list if not current_data[stock].paused]

# 过滤涨停的股票
def filter_limitup_stock(context, stock_list):
	last_prices = history(1, unit='1m', field='close', security_list=stock_list)
	current_data = get_current_data()
	
	# 已存在于持仓的股票即使涨停也不过滤，避免此股票再次可买，但因被过滤而导致选择别的股票
	return [stock for stock in stock_list if stock in context.portfolio.positions.keys()
			or last_prices[stock][-1] < current_data[stock].high_limit]

# 过滤跌停的股票
def filter_limitdown_stock(context, stock_list):
	last_prices = history(1, unit='1m', field='close', security_list=stock_list)
	current_data = get_current_data()
	
	return [stock for stock in stock_list if stock in context.portfolio.positions.keys()
			or last_prices[stock][-1] > current_data[stock].low_limit]
