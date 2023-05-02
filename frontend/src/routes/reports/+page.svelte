<script>
	import api from '$lib/api.js';
	import { onMount } from 'svelte';

	/**
	 * @typedef {Object} report
	 * @property {number} id
	 * @property {string} name
	 * @property {string} email
	 */

	/**
	 * @type {Array<report>}
	 */
	let reports = [];

	onMount(() => {
		api
			.get('/api/reports')
			.then((res) => {
				reports = res.data;
			})
			.catch((err) => {
				//
			});
	});
</script>

<div class="flex flex-col justify-center items-center min-h-screen">
	<h1 class="text-3xl font-bold text-white mb-4">Pending reports</h1>
	<ul class="w-full max-w-md bg-gray-800 shadow-lg rounded-lg px-8 pt-6 pb-8 mb-4">
		{#each reports as report}
			<li class="border-b border-gray-700 py-2">
				<a href="/readreport?id={report.id}" class="text-white font-bold hover:underline"
					>{report.name}</a
				>
				<p class="text-gray-400">{report.email}</p>
			</li>
		{/each}
	</ul>
</div>
