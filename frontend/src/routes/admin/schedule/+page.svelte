<script lang="ts">
    import { onMount } from "svelte";
    import { writable } from "svelte/store";
    import type { Schedule } from "../../../types/global";

    let busSchedule = writable<Schedule[]>([]);
    let selectedSchedules = new Set<number>();

    export let data;
    const { backend_uri, session } = data

    onMount(() => {
        if (data && data.data) {
            busSchedule.set(data.data);
        }
    });

    async function deleteSchedule(id: number) {
    try {
      const response = await fetch(`${backend_uri}:3000/schedules/delete-schedule/${id}`, {
        method: 'DELETE'
      });
      
      if (!response.ok) {
        throw new Error('Failed to delete schedule');
      }

      const newScheduleList = $busSchedule.filter(schedule => schedule.BusScheduleId !== id);
      busSchedule.set(newScheduleList);
      selectedSchedules.delete(id);
      console.log("Deleted Bus Schedule with ID:", id);
    } catch (error) {
      console.error('Error deleting schedule:', error);
    }
  }

    function formatTimestamp(timestamp: string): string {
        const date = new Date(timestamp);
        const year = date.getFullYear();
        const month = (date.getMonth() + 1).toString().padStart(2, '0');
        const day = date.getDate().toString().padStart(2, '0');
        const hours = date.getHours().toString().padStart(2, '0');
        const minutes = date.getMinutes().toString().padStart(2, '0');
        const seconds = date.getSeconds().toString().padStart(2, '0');

        return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    }

    function toggleSelection(id: number) {
        if (selectedSchedules.has(id)) {
            selectedSchedules.delete(id);
        } else {
            selectedSchedules.add(id);
        }
        
        console.log("Selected Bus Schedule IDs:", Array.from(selectedSchedules));
    }

    // function updateSchedule(id: number) {
    //     console.log("Update schedule with ID:", id);
    //     // Add your update logic here
    // }
</script>

<div class="p-6 md:p-12">
    <div class="flex items-center mb-4">
        <h1 class="text-3xl font-semibold mr-4">Bus Schedules</h1>
        <a href="schedule/add-schedule" class="border-2 border-black text-black text-xl px-4 py-2 rounded-full hover:bg-gray-200">
            +
        </a>    
    </div>

    <div class="mt-8">
        <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
                <tr>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Select</th>
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
                        <input type="checkbox" checked={selectedSchedules.has(schedule.BusScheduleId)} on:change={() => toggleSelection(schedule.BusScheduleId)} />
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
