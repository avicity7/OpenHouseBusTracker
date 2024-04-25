<script lang="ts">
    import { onMount } from "svelte";
    import { writable } from 'svelte/store';
    import type { Schedule, Driver } from "../../../../types/global";

    export let data: Schedule[];

    const carplates = writable<string[]>([]);
    const routeNames = writable<string[]>([]);
    const drivers = writable<Driver[]>([]);
    const selectedCarplate = writable<string|null>(null);
    const selectedRouteName = writable<string|null>(null);
    const selectedDriverId = writable<number|null>(null);

    function setDropdownOptions(data: any) {

    if (!data || !Array.isArray(data.data)) {
        console.log("Data is empty or not an array:", data); // Log if data is empty or not an array
        return;
    }

    const uniqueCarplates = new Set<string>();
    const uniqueRouteNames = new Set<string>();
    const uniqueDrivers = new Map<number, string>();

    console.log('unique carplates', uniqueCarplates)
    console.log('data here?', data)
    
    data.data.forEach(({ carplate, route_name, driver }: { carplate: string, route_name: string, driver: any[] }) => {
        console.log("Carplate:", carplate); 
        console.log("RouteName:", route_name); 
        console.log("Driver:", driver); 

        if (carplate) {
            uniqueCarplates.add(carplate);
        }
        if (route_name) {
            uniqueRouteNames.add(route_name);
        }
        if (driver && Array.isArray(driver)) {
            driver.forEach(({ driver_id, driver_name }) => {
                console.log("DriverId:", driver_id);
                console.log("DriverName:", driver_name);
                uniqueDrivers.set(driver_id, driver_name);
            });
        }
    });

    console.log("Unique carplates:", Array.from(uniqueCarplates)); 
    console.log("Unique route names:", Array.from(uniqueRouteNames));
    console.log("Unique drivers:", Array.from(uniqueDrivers.entries())); 

    carplates.set(Array.from(uniqueCarplates));
    routeNames.set(Array.from(uniqueRouteNames));
    drivers.set(Array.from(uniqueDrivers.entries()).map(([DriverId, DriverName]) => ({ DriverId, DriverName })));
}

    onMount(() => {
        console.log("mount is called")
        setDropdownOptions(data);
    });
</script>

<div class="p-6 md:p-12">
    <h1 class="text-3xl font-semibold mb-4">Add New Bus Schedule</h1>
    <form>
        <label for="carplate">Carplate:</label>
        <select id="carplate" bind:value={$selectedCarplate}>
            {#each $carplates as carplate}
                <option value={carplate}>{carplate}</option>
            {/each}
        </select>

        <label for="route_name">Route Name:</label>
        <select id="route_name" bind:value={$selectedRouteName}>
            {#each $routeNames as routeName}
                <option value={routeName}>{routeName}</option>
            {/each}
        </select>

        <label for="driver_id">Driver:</label>
        <select id="driver_id" bind:value={$selectedDriverId}>
            {#each $drivers as { DriverId, DriverName }}
                <option value={DriverId}>{DriverName}</option>
            {/each}
        </select>

        <label for="start_time">Start Time:</label>
        <input type="datetime-local" id="start_time" name="start_time" required/>

        <label for="end_time">End Time:</label>
        <input type="datetime-local" id="end_time" name="end_time" required/>

        <div class="mt-4 flex justify-end">
            <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-md">Add Schedule</button>
        </div>
    </form>
</div>
