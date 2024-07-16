<script>
  export let data
  import ChatBody from "./ChatBody.svelte";
</script>

<div class={`flex flex-col-reverse overflow-auto snap-end row-span-9 w-screen pl-[24em] pr-6`}>
  {#each data.messages as message}
    <ChatBody {data} {message} />
  {/each}
</div>