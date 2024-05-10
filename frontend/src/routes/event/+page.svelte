<script lang="ts">
	import { onMount } from 'svelte';
  import EventTracker from '../../components/eventTracker.svelte';
  import type { EventBus } from '../../types/global.js';
	import type { WebSocketServer } from 'vite';
  export let data
  const { session, followBus, buses, stops, backend_uri } = data

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
    ws = new WebSocket(`ws://${backend_uri.split("//")[1]}:3000/ws`)
  })
</script>

<div class="p-6 md:p-12">
  <h1 class="text-3xl font-semibold">
    Bus Event
  </h1>
  {#if followBus == null || followBus.Carplate == ''}
    <div>
      <form on:submit={createFollowBus}>
        <select
          bind:value={selectedBus}
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
        <button type="submit">
          Follow Bus
        </button>
      </form>
    </div>
  {:else}
    <EventTracker {data} {followBus} {stops} {ws} isEditable={true}/>
  {/if}
</div>