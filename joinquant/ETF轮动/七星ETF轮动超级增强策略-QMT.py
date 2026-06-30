#coding=gbk
"""
七星高照ETF轮动超级增强策略 - QMT 单文件版本
原聚宽策略：七星高照ETF轮动超级增强策略
基于加权动量得分的ETF轮动策略，增强功能包括：
  - A股行情判断与走弱期自动回避
  - 多类别ETF池（海外、商品、A股）动态切换
  - 分钟级当日回撤保护
  - 盈利保护独立检查
  - 溢价率过滤、成交量过滤、短期动量过滤
"""

import math
from datetime import datetime
from typing import Any, Dict, List, Optional, Set

import numpy as np
import pandas as pd


# ============================================================================
# QMTTrader 交易类
# ============================================================================

class HoldingSnapshot:
    """get_holdings 统一返回的持仓快照（与 QMT POSITION 常用字段对齐，便于策略直接访问）。"""
    __slots__ = ('m_nVolume', 'm_nCanUseVolume', 'm_dOpenPrice')

    def __init__(self, vol: Any = 0, can_use: Any = 0, cost: Any = 0):
        self.m_nVolume: int = int(vol or 0)
        self.m_nCanUseVolume: int = int(can_use or 0)
        self.m_dOpenPrice: float = float(cost or 0)

    @staticmethod
    def from_counter_row(dt: Any) -> 'HoldingSnapshot':
        """将柜台 POSITION 单行规范为 HoldingSnapshot。"""
        return HoldingSnapshot(getattr(dt, 'm_nVolume', 0), getattr(dt, 'm_nCanUseVolume', 0), getattr(dt, 'm_dOpenPrice', 0))

class QMTTrader:
    """QMT 通用交易类：下单、行情、柜台查询。（单文件版本：已移除 send_notify 通知逻辑）"""

    def __init__(
        self,
        context_info,
        strategy_name: str = '通用策略',
        *,
        account: str = 'tests',
        accountType: str = 'STOCK',
        daily_fields_preset: Optional[List[str]] = None,
    ):
        """
        初始化：
          - 账号、买卖代码、资金与持仓字段。

        params:
          - context_info: ContextInfo 对象
          - strategy_name: 策略名称
          - account: 资金账号
          - accountType: 资金账号类型
          - daily_fields_preset: 日线数据缓存预设字段列表，默认值为 ['close','amount','volume']。
              设定后 get_daily_data_cached 每次实际拉取时会合并传入字段与预设字段，避免不同字段多次 RPC 调用。
        """
        self.acct_id: str = account
        self.acct_type: str = accountType

        self.is_backtesting: bool = context_info.do_back_test
        self.strategy_name: str = strategy_name

        self.contextInfo = context_info
        self.stock_info_cache = {}
        self._stock_info_cache_date: str = ''  # 缓存对应的自然日期，跨日自动失效

        self.per_amount = 10000
        self.position = 0
        self.positions: Dict[str, int] = {}

        self.waiting_list = []
        # 上一根已处理 K 线对应的自然日（YYYYMMDD），用于 is_new_calendar_day，非「交易日历」语义
        self.last_date: str = ''
        self.stock_code: str = context_info.stockcode + '.' + context_info.market

        if self.is_backtesting:
            download_history_data(self.stock_code, context_info.period, '', '')
            print(f'[init][回测模式][初始资金:{context_info.capital}][回测周期:{context_info.start}~{context_info.end}][{context_info.period}]')
        else:
            print(f'[init][交易模式][账号:{self.acct_id}][账号类型:{self.acct_type}]')

        self.buy_code = 23 if self.acct_type == 'STOCK' else 33
        self.sell_code = 24 if self.acct_type == 'STOCK' else 34

        print(f'[{self.stock_code}][{self.acct_id}][{self.acct_type}] 初始资金:{context_info.capital}, 回测={self.is_backtesting}, {context_info.period}')

        # 日线数据缓存（持久缓存，日线历史不可变，回测全程复用）
        self._daily_data_cache: dict = {}
        self._daily_fields_preset: List[str] = list(daily_fields_preset or ['close', 'amount', 'volume'])

    def get_available_cash(self) -> float:
        """柜台可用资金（核对、提示；数值单位以 QMT get_trade_detail_data(ACCOUNT) 为准，勿与策略现金混用）。"""
        result = 0.0
        resultlist = get_trade_detail_data(self.acct_id, self.acct_type, 'ACCOUNT')
        for obj in resultlist:
            if obj.m_dAvailable > 0:
                result = float(obj.m_dAvailable)
        return result

    def get_strategy_total_value(self, context_info) -> float:
        """策略总市值 = 策略现金 + 持仓市值（需传入 context_info 获取当前价格）。"""
        total_cash = self.get_available_cash()
        holdings = self.get_holdings()
        position_value = 0.0
        for stock_code, snapshot in holdings.items():
            if snapshot.m_nVolume > 0:
                price = self.get_price(context_info, stock_code=stock_code) or 0
                position_value += price * snapshot.m_nVolume
        return total_cash + position_value

    def get_holdings(self, print_holdings=False) -> Dict[str, HoldingSnapshot]:
        """查询持仓明细，统一返回 HoldingSnapshot（总仓/可卖/成本价），数据来自柜台 POSITION。"""
        holdinglist: Dict[str, HoldingSnapshot] = {}
        holdings = get_trade_detail_data(self.acct_id, self.acct_type, 'POSITION')
        if print_holdings:
            print('柜台持仓明细:')
        for dt in holdings:
            stock_key = dt.m_strInstrumentID + '.' + dt.m_strExchangeID
            holdinglist[stock_key] = HoldingSnapshot.from_counter_row(dt)
            if print_holdings:
                h = holdinglist[stock_key]
                print(stock_key, h.m_nVolume, h.m_nCanUseVolume, h.m_dOpenPrice)

        if self.stock_code in holdinglist:
            self.position = holdinglist[self.stock_code].m_nCanUseVolume
        else:
            self.position = 0
        return holdinglist

    def get_stock_info(self, symbol, field=None):
        """获取证券信息（带内存缓存，跨日自动失效）

        注：UpStopPrice/DownStopPrice/PreClose/InstrumentStatus 等字段每日更新，
        缓存按自然日失效，确保跨日运行时获取到当日数据。

        Args:
            symbol: 证券代码（格式： code.market）
            field: 要获取的字段名，如 'InstrumentName'、'ExchangeID' 等。
                   若为 None，返回完整 info dict。
                   常用字段：InstrumentName(名称), ExchangeID(市场), InstrumentID(代码)

        Returns: 指定字段值，或完整 info dict，或 None（获取失败时）
        """
        try:
            # 跨日自动清空缓存（UpStopPrice/DownStopPrice/PreClose 等每日更新）
            today = datetime.now().strftime('%Y%m%d')
            if self._stock_info_cache_date != today:
                self.stock_info_cache.clear()
                self._stock_info_cache_date = today

            cache = self.stock_info_cache
            if symbol not in cache:
                info = self.contextInfo.get_instrument_detail(symbol)
                if info:
                    cache[symbol] = info

            info = cache[symbol]
            if info:
                if field:
                    return info.get(field)
                return info
            return None
        except Exception:
            return None

    # 验证给定的代码和日期，返回该证券在指定日期是否可正常交易
    def is_valid_stock_in_date(self, code: str, date='') -> bool:
        """验证给定证券代码在指定日期是否可正常交易"""
        info = self.get_stock_info(code)
        if not info:
            return False
        open_date = info.get('OpenDate')
        # open_date 为 19700101 时，表示募集中尚未上市
        if (not open_date and info.get('PreClose') == 1.0) or open_date == 19700101:
            return False

        if date == '':
            date = int(datetime.now().strftime("%Y%m%d"))
        else:
            date = int(date)
        if date < open_date:
            return False

        expire_date = info.get('ExpireDate', 0)
        if expire_date and date >= expire_date:
            return False
        # TODO 进一步优化：该日期是否在当日停盘？可通过取行情数据判断
        return True

    def get_valid_stock_list(self, sectors: List[str] = None, date='') -> List[str]:
        """获取有效板块证券列表（带缓存）

        Args:
            sectors: 板块列表，默认 ['沪深ETF']
            date: 当前日期（格式：YYYYMMDD），默认为今天

        Returns:
            list: 有效证券代码列表（code.market 格式）
        """
        if sectors is None:
            sectors = ['沪深ETF']
        if date == '':
            date = datetime.now().strftime("%Y%m%d")

        try:
            # 通过 QMT 遍历板块获取证券列表
            all_codes = []
            for sector in sectors:
                try:
                    codes = self.contextInfo.get_stock_list_in_sector(sector)
                    all_codes.extend(codes)
                except Exception as e:
                    print(f"【ETF列表】获取板块 '{sector}' 证券列表失败: {e}")
                    continue

            # 去重
            all_codes = list(set(all_codes))

            # 过滤有效证券：OpenDate <= date < ExpireDate
            valid_codes = []
            for code in all_codes:
                try:
                    info = self.get_stock_info(code)
                    if not info:
                        continue
                    # 实盘模式下判断：合约已停牌日期（停牌第一天值为0，第二天为1，以此类推。注意，正常交易的股票该值也是0
                    if not self.is_backtesting and info.get('InstrumentStatus', 0) >= 1:
                        continue
                    if self.is_valid_stock_in_date(code, date):
                        valid_codes.append(code)
                except:
                    continue

            return valid_codes
        except Exception as e:
            print(f"【板块证券列表】获取有效证券列表异常: {e}")
            return []

    def get_price(self, context_info, stock_code=None, is_bar_price=True):
        """取价：回测/非最新价用 K 线数据；实盘最新价用 full_tick。"""
        if stock_code is None:
            stock_code = self.stock_code

        if self.is_backtesting or not is_bar_price:
            timetag = context_info.get_bar_timetag(context_info.barpos)
            endtime = timetag_to_datetime(timetag, '%Y%m%d%H%M%S')
            df = context_info.get_market_data_ex(
                ['close', 'lastPrice'],
                [stock_code],
                end_time=endtime,
                count=1,
                period=context_info.period,
                dividend_type=context_info.dividend_type,
                fill_data=True,
                subscribe=True
            )

            if not df or stock_code not in df or df[stock_code].empty:
                print(f'[warning][get_price]df is None [{stock_code} {context_info.period} {endtime}]')
                return None

            current_price = None
            if 'lastPrice' in df[stock_code].columns and not df[stock_code]['lastPrice'].empty:
                current_price = df[stock_code]['lastPrice'].iloc[-1]

            if (current_price is None or math.isnan(current_price)) and not df[stock_code]['close'].empty:
                current_price = df[stock_code]['close'].iloc[-1]

            return float(current_price) if current_price is not None else None
        else:
            tick = context_info.get_full_tick([stock_code])
            tick_data = tick.get(stock_code)
            if tick_data:
                return tick_data['lastPrice']

            print('[warning]获取行情数据失败')
            return None

    def execute_trade(
        self,
        context_info,
        signal: str,
        amount: int = None,
        price: float = -1,
        stock_code: str = None,
    ) -> bool:
        """按信号下单。stock_code 默认 None 表示当前主图标的（self.stock_code），多标的轮动可传入具体代码。
        实盘仅最后一根 K 线执行；回测与实盘相同调用 passorder，便于客户端查看委托/成交。
        """
        if signal not in ['buy', 'sell']:
            return False
        if amount is None:
            amount = self.per_amount
        if amount <= 0:
            return False

        if not self.is_backtesting:
            if not context_info.is_last_bar():
                print('[warning] 非最后一根K线，跳过交易 signal={} amount={} price={} stock_code={}'.format(signal, amount, price, stock_code))
                return False

            if self.waiting_list:
                orders = get_trade_detail_data(self.acct_id, self.acct_type, 'order', self.strategy_name)
                found_list = [o.m_strRemark for o in orders if o.m_strRemark in self.waiting_list]
                self.waiting_list = [i for i in self.waiting_list if i not in found_list]

                if len(self.waiting_list) > 0:
                    print(f'当前有未查到委托 {self.waiting_list} 暂停后续报单!请检查原因后重新运行程序！')
                    return False

        if signal == 'buy':
            return self._execute_buy(context_info, amount, price, stock_code=stock_code)
        return self._execute_sell(context_info, amount, price, stock_code=stock_code)

    def _execute_buy(
        self, context_info, amount: int, price: float = -1, *, stock_code: Optional[str] = None
    ) -> bool:
        """提交买入委托。"""
        code = stock_code or self.stock_code
        cur_price = price if price != -1 else self.get_price(context_info, stock_code=code)
        est_cost = (cur_price * amount) if cur_price else 0.0

        if cur_price:
            avail = self.get_available_cash() if self.is_backtesting else self.available_cash
            if avail < est_cost:
                print(f'资金不足 < {est_cost:.2f}，可能无法买入')
                return False

        if self.is_backtesting and not cur_price:
            print(f'[warning] 回测买入价格获取失败 stock_code={code}，跳过')
            return False

        self.available_cash = self.get_available_cash()
        curdatetime = self.get_current_time(context_info, '%Y-%m-%d %H:%M:%S')
        msg = f'[{curdatetime}][{self.strategy_name}][{code}][buy] {amount}'
        passorder(
            self.buy_code, 1101, self.acct_id, code, 14, price, amount,
            self.strategy_name, 1, msg, context_info
        )
        self.waiting_list.append(msg)
        self.positions[code] = self.positions.get(code, 0) + amount
        if code == self.stock_code:
            self.position += amount
        print(
            f'[{curdatetime}][{self.strategy_name}][买入委托]{code} 数量:{amount} 当前价:{cur_price} 策略持仓:{self.positions.get(code, 0)} 策略现金:{self.get_available_cash():.2f}'
        )
        return True

    def _execute_sell(
        self, context_info, amount, price=-1, *, stock_code: Optional[str] = None
    ):
        """提交卖出委托；校验策略持仓与柜台可卖量。"""
        code = stock_code or self.stock_code

        if self.is_backtesting:
            strat_vol = int(self.positions.get(code, 0))
            if amount > strat_vol:
                print(f'{code} 回测策略持仓不足 {strat_vol} < {amount}')
                return False
        else:
            holdings = self.get_holdings()
            can = holdings[code].m_nCanUseVolume if code in holdings else 0
            if amount > can:
                print(f'{code} 可卖不足 {can} < {amount}')
                return False

        curdatetime = self.get_current_time(context_info, '%Y-%m-%d %H:%M:%S')
        msg = f'[{curdatetime}][{self.strategy_name}][{code}][sell] {amount}'
        passorder(
            self.sell_code, 1101, self.acct_id, code, 14, price, amount,
            self.strategy_name, 1, msg, context_info
        )
        self.waiting_list.append(msg)
        nv = self.positions.get(code, 0) - amount
        if nv <= 0:
            self.positions.pop(code, None)
        else:
            self.positions[code] = nv
        if code == self.stock_code:
            self.position -= amount

        cur_price = price if price != -1 else self.get_price(context_info, stock_code=code)
        print(
            f'[{curdatetime}][{self.strategy_name}][卖出委托]{code} 数量:{amount} 当前价:{cur_price} 策略持仓:{self.positions.get(code, 0)}'
        )
        return True

    def get_current_time(self, context_info: Any, fmt='%H%M%S') -> str:
        """当前 K 线时间转字符串（格式由 fmt 指定）。"""
        return timetag_to_datetime(context_info.get_bar_timetag(context_info.barpos), fmt)

    def get_recent_trading_days(self, *, start_date='', end_date='', count=8000, end_date_inclusive=False) -> List[str]:
        """获取近 count 个交易日历
        - end_date_inclusive 是否包含 end_date 当日。默认 False
        Returns:
            list[str]: 交易日期列表，格式如 ['20240102', '20240103', ...]，按日期升序
        """
        if not end_date_inclusive:
            count += 1
        raw_dates = self.contextInfo.get_trading_dates('SH', start_date, end_date, count, '1d')
        if raw_dates:
            trading_days: List[str] = []
            for d in raw_dates:
                # 排除当日数据，避免未来数据
                if end_date_inclusive or d != end_date:
                    trading_days.append(d)
            return trading_days
        return []

    def get_prev_trade_date(self, date='') -> str:
        """获取 date 的前一个交易日历日期，返回格式：YYYYMMDD（排除 end_date 当日）"""
        trading_days = self.get_recent_trading_days(end_date=date, count=2, end_date_inclusive=False)
        return trading_days[-1] if trading_days else ''

    @staticmethod
    def is_etf(stock_code: str, market: str) -> bool:
        """是否沪深 ETF（is_typed_stock 类型 100013）。"""
        return is_typed_stock(100013, stock_code, market)

    def is_trading_time(self, context_info, curtime='') -> bool:
        """是否 A 股常规交易时段（约 9:30–14:57，按当前 K 线时间判断）。"""
        if not curtime:
            curtime = self.get_current_time(context_info)
        return '093000' <= curtime <= '145700'

    def is_new_calendar_day(self, context_info, curdate='') -> bool:
        """相对上一根 K 线是否进入新自然日"""
        if curdate == '':
            curdate = self.get_current_time(context_info, '%Y%m%d')
        changed = curdate != self.last_date
        if changed:
            self.last_date = curdate
        return changed

    def estimate_today_volume(self, context_info: Any, today_vol: float, cur_dt: str = '') -> float:
        """根据当前K线时间将日内累计成交量投影为全日预估成交量。

        A 股交易日分钟数 = 240（9:30-11:30 120 分钟 + 13:00-15:00 120 分钟）。
        按当前 bar 的 HHMMSS 计算已流逝的交易分钟数，线性推算全日量。

        Args:
            context_info: QMT ContextInfo 对象（需要 barpos 获取当前 bar 时间）
            today_vol: 当日已累计成交量

        Returns:
            float: 预估全日成交量
        """
        if not cur_dt:
            cur_dt = self.get_current_time(context_info, '%Y%m%d%H%M%S')
        hour = int(cur_dt[8:10])
        minute = int(cur_dt[10:12])
        elapsed_minutes = (hour - 9) * 60 + minute - 30
        if hour >= 13:
            elapsed_minutes -= 90
        elapsed_minutes = max(1, min(elapsed_minutes, 240))
        return today_vol * (240.0 / elapsed_minutes)

    def get_volume_ratio(self, ContextInfo: Any, hist_volumes, today_vol: float, lookback_days=5):
        """计算成交量比：当日实时成交量投影到全日 / 历史日均成交量"""
        if hist_volumes is None or len(hist_volumes) < lookback_days:
            print(f"[error]【成交量比计算】历史成交量不足{lookback_days}天，忽略计算")
            return None

        past_n_days_vol = hist_volumes[-lookback_days:]
        if np.any(np.isnan(past_n_days_vol)) or np.any(past_n_days_vol == 0):
            return None

        avg_volume = np.mean(past_n_days_vol)
        if avg_volume == 0:
            return None

        # 盘中时间投影：将当前累计成交量投影到全日
        projected_today_vol = self.estimate_today_volume(ContextInfo, today_vol)
        return projected_today_vol / avg_volume if projected_today_vol > 0 else 0

    def get_premium_rate(self, etf: str, current_price: Optional[float]) -> Optional[float]:
        """获取溢价率"""
        if self.is_backtesting: return 0.0  # 回测无IOPV，跳过溢价率过滤
        if current_price is None or (isinstance(current_price, (int, float)) and (current_price <= 0 or current_price != current_price)): return None
        try:
            iopv: float = get_etf_iopv(etf)
            if iopv <= 0 or (isinstance(iopv, float) and (iopv != iopv or np.isnan(iopv))): return None
            return (float(current_price) - iopv) / iopv
        except Exception as e:
            print(f'[error] 获取 {etf} IOPV 失败！QMT 需在全局作用域提供 get_etf_iopv 函数: {e}')
            return None

    def get_daily_data_cached(
        self, fields: List[str], codes: List[str], count: int,
        end_time: Optional[str] = '', start_time: Optional[str] = '',
        subscribe: bool = False,
    ) -> Dict[str, Dict[str, pd.Series]]:
        """带持久缓存的日线数据获取，替代 ContextInfo.get_market_data_ex(period='1d')

        缓存结构：
          self._daily_data_cache = {code: {
              'data':  {field: pd.Series(index=DatetimeIndex)},  # 缓存数据
              'count': int,    # 缓存数据行数
          }}

        核心策略：
          - 拉取时 end_time 为空（''），让 QMT 一次取到最新/回测终点的数据
          - 回测环境下 QMT 会返回到当前 bar 为止的数据，拉取后整个回测周期可复用
          - 缓存命中只看行数够不够 + 字段是否齐全，不检查 latest
          - 返回时按调用方的 end_time 截断，保证跨代码对齐
        """
        cache = self._daily_data_cache

        # 默认取当前 bar 为结束时间
        if not end_time:
            end_time = self.get_current_time(self.contextInfo, '%Y%m%d')

        # 合并预设字段（一次 RPC 获取所有字段，避免不同字段的多次 RPC 调用）
        fetch_fields = list(set(fields) | set(self._daily_fields_preset))
        # 更新 self._daily_fields_preset，让其总是包含需要的所有字段
        self._daily_fields_preset = fetch_fields

        # ---- 检查缓存是否足够（只看行数+字段，不看 latest） ----
        codes_needed: List[str] = []
        cache_hit_count = 0
        today = datetime.now()
        end_time_dt = datetime.strptime(end_time, '%Y%m%d')

        # 最少需获取的 count 数 = today - end_time 天数 + count
        fetch_count = max(count + (today - end_time_dt).days, 500)

        for code in codes:
            entry = cache.get(code)
            if entry is None:
                codes_needed.append(code)
                continue
            cached_data = entry.get('data', {})
            cached_count = entry.get('count', 0)
            if cached_count >= fetch_count and all(f in cached_data for f in fields):
                cache_hit_count += 1
                continue
            codes_needed.append(code)

        # ---- 拉取缺失数据（end_time='' 取到最新，一次拉够后续复用） ----
        if codes_needed:
            fetch_codes = list(set(codes_needed))

            def _cache_raw_data(raw_dict, field_list):
                """将 get_market_data_ex 返回的原始数据写入缓存"""
                ok = 0
                for code, raw in raw_dict.items():
                    entry = {'data': {}, 'count': 0}
                    for f in field_list:
                        if f in raw and raw[f] is not None and len(raw[f]) > 0:
                            entry['data'][f] = raw[f]
                    if entry['data']:
                        first_ser = next(iter(entry['data'].values()))
                        entry['count'] = len(first_ser)
                        cache[code] = entry
                        ok += 1
                return ok

            try:
                start = datetime.now()
                new_data = self.contextInfo.get_market_data_ex(
                    fetch_fields, fetch_codes,
                    period='1d', count=fetch_count,
                    end_time='', subscribe=False,
                )
                print(f'[get_daily_data_cached] 批量获取{len(fetch_codes)}, 耗时 {(datetime.now() - start).total_seconds()}s')

                cached_ok = _cache_raw_data(new_data, fetch_fields)

                if cached_ok < len(fetch_codes):
                    missing = [c for c in fetch_codes if c not in cache]
                    print(f'[get_daily_data_cached] {len(missing)}/{len(fetch_codes)} 只未缓存 (fetch_count={fetch_count}, fields={fetch_fields})')
                    for code in missing:
                        try:
                            single_data = self.contextInfo.get_market_data_ex(
                                fetch_fields, [code],
                                period='1d', count=fetch_count,
                                end_time='', subscribe=False,
                            )
                            _cache_raw_data(single_data, fetch_fields)
                        except Exception as e2:
                            print(f'[get_daily_data_cached] {code} 单独拉取失败: {e2}')

            except Exception as e:
                print(f'[get_daily_data_cached] 批量获取失败: {e}')
                # fallback：用原始参数重试（保留调用方的 end_time）
                try:
                    fallback_data = self.contextInfo.get_market_data_ex(
                        fields, codes, period='1d', count=count,
                        end_time=end_time, start_time=start_time, subscribe=subscribe,
                    )
                    return fallback_data
                except Exception as e3:
                    print(f'[get_daily_data_cached] fallback也失败[{fields} {count} {end_time} {start_time}]: {e3}')
                    return {}

        # ---- 从缓存返回，按调用方的 end_time 截断 ----
        result: Dict[str, Dict[str, pd.Series]] = {}
        if not fields:
            fields = self._daily_fields_preset
        for code in codes:
            entry = cache.get(code)
            if entry is None:
                continue
            cached_data = entry.get('data', {})
            result[code] = {}
            for f in fields:
                ser = cached_data.get(f)
                if ser is not None and len(ser) > 0:
                    # 先截断到调用方的 end_time，再取末尾 count 行
                    # 顺序不能反：iloc[-count:] 取的是缓存末尾（可能远超 end_time），
                    # 再按 end_time 过滤就会把超出日期的全部丢弃，导致结果为空
                    if end_time:
                        end_str = str(end_time)[:8]
                        ser = ser[[str(x)[:8] <= end_str for x in ser.index]]
                    sliced = ser.iloc[-count:] if len(ser) > 0 else ser
                    result[code][f] = sliced
            # 获取为空则打印日志供调试
            if not result[code]:
                print(f' ~ [get_daily_data_cached] {code} 缓存数据为空')

        return result


# ============================================================================
# 七星高照ETF轮动超级增强策略（来自 run/qixing_etf_rotation_enhanced.py）
# ============================================================================

# ---------------------------------------------------------------------------
# 全局变量
# ---------------------------------------------------------------------------

trader: QMTTrader

# ---------------------------------------------------------------------------
# 【一】策略参数初始化
# ---------------------------------------------------------------------------

def init_global_vars(ContextInfo):
    """初始化全局变量 - 七星高照ETF轮动超级增强策略"""
    global g
    g = {}

    # ==================== 运行时状态变量 ====================
    g['etf_sell_time'] = '1309'               # 卖出时间
    g['etf_buy_time'] = '1310'                # 买入时间
    g['regime_check_time'] = '0940'           # 行情判断时间

    g['dl_history_data_date'] = '1'          # 最近一次下载历史数据的日期. 回测时设置为 '1' 则可关闭数据下载
    g['enable_detailed_log'] = False          # 是否启用详细日志
    g['rankings_cache'] = {'date': None, 'data': None}   # 排名缓存

    # 每日定时任务执行标记
    g['check_positions_done'] = False         # 09:10 持仓检查已执行
    g['regime_check_done'] = False            # 09:40 行情判断已执行
    g['profit_protection_done_times'] = set() # 盈利保护已检查的时间点
    g['etf_sell_done'] = False                # 卖出已执行
    g['etf_buy_done'] = False                 # 买入已执行
    g['daily_summary_done'] = False           # 盘后总结已执行

    # ==================== 核心参数 ====================
    g['lookback_days'] = 25               # 动量计算周期
    g['holdings_num'] = 1                 # 候选数量
    g['defensive_etf'] = '511010.SH'      # 避险资产：10年国债ETF（聚宽原版511010.XSHG）
    g['min_money'] = 5000                 # 最小交易金额

    # ==================== 盈利保护参数 ====================
    g['enable_profit_protection'] = True                      # 盈利保护开关
    g['profit_protection_lookback'] = 1                       # 盈利保护回看周期（天）
    g['profit_protection_threshold'] = 0.05                   # 盈利保护回撤阈值（5%）
    g['profit_protection_check_times'] = ['11:00']            # 盈利保护检查时间点

    g['drawdown_selled_today'] = set()                        # 当日因回撤/盈利保护卖出的标的（禁止日内买回）
    g['buy_date'] = {}                                        # 记录买入日期
    g['trade_log'] = {'sell_records': []}                     # 记录当日卖出

    g['loss'] = 0.97                      # 近3日单日跌幅阈值（排除）

    g['min_score_threshold'] = 0          # 最低得分
    g['max_score_threshold'] = 100.0      # 最高得分

    # ==================== 成交量过滤 ====================
    g['enable_volume_check'] = True
    g['volume_lookback'] = 5
    g['volume_threshold'] = 3 # 由于当日成交量当前计算方式为按240分钟估算，与原始策略比会更大，故这里阈值也需调大一些
    g['volume_return_limit'] = 1          # 年化收益>100%时启用放量过滤

    # ==================== 短期动量过滤 ====================
    g['use_short_momentum_filter'] = True
    g['short_lookback_days'] = 10
    g['short_momentum_threshold'] = 0.0

    # ==================== 溢价率过滤 ====================
    g['enable_premium_filter'] = True      # 是否启用溢价率过滤
    g['premium_threshold'] = 0.20          # 溢价率阈值（20%）

    # ==================== 分钟级当日回撤保护 ====================
    g['intraday_drawdown_threshold'] = 0.02            # 当日回撤阈值（2%）

    # ==================== A股行情判断 ====================
    g['enable_regime_switch'] = True                    # 行情判断开关
    g['weak_period_ma_lookback'] = 10                   # 10日均线
    g['weak_period_max_days'] = 20                      # 走弱期最长持续20个交易日
    g['is_a_share_weak'] = False                        # 当前是否走弱期
    g['weak_period_counter'] = 0                        # 走弱期天数计数器

    # 独立手动开关
    g['enable_avoid_a_share'] = True                    # 走弱期回避A股开关
    g['enable_intraday_drawdown'] = True                # 分钟级回撤保护独立开关

    g['regime_indexes'] = {                             # 监测指数（QMT格式）
        '沪深300': '000300.SH',
        '深证综指': '399101.SZ',
        '创业板指': '399006.SZ',
        '中证A500': '000510.SH',
    }

    # ==================== ETF池（QMT格式：XSHE→SZ, XSHG→SH）====================

    # 海外ETF（走弱期可交易）
    g['overseas_etf_pool'] = [
        "513100.SH",  # 纳指ETF
        "513290.SH",  # 纳指生物ETF
        "513500.SH",  # 标普500ETF
        "159529.SZ",  # 标普消费
        "513400.SH",  # 道琼斯ETF
        "513520.SH",  # 日经225ETF
        "513030.SH",  # 德国30ETF
        "513080.SH",  # 法国ETF
        "513310.SH",  # 中韩半导体ETF
        "513730.SH",  # 东南亚ETF
        "159792.SZ",  # 港股互联ETF
        "513130.SH",  # 恒生科技
        "513050.SH",  # 中概互联网ETF
        "159920.SZ",  # 恒生ETF
        "513690.SH",  # 港股红利
        # 债券ETF
        "511380.SH",  # 可转债ETF
        "511010.SH",  # 国债ETF
        "511220.SH",  # 城投债ETF
    ]

    # 商品ETF（走弱期可交易）
    g['commodity_etf_pool'] = [
        "518880.SH",  # 黄金ETF
        "159980.SZ",  # 有色金属ETF
        "159985.SZ",  # 豆粕ETF
        "501018.SH",  # 南方原油
        '161226.SZ',  # 白银LOF
        "159981.SZ",  # 能源化工ETF
        "512400.SH",  # 工业有色ETF
    ]

    # A股ETF（走弱期回避）
    g['domestic_etf_pool'] = [
        # 指数ETF
        "510300.SH",  # 沪深300ETF
        "510500.SH",  # 中证500ETF
        "510050.SH",  # 上证50ETF
        "510210.SH",  # 上证ETF
        "159915.SZ",  # 创业板ETF
        "588080.SH",  # 科创50
        "512100.SH",  # 中证1000ETF
        "563360.SH",  # A500-ETF
        "563300.SH",  # 中证2000ETF
        # 风格ETF
        "512890.SH",  # 红利低波ETF
        "159967.SZ",  # 创业板成长ETF
        "588020.SH",  # 科创成长ETF
        "512040.SH",  # 价值ETF
        "159201.SZ",  # 自由现金流ETF
        # 行业板块ETF
        "515790.SH",   # 光伏ETF
        "563230.SH",   # 卫星ETF
        "515880.SH",   # 通信ETF
        "512660.SH",   # 军工ETF
        "561380.SH",   # 电网设备ETF
        "159667.SZ",   # 工业母机ETF
        "159559.SZ",   # 机器人ETF
        "159819.SZ",   # 人工智能ETF
        "159381.SZ",   # 创业板人工智能ETF
        "159732.SZ",   # 消费电子ETF
        "159995.SZ",   # 芯片ETF
        "512220.SH",   # TMT(科技传媒通信150）ETF
    ]

    # 完整ETF池（初始化时合并）
    g['etf_pool'] = g['overseas_etf_pool'] + g['commodity_etf_pool'] + g['domestic_etf_pool']

    # 需要订阅行情和下载数据的全量代码（含指数）
    g['all_subscribe_codes'] = list(set(g['etf_pool'] + list(g['regime_indexes'].values())))


# ---------------------------------------------------------------------------
# 【二】策略入口 init
# ---------------------------------------------------------------------------

def init(ContextInfo):
    """策略入口：初始化策略参数、全局变量"""
    global trader

    # 使用 QMTTrader
    trader = QMTTrader(
        ContextInfo,
        strategy_name='七星高照ETF轮动超级增强',
        account=account if 'account' in globals() else 'tests',
        accountType=accountType if 'accountType' in globals() else 'STOCK',
    )

    # 初始化全局变量
    init_global_vars(ContextInfo)

    # 打印策略初始化参数
    print_init_params()

    # 下载历史数据（日线 + 1分钟线）
    dl_history_data(ContextInfo, 'init')

    # 订阅分钟级行情（ETF池 + 指数）
    for code in g['all_subscribe_codes']:
        ContextInfo.subscribe_quote(code, period='1m')

    # 初始化日线数据缓存（拉取一次避免后续单只调用）
    end_date = trader.get_prev_trade_date(
        trader.get_current_time(ContextInfo, '%Y%m%d')[:8]
    )
    _ = trader.get_daily_data_cached(
        ['close'], g['all_subscribe_codes'], count=60, end_time=end_date
    )


def print_init_params():
    """打印策略初始化参数"""
    print(f"""【七星高照ETF轮动超级增强】启动！(QMT版本 - 单文件)
 === 策略参数初始化完成 ===
 === ETF池配置 ===
 - 海外ETF: {len(g['overseas_etf_pool'])}只
 - 商品ETF: {len(g['commodity_etf_pool'])}只
 - A股ETF: {len(g['domestic_etf_pool'])}只
 - 完整池: {len(g['etf_pool'])}只
 - 动量周期: {g['lookback_days']}天
 - 持仓数量: {g['holdings_num']}只
 - 防御ETF: {g['defensive_etf']}
 === 盈利保护 ===
 - 开关: {'开启' if g['enable_profit_protection'] else '关闭'}
 - 回看周期: {g['profit_protection_lookback']}天
 - 回撤阈值: {g['profit_protection_threshold']*100:.0f}%
 - 检查时间点: {g['profit_protection_check_times']}
 === 过滤条件 ===
 - 成交量过滤: {'启用' if g['enable_volume_check'] else '禁用'} (近{g['volume_lookback']}日均量比 < {g['volume_threshold']})
 - 年化收益放量阈值: {g['volume_return_limit']*100:.0f}%
 - 短期动量过滤: {'启用' if g['use_short_momentum_filter'] else '禁用'} ({g['short_lookback_days']}天)
 - 近3日单日跌幅: < {(1-g['loss'])*100:.0f}%
 - 溢价率过滤: {'启用' if g['enable_premium_filter'] else '禁用'} (阈值<={g['premium_threshold']*100:.0f}%)
 === 行情判断 ===
 - 开关: {'启用' if g['enable_regime_switch'] else '关闭'}
 - 走弱期最长: {g['weak_period_max_days']}日
 - 回避A股开关: {'ON' if g['enable_avoid_a_share'] else 'OFF'}
 - 分钟回撤开关: {'ON' if g['enable_intraday_drawdown'] else 'OFF'}
 - 回撤阈值: {g['intraday_drawdown_threshold']*100:.0f}%
 === 日内调度 ===
 - 09:10 持仓检查
 - 09:40 行情判断
 - 盈利保护检查: {g['profit_protection_check_times']}
 - 09:46起 分钟级回撤检查（仅走弱期）
 - 13:09 卖出操作
 - 13:10 买入操作
 - 15:05 盘后总结报告
""")


# ---------------------------------------------------------------------------
# 【三】主循环 handlebar（需设置1分钟K线周期）
# ---------------------------------------------------------------------------

def handlebar(ContextInfo):
    """每根K线执行（1分钟周期）：交易时段过滤 → 定时任务 → 信号 → 下单"""
    if not ContextInfo.do_back_test and not ContextInfo.is_last_bar(): return False

    global trader

    # 获取当前时间
    cur_dt = trader.get_current_time(ContextInfo, '%Y%m%d%H%M')
    cur_date = cur_dt[:8]
    cur_time = cur_dt[8:]

    # 重置每日标志（新交易日）
    if trader.is_new_calendar_day(ContextInfo, cur_date):
        g['check_positions_done'] = False
        g['regime_check_done'] = False
        g['profit_protection_done_times'] = set()
        g['etf_sell_done'] = False
        g['etf_buy_done'] = False
        g['daily_summary_done'] = False
        g['rankings_cache'] = {'date': None, 'data': None}
        g['drawdown_selled_today'] = set()                    # 清空当日回撤卖出缓存
        g['trade_log'] = {'sell_records': []}
        print(f'{"★" * 10} [{cur_date}] {cur_time} 新交易日日开始 {"★" * 10}')

    # ==================== 09:10 持仓检查 ====================
    if not g['check_positions_done'] and cur_time >= '0910':
        check_positions(ContextInfo)
        g['check_positions_done'] = True

    # ==================== 09:40 行情判断 ====================
    if not g['regime_check_done'] and cur_time >= g['regime_check_time']:
        regime_check(ContextInfo, cur_date)
        g['regime_check_done'] = True

    # ==================== 盈利保护独立检查（支持多个时间点） ====================
    for check_time in g['profit_protection_check_times']:
        check_time_str = check_time.replace(':', '')
        if check_time not in g['profit_protection_done_times'] and cur_time >= check_time_str:
            profit_protection_check(ContextInfo)
            g['profit_protection_done_times'].add(check_time)

    # ==================== 分钟级当日回撤检查（走弱期 09:46起） ====================
    if is_intraday_drawdown_enabled() and cur_time >= '0946' and cur_time < '1500':
        intraday_drawdown_check(ContextInfo, cur_date, cur_time)

    # ==================== 13:09 卖出操作 ====================
    if not g['etf_sell_done'] and cur_time >= g['etf_sell_time']:
        dl_history_data(ContextInfo, cur_date)  # 卖出前更新历史数据
        etf_sell_trade(ContextInfo, cur_date, cur_time)
        g['etf_sell_done'] = True

    # ==================== 13:10 买入操作 ====================
    if not g['etf_buy_done'] and cur_time >= g['etf_buy_time']:
        etf_buy_trade(ContextInfo, cur_date, cur_time)
        g['etf_buy_done'] = True

    # ==================== 15:05 盘后总结 ====================
    if not g['daily_summary_done'] and cur_time >= '1505':
        daily_summary_report(ContextInfo, cur_date)
        g['daily_summary_done'] = True


# ---------------------------------------------------------------------------
# 【四】开盘检查模块
# ---------------------------------------------------------------------------

def check_positions(ContextInfo):
    """每日开盘检查持仓状态，打印持仓信息"""
    try:
        holdings = trader.get_holdings()
        for sec, pos in holdings.items():
            if pos.m_nVolume > 0:
                name = get_stock_name(sec)
                cost = pos.m_dOpenPrice
                cur_price = trader.get_price(ContextInfo, sec) or 0
                print(f"  持仓：{sec} {name} 数量{pos.m_nVolume} 成本{cost:.3f} 现价{cur_price:.3f}")
    except Exception as e:
        print(f"【持仓检查】获取持仓失败: {e}")


# ---------------------------------------------------------------------------
# 【五】行情判断模块
# ---------------------------------------------------------------------------

def regime_check(ContextInfo, cur_date: str):
    """每日 09:40 判断A股行情，决定是否允许交易A股ETF，同时自动控制分钟级回撤开关"""
    print("========== 行情判断开始 ==========")

    if not g['enable_regime_switch']:
        g['is_a_share_weak'] = False
        print("  行情判断未启用，始终全市场交易")
        print("========== 行情判断完成 ==========")
        return

    end_date = trader.get_prev_trade_date(cur_date)
    if not end_date:
        print("  [warning] 无法获取前一交易日，跳过行情判断")
        return

    # 获取指数日线数据
    index_codes = list(g['regime_indexes'].values())
    index_data = trader.get_daily_data_cached(
        ['close'], index_codes,
        count=g['weak_period_ma_lookback'] + 2, end_time=end_date
    )

    below_count, above_count = 0, 0
    detail = []

    for name, code in g['regime_indexes'].items():
        try:
            if not index_data or code not in index_data:
                continue
            data = index_data[code]
            if 'close' not in data or len(data['close']) < g['weak_period_ma_lookback']:
                continue
            closes = data['close'].values[-g['weak_period_ma_lookback']:]
            current_price = closes[-1]
            ma_val = closes.mean()
            if current_price < ma_val:
                below_count += 1
                detail.append(f"{name}↓")
            else:
                above_count += 1
                detail.append(f"{name}↑")
        except Exception as e:
            print(f"  [warning] 指数{name}获取失败: {e}")

    old_state = g['is_a_share_weak']

    if not g['is_a_share_weak']:
        # 正常期 → 检查是否进入走弱期
        if below_count >= 3:
            g['is_a_share_weak'] = True
            g['weak_period_counter'] = 0
            print(f"  ?? 进入走弱期 (跌破:{below_count} {detail})")
            if g['enable_avoid_a_share']:
                print(f"     → 将回避A股ETF，仅交易海外+商品ETF")
            else:
                print(f"     → ?? 回避A股开关已关闭，仍交易全市场ETF")
            if g['enable_intraday_drawdown']:
                print(f"     → ??? 分钟级回撤保护已启用（阈值{g['intraday_drawdown_threshold']*100:.0f}%）")
            else:
                print(f"     → ? 分钟级回撤保护已被独立开关关闭，不触发")
    else:
        # 走弱期 → 检查是否恢复
        g['weak_period_counter'] += 1
        if above_count >= 3:
            g['is_a_share_weak'] = False
            g['weak_period_counter'] = 0
            print(f"  ?? 恢复正常期 (站上:{above_count} {detail})")
            if g['enable_avoid_a_share']:
                print(f"     → 恢复交易A股ETF")
            else:
                print(f"     → 回避A股开关关闭，始终交易全市场")
            if g['enable_intraday_drawdown']:
                print(f"     → 关闭分钟级回撤保护")
            else:
                print(f"     → 分钟级回撤保护独立开关已关闭，无变化")
        elif g['weak_period_counter'] >= g['weak_period_max_days']:
            g['is_a_share_weak'] = False
            g['weak_period_counter'] = 0
            print(f"  ? 走弱期满{g['weak_period_max_days']}日强制退出，恢复正常期")
            if g['enable_avoid_a_share']:
                print(f"     → 恢复交易A股ETF")
            else:
                print(f"     → 回避A股开关关闭，始终交易全市场")
            if g['enable_intraday_drawdown']:
                print(f"     → 关闭分钟级回撤保护")
            else:
                print(f"     → 分钟级回撤保护独立开关已关闭，无变化")

    # 状态变化时清除排名缓存
    if old_state != g['is_a_share_weak']:
        g['rankings_cache'] = {'date': None, 'data': None}

    # 输出当前状态
    if g['enable_regime_switch']:
        current_status = '??走弱期' if g['is_a_share_weak'] else '??正常期'
        avoid_status = '(回避A股)' if (g['is_a_share_weak'] and g['enable_avoid_a_share']) else ('(不回避A股)' if g['is_a_share_weak'] else '')
        drawdown_status = '???启用' if (g['is_a_share_weak'] and g['enable_intraday_drawdown']) else ('?关闭' if (g['is_a_share_weak'] and not g['enable_intraday_drawdown']) else '?关闭')
        print(f"  当前状态：{current_status}{avoid_status} 计数:{g['weak_period_counter']}/{g['weak_period_max_days']}")
        print(f"  分钟级回撤保护：{drawdown_status}（阈值{g['intraday_drawdown_threshold']*100:.0f}%）")
    print("========== 行情判断完成 ==========")


def is_intraday_drawdown_enabled() -> bool:
    """判断分钟级回撤保护是否启用（独立开关 + 走弱期）"""
    if not g['enable_intraday_drawdown']:
        return False
    if not g['enable_regime_switch']:
        return False
    return g['is_a_share_weak']


def get_active_etf_pool() -> List[str]:
    """根据行情状态获取当前可交易的ETF池"""
    # 手动关闭回避开关 → 走弱期也不回避A股
    if not g['enable_avoid_a_share']:
        print(f"  【强制】A股回避开关已关闭，使用完整池({len(g['etf_pool'])}只)")
        return g['etf_pool']

    if g['is_a_share_weak']:
        # 走弱期：只交易海外ETF + 商品ETF（回避A股）
        active_pool = g['overseas_etf_pool'] + g['commodity_etf_pool']
        print(f"  【走弱期】使用海外+商品池({len(active_pool)}只)")
        return active_pool
    else:
        # 正常期：交易全部ETF
        print(f"  【正常期】使用完整池({len(g['etf_pool'])}只)")
        return g['etf_pool']


# ---------------------------------------------------------------------------
# 【六】分钟级当日回撤检查
# ---------------------------------------------------------------------------

def intraday_drawdown_check(ContextInfo, cur_date: str, cur_time: str):
    """每分钟执行一次，检查所有持仓从当日盘中最高点的回撤（仅走弱期触发）"""
    holdings = trader.get_holdings()
    for sec, pos in holdings.items():
        if pos.m_nVolume <= 0:
            continue
        if sec not in g['etf_pool'] and sec != g['defensive_etf']:
            continue
        # 跳过当日买入的ETF（避免刚买入就被日内回撤保护卖出）
        if g['buy_date'].get(sec) == cur_date:
            continue

        try:
            # 获取当日分钟线数据（从开盘到当前时刻）
            minute_data = ContextInfo.get_market_data_ex(
                ['high', 'close'],
                [sec],
                period='1m',
                start_time=cur_date + '093000',
                end_time=cur_date + cur_time + '00',
                count=240,
                fill_data=True,
                subscribe=True
            )

            if not minute_data or sec not in minute_data:
                continue

            data = minute_data[sec]
            if 'high' not in data or data['high'] is None or len(data['high']) == 0:
                continue

            day_high = data['high'].max()
            current_price = trader.get_price(ContextInfo, sec)
            if current_price is None or day_high <= 0:
                continue

            drawdown = (day_high - current_price) / day_high
            if drawdown >= g['intraday_drawdown_threshold']:
                name = get_stock_name(sec)
                print(f"  ?? 分钟级回撤触发：{sec} {name} 回撤{drawdown*100:.2f}% > {g['intraday_drawdown_threshold']*100:.0f}%")
                if smart_order_target_value(sec, 0, ContextInfo):
                    print(f"  ?? 分钟级回撤卖出：{sec} {name}")
                    g['drawdown_selled_today'].add(sec)
        except Exception as e:
            pass  # 静默处理分钟级异常


# ---------------------------------------------------------------------------
# 【七】盈利保护独立检查函数
# ---------------------------------------------------------------------------

def profit_protection_check(ContextInfo):
    """独立执行的盈利保护检查函数，遍历所有持仓"""
    if not g['enable_profit_protection']:
        print("[DEBUG] 盈利保护模块已关闭，跳过检查")
        return

    print("========== 盈利保护检查开始 ==========")
    holdings = trader.get_holdings()
    for sec, pos in holdings.items():
        if pos.m_nVolume <= 0:
            continue
        if sec not in g['etf_pool'] and sec != g['defensive_etf']:
            continue
        if check_profit_protection(sec, ContextInfo):
            if smart_order_target_value(sec, 0, ContextInfo):
                name = get_stock_name(sec)
                print(f"  ??? 盈利保护卖出：{sec} {name}")
                g['drawdown_selled_today'].add(sec)
    print("========== 盈利保护检查完成 ==========")


def check_profit_protection(security: str, ContextInfo, lookback=None, threshold=None) -> bool:
    """检查是否触发盈利保护（从最近N日最高点回撤超过阈值）"""
    if not g['enable_profit_protection']:
        return False

    lookback = lookback or g['profit_protection_lookback']
    threshold = threshold or g['profit_protection_threshold']

    # 获取当前日期并计算前一交易日（仅用已完结的日线，避免前视偏差）
    cur_dt = timetag_to_datetime(ContextInfo.get_bar_timetag(ContextInfo.barpos), '%Y%m%d%H%M%S')
    cur_date = cur_dt[:8]
    end_date = trader.get_prev_trade_date(cur_date)

    # 获取最近N日的最高价（不包括当天）
    hist = trader.get_daily_data_cached(['high'], [security], count=lookback, end_time=end_date)
    if not hist or security not in hist:
        return False

    data = hist[security]
    if 'high' not in data or len(data['high']) < lookback:
        return False

    max_high = data['high'].max()
    current_price = trader.get_price(ContextInfo, security)
    if current_price is None:
        return False

    if current_price <= max_high * (1 - threshold):
        name = get_stock_name(security)
        print(f"  ?? {security} {name} 触发盈利保护：当前价{current_price:.3f}，"
              f"最近{lookback}日最高{max_high:.3f}，回撤{(1 - current_price/max_high)*100:.2f}% > {threshold*100:.0f}%")
        return True
    return False


# ---------------------------------------------------------------------------
# 【八】溢价率获取函数
# ---------------------------------------------------------------------------

def get_premium_rate(etf: str, current_price: float, ContextInfo):
    """获取溢价率（使用IOPV替代聚宽的基金净值，适合盘中判断）
    返回: (premium_rate, is_pass)
      - premium_rate: 溢价率数值，None表示无法获取
      - is_pass: True表示通过（溢价率在阈值范围内），False表示超标
    """
    if ContextInfo.do_back_test:
        return None, True  # 回测无IOPV，跳过溢价率过滤

    premium_rate = trader.get_premium_rate(etf, current_price)
    if premium_rate is None:
        return None, False  # 实盘模式下IOPV获取失败，视为不合格
    passed = isinstance(premium_rate, (int, float)) and premium_rate <= g['premium_threshold']
    return premium_rate, passed


# ---------------------------------------------------------------------------
# 【九】核心计算模块 - ETF排名
# ---------------------------------------------------------------------------

def get_cached_rankings(ContextInfo, cur_date: str, cur_time: str) -> List[Dict]:
    """获取缓存的ETF排名，保证同一交易日内多次调用结果一致"""
    today = cur_date
    if g['rankings_cache']['date'] != today:
        print("  重新计算ETF排名...")
        ranked = get_ranked_etfs(ContextInfo, cur_date, cur_time)
        g['rankings_cache'] = {'date': today, 'data': ranked}
    else:
        print("  [DEBUG] 使用缓存的ETF排名")
    return g['rankings_cache']['data']


def get_ranked_etfs(ContextInfo, cur_date: str, cur_time: str) -> List[Dict]:
    """计算所有ETF的动量得分，根据行情状态动态选择ETF池"""
    active_pool = get_active_etf_pool()
    end_date = trader.get_prev_trade_date(cur_date)

    if not end_date:
        print(f"  [排名计算] 无法获取前一交易日，跳过本次计算: cur_date={cur_date}")
        return []

    print(f"  [排名计算] 使用ETF池，合计{len(active_pool)}只ETF")

    lookback = max(g['lookback_days'], g['short_lookback_days'], g['volume_lookback']) + 20

    # 获取历史日线数据（close + volume）
    market_data = trader.get_daily_data_cached(
        ['close', 'volume'], active_pool,
        count=lookback + 20, end_time=end_date
    )

    if not market_data:
        print("  [排名计算] 无法获取历史价格数据")
        return []

    # 获取今日1分钟量价数据（批量获取，避免后续逐只调用 trader.get_price）
    today_vols: Dict[str, float] = {}
    today_prices: Dict[str, float] = {}
    today_suspended: Set[str] = set()

    try:
        minute_data = ContextInfo.get_market_data_ex(
            ['volume', 'close', 'suspendFlag'],
            active_pool,
            period='1m',
            start_time=cur_date + '093000',
            end_time=cur_date + cur_time + '00',
            count=240,
            fill_data=True,
            subscribe=True
        )
        if minute_data:
            for code, data in minute_data.items():
                # 停牌检查: suspendFlag=1 表示停牌
                if 'suspendFlag' in data and data['suspendFlag'] is not None and len(data['suspendFlag']) > 0:
                    if data['suspendFlag'].iloc[-1] == 1:
                        today_suspended.add(code)
                        continue
                if 'volume' in data and data['volume'] is not None and len(data['volume']) > 0:
                    today_vols[code] = data['volume'].sum()
                if 'close' in data and data['close'] is not None and len(data['close']) > 0:
                    close_arr = data['close'].values
                    valid_closes = close_arr[~np.isnan(close_arr)]
                    if len(valid_closes) > 0:
                        today_prices[code] = valid_closes[-1]

        if today_suspended:
            print(f"  ?? 当日停牌ETF({len(today_suspended)}只): {', '.join(sorted(today_suspended))}")

    except Exception as e:
        print(f"  [排名计算] 分钟级数据获取异常: {e}")

    etf_metrics = []
    suspended_count = 0

    for etf in active_pool:
        # 停牌检查
        if etf in today_suspended:
            suspended_count += 1
            continue

        if etf not in market_data:
            continue

        data = market_data[etf]
        if 'close' not in data or data['close'] is None or len(data['close']) == 0:
            continue
        if len(data['close']) < g['lookback_days']:
            continue

        raw_closes = data['close'].values
        raw_volumes = data['volume'].values if 'volume' in data else np.zeros(len(raw_closes))

        valid_mask = (~np.isnan(raw_volumes)) & (raw_volumes > 0)
        hist_closes = raw_closes[valid_mask][-lookback:]
        hist_volumes = raw_volumes[valid_mask][-lookback:]

        if len(hist_closes) < g['lookback_days']:
            continue

        # 从批量获取的1分钟线中取最后一根有效close作为当前价
        current_price = today_prices.get(etf)
        if current_price is None or np.isnan(current_price):
            current_price = trader.get_price(ContextInfo, etf)  # fallback
        today_vol = today_vols.get(etf, 0)

        metrics = calculate_momentum_metrics(
            etf, hist_closes, hist_volumes, current_price, today_vol, ContextInfo
        )

        if metrics is not None:
            if g['min_score_threshold'] < metrics['score'] < g['max_score_threshold']:
                etf_metrics.append(metrics)
            else:
                name = metrics.get('etf_name', '')
                if g['enable_detailed_log']:
                    print(f"  [DEBUG] {etf} {name} 得分{metrics['score']:.2f}超出阈值，过滤")

    # 排名计算汇总日志
    missing_daily = len(active_pool) - suspended_count - len(
        [1 for e in active_pool if e not in today_suspended and e in market_data]
    )
    data_issues = []
    if suspended_count:
        data_issues.append(f"停牌{suspended_count}只")
    if missing_daily:
        data_issues.append(f"日线缺失{missing_daily}只")
    if data_issues:
        print(f"  [排名·数据质量] {', '.join(data_issues)} | 成功计算{len(etf_metrics)}只")

    etf_metrics.sort(key=lambda x: x['score'], reverse=True)
    return etf_metrics


def calculate_momentum_metrics(etf: str, hist_closes, hist_volumes, current_price, today_vol, ContextInfo):
    """计算单只ETF的动量指标（加权线性回归得分 = 年化收益率 × R2），应用所有过滤条件"""
    try:
        name = get_stock_name(etf)

        # 价格序列（历史日线收盘价 + 当天当前价）
        price_series = np.append(hist_closes, current_price)

        # ===== 1. 盈利保护检查（排除） =====
        if check_profit_protection(etf, ContextInfo):
            print(f"  ?? {etf} {name} 触发盈利保护，从排名中排除")
            return None

        # ===== 2. 溢价率过滤（排除） =====
        if g['enable_premium_filter']:
            premium, passed = get_premium_rate(etf, current_price, ContextInfo)
            if premium is not None:
                if not passed:
                    print(f"  ?? {etf} {name} 溢价率{premium*100:.2f}% > 阈值{g['premium_threshold']*100:.0f}%，排除")
                    return None
            else:
                # 实盘模式下无法获取溢价率
                if not ContextInfo.do_back_test:
                    print(f"  ?? {etf} {name} 无法获取溢价率，排除")
                    return None

        # ===== 3. 成交量过滤（排除） =====
        if g['enable_volume_check']:
            vol_ratio = trader.get_volume_ratio(ContextInfo, hist_volumes, today_vol, g['volume_lookback'])
            if vol_ratio is not None and vol_ratio > g['volume_threshold']:
                annualized = get_annualized_returns(price_series, g['lookback_days'])
                if annualized > g['volume_return_limit']:
                    print(f"  ?? {etf} {name} 成交量放量{vol_ratio:.1f}倍，且年化{annualized*100:.1f}% > 阈值{g['volume_return_limit']*100:.1f}%，过滤")
                    return None

        # ===== 4. 短期动量过滤（排除） =====
        if len(price_series) >= g['short_lookback_days'] + 1:
            short_return = price_series[-1] / price_series[-(g['short_lookback_days'] + 1)] - 1
            short_annualized = (1 + short_return) ** (250 / g['short_lookback_days']) - 1
        else:
            short_annualized = 0

        if g['use_short_momentum_filter'] and short_annualized < g['short_momentum_threshold']:
            if g['enable_detailed_log']:
                print(f"  [DEBUG] {etf} {name} 短期动量{short_annualized*100:.1f}% < 阈值{g['short_momentum_threshold']*100:.1f}%，过滤")
            return None

        # ===== 5. 长期动量计算（加权线性回归得分） =====
        recent = price_series[-(g['lookback_days'] + 1):]
        y = np.log(recent)
        x = np.arange(len(y))
        weights = np.linspace(1, 2, len(y))
        slope, intercept = np.polyfit(x, y, 1, w=weights)
        annualized_returns = math.exp(slope * 250) - 1

        # R2（趋势稳定性）
        ss_res = np.sum(weights * (y - (slope * x + intercept)) ** 2)
        ss_tot = np.sum(weights * (y - np.mean(y)) ** 2)
        r_squared = 1 - ss_res / ss_tot if ss_tot != 0 else 0

        score = annualized_returns * r_squared

        # ===== 6. 近3日单日跌幅过滤（排除） =====
        if len(price_series) >= 4:
            day1 = price_series[-1] / price_series[-2]
            day2 = price_series[-2] / price_series[-3]
            day3 = price_series[-3] / price_series[-4]
            if min(day1, day2, day3) < g['loss']:
                print(f"  ?? {etf} {name} 近3日有单日跌幅超{(1-g['loss'])*100:.1f}%，直接排除")
                return None

        return {
            'etf': etf,
            'etf_name': name,
            'annualized_returns': annualized_returns,
            'r_squared': r_squared,
            'score': score,
            'current_price': current_price,
            'short_annualized': short_annualized,
        }

    except Exception as e:
        name = get_stock_name(etf)
        print(f"  [warning] 计算{etf} {name}时出错: {e}")
        return None


def get_annualized_returns(price_series, lookback_days):
    """计算加权年化收益率"""
    recent = price_series[-(lookback_days + 1):]
    y = np.log(recent)
    x = np.arange(len(y))
    weights = np.linspace(1, 2, len(y))
    slope, _ = np.polyfit(x, y, 1, w=weights)
    return math.exp(slope * 250) - 1


# ---------------------------------------------------------------------------
# 【十】卖出模块
# ---------------------------------------------------------------------------

def etf_sell_trade(ContextInfo, cur_date: str, cur_time: str):
    """卖出不符合条件的持仓（排名变化、溢价率过高）"""
    print("========== 卖出操作开始 ==========")

    ranked = get_cached_rankings(ContextInfo, cur_date, cur_time)

    # 确定目标ETF列表（得分前N名且满足得分阈值）
    target_etfs = []
    for m in ranked[:g['holdings_num']]:
        if m['score'] >= g['min_score_threshold']:
            target_etfs.append(m['etf'])

    # 若没有目标ETF且防御可用，则把防御ETF作为目标（供卖出判断用）
    defensive_available = check_defensive_etf_available(ContextInfo)
    if not target_etfs and defensive_available:
        target_etfs = [g['defensive_etf']]
        name = get_stock_name(g['defensive_etf'])
        print(f"  ??? 无目标ETF，防御模式：{g['defensive_etf']} {name}")

    target_set = set(target_etfs)

    # 卖出不在目标列表的持仓
    holdings = trader.get_holdings()
    for sec, pos in holdings.items():
        if pos.m_nVolume <= 0:
            continue
        if sec not in g['etf_pool'] and sec != g['defensive_etf']:
            continue
        if sec not in target_set:
            cost = pos.m_dOpenPrice
            buy_date = g['buy_date'].get(sec)
            hold_days = 0
            if buy_date:
                try:
                    hold_days = (datetime.strptime(cur_date, '%Y%m%d').date()
                                 - datetime.strptime(buy_date, '%Y%m%d').date()).days
                except:
                    pass

            if smart_order_target_value(sec, 0, ContextInfo):
                name = get_stock_name(sec)
                sell_price = trader.get_price(ContextInfo, sec) or 0
                print(f"  ?? 卖出持仓：{sec} {name}")
                g['trade_log']['sell_records'].append({
                    'code': sec,
                    'name': name,
                    'cost': cost,
                    'price': sell_price,
                    'hold_days': hold_days
                })
                if sec in g['buy_date']:
                    del g['buy_date'][sec]

    # 溢价率检查（对仍持有的标的）
    if g['enable_premium_filter']:
        holdings = trader.get_holdings()
        for sec, pos in holdings.items():
            if pos.m_nVolume <= 0:
                continue
            if sec not in g['etf_pool'] and sec != g['defensive_etf']:
                continue
            current_price = trader.get_price(ContextInfo, sec)
            if current_price:
                premium, passed = get_premium_rate(sec, current_price, ContextInfo)
                if premium is not None and not passed:
                    if smart_order_target_value(sec, 0, ContextInfo):
                        name = get_stock_name(sec)
                        print(f"  ?? 溢价率过高 {sec} {name} 溢价率{premium*100:.2f}% > {g['premium_threshold']*100:.0f}%，卖出")

    print("========== 卖出操作完成 ==========")


# ---------------------------------------------------------------------------
# 【十一】买入模块
# ---------------------------------------------------------------------------

def check_intraday_drawdown_for_buy(security: str, ContextInfo, cur_date: str, cur_time: str) -> bool:
    """买入前检查：ETF当前是否处于日内显著回撤状态
    返回 True 表示正在回撤（不宜买入），False 表示正常
    """
    try:
        minute_data = ContextInfo.get_market_data_ex(
            ['high', 'close'],
            [security],
            period='1m',
            start_time=cur_date + '093000',
            end_time=cur_date + cur_time + '00',
            count=240,
            fill_data=True,
            subscribe=True
        )
        if not minute_data or security not in minute_data:
            return False

        data = minute_data[security]
        if 'high' not in data or data['high'] is None or len(data['high']) == 0:
            return False

        day_high = data['high'].max()
        close_arr = data['close'].values
        valid_closes = close_arr[~np.isnan(close_arr)]
        if len(valid_closes) == 0:
            return False
        current = valid_closes[-1]

        if day_high <= 0:
            return False

        drawdown = (day_high - current) / day_high
        return drawdown >= g['intraday_drawdown_threshold']
    except:
        return False


def etf_buy_trade(ContextInfo, cur_date: str, cur_time: str):
    """买入符合条件的ETF，等权分配，按排名顺序逐个尝试直到凑够持仓数量"""
    print("========== 买入操作开始 ==========")

    ranked = get_cached_rankings(ContextInfo, cur_date, cur_time)

    # 打印排名前5的指标
    print("  === ETF排名前5 ===")
    for i, m in enumerate(ranked[:5]):
        print(f"    排名{i+1}: {m['etf']} {m['etf_name']} 得分{m['score']:.4f} 年化{m['annualized_returns']*100:.2f}% R2={m['r_squared']:.4f}")

    # 确定目标ETF列表
    target_etfs = []
    for m in ranked:
        if len(target_etfs) >= g['holdings_num']:
            break

        # 得分过滤
        if m['score'] < g['min_score_threshold']:
            continue

        etf = m['etf']

        # 盈利保护二次检查（防止卖了又买）
        if g['enable_profit_protection'] and check_profit_protection(etf, ContextInfo):
            name = get_stock_name(etf)
            print(f"  ?? {etf} {name} 触发盈利保护，从买入候选列表中排除")
            continue

        # 日内回撤卖出检查（回撤/盈利保护卖出的标的禁止日内买回）
        if etf in g['drawdown_selled_today']:
            name = get_stock_name(etf)
            print(f"  ?? {etf} {name} 今日因回撤/盈利保护卖出，禁止日内买回")
            continue

        # 买入时实时回撤检查（避免"接飞刀"——当前正在大幅回撤的ETF暂不买入）
        if check_intraday_drawdown_for_buy(etf, ContextInfo, cur_date, cur_time):
            name = get_stock_name(etf)
            print(f"  ?? {etf} {name} 当前处于日内回撤状态(>{g['intraday_drawdown_threshold']*100:.0f}%)，暂不买入")
            continue

        # 溢价率过滤
        if g['enable_premium_filter']:
            current_price = m['current_price']
            premium, passed = get_premium_rate(etf, current_price, ContextInfo)
            name = get_stock_name(etf)
            if premium is None and not passed:
                if not ContextInfo.do_back_test:
                    print(f"  ?? {etf} {name} 无法获取溢价率，视为不合格，跳过")
                    continue
            elif premium is not None and not passed:
                print(f"  ?? {etf} {name} 溢价率{premium*100:.2f}% > {g['premium_threshold']*100:.0f}%，跳过")
                continue
            elif premium is not None:
                print(f"  ? {etf} {name} 溢价率{premium*100:.2f}% ≤ {g['premium_threshold']*100:.0f}%，通过")

        # 通过所有检查，加入目标列表
        target_etfs.append(etf)
        name = get_stock_name(etf)
        print(f"  ?? 目标ETF {len(target_etfs)}: {etf} {name} 得分{m['score']:.4f}")

    # 防御模式判断
    if not target_etfs:
        if check_defensive_etf_available(ContextInfo) and g['defensive_etf'] not in g['drawdown_selled_today']:
            target_etfs = [g['defensive_etf']]
            name = get_stock_name(g['defensive_etf'])
            print(f"  ??? 进入防御模式，选择防御ETF：{g['defensive_etf']} {name}")
        else:
            print("  ?? 无目标ETF且防御不可用，保持空仓")
            return

    # 检查是否有持仓需要先卖出（不在目标列表的持仓）
    holdings = trader.get_holdings()
    current_positions = []
    for s, pos in holdings.items():
        if pos.m_nVolume > 0 and (s in g['etf_pool'] or s == g['defensive_etf']):
            current_positions.append(s)
    current_pos_set = set(current_positions)
    target_set = set(target_etfs)
    to_sell = [s for s in current_pos_set if s not in target_set]

    if to_sell:
        to_sell_names = [get_stock_name(s) for s in to_sell]
        print(f"  尚有持仓需要卖出：{list(zip(to_sell, to_sell_names))}，等待卖出完成再买入")
        return

    # 等权分配
    total_val = trader.get_strategy_total_value(ContextInfo)
    target_per_etf = total_val / len(target_etfs)

    for etf in target_etfs:
        current_val = 0
        if etf in holdings:
            pos = holdings[etf]
            if pos.m_nVolume > 0:
                current_price = trader.get_price(ContextInfo, etf)
                current_val = pos.m_nVolume * current_price if current_price else 0

        # 5%容差调仓
        if abs(current_val - target_per_etf) > target_per_etf * 0.05 or current_val == 0:
            if smart_order_target_value(etf, target_per_etf, ContextInfo):
                action = "买入" if current_val < target_per_etf else "调仓"
                name = get_stock_name(etf)
                print(f"  ?? {action}：{etf} {name} 目标金额{target_per_etf:.2f}")

    print("========== 买入操作完成 ==========")


# ---------------------------------------------------------------------------
# 【十二】辅助函数
# ---------------------------------------------------------------------------

def get_stock_name(symbol: str) -> str:
    """获取证券名称"""
    try:
        return trader.get_stock_info(symbol, 'InstrumentName') or '未知'
    except:
        return '未知'


def check_defensive_etf_available(ContextInfo) -> bool:
    """检查防御ETF是否可交易（未停牌、未涨跌停）"""
    etf = g['defensive_etf']
    try:
        if ContextInfo.do_back_test:
            current_price = trader.get_price(ContextInfo, etf)
            if not current_price:
                return False
            return True
        else:
            tick_data = ContextInfo.get_full_tick([etf])
            if etf not in tick_data:
                return False
            tick = tick_data[etf]
            # openInt: 1=停牌, 17=临时停牌
            if tick.get('openInt', 0) in [1, 17]:
                name = get_stock_name(etf)
                print(f"  [DEBUG] 防御ETF {etf} {name} 停牌")
                return False
            last_price = tick.get('lastPrice', 0)
            stock_info = trader.get_stock_info(etf)
            if stock_info:
                high_limit = stock_info.get('UpStopPrice', 0)
                low_limit = stock_info.get('DownStopPrice', 0)
                if high_limit and high_limit > 0 and last_price >= high_limit:
                    name = get_stock_name(etf)
                    print(f"  [DEBUG] 防御ETF {etf} {name} 涨停")
                    return False
                if low_limit and low_limit > 0 and last_price <= low_limit:
                    name = get_stock_name(etf)
                    print(f"  [DEBUG] 防御ETF {etf} {name} 跌停")
                    return False
            return True
    except Exception as e:
        print(f"  检查防御ETF异常: {e}")
        return False


def smart_order_target_value(security: str, target_value: float, ContextInfo) -> bool:
    """智能下单：根据目标市值调整持仓，处理停牌、涨跌停、最小交易金额、T+1"""
    try:
        account_id = getattr(trader, 'acct_id', 'tests')
        stock_name = get_stock_name(security)

        current_price = trader.get_price(ContextInfo, security)
        if not current_price:
            print(f"  {security} {stock_name}: 获取价格失败，跳过交易")
            return False

        # 实盘模式，判断涨跌停
        if not ContextInfo.do_back_test:
            stock_info = trader.get_stock_info(security)
            if stock_info:
                high_limit = stock_info.get('UpStopPrice', 0)
                low_limit = stock_info.get('DownStopPrice', 0)
                if high_limit and high_limit > 0 and current_price >= high_limit:
                    print(f"  {security} {stock_name}: 当前涨停，跳过交易")
                    return False
                if low_limit and low_limit > 0 and current_price <= low_limit:
                    print(f"  {security} {stock_name}: 当前跌停，跳过交易")
                    return False

        # 计算目标数量（按100股整数倍）
        target_amount = int(target_value / current_price)
        target_amount = (target_amount // 100) * 100
        if target_amount <= 0 and target_value > 0:
            target_amount = 100

        # 获取当前持仓
        current_amount = 0
        closeable_amount = 0
        try:
            holdings = trader.get_holdings()
            if security in holdings:
                pos = holdings[security]
                current_amount = pos.m_nVolume
                closeable_amount = pos.m_nCanUseVolume
        except:
            pass

        amount_diff = target_amount - current_amount

        # 最小交易金额检查
        trade_val = abs(amount_diff) * current_price
        if 0 < trade_val < g['min_money']:
            print(f"  {security} {stock_name}: 交易金额{trade_val:.2f} < {g['min_money']}，跳过")
            return False

        # T+1处理：当天买入不可卖出
        if amount_diff < 0:
            if closeable_amount == 0:
                print(f"  {security} {stock_name}: 当天买入不可卖出(T+1)")
                return False
            amount_diff = -min(abs(amount_diff), closeable_amount)

        if amount_diff != 0:
            if amount_diff > 0:
                passorder(23, 1101, account_id, security, 5, -1, amount_diff, ContextInfo)
                print(f"  ?? 买入 {security} {stock_name} 数量{amount_diff} 价格{current_price:.3f}")
                g['buy_date'][security] = timetag_to_datetime(
                    ContextInfo.get_bar_timetag(ContextInfo.barpos), '%Y%m%d'
                )
            else:
                passorder(24, 1101, account_id, security, 5, -1, abs(amount_diff), ContextInfo)
                print(f"  ?? 卖出 {security} {stock_name} 数量{abs(amount_diff)} 价格{current_price:.3f}")
            return True
        return False
    except Exception as e:
        print(f"  下单异常: {security} - {e}")
        return False


# ---------------------------------------------------------------------------
# 【十三】盘后总结报告
# ---------------------------------------------------------------------------

def daily_summary_report(ContextInfo, cur_date: str):
    """盘后总结报告"""
    print("========== 策略运行日报 ==========")
    print(f"  日期: {cur_date[:4]}-{cur_date[4:6]}-{cur_date[6:]}")

    # 市场状态
    if g['enable_regime_switch']:
        status = "??走弱期" if g['is_a_share_weak'] else "??正常期"
        avoid_status = "回避A股" if (g['is_a_share_weak'] and g['enable_avoid_a_share']) else \
                       ("不回避A股" if g['is_a_share_weak'] else "正常交易")
        drawdown_status = "???启用" if (g['is_a_share_weak'] and g['enable_intraday_drawdown']) else \
                          ("?关闭" if (g['is_a_share_weak'] and not g['enable_intraday_drawdown']) else "?关闭")
        print(f"  市场状态：{status} | {avoid_status} 计数:{g['weak_period_counter']}/{g['weak_period_max_days']}")
        print(f"  分钟级回撤：{drawdown_status}（阈值{g['intraday_drawdown_threshold']*100:.0f}%）")
    else:
        print("  行情判断未启用，始终全市场交易")

    # 独立开关汇总
    avoid_switch_status = "ON（走弱期回避A股）" if g['enable_avoid_a_share'] else "OFF（走弱期不回避A股）"
    drawdown_switch_status = "ON（走弱期自动启用）" if g['enable_intraday_drawdown'] else "OFF（不触发）"
    print(f"  独立开关：A股回避={avoid_switch_status} | 分钟回撤={drawdown_switch_status}")

    # 今日卖出记录
    sell_records = g['trade_log'].get('sell_records', [])
    print(f"  今日卖出：{len(sell_records)}只")
    for r in sell_records:
        cost = r.get('cost', 0)
        sell_price = r.get('price', 0)
        profit_pct = (sell_price / cost - 1) * 100 if cost > 0 else 0
        hold_days = r.get('hold_days', 0)
        print(f"    {r['code']} {r['name']} | 成本:{cost:.3f} | 卖出:{sell_price:.3f} | 收益:{profit_pct:+.2f}% | 持有{hold_days}天")

    # 最终持仓
    holdings = trader.get_holdings()
    pos_list = []
    for sec, pos in holdings.items():
        if pos.m_nVolume <= 0:
            continue
        if sec not in g['etf_pool'] and sec != g['defensive_etf']:
            continue
        pos_list.append(sec)
    print(f"  最终持仓：{len(pos_list)}只")
    for sec, pos in holdings.items():
        if pos.m_nVolume <= 0:
            continue
        if sec not in g['etf_pool'] and sec != g['defensive_etf']:
            continue
        current_price = trader.get_price(ContextInfo, sec) or 0
        cost = pos.m_dOpenPrice
        profit_pct = (current_price / cost - 1) * 100 if cost > 0 else 0
        buy_date = g['buy_date'].get(sec)
        hold_days = 0
        if buy_date:
            try:
                hold_days = (datetime.strptime(cur_date, '%Y%m%d').date()
                             - datetime.strptime(buy_date, '%Y%m%d').date()).days
            except:
                pass
        print(f"    {sec} {get_stock_name(sec)} | 成本:{cost:.3f} | 当前:{current_price:.3f} | 收益:{profit_pct:+.2f}% | 持有{hold_days}天")

    # 账户汇总
    total_val = trader.get_strategy_total_value(ContextInfo)
    positions_value = total_val - trader.get_available_cash()
    cash = trader.get_available_cash()
    print(f"  总资产：{total_val:.2f} | 可用：{cash:.2f} | 市值：{positions_value:.2f}")
    print("==========" + "报告结束" + "==========")
    print("")


# ---------------------------------------------------------------------------
# 【十四】历史数据下载
# ---------------------------------------------------------------------------

def dl_history_data(ContextInfo, cur_date: str):
    """下载历史数据（日线 + 1分钟线）"""
    # 回测模式下，只下载一次
    if ContextInfo.do_back_test and g['dl_history_data_date']:
        return

    print(f"  下载ETF池[{len(g['all_subscribe_codes'])}]历史行情...")
    for code in g['all_subscribe_codes']:
        download_history_data(code, '1m', '', '')
        download_history_data(code, '1d', '', '')
    g['dl_history_data_date'] = cur_date