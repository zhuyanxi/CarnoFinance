<script lang="ts">
	import { onMount } from 'svelte';
	import {
		formatOptionTrackDateTime,
		listOptionTrackPlans,
		type OptionTrackPlanListResponse
	} from '$lib/optiontrack';

	let loading = $state(true);
	let error = $state('');
	let page = $state(1);
	let pageSize = 10;
	let data = $state<OptionTrackPlanListResponse>({
		items: [],
		page: 1,
		page_size: pageSize,
		total: 0,
		total_page: 0
	});

	async function loadPlans(targetPage = page) {
		loading = true;
		error = '';
		try {
			data = await listOptionTrackPlans(targetPage, pageSize);
			page = data.page;
		} catch (err: any) {
			error = err.message;
		} finally {
			loading = false;
		}
	}

	onMount(() => {
		loadPlans();
	});
</script>

<div class="space-y-8">
	<section class="hero">
		<div>
			<p class="eyebrow">OptionTrack</p>
			<h1>期权交易计划</h1>
			<p class="desc">管理计划、记录开平仓、关闭计划、查看盈亏。当前列表按创建时间倒序。</p>
		</div>
		<a class="primary" href="/optiontrack/new">新建计划</a>
	</section>

	{#if error}
		<div class="alert">{error}</div>
	{/if}

	<section class="panel">
		<div class="panel-head">
			<h2>计划列表</h2>
			<span>{data.total} 个计划</span>
		</div>

		{#if loading}
			<div class="empty">加载中...</div>
		{:else if data.items.length === 0}
			<div class="empty">暂无计划。先创建第一个交易计划。</div>
		{:else}
			<div class="table-wrap">
				<table>
					<thead>
						<tr>
							<th>计划</th>
							<th>状态</th>
							<th>合约数</th>
							<th>创建时间</th>
							<th>最新交易</th>
							<th></th>
						</tr>
					</thead>
					<tbody>
						{#each data.items as item}
							<tr>
								<td>
									<div class="name">{item.name}</div>
									<div class="sub">ID {item.id}</div>
								</td>
								<td>
									<span class={`status ${item.status === 'CLOSED' ? 'closed' : 'open'}`}>{item.status}</span>
								</td>
								<td>{item.contract_count}</td>
								<td>{formatOptionTrackDateTime(item.created_at)}</td>
								<td>{formatOptionTrackDateTime(item.latest_trade_at)}</td>
								<td><a class="link" href={`/optiontrack/${item.id}`}>查看</a></td>
							</tr>
						{/each}
					</tbody>
				</table>
			</div>
		{/if}

		<div class="pager">
			<button disabled={page <= 1 || loading} onclick={() => loadPlans(page - 1)}>上一页</button>
			<span>第 {data.page} / {Math.max(data.total_page, 1)} 页</span>
			<button disabled={loading || data.page >= data.total_page} onclick={() => loadPlans(page + 1)}>
				下一页
			</button>
		</div>
	</section>
</div>

<style lang="postcss">
	@reference "tailwindcss";

	:global(html) {
		background:
			radial-gradient(circle at top left, rgba(245, 158, 11, 0.2), transparent 28%),
			linear-gradient(180deg, #fff8ed 0%, #f8fafc 52%, #eef2ff 100%);
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

	.primary {
		@apply inline-flex items-center justify-center rounded-2xl bg-amber-500 px-5 py-3 text-sm font-semibold text-white shadow-lg shadow-amber-200 transition hover:bg-amber-600;
	}

	.panel {
		@apply rounded-3xl border border-slate-200 bg-white/90 p-6 shadow-lg shadow-slate-200/70;
	}

	.panel-head {
		@apply mb-6 flex items-center justify-between gap-4;
	}

	.panel-head h2 {
		@apply text-xl font-bold text-slate-900;
	}

	.panel-head span {
		@apply text-sm text-slate-500;
	}

	.alert {
		@apply rounded-2xl border border-rose-200 bg-rose-50 px-4 py-3 text-sm text-rose-700;
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

	.status {
		@apply inline-flex rounded-full px-3 py-1 text-xs font-semibold;
	}

	.open {
		@apply bg-emerald-100 text-emerald-700;
	}

	.closed {
		@apply bg-slate-200 text-slate-700;
	}

	.link {
		@apply text-sm font-semibold text-amber-700 hover:text-amber-900;
	}

	.empty {
		@apply rounded-2xl border border-dashed border-slate-200 bg-slate-50 px-4 py-10 text-center text-sm text-slate-500;
	}

	.pager {
		@apply mt-6 flex items-center justify-between gap-4;
	}

	.pager button {
		@apply rounded-xl border border-slate-200 px-4 py-2 text-sm font-medium text-slate-700 disabled:cursor-not-allowed disabled:opacity-50;
	}

	.pager span {
		@apply text-sm text-slate-500;
	}
</style>