<script lang="ts">
	import { onMount } from 'svelte';
	import type { Schedule, Driver } from '$lib/types/global.js';
	import { page } from '$app/stores';
	import ErrorMessage from '$lib/components/ErrorMessage.svelte';
	import CustomDropdown from '$lib/components/CustomDropdown.svelte';

	export let data
	const { dropdownData } = data

	let carplates: Array<string> = []
	let routeNames: Array<string> = []
	let drivers: Array<Driver> = []
	let selectedCarplate: string | null = null;
	let selectedRouteName: string | null = null;
	let selectedDriverId: number | null = null;
	let errorMessage: string | null = null;

	const setDropdownOptions = () => {
        if (!dropdownData) return;

        const uniqueCarplates = new Set<string>();
        const uniqueRouteNames = new Set<string>();
        const validDrivers: Map<number, string> = new Map();

        dropdownData.forEach(({ Carplate, RouteName, Driver }: Schedule) => {
            if (Carplate) {
                uniqueCarplates.add(Carplate);
            }
            if (RouteName) {
                uniqueRouteNames.add(RouteName);
            }
            if (Driver && Array.isArray(Driver)) {
                Driver.forEach(({ DriverId, DriverName }) => {
                    if (DriverId !== null && DriverName !== null) {
                        validDrivers.set(DriverId, DriverName);
                    }
                });
            }
        });

        carplates = Array.from(uniqueCarplates);
        routeNames = Array.from(uniqueRouteNames);
        drivers = Array.from(validDrivers.entries()).map(([DriverId, DriverName]) => ({
            DriverId,
            DriverName,
        }));
    };

	if ($page.status === 409) {
      errorMessage = $page.error?.message || 'A schedule with the same carplate or driver already exists.';
    }

	// does $: do anything? reactivity is built-in no?

	onMount(() => {
		setDropdownOptions();
	});
</script>

<div class="flex justify-center items-center h-full">
	<div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-20">
		<h1 class="text-2xl font-semibold mb-4">Add New Bus Schedule</h1>
		<ErrorMessage message={errorMessage} />
		<form method="POST" action="?/createBusSchedule">
			
			<!-- <div class="mb-4">
				<label for="carplate" class="block text-sm font-medium mb-1">Carplate:</label>
				<div class="relative">
				  <select
					id="carplate"
					name="carplate"
					bind:value={selectedCarplate}
					class="w-full appearance-none rounded-md border border-gray-300 bg-white text-gray-700 px-3 py-2 shadow-sm focus:outline-none focus:ring-1 focus:ring-red-500"
				  >
					{#if carplates.length === 0}
					  <option value="" disabled>No options available</option>
					{:else}
					  {#each carplates as carplate}
						<option value={carplate} class="bg-light-blue-50 hover:bg-red-100">{carplate}</option>
					  {/each}
					{/if}
				  </select>
				  <div class="absolute inset-y-0 right-0 flex items-center px-2 pointer-events-none">
					<svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
					</svg>
				  </div>
				</div>
			</div>

			<div class="mb-4">
				<label for="route_name" class="block text-sm font-medium mb-1">Route Name:</label>
				<select
					id="route_name"
					name="route_name"
					bind:value={selectedRouteName}
					class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500"
				>
					{#each routeNames as routeName}
						<option value={routeName}>{routeName}</option>
					{/each}
				</select>
			</div> -->

			<div class="mb-4">
				<CustomDropdown
				  label="Carplate"
				  name="carplate"
				  options={carplates}
				  required
				  bind:selected={selectedCarplate}
				/>
			  </div>
			  
			  <div class="mb-4">
				<CustomDropdown
				  label="Route Name"
				  name="route_name"
				  options={routeNames}
				  bind:selected={selectedRouteName}
				/>
			  </div>

			  <div class="mb-4">
				<CustomDropdown
					label="Driver"
					name="driver_id"
					options={drivers}
					bind:selected={selectedDriverId}
					displayOption={(driver) => driver.DriverName}
				/>
			  </div>

				<!-- <div class="mb-4">
					<label for="driver_id" class="block text-sm font-medium mb-1">Driver:</label>
					<select
						id="driver_id"
						name="driver_id"
						bind:value={selectedDriverId}
						class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500"
					>
					{#if drivers === null || drivers.length === 0}
						<option value="" disabled>No options available</option>
					{:else}
						{#each drivers as driver}
							<option value={driver.DriverId}>{driver.DriverName}</option>
						{/each}
					{/if}
					</select>
				</div> -->

			<div class="mb-4">
				<label for="startTime" class="block text-sm font-medium mb-1">Start Time:</label>
				<input
					type="datetime-local"
					id="start_time"
					name="start_time"
					required
					class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500"
				/>
			</div>

			<div class="mb-4">
				<label for="endTime" class="block text-sm font-medium mb-1">End Time:</label>
				<input
					type="datetime-local"
					id="end_time"
					name="end_time"
					required
					class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500"
				/>
			</div>

			<div class="mt-4 flex justify-center">
				<button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800"
					>Add Schedule</button
				>
			</div>
		</form>
	</div>
</div>
