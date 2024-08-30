<script lang="ts">
  export let data
  import ChatBody from "./ChatBody.svelte";
  import type { ChatRoom } from "$lib/types/global";
	import { onMount } from "svelte";

  const { chat_rooms }: { chat_rooms: Array<ChatRoom> } = data
  let chat_room: ChatRoom
  const findRoom = () => {
    for (let i = 0; i < chat_rooms.length; i++) {
      let room = chat_rooms[i]
      if (room.RoomId == data.room_id) {
        chat_room = room
      }
    }
  }  

  onMount(() => {
    findRoom()
  })
</script>

{#if chat_room}
<div class="bg-white border border-neutral-200 border-x-white border-t-white relative p-4 md:px-6 px-10 flex flex-row h-fit row-span-10">
  <p>
    {chat_room.RoomName ? chat_room.RoomName : chat_room.Name}
  </p>

</div>
<div class={`flex flex-col-reverse overflow-y-auto snap-end row-span-10 px-10 pt-16`}>
  {#each data.messages as message}
    <ChatBody {data} {message} />
  {/each}
</div>
{/if}