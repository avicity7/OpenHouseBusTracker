<script lang="ts">
  import type { Schedule } from "$lib/types/global";
  export let bus: Schedule;

  function formatTime(timestamp: string | number | Date) {
    const utcDate = new Date(timestamp);
    return utcDate.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
  }

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

    <div class="mt-14 flex items-center justify-center">
      <div class="text-center">
        <p class="font-semibold text-md text-slate-400">From</p>
        <p class="font-medium text-4xl">{formatTime(bus.StartTime)}</p>
      </div>
      <div class="text-center ml-8">
        <p class="font-semibold text-md text-slate-400">To</p>
        <p class="font-medium text-4xl">{formatTime(bus.EndTime)}</p>
      </div>
    </div>
  </div>