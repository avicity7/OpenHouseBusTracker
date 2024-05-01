<script lang="ts">
  import { onMount } from "svelte";
  import BusScheduleCard from "../../components/busScheduleCard.svelte";
	import type { Schedule } from "../../types/global";

  export let data;
  let { backend_uri } = data;

  let busScheduleList: Schedule[] = [];

  async function fetchData() {
    try {
      const response = await fetch(`${backend_uri}:3000/schedules/get-schedules`);
      if (!response.ok) {
        throw new Error('Failed to fetch data');
      }
      const data = await response.json();
      busScheduleList = data; 
      console.log('Fetched data:', busScheduleList);
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  }

  onMount(() => {
    fetchData();
  });  
</script>

<!-- Should this page be unique to each bus uniquely? (To be asked) -->
<div class="p-6 md:p-12">
  {#if busScheduleList.length > 0}
  <div class="flex justify-center mb-4">
    <div class="p-4 flex justify-between items-center">
      <div class="mr-24">
        <p class="text-3xl font-semibold">Bus {busScheduleList[0].BusScheduleId}</p> 
        <!-- shouldnt be bus Schedule id, should be bus id -->
        <p class="text-md font-medium">{busScheduleList[0].Carplate}</p>
      </div>
      <div class="ml-24">
        <div class="flex flex-row items-center font-semibold text-green-600">
          <svg class="mr-2" xmlns="http://www.w3.org/2000/svg" width="0.75em" height="0.75em" viewBox="0 0 24 24">
            <path fill="currentColor" d="M12 2C6.47 2 2 6.47 2 12s4.47 10 10 10s10-4.47 10-10S17.53 2 12 2"/>
          </svg>
          Touring
        </div>
        <p class="text-xl font-semibold">{busScheduleList[0].RouteName}</p>
      </div>
    </div>
  </div>
{/if}


  <h1 class="text-2xl font-semibold text-center text-slate-400 pt-8">Shifts</h1>
  <div class="md:grid grid-cols-3">
    {#each busScheduleList as bus}
      <BusScheduleCard {bus} />
    {/each}
  </div>
</div>
