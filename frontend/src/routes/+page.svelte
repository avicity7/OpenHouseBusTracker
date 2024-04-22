<script>
  import BusCard from "../components/busCard.svelte";
  let bus_list = [{bus_id: 1, carplate: "S123456789Z", status: 1, route: "East Route"}, {bus_id: 2, carplate: "S987654321Z", status: 1, route: "West Route"}]

  let number = 0
</script>

<div class="p-6 md:p-12">
  <h1 class="text-3xl font-semibold">
    Bus Tracker
  </h1>
  <div class="md:grid grid-cols-3">
    {#each bus_list as bus}
      <BusCard {bus} />
    {/each}
  </div>
</div>