<script lang="ts">
  import type { Schedule } from "$lib/types/global";
  export let bus: Schedule;
  export let shift: Boolean

  const cutoff = 14
  const d = new Date()
  const pastCutoff = d.getHours() > cutoff

  function getRouteColor(routeName: string): string {
    switch (routeName.toLowerCase()) {
      case 'green route':
        return '#2eff58';
      case 'yellow route':
        return '#ffee38';
      default:
        return '#cccccc'; // this is gray
    }
  }
</script>

  <div class="p-8 h-[300px] bg-white rounded-2xl shadow-md mt-6 md:mr-4">
    <div class="flex flex-row justify-between">
      <div>
        <h1 class="text-3xl font-bold">
          {bus.Carplate}
        </h1>
      </div>
      <div class="flex flex-col items-end">
        <div class="flex items-center">
          <div
            class="w-4 h-4 rounded-full mr-2 mb-1"
            style="background-color: {getRouteColor(bus.RouteName)}"
          ></div>
          <p class="font-semibold text-lg">
            {bus.RouteName}
          </p>
        </div>
      </div>
    </div>

    <div class="mt-14 flex flex-col items-center justify-center">
      {#if shift}
        {#if pastCutoff}
          <h2 class="text-lg font-light">Started at</h2>
          <h1 class="text-5xl font-medium mt-2">{cutoff}:00</h1>
        {:else}
          <h2 class="text-lg font-light">until</h2>
          <h1 class="text-5xl font-medium mt-2">{cutoff}:00</h1>
        {/if}
      {:else}
        <h2 class="text-lg font-light">Starts at</h2>
        <h1 class="text-5xl font-medium">{cutoff}:00</h1>
      {/if}
    </div>
  </div>