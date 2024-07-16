<script>
  import ChatThumbnail from '$lib/components/ChatThumbnail.svelte';
  import Plus from '$lib/components/Plus.svelte';

  export let data
  const { chat_rooms } = data
</script>

<div class="h-[90vh] hidden md:flex flex-col items-center justify-center">Select a chat room to get started!</div>

<div class="flex md:hidden fixed h-full bg-white flex-col block px-4 items-center">
<a href="/chat" class="w-full font-semibold text-2xl text-left m-8 ml-16 mb-16">Chats</a>
{#each chat_rooms as chat_room}
  <ChatThumbnail {chat_room} {data}/>
{/each}
<a href="/chat/create-chat" class="text-sm mb-8 mt-4 font-light p-2 hover:bg-red-100 rounded-full">
  <Plus />
</a>
</div>


