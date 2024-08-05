<script lang="ts">
    import { onMount } from "svelte";
    import type { Schedule } from "$lib/types/global.js";
    import ToolTip from "$lib/components/ToolTip.svelte";

    let busSchedule: Schedule[] = [];
    let filteredSchedules: Schedule[] = [];
    let selectedSchedules = new Set<number>();
    let searchTerm = "";
    let selectedRoute = "";
    let selectedCarplate = "";
    let startTime = "";
    let endTime = "";

    let uniqueRoutes: string[] = [];
    let uniqueCarplates: string[] = [];

    export let data;
    const { backend_uri, schedules, routes, carplates } = data

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

        } catch (error) {
            console.error('Error deleting schedule:', error);
        }
    }

    function formatTimestamp(timestamp: string): string {
        const utcDate = new Date(timestamp);
        const formattedDate = utcDate.toLocaleString();
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


    onMount(() => {
        if (data && schedules) {
            busSchedule = schedules;
        }
        
        if (Array.isArray(routes)) {
            uniqueRoutes = routes.map(route => route.RouteName);
        }
        if (Array.isArray(carplates)) {
            uniqueCarplates = carplates.map(carplate => carplate.Carplate);
        }
    });

</script>

<svelte:head>
	<title>Manage - Schedule | SPOH Bus Tracker</title>
</svelte:head>

<div class="p-6 md:p-12">
    <div class="flex items-center justify-between mb-4">
        {#if selectedSchedules.size > 1}
            <button 
                class="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 mr-2 rounded-xl transition duration-300"
                on:click={bulkDelete}
            >
                Bulk Delete
            </button>
        {/if}
        <a href="/admin/schedule/add-schedule" class="border-black text-white font-semibold text-md px-6 py-2 rounded-xl bg-red-700 hover:bg-red-800 mr-2" data-testid="add-schedule-button">
            Add Schedule
        </a>    

        <div class="ml-auto">
            <select class="border border-gray-300 text-sm rounded-xl px-3 py-2" bind:value={selectedRoute} on:change ={filterSchedules} data-testid="search-route">
                <option value="">All Routes</option>
                {#each uniqueRoutes as route}
                    <option value={route}>{route}</option>
                {/each}
            </select>
            <select class="border border-gray-300 text-sm rounded-xl px-3 py-2" bind:value={selectedCarplate} on:change={filterSchedules} data-testid="search-carplate">
                <option value="">All Carplates</option>
                {#each uniqueCarplates as carplate}
                    <option value={carplate}>{carplate}</option>
                {/each}
            </select>
            <input type="datetime" placeholder="Start Time" class="border border-gray-300 text-sm rounded-xl px-3 py-2" bind:value={startTime} on:input={filterSchedules} data-testid="search-start-time">
            <input type="datetime" placeholder="End Time" class="border border-gray-300 text-sm rounded-xl px-3 py-2" bind:value={endTime} on:input={filterSchedules} data-testid="search-end-time">
        </div>
        
        <div class="ml-auto">
            <input type="text" placeholder="Search Drivers..." class="border border-gray-300 rounded-md px-3 py-2 w-60" bind:value={searchTerm} on:input={filterSchedules} data-testid="search-input">
        </div>
    </div>

    <div class="mt-8">
        <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
                <tr>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        <input type="checkbox" on:click={toggleSelectAll} />
                    </th>
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
                        <tr class="hover:bg-gray-100 schedule-item" data-testid= "search-results">
                            <td class="px-6 py-4 whitespace-nowrap">
                                <input type="checkbox" checked={selectedSchedules.has(schedule.BusScheduleId)} on:change={() => {
                                    toggleSelection(schedule.BusScheduleId);
                                    console.log("Updated selected schedules:", selectedSchedules);
                                }} />                     
                            </td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.Carplate}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.RouteName}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.DriverName}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{formatTimestamp(schedule.StartTime)}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{formatTimestamp(schedule.EndTime)}</td>
                            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                                <div class="flex items-center">
                                    <a href={`schedule/update-schedule/${schedule.BusScheduleId}`} class="text-stone-500 hover:text-green-500 mr-8">
                                        <ToolTip text="Update Schedule"> 
                                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                            <path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/>
                                        </svg>
                                        </ToolTip>
                                    </a>
                                    <button class="text-stone-500 hover:text-red-600 text-2xl" on:click={() => deleteSchedule(schedule.BusScheduleId)} data-testid="delete-schedule-button">
                                        <ToolTip text="Delete Schedule"> 
                                        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" {...$$props}>
                                            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" {...$$props}>
                                                <g fill="none">
                                                    <path d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022m-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z" />
                                                    <path fill="currentColor" d="M14.28 2a2 2 0 0 1 1.897 1.368L16.72 5H20a1 1 0 1 1 0 2l-.003.071l-.867 12.143A3 3 0 0 1 16.138 22H7.862a3 3 0 0 1-2.992-2.786L4.003 7.07A1.01 1.01 0 0 1 4 7a1 1 0 0 1 0-2h3.28l.543-1.632A2 2 0 0 1 9.721 2zm3.717 5H6.003l.862 12.071a1 1 0 0 0 .997.929h8.276a1 1 0 0 0 .997-.929zM10 10a1 1 0 0 1 .993.883L11 11v5a1 1 0 0 1-1.993.117L9 16v-5a1 1 0 0 1 1-1m4 0a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1m.28-6H9.72l-.333 1h5.226z" />
                                                </g>
                                            </svg>
                                        </svg>
                                        </ToolTip>
                                    </button>
                                </div>
                            </td>
                        </tr>
                    {/each} 
                {:else if searchTerm !== '' || selectedCarplate !== '' || selectedRoute !== '' || startTime !== '' || endTime !== ''}
                    <tr>
                        <td colspan="8" class="px-6 py-4 whitespace-nowrap text-center text-gray-400">No matching schedules found</td>
                    </tr>
                {:else}
                    {#each busSchedule as schedule (schedule.BusScheduleId)}
                        <tr class="hover:bg-gray-100">
                            <td class="px-6 py-4 whitespace-nowrap">
                                <input type="checkbox" checked={selectedSchedules.has(schedule.BusScheduleId)} on:change={() => {
                                    toggleSelection(schedule.BusScheduleId);
                                    console.log("Updated selected schedules:", selectedSchedules);
                                }} />                     
                            </td>
                            
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.Carplate}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.RouteName}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{schedule.DriverName}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{formatTimestamp(schedule.StartTime)}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{formatTimestamp(schedule.EndTime)}</td>
                            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                                <div class="flex items-center">
                                    <a href={`schedule/update-schedule/${schedule.BusScheduleId}`} class="text-stone-500 hover:text-green-500 mr-8">
                                        <ToolTip text="Update Schedule"> 
                                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                            <path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/>
                                        </svg>
                                        </ToolTip>
                                    </a>
                                    <button class="text-stone-500 hover:text-red-600 text-2xl" on:click={() => deleteSchedule(schedule.BusScheduleId)}>
                                        <ToolTip text="Delete Schedule"> 
                                        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" {...$$props}>
                                            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" {...$$props}>
                                                <g fill="none">
                                                    <path d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022m-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z" />
                                                    <path fill="currentColor" d="M14.28 2a2 2 0 0 1 1.897 1.368L16.72 5H20a1 1 0 1 1 0 2l-.003.071l-.867 12.143A3 3 0 0 1 16.138 22H7.862a3 3 0 0 1-2.992-2.786L4.003 7.07A1.01 1.01 0 0 1 4 7a1 1 0 0 1 0-2h3.28l.543-1.632A2 2 0 0 1 9.721 2zm3.717 5H6.003l.862 12.071a1 1 0 0 0 .997.929h8.276a1 1 0 0 0 .997-.929zM10 10a1 1 0 0 1 .993.883L11 11v5a1 1 0 0 1-1.993.117L9 16v-5a1 1 0 0 1 1-1m4 0a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1m.28-6H9.72l-.333 1h5.226z" />
                                                </g>
                                            </svg>
                                        </svg>
                                        </ToolTip>
                                    </button>
                                </div>
                            </td>                 
                        </tr>
                    {/each}
                {/if}
            </tbody>
        </table>
    </div>
</div>