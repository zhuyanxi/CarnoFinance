<script lang="ts">
	import { goto } from '$app/navigation';
	import {
		calcTradeAmount,
		createOptionTrackPlan,
		isOptionTrackOption,
		toDateTimeLocalValue,
		type OptionTrackTradeInput
	} from '$lib/optiontrack';

	type EditableTrade = OptionTrackTradeInput;

	let planName = $state('');
	let saving = $state(false);
	let error = $state('');
	let positions = $state<EditableTrade[]>([createEmptyTrade()]);

	function createEmptyTrade(): EditableTrade {
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

	function addPosition() {
		positions = [...positions, createEmptyTrade()];
	}

	function removePosition(index: number) {
		if (positions.length === 1) {
			return;
		}
		positions = positions.filter((_, currentIndex) => currentIndex !== index);
	}

	async function handleSubmit(event: Event) {
		event.preventDefault();
		saving = true;
		error = '';
		try {
			const detail = await createOptionTrackPlan({
				name: planName,
				positions
			});
			await goto(`/optiontrack/${detail.plan.id}`);
		} catch (err: any) {
			error = err.message;
		} finally {
			saving = false;
		}
	}
</script>

<div class="space-y-8">
	<section class="hero">
		<div>
			<p class="eyebrow">S01</p>
			<h1>创建交易计划</h1>
			<p class="desc">录入计划名称与初始建仓合约。金额自动按价格 × 张数 × 合约乘数计算。</p>
		</div>
		<a class="ghost" href="/optiontrack">返回列表</a>
	</section>

	{#if error}
		<div class="alert">{error}</div>
	{/if}

	<form class="stack" onsubmit={handleSubmit}>
		<section class="panel">
			<h2>计划信息</h2>
			<label>
				<span>计划名称</span>
				<input bind:value={planName} placeholder="例如: 6月认购波段计划" required />
			</label>
		</section>

		<section class="panel">
			<div class="panel-head">
				<h2>初始建仓</h2>
				<button class="secondary" type="button" onclick={addPosition}>新增合约</button>
			</div>

			<div class="cards">
				{#each positions as position, index}
					<div class="card">
						<div class="card-head">
							<h3>合约 {index + 1}</h3>
							<button type="button" class="danger-link" onclick={() => removePosition(index)}>
								删除
							</button>
						</div>
						<div class="grid">
							<label>
								<span>品种类型</span>
								<select bind:value={position.instrument_type}>
									<option value="OPTION">期权 / OPTION</option>
									<option value="FUTURE">期货 / FUTURE</option>
								</select>
							</label>
							<label><span>合约名称</span><input bind:value={position.contract_name} required /></label>
							<label><span>合约代码</span><input bind:value={position.contract_code} /></label>
							<label>
								<span>{isOptionTrackOption(position.instrument_type) ? '到期日' : '交割日 / 合约月'}</span>
								<input
									type="date"
									bind:value={position.expiry_date}
									required={isOptionTrackOption(position.instrument_type)}
								/>
							</label>
							{#if isOptionTrackOption(position.instrument_type)}
								<label><span>执行价</span><input type="number" min="0" step="0.01" bind:value={position.strike_price} required /></label>
								<label>
									<span>期权方向</span>
									<select bind:value={position.option_side}>
										<option value="CALL">认购 / CALL</option>
										<option value="PUT">认沽 / PUT</option>
									</select>
								</label>
							{/if}
							<label>
								<span>交易类型</span>
								<select bind:value={position.trade_type}>
									<option value="BUY_OPEN">买开</option>
									<option value="SELL_OPEN">卖开</option>
								</select>
							</label>
							<label><span>合约乘数</span><input type="number" min="1" bind:value={position.contract_multiplier} required /></label>
							<label><span>建仓时间</span><input type="datetime-local" bind:value={position.trade_time} required /></label>
							<label><span>{isOptionTrackOption(position.instrument_type) ? '成本' : '期货价格'}</span><input type="number" min="0" step="0.0001" bind:value={position.price} required /></label>
							<label><span>张数</span><input type="number" min="1" bind:value={position.quantity} required /></label>
							<label class="full"><span>备注</span><input bind:value={position.remark} /></label>
						</div>
						<div class="amount">金额: {calcTradeAmount(position.price, position.quantity, position.contract_multiplier).toFixed(2)}</div>
					</div>
				{/each}
			</div>
		</section>

		<div class="actions">
			<a class="ghost" href="/optiontrack">取消</a>
			<button class="primary" type="submit" disabled={saving}>{saving ? '保存中...' : '保存计划'}</button>
		</div>
	</form>
</div>

<style lang="postcss">
	@reference "tailwindcss";

	:global(html) {
		background:
			radial-gradient(circle at top right, rgba(251, 191, 36, 0.18), transparent 24%),
			linear-gradient(180deg, #fffaf0 0%, #f8fafc 58%, #eef2ff 100%);
	}

	.hero {
		@apply flex flex-col gap-4 rounded-3xl border border-amber-200/70 bg-white/85 p-8 shadow-xl shadow-amber-100/60 md:flex-row md:items-end md:justify-between;
	}

	.eyebrow {
		@apply text-xs font-semibold uppercase tracking-[0.3em] text-amber-700;
	}

	h1 {
		@apply mt-2 text-4xl font-black tracking-tight text-slate-900;
	}

	.desc {
		@apply mt-3 max-w-2xl text-sm leading-6 text-slate-600;
	}

	.stack {
		@apply space-y-6;
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

	.cards {
		@apply space-y-4;
	}

	.card {
		@apply rounded-2xl border border-slate-200 bg-slate-50/80 p-5;
	}

	.card-head {
		@apply mb-4 flex items-center justify-between;
	}

	.card-head h3 {
		@apply text-base font-semibold text-slate-900;
	}

	.grid {
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
		@apply rounded-xl border-slate-200 bg-white px-4 py-3 text-sm text-slate-900 shadow-sm focus:border-amber-400 focus:ring-amber-400;
	}

	.full {
		@apply md:col-span-2 xl:col-span-4;
	}

	.amount {
		@apply mt-4 rounded-2xl bg-amber-50 px-4 py-3 text-sm font-semibold text-amber-700;
	}

	.actions {
		@apply flex items-center justify-end gap-3;
	}

	.primary,
	.secondary,
	.ghost {
		@apply inline-flex items-center justify-center rounded-2xl px-5 py-3 text-sm font-semibold transition;
	}

	.primary {
		@apply bg-amber-500 text-white shadow-lg shadow-amber-200 hover:bg-amber-600 disabled:cursor-not-allowed disabled:opacity-50;
	}

	.secondary {
		@apply border border-slate-200 bg-white text-slate-700 hover:bg-slate-50;
	}

	.ghost {
		@apply border border-slate-200 bg-white text-slate-700 hover:bg-slate-50;
	}

	.danger-link {
		@apply text-sm font-semibold text-rose-600 hover:text-rose-800;
	}

	.alert {
		@apply rounded-2xl border border-rose-200 bg-rose-50 px-4 py-3 text-sm text-rose-700;
	}
</style>