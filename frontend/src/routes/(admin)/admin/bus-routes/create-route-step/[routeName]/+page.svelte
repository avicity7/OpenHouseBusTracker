<script lang="ts">
	export let data;
	const { backend_uri, stops } = data;
	import { page } from "$app/stores";
	const { routeName } = $page.params

	let selectedStop: string

	const createRoute = async () => {
		const payload = {
			RouteName: routeName,
			StopName: selectedStop,
			Order: 0
		};
		await fetch(`${backend_uri}:3000/route-step/create-route-step`, {
			method: 'POST',
			body: JSON.stringify(payload)
		});
		location.replace('/admin/bus-routes');
	};
</script>

<div class="flex justify-center items-center h-full">
	<div class="bg-white shadow rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-12">
		<h1 class="text-2xl font-semibold mb-4">Add Route Step</h1>
		<form on:submit|preventDefault={createRoute}>
			<div class="mb-4">
				<label for="AddingRoute" class="block text-sm font-medium mb-1">Stop Name</label>
        <select
          bind:value={selectedStop}
          class="rounded-lg border-2 border-stone-400"
        >
          {#each stops as stop}
            <option 
              value={stop}
            >
              {stop}
            </option>
          {/each}
        </select>
			</div>

			<div class="mt-4 flex justify-center">
				<button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800"
					>Add Route Step</button
				>
			</div>
		</form>
	</div>
</div>
