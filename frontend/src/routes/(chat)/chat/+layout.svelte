<script lang="ts">
  export let data
  import ChatThumbnail from '$lib/components/ChatThumbnail.svelte';
  import Plus from '$lib/components/Plus.svelte';
	import { onMount } from 'svelte';
  import type { ChatRoom } from '$lib/types/global.js';

  let { account, chat_rooms, env, backend_uri } = data
  
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

    ws.onmessage = async (msg) => {
      let parts = msg.data.split(' ');
      console.log(parts)
      console.log(rooms)
      if (rooms.includes(parts[0])) {
        getRooms()
      }
    };
  })
</script>

<aside class="hidden md:flex fixed h-full bg-white flex-col block px-4 items-center">
  <a href="/chat" class="w-full font-semibold text-2xl text-left m-8 ml-16 mb-16">Chats</a>
  {#each chat_rooms as chat_room}
    <ChatThumbnail {chat_room} {data}/>
  {/each}
  <a href="/chat/create-chat" class="text-sm mb-8 mt-4 font-light p-2 hover:bg-red-100 rounded-full">
    <Plus />
  </a>
</aside>

<main>
  <slot />
</main>



