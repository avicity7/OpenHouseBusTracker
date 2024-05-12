<script lang="ts">
	import { onMount } from 'svelte';
  import EventTracker from '../../components/eventTracker.svelte';
  import type { EventBus } from '../../types/global.js';
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
  <h1 class="text-3xl font-semibold">
    Bus Event
  </h1>
  {#if followBus == null || followBus.Carplate == ''}
    <div class="w-full mt-28 flex">
      <form on:submit={createFollowBus} class="mx-auto flex flex-col">
        <select
          bind:value={selectedBus}
          class="rounded-lg border-2 border-stone-400"
        >
          {#each buses as bus}
            <option 
              value={bus.Carplate}
              selected={selectedBus.Carplate == bus.Carplate}
            >
              {bus.Carplate}
            </option>
          {/each}
        </select>
        <button type="submit" class="mt-12 bg-red-700 hover:bg-red-800 text-white rounded-md">
          Follow Bus
        </button>
      </form>
    </div>
  {:else}
    <EventTracker {data} {followBus} {stops} {ws} isEditable={true}/>
  {/if}
</div>