<script lang="ts">
	import { onMount } from "svelte";
  import { page } from "$app/stores";

  export let data

  const { session, account } = data

  let menuOpen = false
  let avatar:HTMLElement

  onMount(() => {
    document.addEventListener('click', (e: MouseEvent) => {
      if (!avatar.contains(e.target as Node)) {
        menuOpen = false
      }
    })
  })

</script>

<div bind:this={avatar}>
  <button on:click={() => menuOpen = !menuOpen}>
    <div class={`relative inline-flex items-center justify-center w-8 h-8 overflow-hidden ${session?.Role == "admin" ? "bg-red-700" : "bg-orange-800"} rounded-full`}>
        {#if session?.Role == "admin"}
          <span class="font-medium text-white">A</span>
        {:else}
          <span class="font-medium text-white">SH</span>
        {/if}
    </div>
  </button>
  
  {#if menuOpen}
    <div class={`absolute z-10 bg-white divide-y divide-gray-100 rounded-lg shadow w-44 right-0 mr-8 ${(session?.Role == "admin" && $page.url.pathname == '/') ? "mt-4" : "mt-2"}`}>
        <div class="px-4 py-3 text-sm text-gray-900">
          <div>{account.Name}</div>
          <div class="font-medium truncate">{account.Email}</div>
          <div>{session.Role}</div>
        </div>
        <ul class="py-2 text-sm text-gray-700 dark:text-gray-200">
          <li>
            <a href="/profile/settings" class="block px-4 py-2 hover:bg-gray-100" on:click={() => menuOpen = false}>Profile Settings</a>
          </li>
        </ul>
        <form method="POST" action="/?/signOut">
          <button class="w-full py-1">
            <div class="text-left block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">Sign out</div>
          </button>
        </form>
    </div>
  {/if}
</div>