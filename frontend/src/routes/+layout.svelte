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


<nav class={`z-10 text-sm justify-between flex flex-row items-center ${white ? "bg-white/40 backdrop-blur" : ""} ${$page.url.pathname == '/' ? 'w-screen fixed pl-2 md:pl-6 pr-6 py-6' : 'sticky p-2 px-4 bg-white/40 backdrop-blur'} top-0 transition-all ease-out`}>
  <div class={`flex flex-row items-center ${(session?.Role == "admin" && $page.url.pathname == '/') ? "p-2 px-4 rounded bg-white" : "px-6"}`}>
    <a href="/" class="flex flex-row items-center">
      <div class={`text-red-600 font-bold text-2xl mr-4`}>
        SP
      </div>
      <span class="font-semibold font-stone-800 text-md hidden md:block">Bus Tracker</span>
    </a>
  </div>
  <div class={`flex flex-row items-center ${session?.Role == "admin" && $page.url.pathname == '/' ? "bg-white p-2 rounded" : ""}`}>
    {#if session?.Role == "admin"}
      <div class="flex flex-row items-center">
        <a href="/admin/users" class={"ml-2 font-medium hidden md:block "+($page.url.pathname.includes('admin') ? "text-red-700" : "hover:text-red-700")}>Manage</a>
        <a href="/chat" class={"mx-8 font-medium hidden md:block "+($page.url.pathname.includes('chat') ? "text-red-700" : "hover:text-red-700")}>Chat</a>
        <Avatar {data} />
      </div>
    {/if}

    {#if session?.Role == "user"}
      <div class="flex flex-row items-center">
        <a href="/bus-routes" class={"font-medium mr-8 hidden md:block "+($page.url.pathname.includes('bus-routes') ? "text-red-700" : "hover:text-red-700")}>Routes</a>
        <a href="/schedule" class={"font-medium hidden md:block "+($page.url.pathname.includes('schedule') ? "text-red-700" : "hover:text-red-700")}>My Shifts</a>
        <a href="/chat" class={"mx-8 font-medium hidden md:block "+($page.url.pathname.includes('chat') ? "text-red-700" : "hover:text-red-700")}>Chat</a>
        <Avatar {data} />
      </div>
    {/if}

    {#if !session}
      <div class="flex flex-row items-center">
        <a href="/bus-routes" class={"font-medium mr-8 "+($page.url.pathname.includes('bus-routes') ? "text-red-700" : "hover:text-red-700")}>Routes</a>
        <a href="/help" class={"font-medium "+($page.url.pathname.includes('/help') ? "text-red-700" : "hover:text-red-700")}>Help</a>
        <a href="/profile" class={"mx-8 font-medium "+($page.url.pathname.includes('/profile') ? "text-red-700" : "hover:text-red-700")}>Login</a>
      </div>
    {/if}
  </div>
</nav>
{#if session && account?.VerificationToken != "" && !$page.url.pathname.includes("verify") && !account?.VerificationToken.startsWith("reset")}
  <div class="flex flex-col text-center items-center my-auto mx-8">
    <h1 class="mt-16 text-3xl font-semibold">You're signed up!</h1>
    <p class="mt-8 text-xl">But before you use the Tracker, you need to verify your email.</p>
  </div>
{:else if session && account?.Role == "public"}
  <div class="flex flex-col text-center items-center my-auto mx-8">
    <h1 class="mt-16 text-3xl font-semibold">You've been registered!</h1>
    <p class="mt-8 text-xl">Before you can access the rest of the platform, ask an admin to verify you!</p>
    <p class="mt-8 text-xl">Once they've done so, you can refresh this page.</p>
  </div>
{:else}
  <slot />
{/if}
