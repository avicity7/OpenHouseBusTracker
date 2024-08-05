<script lang="ts">
  import ChatThumbnail from '$lib/components/ChatThumbnail.svelte';
  import Plus from '$lib/components/Plus.svelte';
	import { onMount } from 'svelte';

  export let data
  const { chat_rooms } = data

  let button: HTMLElement
  let menu = false

  onMount(() => {
    document.addEventListener('click', (e: MouseEvent) => {
      if (!button.contains(e.target as Node)) {
        menu = false
      }
    })
  })
</script>

<svelte:head>
	<title>Chat | SPOH Bus Tracker</title>
</svelte:head>

<div class="h-[90vh] hidden md:flex flex-col items-center justify-center">Select a chat room to get started!</div>

<div class="block md:hidden flex flex-row justify-between min-w-full items-center p-8">
  <a href="/chat" class="w-full font-semibold text-2xl text-left">Chats</a>
  <div class="flex flex-col">
    <button class="text-sm font-light hover:bg-red-100 rounded-full p-2" on:click={() => menu = !menu} bind:this={button}>
      <Plus />
    </button>
    {#if menu}
      <div class="bg-white absolute right-0 mr-10 mt-10 w-[10em] shadow-lg flex flex-col items-center rounded-lg overflow-hidden text-sm font-medium">
        <a href="/chat/create-chat" class="py-4 hover:bg-red-100 flex w-full justify-center">
          Create chat
        </a>
        <a href="/chat/create-group" class="py-4 hover:bg-red-100 flex w-full justify-center">
          Create group
        </a>
      </div>
    {/if}
  </div>
</div>
<div class="px-8">
  {#each chat_rooms as chat_room}
    <ChatThumbnail {chat_room} />
  {/each}
</div>


