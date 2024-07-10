<script lang="ts">
	import { onMount } from 'svelte';
  import EventTracker from '$lib/components/EventTracker.svelte'
  export let data
  const { followBus, stops, backend_uri, env } = data

  let ws: WebSocket

  onMount(() => {
    ws = new WebSocket(`${env == "PROD" ? "wss" : "ws"}://${backend_uri.split("//")[1]}:3000/ws`)
  })
</script>

<div class="p-6 md:p-12">
  {#if followBus == null || followBus.Carplate == ''}
    <div class="mt-24 flex flex-col m-auto items-center">
      <p class="text-stone-500 font-semibold">You're not assigned to any bus yet!</p>
    </div>
  {:else}
    <EventTracker {data} {followBus} {stops} {ws} isEditable={true}/>
  {/if}
</div>