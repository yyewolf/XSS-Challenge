<script>
	import api from '$lib/api.js';
	import { onMount } from 'svelte';

	/**
	 * @typedef {Object} report
	 * @property {number} id
	 * @property {string} name
	 * @property {string} email
     * @property {string} issue
	 */

	/**
	 * @type {report}
	 */
	let r = {
        id: -1,
        name: '',
        email: '',
        issue: '',
    };

    let notfound = false;


	onMount(() => {
		api
			.get('/api/report/' + window.location.search.split('=')[1])
			.then((res) => {
				r = res.data;
			})
			.catch((err) => {
				notfound = true;
			});
	});

    function deleteReport() {
        api
            .delete('/api/report/' + window.location.search.split('=')[1])
            .then((res) => {
                window.location.href = '/reports';
            })
            .catch((err) => {
                //
            });
    }
</script>

<div class="flex flex-col justify-center items-center min-h-screen ">
    {#if r.id == -1}
        {#if notfound}
            <h1 class="text-3xl font-bold text-white mb-4">Report not found</h1>
        {:else}
            <h1 class="text-3xl font-bold text-white mb-4">Loading...</h1>
        {/if}
    {:else}
	<h1 class="text-3xl font-bold text-white mb-4">Report #{r.id} Details</h1>
	<div class="w-full max-w-md bg-gray-800 shadow-lg rounded-lg px-8 pt-6 pb-8 mb-4">
		<h2 class="text-lg font-bold text-white mb-4">{r.name}</h2>
		<p class="text-gray-400 mb-4">{r.email}</p>
		<p class="text-white mb-4">
            <!-- UNSAFE THING HERE : -->
            {@html r.issue}
		</p>
		<button
			class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
			id="qsHJGDJQHSGDJHQGSDJFQSGSD"
            on:click={deleteReport}
		>
			Delete Report
		</button>
	</div>
    {/if}
</div>
