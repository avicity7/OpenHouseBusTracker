<script lang="ts">
    import { onMount } from "svelte";
    import { writable } from 'svelte/store';
    import type { Schedule, Driver } from "../../../../../types/global";

    export let data: {
        dropdownData: { data: { carplate: string, route_name: string, driver: { driver_id: number, driver_name: string }[] }[] };
        scheduleData: { schedule: Schedule };
    };
    
    let { dropdownData, scheduleData } = data;

    const carplates = writable<string[]>([]);
    const routeNames = writable<string[]>([]);
    const drivers = writable<Driver[]>([]);
    const selectedCarplate = writable<string|null>(null);
    const selectedRouteName = writable<string|null>(null);
    const selectedDriverId = writable<number|null>(null);
    let selectedStartTime = "";
    let selectedEndTime = "";

    function setDropdownOptions(data: { carplate: string, route_name: string, driver: { driver_id: number, driver_name: string }[] }[]) {
        const uniqueCarplates = new Set<string>();
        const uniqueRouteNames = new Set<string>();
        const uniqueDrivers = new Map<number, string>();

        data.forEach(({ carplate, route_name, driver }) => {
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
        setDropdownOptions(dropdownData.data);

        if (scheduleData && scheduleData.schedule) {
            const { Carplate, RouteName, DriverId, StartTime, EndTime } = scheduleData.schedule;
            selectedCarplate.set(Carplate);
            selectedRouteName.set(RouteName);
            selectedDriverId.set(DriverId);
            selectedStartTime = new Date(StartTime).toISOString().slice(0, 16);
            selectedEndTime = new Date(EndTime).toISOString().slice(0, 16);
        }
    });

</script>

<div class="p-6 md:p-12">
    <h1 class="text-3xl font-semibold mb-4">Update Select Bus Schedule</h1>
    <form method="POST" action="?/updateBusSchedule">
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
        <input type="datetime-local" id="start_time" name="start_time" required bind:value={selectedStartTime}/>

        <label for="endTime">End Time:</label>
        <input type="datetime-local" id="end_time" name="end_time" required bind:value={selectedEndTime}/>

        <div class="mt-4 flex justify-end">
            <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-md">Update Schedule</button>
        </div>
    </form>
</div>
