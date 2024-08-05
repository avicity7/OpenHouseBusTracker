<script lang="ts">
  export let data
  import ChatThumbnail from '$lib/components/ChatThumbnail.svelte';
  import Plus from '$lib/components/Plus.svelte';
	import { onMount } from 'svelte';
  import type { ChatRoom } from '$lib/types/global.js';

  let { account, chat_rooms, env, backend_uri } = data

  let menu = false
  let button:HTMLElement
  
  let rooms: Array<string> = []
  chat_rooms.forEach(room => {
    rooms.push(room.RoomId)
  })

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
      let parts = msg.data.split(' ');
      if (rooms.includes(parts[0])) {
        getRooms()
      }
    };
  })
</script>

<aside class="hidden md:flex fixed h-full bg-white flex-col block px-4 items-center min-w-[20em]">
  <div class="flex flex-row justify-between min-w-full items-center p-8">
    <a href="/chat" class="w-full font-semibold text-2xl text-left">Chats</a>
    <div class="flex flex-col">
      <button class="text-sm font-light hover:bg-red-100 rounded-full p-2" on:click={() => menu = !menu} bind:this={button}>
        <Plus />
      </button>
      {#if menu}
        <div class="bg-white absolute mt-10 w-[10em] shadow-lg flex flex-col items-center rounded-lg overflow-hidden text-sm font-medium">
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
  {#each chat_rooms as chat_room}
    <ChatThumbnail {chat_room} />
  {/each}
</aside>

<main>
  <slot />
</main>



