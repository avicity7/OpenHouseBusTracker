<script lang="ts">
  import '../app.css'
  import '../../node_modules/mapbox-gl/dist/mapbox-gl.css';
  import { page } from '$app/stores';  
  import { onMount } from 'svelte';
  import { injectSpeedInsights } from '@vercel/speed-insights/sveltekit';
  import Avatar from '$lib/components/Avatar.svelte';
  export let data
  let { session, account } = data
  let menu = false,
  white = false,
  offset = 500

  injectSpeedInsights();

  const detectColorScheme = () => {
    let theme = "light"
    if (localStorage.getItem("theme") == "dark") theme = "dark"
    if (theme == "dark") document.documentElement.setAttribute("class", "dark");
    else document.documentElement.setAttribute("class", "light");
  }

  onMount(() => {
    detectColorScheme()
    if (screen.width < 500) offset = 200
    document.addEventListener("scroll", () => { window.scrollY > offset ? white = true : white = false })
  })
</script>


<nav class={`z-10 text-sm justify-between flex flex-row items-center ${white ? "bg-white/40 backdrop-blur" : ""} ${$page.url.pathname == '/' ? 'w-screen fixed p-6' : 'sticky p-2 px-4 bg-white/40 backdrop-blur'} top-0 transition-all ease-out`}>
  <div class={`flex flex-row items-center ${(session?.Role == "admin" && $page.url.pathname == '/') ? "p-2 px-4 rounded bg-white" : "px-6"}`}>
    <a href="/" class="flex flex-row items-center">
      <div class={`text-red-600 font-bold text-2xl mr-4`}>
        SP
      </div>
      <span class="font-semibold font-stone-800 text-md">Bus Tracker</span>
    </a>
  </div>
  <button class="block md:hidden" on:click={() => {menu = !menu}}>
    <svg xmlns="http://www.w3.org/2000/svg" width="2em" height="2em" viewBox="0 0 24 24" {...$$props}><path fill="currentColor" d="M4 18h16c.55 0 1-.45 1-1s-.45-1-1-1H4c-.55 0-1 .45-1 1s.45 1 1 1m0-5h16c.55 0 1-.45 1-1s-.45-1-1-1H4c-.55 0-1 .45-1 1s.45 1 1 1M3 7c0 .55.45 1 1 1h16c.55 0 1-.45 1-1s-.45-1-1-1H4c-.55 0-1 .45-1 1"/></svg>
  </button>
  <div class={`hidden md:block flex flex-row items-center px-6 ${session?.Role == "admin" && $page.url.pathname == '/' ? "bg-white p-2 rounded" : ""}`}>
    {#if session?.Role == "admin"}
      <div class="flex flex-row items-center">
        <a href="/admin/users" class={"font-medium "+($page.url.pathname.includes('admin') ? "text-red-700" : "hover:text-red-700")}>Manage</a>
        <a href="/chat" class={"mx-8 font-medium "+($page.url.pathname == '/chat' ? "text-red-700" : "hover:text-red-700")}>Chat</a>
        <Avatar {data} />
      </div>
    {/if}

    {#if session?.Role == "user"}
      <div class="flex flex-row items-center">
        <a href="/schedule" class={"font-medium "+($page.url.pathname.includes('schedule') ? "text-red-700" : "hover:text-red-700")}>My Shifts</a>
        <a href="/chat" class={"mx-8 font-medium "+($page.url.pathname == '/chat' ? "text-red-700" : "hover:text-red-700")}>Chat</a>
        <Avatar {data} />
      </div>
    {/if}

    {#if !session}
      <a href="/help" class={"ml-6 font-medium "+($page.url.pathname == '/help' ? "text-red-700" : "hover:text-red-700")}>Help</a>
      <a href="/profile" class={"ml-6 font-medium "+($page.url.pathname == '/profile' ? "text-red-700" : "hover:text-red-700")}>Login</a>
    {/if}
  </div>
</nav>
{#if session && account?.VerificationToken != "" && !$page.url.pathname.includes("verify") && !account?.VerificationToken.startsWith("reset")}
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
          <a href="/admin/users" class={"mb-5 font-medium "+($page.url.pathname == '/admin/users' ? "text-red-700" : "hover:text-red-700")} data-testid="manage">Manage</a>
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
  <div class={`${menu ? "hidden" : "block"}`}>
    <slot />
  </div>
{/if}
