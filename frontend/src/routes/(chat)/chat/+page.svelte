<script lang="ts">
  import ChatThumbnail from '$lib/components/ChatThumbnail.svelte';
  import Plus from '$lib/components/Plus.svelte';
	import { onMount } from 'svelte';
  import type { ChatRoom } from '$lib/types/global.js';
	import { onNavigate } from '$app/navigation';

  export let data
  let { chat_rooms, backend_uri, account, env } = data

  let button: HTMLElement
  let menu = false

  const getRooms = async () => {
    const response = await fetch(`${backend_uri}:3000/chat/get-chat-rooms/${account?.Email}`)
    chat_rooms = await response.json() as Array<ChatRoom>
  }

  onMount(() => {
    const ws = new WebSocket(`${env == 'PROD' ? 'wss' : 'ws'}://${backend_uri.split('//')[1]}:3000/ws`);
    document.addEventListener('click', (e: MouseEvent) => {
      if (!button.contains(e.target as Node)) {
        menu = false
      }
    })

    ws.onmessage = async (msg) => {
      let refresh = false
      let parts = msg.data.split(' ');
      if (parts[0] == account?.Email) {
        await getRooms()
        refresh = true
      }
      for (let i = 0; i < chat_rooms.length; i++) {
        const room = chat_rooms[i]
        if (room.RoomId == parts[0] && !refresh) {
          await getRooms()
          break
        } 
      }
    };
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


