<script lang="ts">
    import { onMount } from "svelte";
    import type { Schedule } from "../../../types/global";

    let busSchedule: Schedule[] = [];
    let filteredSchedules: Schedule[] = [];
    let selectedSchedules = new Set<number>();
    let searchTerm = "";
    let selectedRoute = "";
    let selectedCarplate = "";
    let startTime = "";
    let endTime = "";

    let uniqueRoutes = [];
    let uniqueCarplates = [];

    export let data;
    const { backend_uri, schedules, dropdownData } = data

    async function deleteSchedule(id: number | number[]) {
        try {
            const response = await fetch(`${backend_uri}:3000/schedules/delete-schedule`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(Array.isArray(id) ? id : [id])
            });
            
            if (!response.ok) {
                throw new Error('Failed to delete schedule');
            }

            const updatedSchedules = busSchedule.filter(schedule => {
                if (Array.isArray(id)) {
                    return !id.includes(schedule.BusScheduleId);
                } else {
                    return schedule.BusScheduleId !== id;
                }
            });
            busSchedule = updatedSchedules;

            console.log("Deleted Bus Schedule with ID:", id);
        } catch (error) {
            console.error('Error deleting schedule:', error);
        }
    }

    // using toLocaleString() is much cleaner
    function formatTimestamp(timestamp: string): string {
        const utcDate = new Date(timestamp);
        const localDate = new Date(utcDate.getTime() + (utcDate.getTimezoneOffset() * 60000)); 
        const formattedDate = localDate.toLocaleString();
        return formattedDate;
    }

    function toggleSelection(id: number) {
        if (selectedSchedules.has(id)) {
            selectedSchedules.delete(id);
        } else {
            selectedSchedules.add(id);
        }
        selectedSchedules = selectedSchedules;
    }

    function toggleSelectAll(event: MouseEvent) {
        const isChecked = (event.target as HTMLInputElement).checked;
        if (isChecked) {
            selectedSchedules.clear();
            busSchedule.forEach(schedule => {
                selectedSchedules.add(schedule.BusScheduleId);
            });
        } else {
            selectedSchedules.clear();
        }
        selectedSchedules = selectedSchedules;

        const checkboxes = document.querySelectorAll('input[type="checkbox"]');
        checkboxes.forEach(checkbox => {
            (checkbox as HTMLInputElement).checked = isChecked;
        });
    }

    async function bulkDelete() {
        const idsToDelete = Array.from(selectedSchedules);
        console.log(idsToDelete)
        try {
            await deleteSchedule(idsToDelete);
            const selectAllCheckbox = document.querySelector('input[type="checkbox"]');
            if (selectAllCheckbox) {
                (selectAllCheckbox as HTMLInputElement).checked = false;
            }
        
            console.log("Bulk delete successful.");
        } catch (error) {
            console.error('Error during bulk delete:', error);
        }
    }

    function filterSchedules() {
        console.log(selectedCarplate)
        filteredSchedules = busSchedule.filter(schedule => {
            const driverName = schedule.DriverName.toLowerCase();
            const routeMatch = !selectedRoute || schedule.RouteName.toLowerCase() === selectedRoute.toLowerCase();
            const carplateMatch = !selectedCarplate || schedule.Carplate.toLowerCase() === selectedCarplate.toLowerCase();
            const startTimeMatch = !startTime || formatTimestamp(schedule.StartTime).toLowerCase().includes(startTime.toLowerCase());
            const endTimeMatch = !endTime || formatTimestamp(schedule.EndTime).toLowerCase().includes(endTime.toLowerCase());
            const searchTermMatch = !searchTerm || driverName.includes(searchTerm.toLowerCase());
            
        return searchTermMatch && routeMatch && carplateMatch && startTimeMatch && endTimeMatch;
        });
    }

        function getUniqueRoutes() {
            if (dropdownData) {
                return [...new Set(dropdownData.map(item => item.RouteName))];
            } else {
                return [];
            }
        }

        function getUniqueCarplates() {
            if (dropdownData) {
                return [...new Set(dropdownData.map(item => item.Carplate))];
            } else {
                return [];
            }
        }

    onMount(() => {
        if (data && schedules) {
            busSchedule = schedules;
        }
        uniqueRoutes = getUniqueRoutes();
        uniqueCarplates = getUniqueCarplates();
    });

</script>

<div class="p-6 md:p-12">
    <div class="flex items-center justify-between mb-4">
        <h1 class="text-3xl font-semibold mr-4">Bus Schedules</h1>
    
        {#if selectedSchedules.size > 1}
            <button 
                class="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 mr-2 rounded transition duration-300"
                on:click={bulkDelete}
            >
                Bulk Delete
            </button>
        {/if}
        <a href="schedule/add-schedule" class="border-2 border-black text-black text-xl px-4 py-2 rounded-full hover:bg-gray-200 mr-2">
            +
        </a>    

        <div class="ml-auto">
            <select class="border border-gray-300 text-sm rounded-xl px-3 py-2" bind:value={selectedRoute} on:change ={filterSchedules}>
                <option value="">All Routes</option>
                {#each getUniqueRoutes() as route}
                    <option value={route}>{route}</option>
                {/each}
            </select>
            <select class="border border-gray-300 text-sm rounded-xl px-3 py-2" bind:value={selectedCarplate} on:change={filterSchedules}>
                <option value="">All Carplates</option>
                {#each getUniqueCarplates() as carplate}
                    <option value={carplate}>{carplate}</option>
                {/each}
            </select>
            <input type="datetime" placeholder="Start Time" class="border border-gray-300 text-sm rounded-xl px-3 py-2" bind:value={startTime} on:input={filterSchedules}>
            <input type="datetime" placeholder="End Time" class="border border-gray-300 text-sm rounded-xl px-3 py-2" bind:value={endTime} on:input={filterSchedules}>
        </div>
        
        <div class="ml-auto">
            <input type="text" placeholder="Search..." class="border border-gray-300 rounded-md px-3 py-2 w-60" bind:value={searchTerm} on:input={filterSchedules}>
        </div>
    </div>

    <div class="mt-8">
        <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
                <tr>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        <input type="checkbox" on:click={toggleSelectAll} />
                    </th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Bus Schedule ID</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Bus Carplate</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Route Description</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Driver Name</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Start Time</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">End Time</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
                </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
                {#if filteredSchedules.length > 0}
                    {#each filteredSchedules as schedule (schedule.BusScheduleId)}
                        <tr>
                            <td class="px-6 py-4 whitespace-nowrap">
                                <input type="checkbox" checked={selectedSchedules.has(schedule.BusScheduleId)} on:change={() => {
                                    toggleSelection(schedule.BusScheduleId);
                                    console.log("Updated selected schedules:", selectedSchedules);
                                }} />                     
                            </td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.BusScheduleId}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.Carplate}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.RouteName}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.DriverName}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{formatTimestamp(schedule.StartTime)}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{formatTimestamp(schedule.EndTime)}</td>
                            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                                <a href={`schedule/update-schedule/${schedule.BusScheduleId}`} class="text-green-600 hover:text-green-900 mr-5">Update</a>
                                <button class="text-red-600 hover:text-red-900" on:click={() => deleteSchedule(schedule.BusScheduleId)}>Delete</button>
                            </td>
                        </tr>
                    {/each} 
                {:else if searchTerm !== '' || selectedCarplate !== '' || selectedRoute !== '' || startTime !== '' || endTime !== ''}
                    <tr>
                        <td colspan="8" class="px-6 py-4 whitespace-nowrap text-center">No matching schedules found</td>
                    </tr>
                {:else}
                    {#each busSchedule as schedule (schedule.BusScheduleId)}
                        <tr>
                            <td class="px-6 py-4 whitespace-nowrap">
                                <input type="checkbox" checked={selectedSchedules.has(schedule.BusScheduleId)} on:change={() => {
                                    toggleSelection(schedule.BusScheduleId);
                                    console.log("Updated selected schedules:", selectedSchedules);
                                }} />                     
                            </td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.BusScheduleId}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.Carplate}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.RouteName}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.DriverName}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{formatTimestamp(schedule.StartTime)}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{formatTimestamp(schedule.EndTime)}</td>
                            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                                <a href={`schedule/update-schedule/${schedule.BusScheduleId}`} class="text-green-600 hover:text-green-900 mr-5">Update</a>
                                <button class="text-red-600 hover:text-red-900" on:click={() => deleteSchedule(schedule.BusScheduleId)}>Delete</button>
                            </td>
                        </tr>
                    {/each}
                {/if}
            </tbody>
        </table>
    </div>
</div>