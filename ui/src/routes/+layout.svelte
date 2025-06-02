<script lang="ts">
	import '../app.css';
	import { ChartNoAxesCombined, ChartPie, House, Settings } from '@lucide/svelte';

	let { children } = $props();
</script>

<div class="flex h-screen">
	<aside class="sidebar">
		<a href="/" class="home">
			<House color="#ffffff" strokeWidth={1} />
			<span>Home</span>
		</a>

		<nav class="nav">
			<a
				href="/rsrs"
				class="bg-indigo-600 text-white"
			>
				<ChartNoAxesCombined color="#ffffff" strokeWidth={1} />
				<span>ETF RSRS Data</span>
			</a>
			<a
				href="/kalmanfilter"
				class="text-indigo-200 hover:bg-indigo-700 hover:text-white"
			>
				<ChartPie color="#ffffff" strokeWidth={1} />
				<span>Kalmanfilter</span>
			</a>
			<a
				href="#top"
				class="text-indigo-200 hover:bg-indigo-700 hover:text-white"
			>
				<Settings color="#ffffff" strokeWidth={1} />
				<span>Settings</span>
			</a>
		</nav>

		<!-- <div class="mt-auto">
			<a
				href="#"
				class="flex items-center space-x-3 rtl:space-x-reverse text-indigo-200 hover:bg-indigo-700 hover:text-white p-3 rounded-lg transition-colors duration-150"
			>
				<svg
					class="w-6 h-6"
					fill="none"
					stroke="currentColor"
					viewBox="0 0 24 24"
					xmlns="http://www.w3.org/2000/svg"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
					></path></svg
				>
				<span>Logout</span>
			</a>
		</div> -->
	</aside>
	{@render children()}
</div>

<style lang="postcss">
	@reference "tailwindcss";
	.sidebar {
		@apply fixed inset-y-0 left-0 z-30 w-64 bg-indigo-800 text-white p-6 space-y-6 transform -translate-x-full md:translate-x-0 md:relative md:inset-0 md:flex md:flex-col shadow-lg;
	}

	.home {
		@apply flex items-center space-x-3 rtl:space-x-reverse px-2 pb-4 border-b border-indigo-700;
	}
	.home span {
		@apply text-2xl font-semibold text-white;
	}

	.nav {
		@apply flex-1 space-y-1;
	}
	.nav a {
		@apply flex items-center space-x-3 rtl:space-x-reverse p-3 rounded-lg transition-colors duration-150;
	}
</style>
