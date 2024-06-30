<script lang="ts">
	import { onMount } from "svelte";
  import HeroBanner from "./HeroBanner.svelte";

  export let data;
  const { session } = data;

  let loaded = false

  onMount(() => {
    loaded = true
  })
</script>

<div class={`h-lvh bg-gradient-to-b from-red-100 to-transparent flex flex-row justify-center transition-opacity ease-in duration-700 ${loaded ? "opacity-100" : "opacity-0"}`}>
  <div>
    <div class="flex flex-col items-center mt-72">
      <h1 class="font-bold text-stone-700 text-5xl md:text-7xl mb-4">Open House</h1>
      <h2 class="font-bold text-stone-500 text-4xl md:text-6xl">Bus Tracker</h2>
      {#if session?.Role == "user"}
        <a href="/event" class="ml-6 rounded-full px-6 py-2 text-stone-800 transition-colors duration-200 ease-in-out hover:text-red-500 font-semibold mt-12 bg-white drop-shadow text-sm">Follow Bus</a>
      {:else}
        <a href="/profile/sign-up" class="ml-6 rounded-full px-6 py-2 text-stone-800 transition-colors duration-200 ease-in-out hover:text-red-500 font-semibold mt-12 bg-white drop-shadow-sm text-sm">Get Started</a>
      {/if}
    </div>
    <div class="flex flex-col items-center mt-72">
      <h1 class="text-4xl font-light text-stone-800 mb-2">Supporting Open House Operations</h1>

      <HeroBanner bold="Replacing" text="manual tracking" source="table.jpeg" />
      <HeroBanner bold="Helping" text="student helpers" source="student1.jpeg" />
      <HeroBanner bold="Driving" text="success" source="event.jpeg" />

      <h1 class="text-4xl font-light text-stone-800 mb-2 mt-20 whitespace-pre">Our <span class="font-semibold text-stone-700">impact</span>
      </h1>
    </div>
  </div>
</div>
