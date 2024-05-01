<script lang="ts">
  export let bus;

  function formatTime(timestamp: string | number | Date) {
    const utcDate = new Date(timestamp);
    const localDate = new Date(utcDate.getTime() + (utcDate.getTimezoneOffset() * 60000)); 
    return localDate.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
  }
</script>

  <div class="p-8 h-[300px] bg-white rounded-2xl shadow-md mt-6 md:mr-4">
    <div class="flex flex-row justify-between">
      <div>
        <h1 class="text-3xl font-bold">
          {bus.DriverName}
        </h1>
      </div>
      <div class="flex flex-col items-end">
        <p class="mt-1 font-semibold text-lg">
          {bus.RouteName}
        </p>
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


 <!-- <h1 class="text-2xl font-semibold">
        Bus {bus.BusScheduleId}
      </h1>
      <p class="text-lg mt-1">
        {bus.Carplate}
      </p> -->

  <!-- <div class={"flex flex-row items-center font-semibold " + (bus.Status ? "text-green-600" : "text-red-700")}> -->
    <!-- <svg class="mr-2" xmlns="http://www.w3.org/2000/svg" width="0.75em" height="0.75em" viewBox="0 0 24 24" {...$$props}><path fill="currentColor" d="M12 2C6.47 2 2 6.47 2 12s4.47 10 10 10s10-4.47 10-10S17.53 2 12 2"/></svg> -->
  <!-- {bus.Status ? "Touring" : "Inactive"} -->
<!-- </div> -->