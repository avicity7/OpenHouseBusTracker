<script lang="ts">
  import { PUBLIC_BACKEND_URL } from '$env/static/public'
	import { onMount } from 'svelte';
  export let data
  const { session, followBus, buses, stops } = data
  import type { EventBus, Event } from '../../types/global.js'
  let events: Array<Event> = [{'StopName': '', 'Order': 0, 'EventId': 0, 'Timestamp': ''}]
  let width = 0,
    loaded = false

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
  
  const getEvents = async() => {
    const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event/get-events/${followBus?.ScheduleId}`)
    let parsed = await response.json() as Array<Event>
      if (parsed !== null) {
        events = parsed
      }
    let baseWidth = Math.floor(100 / (stops.length))
    if (events[0].EventId == 1 || events[0].EventId == 3 && events[0].Order == stops.length) {
      width = baseWidth / 4
    } else if (events[0].EventId == 2) {
      width = baseWidth * events[0].Order - (baseWidth / 2)
    } else {
      width = baseWidth * events[0].Order
    }
  }

  const createEvent = async() => {
    let payload = {
      Carplate: followBus?.Carplate,
      RouteName: followBus?.RouteName,
      EventId: 0,
      StopName: '' 
    }
    if (events[0].EventId == 3) {
      const stopIndex = stops.findIndex((stop) => stop.StopName == events[0].StopName)
      if (stopIndex != stops.length - 1) {
        payload.EventId = 2
        payload.StopName = stops[stopIndex+1].StopName
      } else {
        payload.EventId = 2
        payload.StopName = stops[0].StopName
      }
    } else if (events[0].EventId == 2) {
      payload.EventId = 3
      payload.StopName = events[0].StopName
    } else if (events[0].EventId == 1) {
      payload.EventId = 2
      payload.StopName = stops[0].StopName
    } else {
      payload.EventId = 1
      payload.StopName = stops[0].StopName
    }

    await fetch(`${PUBLIC_BACKEND_URL}:3000/event/create-event`, {
      method: 'PUT',
      body: JSON.stringify(payload),
      headers: {
        'content-type': 'application/json'
      },
    })

    getEvents()
  }

  onMount(async () => {
    await getEvents()
    loaded = true
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
  {:else if loaded}
    <div class="bg-white max-w-4xl mx-auto flex flex-col rounded-lg p-8 mt-8">
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
      <div class="h-1 bg-stone-300 w-full rounded-full mt-12 flex items-center relative z-0">
        <div class="h-1 bg-red-600 absolute rounded-full" style={`width: ${width}%`}></div>
        <div class={"grid w-full "+"grid-cols-"+stops.length}>
          {#each stops as stop}
            <div class="mx-auto">
              <div class={"h-6 w-6 rounded-full " + (stop.Order <= events[0].Order && events[0].EventId != 1 && !(events[0].Order == stops.length && events[0].EventId == 3)? "bg-red-600" : "bg-stone-300")}></div> 
              <div class="absolute">{stop.StopName}</div>
            </div>
          {/each}
        </div>
      </div>
      <div class="mt-20 flex flex-col items-center">
        {#if events[0].EventId == 3 || events[0].Order == 0 || events[0].EventId == 1}
          <h2>Next stop:</h2>
        {:else}
          <h2>Current stop:</h2>
        {/if}
        <div class="text-4xl font-semibold mt-4">
          {#if events[0].Order != 0}
            {#if events[0].EventId == 3}
              {#if events[0].Order != stops.length}
                <h1>{stops[events[0].Order].StopName}</h1>
              {:else}
                <h1>{stops[0].StopName}</h1>
              {/if}
            {:else}
              <h1>{stops[events[0].Order - 1].StopName}</h1>
            {/if}
          {:else}
            <h1>{stops[0].StopName}</h1>
          {/if}
        </div>
      </div>
      <div class="mt-8 items-center mx-auto">
        {#if events[0].Order == 0}
          <button class="bg-red-600 hover:bg-red-700 px-12 py-2 rounded-lg text-white" on:click={createEvent}>
            Start Tour
          </button>
        {:else if events[0].EventId == 1 || events[0].EventId == 3}
          <button class="bg-red-700 hover:bg-red-800 px-12 py-2 rounded-lg text-white" on:click={createEvent}>
            Arrived
          </button>
        {:else if events[0].EventId == 2}
          <button class="bg-blue-700 hover:bg-blue-800 px-12 py-2 rounded-lg text-white" on:click={createEvent}>
            Departed
          </button>
        {/if}
      </div>
    </div>
  {/if}
</div>