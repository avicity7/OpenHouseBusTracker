<script lang="ts">
  import BusScheduleCard from "$lib/components/BusScheduleCard.svelte"

  export let data;

  const { currentSchedules = [], futureSchedules = [] } = data;

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
  {#if futureSchedules.length > 0}
    <div class="md:grid grid-cols-3">
      {#each futureSchedules as bus}
        <BusScheduleCard {bus} />
      {/each}
    </div>
  {:else}
    <p class="text-gray-500">No upcoming schedules</p>
  {/if}
</div>