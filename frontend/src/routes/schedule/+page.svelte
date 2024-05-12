<script lang="ts">
  import { onMount } from "svelte";
  import BusScheduleCard from "$lib/components/BusScheduleCard.svelte"
  import type { Schedule } from "$lib/types/global";

  export let data: { data: Schedule[] };

// let busScheduleList: Schedule[] = [];

let currentSchedules: Schedule[] = [];
let upcomingSchedules: Schedule[] = [];

onMount(() => {
    if (data && data.data) {
      const currentTime = new Date().getTime();
      
      data.data.forEach(schedule => {
        const startTime = new Date(schedule.StartTime).getTime();
        const endTime = new Date(schedule.EndTime).getTime();
        if (currentTime >= startTime && currentTime <= endTime) {
          currentSchedules.push(schedule);
        } else {
          upcomingSchedules.push(schedule);
        }
        currentSchedules = currentSchedules
        upcomingSchedules = upcomingSchedules 
      });

      // busScheduleList = data.data;
      // console.log(busScheduleList)
    }
  });
</script>

  <div class="p-6 md:p-12">
    <h1 class="text-3xl font-semibold pt-8 mb-6">My Shifts</h1>

    <h1 class="text-xl font-semibold text-slate-400 pt-8">Current Shifts</h1>
    {#if currentSchedules.length > 0}
      <div class="md:grid grid-cols-3">
        {#each currentSchedules as bus}
          <BusScheduleCard {bus} />
        {/each}
      </div>
    {:else}
      <p class="text-gray-500">No current schedules</p>
    {/if}

    <hr class="my-8 border-gray-200">

    <h1 class="text-xl font-semibold text-slate-400 pt-8">Upcoming Shifts</h1>
    {#if upcomingSchedules.length > 0}
      <div class="md:grid grid-cols-3">
        {#each upcomingSchedules as bus}
          <BusScheduleCard {bus} />
        {/each}
      </div>
    {:else}
      <p class="text-gray-500">No upcoming schedules</p>
    {/if}
  </div>