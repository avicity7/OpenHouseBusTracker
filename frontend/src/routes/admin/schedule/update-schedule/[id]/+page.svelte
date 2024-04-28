<script lang="ts">
    import { onMount } from "svelte";
    import { writable } from 'svelte/store';
    import type { Schedule, Driver } from "../../../../../types/global";

    export let dropdownData: { data: Schedule[] };
    export let scheduleData: { schedule: Schedule };

    const carplates = writable<string[]>([]);
    const routeNames = writable<string[]>([]);
    const drivers = writable<Driver[]>([]);
    const selectedCarplate = writable<string|null>(null);
    const selectedRouteName = writable<string|null>(null);
    const selectedDriverId = writable<number|null>(null);

    function setDropdownOptions(data: any) {

    if (!data || !Array.isArray(dropdownData)) {
        console.log("Data is empty or not an array:", data);
        return;
    }

    const uniqueCarplates = new Set<string>();
    const uniqueRouteNames = new Set<string>();
    const uniqueDrivers = new Map<number, string>();

    
    data.data.forEach(({ carplate, route_name, driver }: { carplate: string, route_name: string, driver: any[] }) => {

        if (carplate) {
            uniqueCarplates.add(carplate);
        }
        if (route_name) {
            uniqueRouteNames.add(route_name);
        }
        if (driver && Array.isArray(driver)) {
            driver.forEach(({ driver_id, driver_name }) => {
                uniqueDrivers.set(driver_id, driver_name);
            });
        }
    });

    carplates.set(Array.from(uniqueCarplates));
    routeNames.set(Array.from(uniqueRouteNames));
    drivers.set(Array.from(uniqueDrivers.entries()).map(([DriverId, DriverName]) => ({ DriverId, DriverName })));
}

    onMount(() => {
        console.log("Dropdown Data:", dropdownData);
        console.log("Schedule Data:", scheduleData);
        // setDropdownOptions(data);
    });

</script>

<div class="p-6 md:p-12">
    <h1 class="text-3xl font-semibold mb-4">Update Select Bus Schedule</h1>
    <form method="POST" action="?/createBusSchedule">
        <label for="carplate">Carplate:</label>
        <select id="carplate" name="carplate" bind:value={$selectedCarplate}>
            {#each $carplates as carplate}
                <option value={carplate}>{carplate}</option>
            {/each}
        </select>

        <label for="route_name">Route Name:</label>
        <select id="route_name" name="route_name" bind:value={$selectedRouteName}>
            {#each $routeNames as routeName}
                <option value={routeName}>{routeName}</option>
            {/each}
        </select>

        <label for="driver_id">Driver:</label>
        <select id="driver_id" name="driver_id" bind:value={$selectedDriverId}>
            {#each $drivers as { DriverId, DriverName }}
                <option value={DriverId}>{DriverName}</option>
            {/each}
        </select>

        <label for="startTime">Start Time:</label>
        <input type="datetime-local" id="start_time" name="start_time" required/>

        <label for="endTime">End Time:</label>
        <input type="datetime-local" id="end_time" name="end_time" required/>

        <div class="mt-4 flex justify-end">
            <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-md">Update Schedule</button>
        </div>
    </form>
</div>
