<script lang="ts">
	import type { RSRS } from '$lib/etf';
	import { getFormattedCurrentDate, type ErrMsg } from '$lib';
	let period = $state('25');
	let date = $state(getFormattedCurrentDate());
	let rsrsTable = $state<RSRS[]>([]);
	let errorMessage = $state<ErrMsg>({
		code: 0,
		message: '',
		hidden: true
	});
	const fetchData = async (period: string, date: string) => {
		try {
			date = date.replace(/-/g, ''); // Convert date to YYYYMMDD format
			const response = await fetch(
				`http://localhost:8080/etf/rsrslist?period=${period}&date=${date}`
			);

			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}

			rsrsTable = (await response.json()) as RSRS[];
		} catch (error: any) {
			console.error('Error fetching data:', error);
			errorMessage.code = error.code || 500;
			errorMessage.message = `Error: ${error.message}`;
			errorMessage.hidden = false;
		} finally {
		}
	};
	function handleFormSubmit(event: Event) {
		event.preventDefault();
		// const form = event.target as HTMLFormElement;
		// const period = (form.period as HTMLInputElement).value;
		// const date = (form.date as HTMLInputElement).value;

		// fetchData(period, date);
		console.log(`Fetching data for period: ${period}, date: ${date}`);

		fetchData(period, date);
	}

	const tableHeaders = [
		{ label: 'TS Code', key: 'tsCode' },
		{ label: 'Name', key: 'name' },
		{ label: 'RSRS', key: 'rsrs' }
	];
    fetchData(period, date);
</script>

<div class="max-w-4xl mx-auto">
	<h1 class="text-3xl font-bold text-center text-indigo-700 mb-8">ETF Stratedy Using RSRS</h1>

	<form id="rsrsForm" class="bg-white p-6 rounded-lg shadow-lg mb-8" onsubmit={handleFormSubmit}>
		<div class="grid grid-cols-1 md:grid-cols-3 gap-6">
			<div>
				<label for="period" class="block text-sm font-medium text-gray-700 mb-1">Period:</label>
				<input
					type="number"
					id="period"
					name="period"
					bind:value={period}
					required
					class="mt-1 block w-full p-3 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
				/>
			</div>
			<div>
				<label for="date" class="block text-sm font-medium text-gray-700 mb-1"
					>Date (YYYYMMDD):</label
				>
				<input
					id="date"
					name="date"
					type="date"
					placeholder="YYYY-MM-DD"
					bind:value={date}
					required
					class="mt-1 block w-full p-3 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
				/>
			</div>
			<div class="md:mt-7">
				<button
					type="submit"
					class="w-full bg-gradient-to-r from-slate-300 to-slate-500 text-white py-3 px-4 rounded-md shadow-sm text-sm font-medium"
				>
					Fetch Data
				</button>
			</div>
		</div>
	</form>

	<div class="bg-white p-6 rounded-lg shadow-lg">
		<h2 class="text-xl font-semibold text-gray-800 mb-4">Results</h2>
		{#if rsrsTable.length > 0}
			<div class="overflow-x-auto bg-white rounded-lg shadow-lg">
				<table class="table-auto min-w-full divide-y divide-gray-300">
					<thead class="bg-gradient-to-r from-slate-300 to-slate-500">
						<tr>
							{#each tableHeaders as theader}
								<th scope="col" class="th">{theader.label}</th>
							{/each}</tr
						>
					</thead>
					<tbody class="divide-y divide-gray-200">
						{#each rsrsTable as row}
							<tr class="hover:bg-gray-50 transition-colors duration-150">
								<td class="td">{row.ts_code}</td>
								<td class="td">{row.name}</td>
								<td class="td">{row.rsrs}</td>
							</tr>
						{/each}
					</tbody>
				</table>
			</div>
		{:else}
			<div class="text-gray-500 text-center mb-4 py-4">
				No data available. Please enter parameters and click "Fetch Data".
			</div>
		{/if}
	</div>
</div>

<style lang="postcss">
	@reference "tailwindcss";
	:global(html) {
		background-color: theme(--color-gray-100);
	}

	.th {
		@apply px-6 py-4 text-left text-xs sm:text-sm font-semibold text-white uppercase tracking-wider;
	}

	.td {
		@apply px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900;
	}
</style>
