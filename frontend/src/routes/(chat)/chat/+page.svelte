<script>
  import ChatThumbnail from '$lib/components/ChatThumbnail.svelte';
  import Plus from '$lib/components/Plus.svelte';

  export let data
  const { chat_rooms } = data
</script>

<div class="h-[90vh] hidden md:flex flex-col items-center justify-center">Select a chat room to get started!</div>

<div class="flex md:hidden fixed h-full bg-white flex-col block md:px-4 w-full">
  <a href="/chat" class="font-semibold text-2xl text-left m-8 md:ml-16 mb-16">Chats</a>
  <div class="mx-8 md:mx-0 flex flex-col items-center">
    {#each chat_rooms as chat_room}
      <ChatThumbnail {chat_room} {data}/>
    {/each}
    <a href="/chat/create-chat" class="text-sm mt-4 font-light hover:bg-red-100 rounded-full">
      <Plus />
    </a>
  </div>
</div>


