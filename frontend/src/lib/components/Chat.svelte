<script>
  export let data
  import ChatBody from "./ChatBody.svelte";
</script>
{#each data.messages as message}
  <ChatBody {data} {message} />
{/each}