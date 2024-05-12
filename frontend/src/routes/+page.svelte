<script lang="ts">
  export let data 
  const { backend_uri, stops, env } = data
	import { onMount } from "svelte";
  import EventTracker from "../components/eventTracker.svelte";
  import type { FollowBusEvent } from "../types/global";

  let busList: Array<FollowBusEvent> = []
  let ws: WebSocket

  let loaded = false
  
  onMount(async() => {
    const getAllFollowBus = async () => {
      const response = await fetch(`${backend_uri}:3000/event/get-all-follow-bus`)
      let parsed = await response.json()
      if (parsed != null) {
        busList = parsed
      }
    }

    ws = new WebSocket(`${env == "PROD" ? "wss" : "ws"}://${backend_uri.split("//")[1]}:3000/ws`)
  
    getAllFollowBus()
  }) 
</script>

<div class="p-6 md:p-12">
  <h1 class="text-3xl font-semibold">
    Bus Tracker
  </h1>
  <div>
    {#if loaded && busList.length != 0}
      {#each busList as followBus}
        <EventTracker {data} {followBus} {ws} stops={stops.find((stop) => stop != null ? stop[0].RouteName == followBus.RouteName : false) ?? []} isEditable={false}/>
      {/each}
    {:else}
    <div class="w-full flex mt-24">
      <p class="m-auto text-stone-500 font-semibold">There are no buses to display.</p>
    </div>
    {/if}
  </div>
</div>