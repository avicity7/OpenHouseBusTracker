<script lang="ts">
	export let data;
	const { backend_uri } = data;

	let carplate: string;

	const createBus = async () => {
		await fetch(`${backend_uri}:3000/bus/create-bus/${carplate}`, {
			method: 'PUT'
		});
		location.replace('/admin/buses');
	};
</script>

<div class="flex justify-center items-center h-full">
	<div class="bg-white shadow rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-20">
		<h1 class="text-2xl font-semibold mb-4">Add New Bus</h1>
		<form on:submit|preventDefault={createBus}>
			<div class="mb-4">
				<label for="startTime" class="block text-sm font-medium mb-1">Carplate</label>
				<input
					type="text"
					required
					bind:value={carplate}
					class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500" data-testid="bus-name-input"
				/>
			</div>

			<div class="mt-4 flex justify-center">
				<button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800" data-testid="submit-add-bus">
					Add Bus
				</button>
			</div>
		</form>
	</div>
</div>
