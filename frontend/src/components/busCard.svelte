<script lang="ts">
  export let bus: {bus_id: number, carplate: string, status: number, route: string}
</script>

<div class="p-6 bg-white rounded-xl border border-stone-200 shadow-sm mt-6 md:mr-4">
  <div class="flex flex-row justify-between">
    <div>
      <h1 class="text-2xl font-semibold">
        Bus {bus.bus_id}
      </h1>
      <p class="text-lg mt-1">
        {bus.carplate}
      </p>
    </div>
    <div class="flex flex-col items-end">
      <div class={"flex flex-row items-center font-semibold " + (bus.status ? "text-green-600" : "text-red-700")}>
        <svg class="mr-2" xmlns="http://www.w3.org/2000/svg" width="0.75em" height="0.75em" viewBox="0 0 24 24" {...$$props}><path fill="currentColor" d="M12 2C6.47 2 2 6.47 2 12s4.47 10 10 10s10-4.47 10-10S17.53 2 12 2"/></svg>
        {bus.status ? "Touring" : "Inactive"}
      </div>
      <p class="mt-1">
        {bus.route}
      </p>
    </div>
  </div>
  <div class="w-full mt-6 text-center">
    <p class="text-stone-400">
      Next stop
    </p>
    <p class="text-4xl font-semibold mt-2">
      Block T19
    </p>
    <p class="text-2xl mt-2">
      12:34 pm
    </p>
  </div>
</div>