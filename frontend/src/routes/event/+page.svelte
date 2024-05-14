<script lang="ts">
	import { onMount } from 'svelte';
  import EventTracker from '$lib/components/EventTracker.svelte'
  import type { EventBus } from '$lib/types/global.js'
	import type { WebSocketServer } from 'vite';
  export let data
  const { session, followBus, buses, stops, backend_uri, env } = data

  let selectedBus: EventBus = {
    Carplate: '',
    Status: false
  }

  let ws: WebSocket

  const createFollowBus = async() => {
    await fetch(`${backend_uri}:3000/event/create-follow-bus`, {
      method: 'PUT',
      body: JSON.stringify({
        Carplate: selectedBus,
        Email: session!.Email
      }),
      headers: {
        'content-type': 'application/json'
      },
    })
  }

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