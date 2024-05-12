<script lang="ts">
  export let data 
  const { backend_uri, stops, env } = data
	import { onMount } from "svelte";
  import EventTracker from "../components/eventTracker.svelte";
  import type { FollowBusEvent } from "../types/global";

  let busList: Array<FollowBusEvent> = []
  let ws: WebSocket
  
  onMount(async() => {
    const getAllFollowBus = async () => {
      const response = await fetch(`${backend_uri}:3000/event/get-all-follow-bus`)
      busList = await response.json()
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
    {#each busList as followBus}
      <EventTracker {data} {followBus} {ws} stops={stops.find((stop) => stop != null ? stop[0].RouteName == followBus.RouteName : false) ?? []} isEditable={false}/>
    {/each}
  </div>
</div>