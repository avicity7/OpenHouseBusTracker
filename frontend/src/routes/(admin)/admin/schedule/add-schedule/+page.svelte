<script lang="ts">
	import { onMount } from 'svelte';
	import type { Schedule, Driver, EventBus } from '$lib/types/global.js';
	import { page } from '$app/stores';
	import ErrorMessage from '$lib/components/ErrorMessage.svelte';
	import CustomDropdown from '$lib/components/CustomDropdown.svelte';

	export let data
	const { dropdownData } = data

	let buses: Array<EventBus> = []
	let routeNames: Array<string> = []
	let drivers: Array<Driver> = []
	let selectedBus: EventBus
	let selectedRouteName: string | null = null;
	let selectedDriverId: number | null = null;
	let errorMessage: string | null = null;

	const setDropdownOptions = () => {
        if (!dropdownData) return;

        const uniqueBusId = new Set<string>();
        const uniqueRouteNames = new Set<string>();
        const validDrivers: Map<number, string> = new Map();

        dropdownData.forEach(({ BusId, Carplate, RouteName, Driver }: Schedule) => {
			let bus = {BusId: BusId, Carplate: Carplate, Status: false, Hidden: false}
            if (Carplate && BusId) {	
				let ol = uniqueBusId.size
                uniqueBusId.add(BusId)
				if (uniqueBusId.size != ol) buses.push(bus)
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

        routeNames = Array.from(uniqueRouteNames);
        drivers = Array.from(validDrivers.entries()).map(([DriverId, DriverName]) => ({
            DriverId,
            DriverName,
        }));
    };

	$: selectedBus, () => {
		console.log('Selected bus:', selectedBus);
	};
	 	
	if ($page.status === 409) {
      errorMessage = $page.error?.message || 'A schedule with the same carplate or driver already exists.';
    }

	onMount(() => {
		console.log(dropdownData)
		setDropdownOptions();
	});
</script>

<div class="flex justify-center items-center h-full">
	<div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-20">
		<h1 class="text-2xl font-semibold mb-4">Add New Bus Schedule</h1>
		<ErrorMessage message={errorMessage} />
		<form method="POST" action="?/createBusSchedule">
			<div class="mb-4 carplate">
				<CustomDropdown
				  label="Bus"
				  name="bus"
				  options={buses}
				  required
				  bind:selected={selectedBus}
				/>
			  </div>
			  
			  <div class="mb-4 route-name">
				<CustomDropdown
				  label="Route Name"
				  name="route_name"
				  options={routeNames}
				  bind:selected={selectedRouteName}
				/>
			  </div>

			  <div class="mb-4 driver">
				<CustomDropdown
					label="Driver"
					name="driver_id"
					options={drivers}
					bind:selected={selectedDriverId}
				/>
			  </div>

			<div class="mb-4">
				<label for="startTime" class="block text-sm font-medium mb-1">Start Time:</label>
				<input
					type="datetime-local"
					id="start_time"
					name="start_time"
					required
					class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500"
					data-testid="start-time"
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
					data-testid="end-time"
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
