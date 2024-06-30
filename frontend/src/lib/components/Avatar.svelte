<script lang="ts">
  export let data

  const { session, account } = data

  let menuOpen = false
  let avatar:HTMLElement

</script>

<svelte:window on:click={(e) => {
  if (!avatar.contains(e.target)) {
    menuOpen = false
  }
}}/>

<div>
  <button on:click={() => menuOpen = !menuOpen} bind:this={avatar}>
    <div class={`relative inline-flex items-center justify-center w-8 h-8 overflow-hidden ${session?.Role == "admin" ? "bg-red-700" : "bg-orange-800"} rounded-full`}>
        {#if session?.Role == "admin"}
          <span class="font-medium text-white">A</span>
        {:else}
          <span class="font-medium text-white">SH</span>
        {/if}
    </div>
  </button>
  
  {#if menuOpen}
    <div class="absolute z-10 bg-white divide-y divide-gray-100 rounded-lg shadow w-44 right-0 mr-8 mt-2">
        <div class="px-4 py-3 text-sm text-gray-900">
          <div>{account.Name}</div>
          <div class="font-medium truncate">{account.Email}</div>
          <div>{session.Role}</div>
        </div>
        <ul class="py-2 text-sm text-gray-700 dark:text-gray-200">
          <li>
            <a href="/profile" class="block px-4 py-2 hover:bg-gray-100">Dashboard</a>
          </li>
        </ul>
        <form method="POST" action="/profile?/signOut">
          <button class="w-full py-1">
            <div class="text-left block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">Sign out</div>
          </button>
        </form>
    </div>
  {/if}
</div>