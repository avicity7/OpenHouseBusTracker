<script lang="ts">
	import type { ChatRoom } from "$lib/types/global";
  import { page } from "$app/stores";
  export let chat_room: ChatRoom
  export let data

  const { account } = data
</script>

<a href={`/chat/${chat_room.RoomId}`} class={`${$page.url.pathname.includes(chat_room.RoomId) ? "bg-red-100" : "hover:bg-gray-100"} m-2 rounded-md px-3 py-2`}>
  {chat_room.User1 == account.Name ? chat_room.User2 : chat_room.User1}
</a>
