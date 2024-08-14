<script lang="ts">
	import type { ChatRoom } from "$lib/types/global";
  import { PUBLIC_BACKEND_URL } from "$env/static/public";
  import { page } from "$app/stores";
	import Trash from "./Trash.svelte";
  export let chat_room: ChatRoom

  const formatTimestamp = (timestamp: string) => {
    const utcDate = new Date(timestamp);
    const formattedDate = utcDate.toLocaleString();
    return formattedDate;
  }

  const deleteRoom = async () => {
    await fetch(`${PUBLIC_BACKEND_URL}:3000/chat/delete-room/${chat_room.RoomId}`, {
      method: "DELETE",
      headers: {
        'content-type': 'application/json'
      },
    })
  }
</script>

<a href={`/chat/${chat_room.RoomId}`} class={`group ${$page.url.pathname.includes(chat_room.RoomId) ? "bg-red-100" : "hover:bg-stone-100"} rounded-md p-4 w-80`}>
  <div class="mb-2 flex flex-row justify-between items-center">
    <div class="font-medium">
      {chat_room.RoomName != "" ? chat_room.RoomName : chat_room.Name}
    </div>
    {#if chat_room.LatestMessage.Body != ''}
      <div class="text-xs">
        {formatTimestamp(chat_room.LatestMessage.Timestamp)}
      </div>
    {/if}
  </div>

  <div class="w-full flex justify-between">
    <div class="text-sm flex flex-row">
      <div class="mr-1">
        {chat_room.LatestMessage.FromName + (chat_room.LatestMessage.FromName ? ":" : "")}
      </div>
      <div class="text-ellipsis overflow-hidden text-nowrap">
        {chat_room.LatestMessage.Body}
      </div>
    </div>
    <button class="invisible group-hover:visible text-gray-400 hover:text-red-400" on:click={deleteRoom}>
      <Trash />
    </button>
  </div>
</a>
