<script>
	import api from '$lib/api.js';

	let report = {
		name: '',
		email: '',
		issue: ''
	};

	let alert = {
		type: '',
		message: ''
	};
    let alertStop = -1;

	function submitReport() {
		alert.type = '';
		alert.message = '';
        if (alertStop >= 0) clearTimeout(alertStop);

		setTimeout(() => {
			api
				.post('/api/report', report)
				.then((res) => {
					alert.type = 'success';
					alert.message = 'Your report has been submitted successfully!';
				})
				.catch((err) => {
					alert.type = 'error';
					alert.message = 'There was an error submitting your report. Please try again later.';
				})
				.finally(() => {
					alertStop = setTimeout(() => {
						alert.type = '';
						alert.message = '';
					}, 5000);
				});
		}, 200);
	}
</script>

<div class="flex flex-col justify-center items-center min-h-screen">
	<h1 class="text-3xl font-bold text-white mb-4">Report an Issue</h1>
	{#if alert.type === 'success'}
		<div class="bg-green-500 text-white font-bold rounded-lg border shadow-lg p-4 mb-4">
			{alert.message}
		</div>
	{:else if alert.type === 'error'}
		<div class="bg-red-500 text-white font-bold rounded-lg border shadow-lg p-4 mb-4">
			{alert.message}
		</div>
	{/if}
	<form class="w-full max-w-md bg-slate-700 shadow-lg rounded-lg px-8 pt-6 pb-8 mb-4">
		<div class="mb-4">
			<label class="block text-white font-bold mb-2" for="name"> Name </label>
			<input
				class="appearance-none bg-gray-700 border rounded w-full py-2 px-3 text-white leading-tight focus:outline-none focus:shadow-outline"
				id="name"
				type="text"
				bind:value={report.name}
				placeholder="Enter your name"
			/>
		</div>
		<div class="mb-4">
			<label class="block text-white font-bold mb-2" for="email"> Email </label>
			<input
				class="appearance-none bg-gray-700 border rounded w-full py-2 px-3 text-white leading-tight focus:outline-none focus:shadow-outline"
				id="email"
				type="email"
				bind:value={report.email}
				placeholder="Enter your email"
			/>
		</div>
		<div class="mb-4">
			<label class="block text-white font-bold mb-2" for="issue"> Issue </label>
			<textarea
				class="appearance-none bg-gray-700 border rounded w-full py-2 px-3 text-white leading-tight focus:outline-none focus:shadow-outline"
				id="issue"
				bind:value={report.issue}
				rows="6"
				placeholder="Describe the issue you encountered"
			/>
		</div>
		<div class="flex items-center justify-end">
			<button
				class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
				type="submit"
				on:click|preventDefault={submitReport}
			>
				Submit Report
			</button>
		</div>
	</form>
</div>
