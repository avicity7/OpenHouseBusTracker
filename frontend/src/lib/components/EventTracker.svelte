<script lang="ts">
  import type { EventBus, Event, RouteStep, FollowBusEvent } from '$lib/types/global.js'
  import { PUBLIC_BACKEND_URL } from '$env/static/public'
	import { onMount } from 'svelte';
  export let data
  const { session } = data
  export let stops: Array<RouteStep>
  export let followBus: FollowBusEvent 
  export let isEditable: boolean
  export let ws: WebSocket

  let events: Array<Event> = [{'StopName': '', 'Order': 0, 'EventId': 0, 'Timestamp': ''}]
  let width = 0,
    loaded = false
  
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

  const createRestartEvent = async() => {
    let payload = {
      Carplate: followBus?.Carplate,
      RouteName: followBus?.RouteName,
      EventId: 6,
      StopName: events[0].StopName 
    }
    await fetch(`${PUBLIC_BACKEND_URL}:3000/event/create-event`, {
      method: 'PUT',
      body: JSON.stringify(payload),
      headers: {
        'content-type': 'application/json'
      },
    })

    payload = {
      Carplate: followBus?.Carplate,
      RouteName: followBus?.RouteName,
      EventId: 1,
      StopName: stops[0].StopName 
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

  const createEmergency = async() => {
    let payload = {
      Carplate: followBus?.Carplate,
      RouteName: followBus?.RouteName,
      EventId: 5,
      StopName: events[0].StopName 
    }
    await fetch(`${PUBLIC_BACKEND_URL}:3000/event/create-event`, {
      method: 'PUT',
      body: JSON.stringify(payload),
      headers: {
        'content-type': 'application/json'
      },
    })
  }

  onMount(async () => {
    await getEvents()
    loaded = true
    ws.onmessage = (msg) => {
      let parts = msg.data.split(" ")
      if (parts[0] == 'refresh') {
        getEvents()
      }
    }
  })
</script>

<div class="p-6 md:p-12">
  <div class="bg-white max-w-4xl mx-auto flex flex-col rounded-lg p-8 mt-8">
    <div class="flex justify-between">
      <h1 class="text-2xl">{followBus?.Carplate}</h1>
    </div>
    <div class="flex justify-between mt-1">
      <h2 class="text-xl font-semibold">{followBus.DriverName}</h2>
      <h2 class="text-xl font-semibold">{followBus.RouteName}</h2>
    </div>
    <div class="flex justify-between mt-4">
      <div class="flex flex-col">
        <p class="font-semibold text-stone-600">Start Time</p>
        <p class="font-semibold">{new Date(followBus.BusStartTime).toLocaleString('en-gb', { weekday: 'long', year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' })}</p>
      </div>
      <div class="flex flex-col text-end">
        <p class="font-semibold text-stone-600">End Time</p>
        <p class="font-semibold">{new Date(followBus.BusEndTime).toLocaleString('en-gb', { weekday: 'long', year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' })}</p>
      </div>
    </div>
    <div class="h-1 bg-stone-300 w-full rounded-full mt-12 flex items-center relative z-0">
      <div class="h-1 bg-red-600 absolute rounded-full" style={`width: ${width}%`}></div>
      <div class="grid w-full" style={`grid-template-columns: repeat(${stops.length}, minmax(0, 1fr))`}>
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
        <h2 class="text-orange-800 font-semibold">On the way to:</h2>
      {:else}
        <h2 class="text-green-800 font-semibold">Picking up passengers at:</h2>
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
    {#if isEditable}
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
      <button on:click|preventDefault={createRestartEvent} class="mx-auto mt-8 text-stone-400 hover:text-red-600">
        <svg xmlns="http://www.w3.org/2000/svg" width="2em" height="2em" viewBox="0 0 24 24"><path fill="currentColor" d="M18.258 3.508a.75.75 0 0 1 .463.693v4.243a.75.75 0 0 1-.75.75h-4.243a.75.75 0 0 1-.53-1.28L14.8 6.31a7.25 7.25 0 1 0 4.393 5.783a.75.75 0 0 1 1.488-.187A8.75 8.75 0 1 1 15.93 5.18l1.51-1.51a.75.75 0 0 1 .817-.162"/></svg>
      </button>
      <div class="flex mx-auto mt-12 text-stone-700 font-medium">
        <p>Have an emergency?</p>
        <button class="ml-1 text-red-700 hover:text-red-800" on:click|preventDefault={createEmergency}>
          Alert admins
        </button>
      </div>
    {/if}
  </div>
</div>