<script>
  import '../app.css'
  import '../../node_modules/mapbox-gl/dist/mapbox-gl.css';
  import { onMount } from 'svelte';
  export let data
  let { session } = data;
  import { injectSpeedInsights } from '@vercel/speed-insights/sveltekit';

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

<nav class="bg-white p-4 px-6 justify-between flex flex-row items-center">
  <a href="/" class="text-red-600 font-bold text-2xl">
    SP
  </a>
  <button class="block md:hidden">
    <svg xmlns="http://www.w3.org/2000/svg" width="2em" height="2em" viewBox="0 0 24 24" {...$$props}><path fill="currentColor" d="M4 18h16c.55 0 1-.45 1-1s-.45-1-1-1H4c-.55 0-1 .45-1 1s.45 1 1 1m0-5h16c.55 0 1-.45 1-1s-.45-1-1-1H4c-.55 0-1 .45-1 1s.45 1 1 1M3 7c0 .55.45 1 1 1h16c.55 0 1-.45 1-1s-.45-1-1-1H4c-.55 0-1 .45-1 1"/></svg>
  </button>
  <div class="hidden md:block flex flex-row items-center">
    <a href="/bus-routes" class="text-lg font-medium hover:text-red-600">Routes</a>
    {#if session?.Role == "admin"}
      <a href="/admin/users" class="ml-6 text-lg font-medium hover:text-red-600">Users</a>
      <!-- <a href="/admin/schedule" class="ml-6 text-lg font-medium hover:text-red-600">Schedule</a> -->
    {/if}
    <a href="/profile" class="ml-6 text-lg font-medium hover:text-red-600">Profile</a>
    {#if session?.Role == "user"}
      <a href="/event" class="ml-6 bg-red-700 hover:bg-red-800 rounded-full px-3 py-1 text-white ">Follow Bus</a>
    {/if}
  </div>
</nav>
<slot />