<script lang="ts">
  export let data
  import ChatThumbnail from '$lib/components/ChatThumbnail.svelte';
  import Plus from '$lib/components/Plus.svelte';

  const { account, chat_rooms } = data
</script>

<aside class="fixed h-full bg-white flex flex-col block px-4 items-center">
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



