<script lang="ts">
  import type { SwapRequest } from "$lib/types/global"
  import BusScheduleCard from "$lib/components/BusScheduleCard.svelte"

  export let data;

  const { currentSchedules = [], futureSchedules = [], swapRequests, account, backend_uri } = data;

  const acceptRequest = async (sr: SwapRequest) => {
    await fetch(`${backend_uri}:3000/event-helpers/accept-swap-request`, {
      method: "PUT",
      body: JSON.stringify(sr)
    })

    location.reload()
  }

  const deleteRequest = async (sr: SwapRequest) => {
    await fetch(`${backend_uri}:3000/event-helpers/delete-swap-request`, {
      method: "DELETE",
      body: JSON.stringify(sr)
    })

    location.reload()
  }
</script>

<svelte:head>
	<title>My Shifts | SPOH Bus Tracker</title>
</svelte:head>

<div class="p-6 md:p-12">
  <h1 class="text-3xl font-semibold pt-8 mb-6">My Shifts</h1>

  <h1 class="text-xl font-semibold text-stone-400 pt-8">Current Shifts</h1>
  {#if currentSchedules && currentSchedules.length > 0}
    <div class="md:grid grid-cols-3">
      {#each currentSchedules as bus}
        <BusScheduleCard {bus} />
      {/each}
    </div>
  {:else}
    <p class="text-gray-500">No current schedules</p>
  {/if}

  <hr class="my-8 border-gray-200">

  <h1 class="text-xl font-semibold text-stone-400 pt-8">Upcoming Shifts</h1>
  {#if futureSchedules && futureSchedules.length > 0}
    <div class="md:grid grid-cols-3">
      {#each futureSchedules as bus}
        <BusScheduleCard {bus} />
      {/each}
    </div>
  {:else}
    <p class="text-gray-500">No upcoming schedules</p>
  {/if}

  <div class="flex flex-row items-center align-center my-12">
    <h1 class="text-3xl font-semibold">Shift Swap Requests</h1>

    <a href="/schedule/swap" class="ml-4 text-neutral-400 hover:text-red-600">
      + Make a request
    </a>
  </div>

  {#if swapRequests}
    {#each swapRequests as swapRequest}
      {#if swapRequest.From == account?.Email}
        <div class="bg-white shadow-md max-w-md flex flex-col items-center">
          <h3 class="text-neutral-400 my-8">Outgoing request to</h3>
          <h1 class="text-2xl font-semibold">{swapRequest.WithName}</h1>
          <h2 class="text-xl font-medium mt-8">for {swapRequest.TargetShift ? "an AM" : "a PM"} shift</h2>
          <button class="px-2 py-1 bg-red-700 hover:bg-red-800 text-white rounded my-8" on:click={() => deleteRequest(swapRequest)}>
            Cancel request
          </button>
        </div>
      {:else}
        <div class="bg-white shadow-md max-w-md flex flex-col items-center">
          <h3 class="text-neutral-400 my-8">Incoming request from</h3>
          <h1 class="text-2xl font-semibold">{swapRequest.FromName}</h1>
          <h2 class="text-xl font-medium mt-8">for your {swapRequest.TargetShift ? " AM" : "PM"} shift</h2>
          <div class="flex flex-row items-center">
            <button class="px-3 py-2 bg-red-700 hover:bg-red-800 text-white rounded my-8 mx-2" on:click={() => deleteRequest(swapRequest)}>
              Cancel request
            </button>
            <button class="px-3 py-2 bg-green-800 hover:bg-green-900 text-white rounded my-8 mx-2" on:click={() => acceptRequest(swapRequest)}>
              Accept request
            </button>
          </div>
        </div>
      {/if}
    {/each}
  {/if}
</div>