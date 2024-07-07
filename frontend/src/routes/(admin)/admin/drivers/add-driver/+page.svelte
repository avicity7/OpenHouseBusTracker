<script lang="ts">
	import { PUBLIC_BACKEND_URL } from "$env/static/public";
    import { goto } from '$app/navigation';

    let newDriverName = '';

    async function addDriver(event: { preventDefault: () => void; }) {
		event.preventDefault();

		const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/driver/add-driver`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ DriverName: newDriverName })
		});

		if (response.ok) {
			goto('/admin/drivers');
		} else {
			console.log(`Failed to add ${newDriverName}`)
		}
	}

</script>

<div class="flex justify-center items-center h-full">
	<div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-20">
		<h1 class="text-2xl font-semibold mb-4">Add New Driver</h1>
		<form on:submit={addDriver}>
			<div class="mb-4">
				<label for="driverName" class="block text-sm font-medium mb-1">Driver Name:</label>
				<input
					id="driverName"
					name="driverName"
					bind:value={newDriverName}
					class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500"
                    required
                />
			</div>

			<div class="mt-4 flex justify-center">
				<button type="submit" class="px-4 py-2 bg-red-700 text-white rounded-md hover:bg-red-800">
                    Add Driver
                </button>
			</div>
		</form>
	</div>
</div>
