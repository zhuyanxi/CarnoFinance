<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/state';
	import {
		calcTradeAmount,
		closeOptionTrackPlan,
		createOptionTrackTrade,
		deleteOptionTrackTrade,
		formatOptionTrackContractMeta,
		formatMoney,
		formatOptionTrackDateTime,
		getOptionTrackPlanDetail,
		getOptionTrackPnL,
		isOptionTrackOption,
		previewOptionTrackPnL,
		toDateTimeLocalValue,
		updateOptionTrackTrade,
		type OptionTrackPlanDetail,
		type OptionTrackPnLResponse,
		type OptionTrackTradeInput,
		type OptionTrackTradeRecord
	} from '$lib/optiontrack';

	let loading = $state(true);
	let savingTrade = $state(false);
	let closingPlan = $state(false);
	let previewing = $state(false);
	let error = $state('');
	let detail = $state<OptionTrackPlanDetail | null>(null);
	let pnl = $state<OptionTrackPnLResponse | null>(null);
	let editingTradeId = $state<number | null>(null);
	let priceOverrides = $state<Record<string, number>>({});
	let tradeForm = $state<OptionTrackTradeInput>(createEmptyTrade());

	function createEmptyTrade(): OptionTrackTradeInput {
		return {
			instrument_type: 'OPTION',
			contract_name: '',
			contract_code: '',
			expiry_date: '',
			strike_price: 0,
			option_side: 'CALL',
			contract_multiplier: 10000,
			trade_type: 'BUY_OPEN',
			trade_time: toDateTimeLocalValue(),
			price: 0,
			quantity: 1,
			remark: ''
		};
	}

	const planId = Number(page.params.id);

	async function loadAll() {
		loading = true;
		error = '';
		try {
			detail = await getOptionTrackPlanDetail(planId);
			pnl = await getOptionTrackPnL(planId);
			const nextOverrides: Record<string, number> = {};
			for (const contract of pnl.contracts) {
				nextOverrides[contract.contract_key] = contract.current_price;
			}
			priceOverrides = nextOverrides;
		} catch (err: any) {
			error = err.message;
		} finally {
			loading = false;
		}
	}

	function startEdit(trade: OptionTrackTradeRecord) {
		editingTradeId = trade.id;
		tradeForm = {
			instrument_type: trade.instrument_type,
			contract_name: trade.contract_name,
			contract_code: trade.contract_code ?? '',
			expiry_date: trade.expiry_date,
			strike_price: trade.strike_price,
			option_side: trade.option_side,
			contract_multiplier: trade.contract_multiplier,
			trade_type: trade.trade_type,
			trade_time: toDateTimeLocalValue(trade.trade_time),
			price: trade.price,
			quantity: trade.quantity,
			remark: trade.remark ?? ''
		};
	}

	function resetTradeForm() {
		editingTradeId = null;
		tradeForm = createEmptyTrade();
	}

	async function handleTradeSubmit(event: Event) {
		event.preventDefault();
		savingTrade = true;
		error = '';
		try {
			if (editingTradeId) {
				await updateOptionTrackTrade(planId, editingTradeId, tradeForm);
			} else {
				await createOptionTrackTrade(planId, tradeForm);
			}
			resetTradeForm();
			await loadAll();
		} catch (err: any) {
			error = err.message;
		} finally {
			savingTrade = false;
		}
	}

	async function handleDeleteTrade(tradeId: number) {
		if (!window.confirm('确认删除这条交易记录?')) {
			return;
		}
		error = '';
		try {
			await deleteOptionTrackTrade(planId, tradeId);
			if (editingTradeId === tradeId) {
				resetTradeForm();
			}
			await loadAll();
		} catch (err: any) {
			error = err.message;
		}
	}

	async function handleClosePlan() {
		if (!window.confirm('确认关闭计划? 仅当全部持仓已清仓时成功。')) {
			return;
		}
		closingPlan = true;
		error = '';
		try {
			await closeOptionTrackPlan(planId);
			await loadAll();
		} catch (err: any) {
			error = err.message;
		} finally {
			closingPlan = false;
		}
	}

	async function handlePreviewPnL() {
		previewing = true;
		error = '';
		try {
			pnl = await previewOptionTrackPnL(planId, {
				prices: Object.entries(priceOverrides)
					.filter(([, currentPrice]) => currentPrice > 0)
					.map(([contract_key, current_price]) => ({ contract_key, current_price }))
			});
		} catch (err: any) {
			error = err.message;
		} finally {
			previewing = false;
		}
	}

	onMount(() => {
		loadAll();
	});
</script>

<div class="space-y-8">
	<section class="hero">
		<div>
			<p class="eyebrow">S02 / S03 / S05</p>
			<h1>{detail?.plan.name ?? `计划 #${planId}`}</h1>
			<p class="desc">交易明细、当前持仓、计划关闭、盈亏试算都在这里。</p>
		</div>
		<div class="hero-actions">
			<a class="ghost" href="/optiontrack">返回列表</a>
			{#if detail?.plan.status === 'IN_PROGRESS'}
				<button class="primary" disabled={closingPlan} onclick={handleClosePlan}>
					{closingPlan ? '关闭中...' : '关闭计划'}
				</button>
			{/if}
		</div>
	</section>

	{#if error}
		<div class="alert">{error}</div>
	{/if}

	{#if loading}
		<div class="empty">加载中...</div>
	{:else if detail && pnl}
		<section class="stats-grid">
			<div class="stat-card">
				<span>计划状态</span>
				<strong>{detail.plan.status}</strong>
				<small>创建于 {formatOptionTrackDateTime(detail.plan.created_at)}</small>
			</div>
			<div class="stat-card">
				<span>总盈亏</span>
				<strong class:negative={pnl.total_pnl < 0}>{formatMoney(pnl.total_pnl)}</strong>
				<small>已实现 {formatMoney(pnl.realized_pnl)} / 未实现 {formatMoney(pnl.unrealized_pnl)}</small>
			</div>
			<div class="stat-card">
				<span>合约数量</span>
				<strong>{detail.plan.contract_count}</strong>
				<small>最新交易 {formatOptionTrackDateTime(detail.plan.latest_trade_at)}</small>
			</div>
		</section>

		<section class="panel">
			<div class="panel-head">
				<h2>当前持仓</h2>
				<button class="secondary" disabled={previewing} onclick={handlePreviewPnL}>
					{previewing ? '试算中...' : '刷新试算'}
				</button>
			</div>

			{#if detail.positions.length === 0}
				<div class="empty">暂无持仓或历史交易。</div>
			{:else}
				<div class="table-wrap">
					<table>
						<thead>
							<tr>
								<th>合约</th>
								<th>类型</th>
								<th>多头</th>
								<th>空头</th>
								<th>默认价</th>
								<th>试算价</th>
							</tr>
						</thead>
						<tbody>
							{#each detail.positions as position}
								<tr>
									<td>
										<div class="name">{position.contract_name}</div>
										<div class="sub">{formatOptionTrackContractMeta(position)}</div>
									</td>
									<td>{position.instrument_type}</td>
									<td>{position.long_quantity}</td>
									<td>{position.short_quantity}</td>
									<td>{formatMoney(position.last_trade_price)}</td>
									<td>
										<input
											type="number"
											step="0.0001"
											bind:value={priceOverrides[position.contract_key]}
											disabled={position.long_quantity === 0 && position.short_quantity === 0}
										/>
									</td>
								</tr>
							{/each}
						</tbody>
					</table>
				</div>
			{/if}
		</section>

		<section class="panel">
			<h2>分合约盈亏</h2>
			<div class="pnl-grid">
				{#each pnl.contracts as contract}
					<div class="pnl-card">
						<div class="name">{contract.contract_name}</div>
						<div class="sub">{formatOptionTrackContractMeta(contract)}</div>
						<div class="pnl-row"><span>当前价</span><strong>{formatMoney(contract.current_price)}</strong></div>
						<div class="pnl-row"><span>已实现</span><strong class:negative={contract.realized_pnl < 0}>{formatMoney(contract.realized_pnl)}</strong></div>
						<div class="pnl-row"><span>未实现</span><strong class:negative={contract.unrealized_pnl < 0}>{formatMoney(contract.unrealized_pnl)}</strong></div>
						<div class="pnl-row total"><span>总盈亏</span><strong class:negative={contract.total_pnl < 0}>{formatMoney(contract.total_pnl)}</strong></div>
					</div>
				{/each}
			</div>
		</section>

		<section class="panel">
			<div class="panel-head">
				<h2>{editingTradeId ? '编辑交易记录' : '新增交易记录'}</h2>
				{#if editingTradeId}
					<button class="ghost" type="button" onclick={resetTradeForm}>取消编辑</button>
				{/if}
			</div>

			<form class="trade-grid" onsubmit={handleTradeSubmit}>
				<label>
					<span>品种类型</span>
					<select bind:value={tradeForm.instrument_type}>
						<option value="OPTION">期权 / OPTION</option>
						<option value="FUTURE">期货 / FUTURE</option>
					</select>
				</label>
				<label><span>合约名称</span><input bind:value={tradeForm.contract_name} required /></label>
				<label><span>合约代码</span><input bind:value={tradeForm.contract_code} /></label>
				<label>
					<span>{isOptionTrackOption(tradeForm.instrument_type) ? '到期日' : '交割日 / 合约月'}</span>
					<input type="date" bind:value={tradeForm.expiry_date} required={isOptionTrackOption(tradeForm.instrument_type)} />
				</label>
				{#if isOptionTrackOption(tradeForm.instrument_type)}
					<label><span>执行价</span><input type="number" min="0" step="0.01" bind:value={tradeForm.strike_price} required /></label>
					<label>
						<span>期权方向</span>
						<select bind:value={tradeForm.option_side}>
							<option value="CALL">认购 / CALL</option>
							<option value="PUT">认沽 / PUT</option>
						</select>
					</label>
				{/if}
				<label>
					<span>交易类型</span>
					<select bind:value={tradeForm.trade_type} disabled={detail.plan.status === 'CLOSED'}>
						<option value="BUY_OPEN">买开</option>
						<option value="SELL_OPEN">卖开</option>
						<option value="BUY_CLOSE">买平</option>
						<option value="SELL_CLOSE">卖平</option>
					</select>
				</label>
				<label><span>合约乘数</span><input type="number" min="1" bind:value={tradeForm.contract_multiplier} required /></label>
				<label><span>交易时间</span><input type="datetime-local" bind:value={tradeForm.trade_time} required /></label>
				<label><span>{isOptionTrackOption(tradeForm.instrument_type) ? '交易价格' : '期货价格'}</span><input type="number" min="0" step="0.0001" bind:value={tradeForm.price} required /></label>
				<label><span>张数</span><input type="number" min="1" bind:value={tradeForm.quantity} required /></label>
				<label class="full"><span>备注</span><input bind:value={tradeForm.remark} /></label>
				<div class="amount">金额: {calcTradeAmount(tradeForm.price, tradeForm.quantity, tradeForm.contract_multiplier).toFixed(2)}</div>
				<div class="actions full">
					<button class="primary" type="submit" disabled={savingTrade || detail.plan.status === 'CLOSED'}>
						{savingTrade ? '提交中...' : editingTradeId ? '保存修改' : '新增交易'}
					</button>
				</div>
			</form>
		</section>

		<section class="panel">
			<h2>交易记录</h2>
			{#if detail.trades.length === 0}
				<div class="empty">暂无交易记录。</div>
			{:else}
				<div class="table-wrap">
					<table>
						<thead>
							<tr>
								<th>时间</th>
								<th>合约</th>
								<th>类型</th>
								<th>价格</th>
								<th>张数</th>
								<th>金额</th>
								<th>来源</th>
								<th></th>
							</tr>
						</thead>
						<tbody>
							{#each detail.trades as trade}
								<tr>
									<td>{formatOptionTrackDateTime(trade.trade_time)}</td>
									<td>
										<div class="name">{trade.contract_name}</div>
										<div class="sub">{formatOptionTrackContractMeta(trade)}</div>
									</td>
									<td>{trade.trade_type}</td>
									<td>{formatMoney(trade.price)}</td>
									<td>{trade.quantity}</td>
									<td>{formatMoney(trade.amount)}</td>
									<td>{trade.entry_source}</td>
									<td>
										<div class="row-actions">
											<button class="ghost small" type="button" onclick={() => startEdit(trade)}>编辑</button>
											<button class="danger small" type="button" onclick={() => handleDeleteTrade(trade.id)}>删除</button>
										</div>
									</td>
								</tr>
							{/each}
						</tbody>
					</table>
				</div>
			{/if}
		</section>
	{/if}
</div>

<style lang="postcss">
	@reference "tailwindcss";

	:global(html) {
		background:
			radial-gradient(circle at top left, rgba(16, 185, 129, 0.14), transparent 24%),
			linear-gradient(180deg, #f8fffb 0%, #f8fafc 54%, #eef2ff 100%);
	}

	.hero {
		@apply flex flex-col gap-4 rounded-3xl border border-emerald-200/70 bg-white/85 p-8 shadow-xl shadow-emerald-100/60 md:flex-row md:items-end md:justify-between;
	}

	.hero-actions {
		@apply flex flex-wrap items-center gap-3;
	}

	.eyebrow {
		@apply text-xs font-semibold uppercase tracking-[0.3em] text-emerald-700;
	}

	h1 {
		@apply mt-2 text-4xl font-black tracking-tight text-slate-900;
	}

	.desc {
		@apply mt-3 max-w-2xl text-sm leading-6 text-slate-600;
	}

	.panel {
		@apply rounded-3xl border border-slate-200 bg-white/90 p-6 shadow-lg shadow-slate-200/70;
	}

	.panel h2 {
		@apply text-xl font-bold text-slate-900;
	}

	.panel-head {
		@apply mb-4 flex items-center justify-between gap-4;
	}

	.stats-grid {
		@apply grid grid-cols-1 gap-4 lg:grid-cols-3;
	}

	.stat-card {
		@apply rounded-3xl border border-slate-200 bg-white/90 p-6 shadow-lg shadow-slate-200/70;
	}

	.stat-card span {
		@apply text-xs font-semibold uppercase tracking-[0.2em] text-slate-500;
	}

	.stat-card strong {
		@apply mt-3 block text-3xl font-black text-slate-900;
	}

	.stat-card small {
		@apply mt-2 block text-sm text-slate-500;
	}

	.pnl-grid {
		@apply mt-4 grid grid-cols-1 gap-4 xl:grid-cols-2;
	}

	.pnl-card {
		@apply rounded-2xl border border-slate-200 bg-slate-50/80 p-5;
	}

	.pnl-row {
		@apply mt-3 flex items-center justify-between text-sm text-slate-600;
	}

	.pnl-row strong {
		@apply text-base font-semibold text-slate-900;
	}

	.total {
		@apply border-t border-slate-200 pt-3;
	}

	.trade-grid {
		@apply grid grid-cols-1 gap-4 md:grid-cols-2 xl:grid-cols-4;
	}

	label {
		@apply flex flex-col gap-2 text-sm font-medium text-slate-700;
	}

	label span {
		@apply text-xs uppercase tracking-[0.2em] text-slate-500;
	}

	input,
	select {
		@apply rounded-xl border-slate-200 bg-white px-4 py-3 text-sm text-slate-900 shadow-sm focus:border-emerald-400 focus:ring-emerald-400;
	}

	.table-wrap {
		@apply overflow-x-auto;
	}

	table {
		@apply min-w-full border-separate border-spacing-0;
	}

	th {
		@apply border-b border-slate-200 px-4 py-3 text-left text-xs font-semibold uppercase tracking-[0.2em] text-slate-500;
	}

	td {
		@apply border-b border-slate-100 px-4 py-4 text-sm text-slate-700;
	}

	.name {
		@apply font-semibold text-slate-900;
	}

	.sub {
		@apply mt-1 text-xs text-slate-500;
	}

	.amount {
		@apply rounded-2xl bg-emerald-50 px-4 py-3 text-sm font-semibold text-emerald-700;
	}

	.actions {
		@apply flex items-center justify-end;
	}

	.row-actions {
		@apply flex flex-wrap items-center justify-end gap-2;
	}

	.primary,
	.secondary,
	.ghost,
	.danger {
		@apply inline-flex items-center justify-center rounded-2xl px-5 py-3 text-sm font-semibold transition;
	}

	.small {
		@apply px-3 py-2 text-xs;
	}

	.primary {
		@apply bg-emerald-500 text-white shadow-lg shadow-emerald-200 hover:bg-emerald-600 disabled:cursor-not-allowed disabled:opacity-50;
	}

	.secondary,
	.ghost {
		@apply border border-slate-200 bg-white text-slate-700 hover:bg-slate-50;
	}

	.danger {
		@apply bg-rose-50 text-rose-700 hover:bg-rose-100;
	}

	.alert {
		@apply rounded-2xl border border-rose-200 bg-rose-50 px-4 py-3 text-sm text-rose-700;
	}

	.empty {
		@apply rounded-2xl border border-dashed border-slate-200 bg-slate-50 px-4 py-10 text-center text-sm text-slate-500;
	}

	.full {
		@apply md:col-span-2 xl:col-span-4;
	}

	.negative {
		@apply text-rose-600;
	}
</style>