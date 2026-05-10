const OPTION_TRACK_API_BASE = 'http://localhost:8080/optiontrack';

export type OptionTrackInstrumentType = 'OPTION' | 'FUTURE';

export interface OptionTrackTradeInput {
	instrument_type: OptionTrackInstrumentType;
	contract_name: string;
	contract_code: string;
	expiry_date: string;
	strike_price: number;
	option_side: string;
	contract_multiplier: number;
	trade_type: string;
	trade_time: string;
	price: number;
	quantity: number;
	remark: string;
}

export interface OptionTrackPlanCreateInput {
	name: string;
	positions: OptionTrackTradeInput[];
}

export interface OptionTrackPlanSummary {
	id: number;
	name: string;
	status: string;
	created_at: string;
	closed_at?: string;
	latest_trade_at?: string;
	contract_count: number;
}

export interface OptionTrackTradeRecord {
	id: number;
	plan_id: number;
	instrument_type: OptionTrackInstrumentType;
	contract_name: string;
	contract_code?: string;
	expiry_date: string;
	strike_price: number;
	option_side: string;
	contract_multiplier: number;
	trade_type: string;
	trade_time: string;
	price: number;
	quantity: number;
	amount: number;
	entry_source: string;
	remark?: string;
	created_at: string;
	updated_at: string;
}

export interface OptionTrackPosition {
	contract_key: string;
	instrument_type: OptionTrackInstrumentType;
	contract_name: string;
	contract_code?: string;
	expiry_date: string;
	strike_price: number;
	option_side: string;
	contract_multiplier: number;
	long_quantity: number;
	short_quantity: number;
	net_quantity: number;
	can_sell_close: number;
	can_buy_close: number;
	last_trade_price: number;
	average_long_price: number;
	average_short_price: number;
}

export interface OptionTrackPlanDetail {
	plan: OptionTrackPlanSummary;
	trades: OptionTrackTradeRecord[];
	positions: OptionTrackPosition[];
}

export interface OptionTrackPlanListResponse {
	items: OptionTrackPlanSummary[];
	page: number;
	page_size: number;
	total: number;
	total_page: number;
}

export interface OptionTrackPnLContract {
	contract_key: string;
	instrument_type: OptionTrackInstrumentType;
	contract_name: string;
	contract_code?: string;
	expiry_date: string;
	strike_price: number;
	option_side: string;
	contract_multiplier: number;
	current_price: number;
	realized_pnl: number;
	unrealized_pnl: number;
	total_pnl: number;
	long_quantity: number;
	short_quantity: number;
	average_long_price: number;
	average_short_price: number;
}

export interface OptionTrackPnLResponse {
	plan_id: number;
	status: string;
	realized_pnl: number;
	unrealized_pnl: number;
	total_pnl: number;
	contracts: OptionTrackPnLContract[];
}

export interface OptionTrackPnLPreviewRequest {
	prices: Array<{
		contract_key: string;
		current_price: number;
	}>;
}

async function request<T>(path: string, init?: RequestInit): Promise<T> {
	const response = await fetch(`${OPTION_TRACK_API_BASE}${path}`, {
		headers: {
			'Content-Type': 'application/json',
			...(init?.headers ?? {})
		},
		...init
	});

	if (!response.ok) {
		const data = (await response.json().catch(() => null)) as { error?: string } | null;
		throw new Error(data?.error || `HTTP ${response.status}`);
	}

	if (response.status === 204) {
		return undefined as T;
	}

	return (await response.json()) as T;
}

export function listOptionTrackPlans(page = 1, pageSize = 10) {
	return request<OptionTrackPlanListResponse>(`/plans?page=${page}&page_size=${pageSize}`);
}

export function createOptionTrackPlan(payload: OptionTrackPlanCreateInput) {
	return request<OptionTrackPlanDetail>('/plans', {
		method: 'POST',
		body: JSON.stringify(payload)
	});
}

export function getOptionTrackPlanDetail(planId: number) {
	return request<OptionTrackPlanDetail>(`/plans/${planId}`);
}

export function createOptionTrackTrade(planId: number, payload: OptionTrackTradeInput) {
	return request<OptionTrackPlanDetail>(`/plans/${planId}/trades`, {
		method: 'POST',
		body: JSON.stringify(payload)
	});
}

export function updateOptionTrackTrade(
	planId: number,
	tradeId: number,
	payload: OptionTrackTradeInput
) {
	return request<OptionTrackPlanDetail>(`/plans/${planId}/trades/${tradeId}`, {
		method: 'PUT',
		body: JSON.stringify(payload)
	});
}

export function deleteOptionTrackTrade(planId: number, tradeId: number) {
	return request<OptionTrackPlanDetail>(`/plans/${planId}/trades/${tradeId}`, {
		method: 'DELETE'
	});
}

export function closeOptionTrackPlan(planId: number) {
	return request<OptionTrackPlanSummary>(`/plans/${planId}/close`, {
		method: 'POST'
	});
}

export function getOptionTrackPnL(planId: number) {
	return request<OptionTrackPnLResponse>(`/plans/${planId}/pnl`);
}

export function previewOptionTrackPnL(planId: number, payload: OptionTrackPnLPreviewRequest) {
	return request<OptionTrackPnLResponse>(`/plans/${planId}/pnl/preview`, {
		method: 'POST',
		body: JSON.stringify(payload)
	});
}

export function formatOptionTrackDateTime(value?: string) {
	if (!value) {
		return '-';
	}
	return new Date(value).toLocaleString();
}

export function toDateTimeLocalValue(value?: string) {
	const date = value ? new Date(value) : new Date();
	const pad = (input: number) => `${input}`.padStart(2, '0');
	return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(date.getDate())}T${pad(date.getHours())}:${pad(date.getMinutes())}`;
}

export function formatMoney(value: number) {
	return value.toFixed(2);
}

export function calcTradeAmount(price: number, quantity: number, multiplier: number) {
	return Number.isFinite(price * quantity * multiplier) ? price * quantity * multiplier : 0;
}

export function isOptionTrackOption(instrumentType?: string) {
	return instrumentType !== 'FUTURE';
}

export function formatOptionTrackContractMeta(input: {
	instrument_type?: string;
	contract_code?: string;
	expiry_date?: string;
	strike_price?: number;
	option_side?: string;
}) {
	const parts: string[] = [];
	if (input.contract_code) {
		parts.push(input.contract_code);
	}
	if (isOptionTrackOption(input.instrument_type)) {
		if (input.expiry_date) {
			parts.push(input.expiry_date);
		}
		if ((input.strike_price ?? 0) > 0) {
			parts.push(`${input.strike_price}`);
		}
		if (input.option_side) {
			parts.push(input.option_side);
		}
	} else {
		parts.push('FUTURE');
		if (input.expiry_date) {
			parts.push(input.expiry_date);
		}
	}
	return parts.join(' / ');
}