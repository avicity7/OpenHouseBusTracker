<script>
  export let data
  import ChatBody from "./ChatBody.svelte";
</script>

<div class={`flex flex-col-reverse overflow-y-auto snap-end row-span-10 md:pl-[24em] px-10`}>
  {#each data.messages as message}
    <ChatBody {data} {message} />
  {/each}
</div>