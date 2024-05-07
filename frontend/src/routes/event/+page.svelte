<script lang="ts">
  import { PUBLIC_BACKEND_URL } from '$env/static/public'
  export let data
  const { session, followBus, buses, stops } = data
  import type { EventBus } from '../../types/global.js'

  let selectedBus: EventBus = {
    Carplate: '',
    Status: false
  }

  const createFollowBus = async() => {
    await fetch(`${PUBLIC_BACKEND_URL}:3000/event/create-follow-bus`, {
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
    <div class="bg-white max-w-4xl mx-auto flex flex-col rounded p-8">
      <div class="flex justify-between">
        <h1 class="text-2xl">{followBus.Carplate}</h1>
      </div>
      <div class="flex justify-between mt-1">
        <h2 class="text-xl font-semibold">{followBus.DriverName}</h2>
        <h2 class="text-xl font-semibold">{followBus.RouteName}</h2>
      </div>
      <div class="flex justify-between mt-2">
        <div>{new Date(followBus.StartTime).toLocaleString()}</div>
        <div>{new Date(followBus.EndTime).toLocaleString()}</div>
      </div>
      <div class="h-1 bg-stone-300 w-full rounded-full mt-12 flex items-center relative">
        <div class="bg-red-600 h-1 absolute rounded-full w-full"></div>
        {#each stops as stop}
          <div class={(stop.Order == 1 ? "" : `ml-[${Math.round(100 / stops.length)}%] `)+"mx-auto relative"}>
            <div class="h-6 w-6 bg-red-600 rounded-full"></div> 
            <div class="absolute">{stop.StopName}</div>
          </div>
        {/each} 
      </div>
      <div class="mt-8 items-center">

      </div>
    </div>
  {/if}
</div>