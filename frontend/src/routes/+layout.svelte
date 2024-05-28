<script lang="ts">
  import '../app.css'
  import '../../node_modules/mapbox-gl/dist/mapbox-gl.css';
  import { page } from '$app/stores';  
  import { onMount } from 'svelte';
  import type { User } from '$lib/types/global';
  import { injectSpeedInsights } from '@vercel/speed-insights/sveltekit';
  export let data
  let { session, backend_uri, account } = data
  let menu = false

  injectSpeedInsights();

  const detectColorScheme = () => {
    let theme = "light"

    if (localStorage.getItem("theme") == "dark"){
      theme = "dark"
    }

    if (theme == "dark") {
      document.documentElement.setAttribute("class", "dark");
    } else {
      document.documentElement.setAttribute("class", "light");
    }
  }

  onMount(() => {
    detectColorScheme()
  })
</script>

<nav class="z-10 text-md p-3 px-6 justify-between flex flex-row items-center sticky top-0 bg-white border-0 border-b-2 border-stone-200">
  <a href="/" class="text-red-600 font-bold text-2xl">
    SP
  </a>
  <button class="block md:hidden" on:click={() => {menu = !menu}}>
    <svg xmlns="http://www.w3.org/2000/svg" width="2em" height="2em" viewBox="0 0 24 24" {...$$props}><path fill="currentColor" d="M4 18h16c.55 0 1-.45 1-1s-.45-1-1-1H4c-.55 0-1 .45-1 1s.45 1 1 1m0-5h16c.55 0 1-.45 1-1s-.45-1-1-1H4c-.55 0-1 .45-1 1s.45 1 1 1M3 7c0 .55.45 1 1 1h16c.55 0 1-.45 1-1s-.45-1-1-1H4c-.55 0-1 .45-1 1"/></svg>
  </button>
  <div class="hidden md:block flex flex-row items-center">
    {#if session?.Role == "admin"}
      <a href="/bus-routes" class={"font-medium "+($page.url.pathname == '/bus-routes' ? "text-red-700" : "hover:text-red-700")}>Routes</a>
    {/if}
    {#if session?.Role == "admin"}
      <a href="/admin/users" class={"ml-6 font-medium "+($page.url.pathname.includes('admin') ? "text-red-700" : "hover:text-red-700")}>Manage</a>
      <!-- <a href="/admin/users" class={"ml-6 font-medium "+($page.url.pathname == '/admin/users' ? "text-red-700" : "hover:text-red-700")}>Users</a>
      <a href="/admin/schedule" class={"ml-6 font-medium "+($page.url.pathname == '/admin/schedule' ? "text-red-700" : "hover:text-red-700")}>Schedule</a>
      <a href="/admin/drivers" class={"ml-6 font-medium "+($page.url.pathname == '/admin/drivers' ? "text-red-700" : "hover:text-red-700")}>Drivers</a>
      <a href="/admin/buses" class={"ml-6 font-medium "+($page.url.pathname == '/admin/buses' ? "text-red-700" : "hover:text-red-700")}>Buses</a> -->
    {/if}
    <!-- should it be viewable by everyone even if they're not signed up? -->
  <a href="/profile" class={"ml-6 font-medium "+($page.url.pathname == '/profile' ? "text-red-700" : "hover:text-red-700")}>Profile</a>
    {#if session?.Role == "user"}
    <a href="/schedule" class={"ml-6 font-medium "+($page.url.pathname == '/schedule' ? "text-red-700" : "hover:text-red-700")}>Bus Schedule</a>
    <a href="/event" class="ml-6 bg-red-700 hover:bg-red-800 rounded-full px-3 py-1 text-white ">Follow Bus</a>
    {/if}
  </div>
</nav>
{#if session && account.VerificationToken != "" && !$page.url.pathname.includes("verify")}
  <div class="flex flex-col text-center items-center my-auto mx-8">
    <h1 class="mt-16 text-3xl font-semibold">You're signed up!</h1>
    <p class="mt-8 text-xl">But before you use the Tracker, you need to verify your email.</p>
  </div>
{:else}
  {#if menu}
    <div class="w-screen h-screen flex z-10 p-12 text-3xl flex-col items-end">
      <button class="mb-6" on:click={() => { menu = false }}>
        <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 10 24"><path fill="currentColor" d="m12 13.4l-2.917 2.925q-.277.275-.704.275t-.704-.275q-.275-.275-.275-.7t.275-.7L10.6 12L7.675 9.108Q7.4 8.831 7.4 8.404t.275-.704q.275-.275.7-.275t.7.275L12 10.625L14.892 7.7q.277-.275.704-.275t.704.275q.3.3.3.713t-.3.687L13.375 12l2.925 2.917q.275.277.275.704t-.275.704q-.3.3-.712.3t-.688-.3z"/></svg>
      </button>
      <button class="flex flex-col items-end" on:click={() => { menu = false }}>
        <a href="/bus-routes" class={"mb-5 font-medium "+($page.url.pathname == '/bus-routes' ? "text-red-700" : "hover:text-red-700")}>Routes</a>
        <!-- might not need routes  -->
        {#if session?.Role == "admin"}
          <a href="/admin/users" class={"mb-5 font-medium "+($page.url.pathname == '/admin/users' ? "text-red-700" : "hover:text-red-700")}>Manage</a>
          <!-- <a href="/admin/schedule" class={"mb-5 font-medium "+($page.url.pathname == '/admin/schedule' ? "text-red-700" : "hover:text-red-700")}>Schedule</a> -->
        {/if}
        <a href="/profile" class={"mb-5 font-medium "+($page.url.pathname == '/profile' ? "text-red-700" : "hover:text-red-700")}>Profile</a>
        {#if session?.Role == "user"}
          <a href="/schedule" class={"mb-5 font-medium "+($page.url.pathname == '/schedule' ? "text-red-700" : "hover:text-red-700")}>Bus Schedule</a>
          <a href="/event" class="bg-red-700 hover:bg-red-800 rounded-full px-3 py-1 text-white ">Follow Bus</a>
        {/if}
      </button>
    </div>
  {/if}
  <div class={menu ? "hidden" : "block"}>
    <slot />
  </div>
{/if}
