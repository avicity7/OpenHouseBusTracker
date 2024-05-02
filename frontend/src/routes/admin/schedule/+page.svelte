<script lang="ts">
    import { onMount } from "svelte";
    import { writable } from "svelte/store";
    import type { Schedule } from "../../../types/global";

    let busSchedule = writable<Schedule[]>([]);
    let selectedSchedules = new Set<number>();

    export let data;
    const { backend_uri } = data

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

            const updatedSchedules = $busSchedule.filter(schedule => {
                if (Array.isArray(id)) {
                    return !id.includes(schedule.BusScheduleId);
                } else {
                    return schedule.BusScheduleId !== id;
                }
            });
            busSchedule.set(updatedSchedules);

            console.log("Deleted Bus Schedule with ID:", id);
        } catch (error) {
            console.error('Error deleting schedule:', error);
        }
    }

    // might be cleaner and easier using a library
    function formatTimestamp(timestamp: string): string { 
        const utcDate = new Date(timestamp);
        const localDate = new Date(utcDate.getTime() + (utcDate.getTimezoneOffset() * 60000)); 
        const year = localDate.getFullYear();
        const month = (localDate.getMonth() + 1).toString().padStart(2, '0');
        const day = localDate.getDate().toString().padStart(2, '0');
        let hours = localDate.getHours();
        const ampm = hours >= 12 ? 'PM' : 'AM';
        hours = hours % 12;
        hours = hours ? hours : 12;
        const hoursString = hours.toString().padStart(2, '0');
        const minutes = localDate.getMinutes().toString().padStart(2, '0');
        const seconds = localDate.getSeconds().toString().padStart(2, '0');

        return `${year}-${month}-${day} ${hoursString}:${minutes}:${seconds} ${ampm}`;
    }

    function toggleSelection(id: number) {
        if (selectedSchedules.has(id)) {
            selectedSchedules.delete(id);
        } else {
            selectedSchedules.add(id);
        }
        // console.log("Number of selected schedules:", selectedSchedules.size);
        // console.log("Selected Bus Schedule IDs:", Array.from(selectedSchedules));
        selectedSchedules = selectedSchedules;
    }

    function toggleSelectAll(event: MouseEvent) {
        const isChecked = (event.target as HTMLInputElement).checked;
        if (isChecked) {
            selectedSchedules.clear();
            $busSchedule.forEach(schedule => {
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

    onMount(() => {
        if (data && data.data) {
            busSchedule.set(data.data);
        }
    });
</script>

<div class="p-6 md:p-12">
    <div class="flex items-center mb-4">
        <h1 class="text-3xl font-semibold mr-4">Bus Schedules</h1>

        {#if selectedSchedules.size > 1}
            <button 
                class="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 mr-2 rounded transition duration-300"
                on:click={bulkDelete}
            >
                Bulk Delete
            </button>
        {/if}
        <a href="schedule/add-schedule" class="border-2 border-black text-black text-xl px-4 py-2 rounded-full hover:bg-gray-200">
            +
        </a>    
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
                {#each $busSchedule as schedule (schedule.BusScheduleId)}
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
            </tbody>
        </table>
    </div>
</div>
