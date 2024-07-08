<script lang="ts">
    import { onMount } from "svelte";
    import type { Schedule, Driver } from "$lib/types/global";
	import CustomDropdown from "$lib/components/CustomDropdown.svelte";

    export let data: {
        dropdownData: { data: Schedule[] } | undefined;
        scheduleData: { schedule: Schedule } | undefined;
    };
    
    let { dropdownData, scheduleData } = data;
    
    let carplates: string[] = [];
    let routeNames: string[] = [];
    let drivers: Driver[] = [];
    let selectedCarplate: string | null = null;
    let selectedRouteName: string | null = null;
    let selectedDriverId: number | null = null;
    let selectedStartTime = "";
    let selectedEndTime = "";

    function setDropdownOptions(data: Schedule[]) {
        if (!data) return;
        
        const uniqueCarplates = new Set<string>();
        const uniqueRouteNames = new Set<string>();
        const uniqueDrivers = new Map<number, string>();

        data.forEach(({ Carplate, RouteName, Driver }) => {
            if (Carplate) {
                uniqueCarplates.add(Carplate);
            }
            if (RouteName) {
                uniqueRouteNames.add(RouteName);
            }
            if (Driver && Array.isArray(Driver)) {
                Driver.forEach(({ DriverId, DriverName }) => {
                    uniqueDrivers.set(DriverId, DriverName);
                });
            }
        });

        carplates = Array.from(uniqueCarplates);
        routeNames = Array.from(uniqueRouteNames);
        drivers = Array.from(uniqueDrivers.entries()).map(([DriverId, DriverName]) => ({ DriverId, DriverName }));
    }

    onMount(() => {
        if (data && dropdownData) {
            setDropdownOptions(dropdownData.data);
        }

        if (scheduleData && scheduleData.schedule) {
            const { Carplate, RouteName, DriverId, StartTime, EndTime } = scheduleData.schedule;
            selectedCarplate = Carplate;
            selectedRouteName = RouteName;
            selectedDriverId = DriverId;
            
            selectedStartTime =StartTime.split('+')[0];
        
            selectedEndTime = EndTime.split('+')[0];
        }
    });

</script>

<div class="flex justify-center items-center h-full">
    <div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-24">
        <h1 class="text-2xl font-semibold mb-4">Update Select Bus Schedule</h1>
        <form method="POST" action="?/updateBusSchedule">
            <!-- <div class="mb-4">
                <label for="carplate" class="block text-sm font-medium mb-1">Carplate:</label>
                <select id="carplate" name="carplate" bind:value={selectedCarplate} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500">
                    {#each carplates as carplate}
                        <option value={carplate}>{carplate}</option>
                    {/each}
                </select>
            </div> -->
            <div class="mb-4">
                <CustomDropdown
                  label="Carplate"
                  options={carplates}
                  bind:selected={selectedCarplate}
                />
            </div>

            <!-- <div class="mb-4">
                <label for="route_name" class="block text-sm font-medium mb-1">Route Name:</label>
                <select id="route_name" name="route_name" bind:value={selectedRouteName} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500">
                    {#each routeNames as routeName}
                        <option value={routeName}>{routeName}</option>
                    {/each}
                </select>
            </div> -->
            <div class="mb-4">
                <CustomDropdown
                  label="Route Name"
                  options={routeNames}
                  bind:selected={selectedRouteName}
                />
            </div>

            <!-- same issue need to account for int -->
            <div class="mb-4">
                <label for="driver_id" class="block text-sm font-medium mb-1">Driver:</label>
                <select id="driver_id" name="driver_id" bind:value={selectedDriverId} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500">
                    {#each drivers as { DriverId, DriverName }}
                        <option value={DriverId}>{DriverName}</option>
                    {/each}
                </select>
            </div>

            <div class="mb-4">
                <label for="startTime" class="block text-sm font-medium mb-1">Start Time:</label>
                <input type="datetime-local" id="start_time" name="start_time" required bind:value={selectedStartTime} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500" />
            </div>

            <div class="mb-4">
                <label for="endTime" class="block text-sm font-medium mb-1">End Time:</label>
                <input type="datetime-local" id="end_time" name="end_time" required bind:value={selectedEndTime} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500" />
            </div>

            <div class="mt-4 flex justify-center">
                <button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800">Update Schedule</button>
            </div>
        </form>
    </div>
</div>

