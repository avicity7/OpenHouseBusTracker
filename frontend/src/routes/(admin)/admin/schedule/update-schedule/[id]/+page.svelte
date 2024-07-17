<script lang="ts">
	import { onMount } from 'svelte';
	import type { Route, Driver, EventBus } from '$lib/types/global';
	import CustomDropdown from '$lib/components/CustomDropdown.svelte';

	export let data;

	let { drivers, buses, routes, schedule } = data;
	let selectedBus: EventBus;
	let selectedRoute: Route;
	let selectedDriverId: number;
	let selectedStartTime = '';
	let selectedEndTime = '';
	let selectedDriver: Driver;
	let loaded = false;

	onMount(() => {
		if (schedule) {
			const { BusId, RouteName, DriverId, StartTime, EndTime } = schedule;
			selectedBus = buses.find((bus) => bus.BusId == BusId)!;
			selectedRoute = routes.find((route) => route.RouteName == RouteName)!;
			selectedDriverId = DriverId;
			selectedStartTime = StartTime.split('+')[0];
			selectedEndTime = EndTime.split('+')[0];
			selectedDriver = drivers.find((driver) => selectedDriverId == driver.DriverId)!;
			loaded = true;
		}
	});
</script>

{#if loaded}
	<div class="flex justify-center items-center h-full">
		<div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-24">
			<h1 class="text-2xl font-semibold mb-4">Update Select Bus Schedule</h1>
			<form method="POST" action="?/updateBusSchedule">
				<div class="mb-4">
					<CustomDropdown label="Bus" name="bus" options={buses} bind:selected={selectedBus} />
				</div>
				<div class="mb-4">
					<CustomDropdown
						label="Route"
						name="route"
						options={routes}
						bind:selected={selectedRoute}
					/>
				</div>

				{#if selectedDriver}
					<div class="mb-4">
						<CustomDropdown
							label="Driver"
							name="driver"
							options={drivers}
							bind:selected={selectedDriver}
						/>
					</div>
				{/if}
				<div class="mb-4">
					<label for="startTime" class="block text-sm font-medium mb-1">Start Time:</label>
					<input
						type="datetime-local"
						id="start_time"
						name="start_time"
						required
						bind:value={selectedStartTime}
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
						bind:value={selectedEndTime}
						class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500"
					/>
				</div>

				<div class="mt-4 flex justify-center">
					<button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800"
						>Update Schedule</button
					>
				</div>
			</form>
		</div>
	</div>
{/if}
