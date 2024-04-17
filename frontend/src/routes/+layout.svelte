<script>
  import '../app.css'
</script>

<nav class="bg-white p-4 px-6 justify-between flex flex-row items-center">
  <div class="text-red-600 font-bold text-2xl">
    SP
  </div>
  <button class="block md:hidden">
    <svg xmlns="http://www.w3.org/2000/svg" width="2em" height="2em" viewBox="0 0 24 24" {...$$props}><path fill="currentColor" d="M4 18h16c.55 0 1-.45 1-1s-.45-1-1-1H4c-.55 0-1 .45-1 1s.45 1 1 1m0-5h16c.55 0 1-.45 1-1s-.45-1-1-1H4c-.55 0-1 .45-1 1s.45 1 1 1M3 7c0 .55.45 1 1 1h16c.55 0 1-.45 1-1s-.45-1-1-1H4c-.55 0-1 .45-1 1"/></svg>
  </button>
  <div class="hidden md:block">
    <div class="flex flex-row">
      <a href="/routes" class="hidden text-lg font-medium hover:text-red-600 md:block">Routes</a>
      <a href="/account" class="ml-6 hidden text-lg font-medium hover:text-red-600 md:block">Account</a>
    </div>
  </div>
</nav>
<slot />