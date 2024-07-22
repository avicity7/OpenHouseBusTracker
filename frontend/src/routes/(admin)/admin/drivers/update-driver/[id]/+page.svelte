<script lang="ts">
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import type { Driver } from '$lib/types/global';
	import { goto } from '$app/navigation';

	let driver: Driver = { DriverId: 0, DriverName: '' };

	async function updateDriver() {
		const response = await fetch(
			`${PUBLIC_BACKEND_URL}:3000/driver/update-driver`,
			{
				method: 'PUT',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(driver)
			}
		);

		if (response.ok) {
			goto('/admin/drivers');
		}
	}

	onMount(() => {
		const { id } = $page.params;
		driver = JSON.parse(decodeURIComponent(id));
	});
</script>

<div class="flex justify-center items-center h-full">
	<div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-20">
		<h1 class="text-2xl font-semibold mb-4">Update Driver</h1>
		<form on:submit={updateDriver}>
			<div class="mb-4">
				<label for="driverName" class="block text-sm font-medium mb-1">Driver Name:</label>
				<input
					id="driverName"
					name="driverName"
					bind:value={driver.DriverName}
					class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500"
					required
				/>
			</div>

			<div class="mt-4 flex justify-center">
				<button type="submit" class="px-4 py-2 bg-red-700 text-white rounded-md hover:bg-red-800">
					Update Driver
				</button>
			</div>
		</form>
	</div>
</div>
