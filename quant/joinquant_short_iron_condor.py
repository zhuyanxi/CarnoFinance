"""
JoinQuant Short Iron Condor strategy template

Usage:
- Adapt option API helper functions (placeholders) to JoinQuant environment.
- Upload this script to JoinQuant backtest, implement missing API calls if needed.

Features encoded from report:
- Short legs target delta ~0.10-0.16
- Wing spread configurable (0.10/0.20/0.30 price or strike steps)
- Entry DTE: 30-45 days
- Roll when remaining DTE <= 21 or profit >=50-75% or short leg delta > 0.30
- Risk rules: profit take, stop loss, max position sizing

This file is a runnable template; replace placeholder functions with actual JoinQuant option APIs.
"""
from datetime import datetime, timedelta
import math

# ----- Strategy parameters -----
TARGET_SHORT_DELTA = 0.14  # target absolute delta for short legs (0.10-0.16)
ENTRY_DTE_MIN = 30
ENTRY_DTE_MAX = 45
ROLL_DTE_THRESHOLD = 21
WING_SPREADS = [0.10, 0.20, 0.30]  # selectable
PROFIT_TAKE_RATIO = 0.6  # 60% of initial credit
STOP_LOSS_RATIO = 1.0  # 100% of initial credit
MAX_ACCOUNT_RISK = 0.2  # max % of account used for all credit spreads
SINGLE_POSITION_RISK = 0.1  # max % of account per iron condor


# ----- Placeholders for JoinQuant option APIs -----
def get_underlying_price(context, security):
    """Return current underlying price. Replace with JoinQuant API."""
    # Example placeholder call. Replace with real jqdata call if available.
    try:
        return context.run_query_get_price(security)
    except Exception:
        raise NotImplementedError("Implement get_underlying_price using JoinQuant API")


def get_option_chain(security, as_of_date):
    """Return list of option contracts near maturity for underlying security.
    Each contract dict must include: code, strike, right ('call'|'put'), expiry (date), delta, bid, ask
    Implement using JoinQuant option chain query.
    """
    raise NotImplementedError("Implement get_option_chain using JoinQuant jqdata option API")


def place_option_order(context, contract_code, quantity, side):
    """Place order for option contract. side: 'buy' or 'sell'.
    Replace with JoinQuant order_option or equivalent API.
    """
    raise NotImplementedError("Implement place_option_order with JoinQuant API")


# ----- Helpers -----
def find_contract_by_delta(chain, right, target_delta, expiry_window=(ENTRY_DTE_MIN, ENTRY_DTE_MAX)):
    """Find option contract in chain matching right and closest to target_delta and expiry window."""
    today = datetime.today().date()
    candidates = []
    for c in chain:
        if c['right'] != right:
            continue
        dte = (c['expiry'] - today).days
        if dte < expiry_window[0] or dte > expiry_window[1]:
            continue
        # delta provided as absolute for puts/calls
        candidates.append((abs(c['delta'] - target_delta), dte, c))
    if not candidates:
        return None
    candidates.sort(key=lambda x: (x[0], x[1]))
    return candidates[0][2]


def select_iron_condor(context, underlying):
    """Select short iron condor legs given underlying and parameters.
    Returns dict with leg contract codes and initial credit estimate.
    """
    chain = get_option_chain(underlying, datetime.today().date())
    # find short put and short call by target delta
    short_put = find_contract_by_delta(chain, 'put', TARGET_SHORT_DELTA)
    short_call = find_contract_by_delta(chain, 'call', TARGET_SHORT_DELTA)
    if not short_put or not short_call:
        return None

    # select wing strikes: move by nearest wing spread distance in price units
    # choose smallest wing that meets risk rules by default
    chosen_wing = WING_SPREADS[0]

    # find long protection legs: farther OTM by chosen_wing (strike difference)
    # naive approach: find put with strike = short_put.strike - chosen_wing
    target_put_strike = short_put['strike'] - chosen_wing
    target_call_strike = short_call['strike'] + chosen_wing

    long_put = min((c for c in chain if c['right'] == 'put'), key=lambda x: abs(x['strike'] - target_put_strike), default=None)
    long_call = min((c for c in chain if c['right'] == 'call'), key=lambda x: abs(x['strike'] - target_call_strike), default=None)

    if not long_put or not long_call:
        return None

    # estimate credit: (short_put.bid + short_call.bid) - (long_put.ask + long_call.ask)
    credit = (short_put.get('bid', 0) + short_call.get('bid', 0)) - (long_put.get('ask', 0) + long_call.get('ask', 0))

    return {
        'short_put': short_put,
        'long_put': long_put,
        'short_call': short_call,
        'long_call': long_call,
        'credit': credit,
        'wing': chosen_wing,
    }


# ----- Strategy lifecycle (JoinQuant hooks) -----
def initialize(context):
    """Called once at strategy start. Set context globals and schedule daily checks."""
    context.params = {
        'underlying': '510050.XSHG',  # 50ETF underlying code in JoinQuant
    }
    context.positions = []
    # schedule daily rebalance at market open
    try:
        schedule_function(daily_check, date_rule='every_day', time_rule='open')
    except Exception:
        # some environments use different schedule API; user adapt as needed
        pass


def daily_check(context):
    """Daily management: entry, roll, take profit, stop loss."""
    underlying = context.params['underlying']
    # entry logic: if no position, try to open
    if not context.positions:
        ic = select_iron_condor(context, underlying)
        if ic and ic['credit'] > 0:
            notional_risk = estimate_position_risk(context, ic)
            try:
                total_value = context.portfolio.total_value
            except Exception:
                total_value = 1
            if notional_risk / total_value <= SINGLE_POSITION_RISK:
                open_iron_condor(context, ic)
                context.positions.append({'ic': ic, 'entry_credit': ic['credit'], 'entry_date': datetime.today().date()})

    # manage existing positions
    for pos in list(context.positions):
        manage_position(context, pos)


def estimate_position_risk(context, ic):
    """Estimate potential max loss of iron condor in account currency.
    Simplified: wing width - credit, times contract multiplier (placeholder 1).
    """
    wing = ic['wing']
    credit = ic['credit']
    # contract multiplier placeholder
    multiplier = 1
    max_loss = (wing - credit) * multiplier
    return max_loss


def open_iron_condor(context, ic):
    """Place orders to open iron condor. Sell short legs, buy long legs.
    Quantity and sizing left to implement per account rules.
    """
    qty = 1  # placeholder quantity
    place_option_order(context, ic['short_put']['code'], qty, 'sell')
    place_option_order(context, ic['long_put']['code'], qty, 'buy')
    place_option_order(context, ic['short_call']['code'], qty, 'sell')
    place_option_order(context, ic['long_call']['code'], qty, 'buy')


def manage_position(context, pos):
    """Manage single iron condor: roll, take profit, stop loss."""
    ic = pos['ic']
    entry_credit = pos['entry_credit']
    current_credit = quote_iron_condor(context, ic)
    # profit take
    if current_credit <= (1 - PROFIT_TAKE_RATIO) * entry_credit:
        # realized credit shrink → take profit
        close_iron_condor(context, ic)
        context.positions.remove(pos)
        return

    # stop loss: if loss reaches STOP_LOSS_RATIO
    if current_credit >= (1 + STOP_LOSS_RATIO) * entry_credit:
        close_iron_condor(context, ic)
        context.positions.remove(pos)
        return

    # roll conditions: short leg delta too large or DTE <= ROLL_DTE_THRESHOLD
    short_put = ic['short_put']
    short_call = ic['short_call']
    today = datetime.today().date()
    min_dte = min((short_put['expiry'] - today).days, (short_call['expiry'] - today).days)
    if min_dte <= ROLL_DTE_THRESHOLD or short_put['delta'] > 0.30 or short_call['delta'] > 0.30:
        # simple strategy: close and attempt to open new iron condor
        close_iron_condor(context, ic)
        context.positions.remove(pos)
        new_ic = select_iron_condor(context, context.params['underlying'])
        if new_ic and new_ic['credit'] > 0:
            open_iron_condor(context, new_ic)
            context.positions.append({'ic': new_ic, 'entry_credit': new_ic['credit'], 'entry_date': datetime.today().date()})


def quote_iron_condor(context, ic):
    """Return current net credit for iron condor using mid prices."""
    short_put = ic['short_put']
    short_call = ic['short_call']
    long_put = ic['long_put']
    long_call = ic['long_call']
    mid_short_put = (short_put.get('bid', 0) + short_put.get('ask', 0)) / 2
    mid_short_call = (short_call.get('bid', 0) + short_call.get('ask', 0)) / 2
    mid_long_put = (long_put.get('bid', 0) + long_put.get('ask', 0)) / 2
    mid_long_call = (long_call.get('bid', 0) + long_call.get('ask', 0)) / 2
    return (mid_short_put + mid_short_call) - (mid_long_put + mid_long_call)


def close_iron_condor(context, ic):
    """Close iron condor by reversing opening trades."""
    qty = 1
    place_option_order(context, ic['short_put']['code'], qty, 'buy')
    place_option_order(context, ic['long_put']['code'], qty, 'sell')
    place_option_order(context, ic['short_call']['code'], qty, 'buy')
    place_option_order(context, ic['long_call']['code'], qty, 'sell')


if __name__ == '__main__':
    print('This file is a template. Deploy on JoinQuant and implement API placeholders.')
