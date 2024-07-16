<script lang="ts">
  import type { Message } from "$lib/types/global";
  export let data
  export let message: Message
  
  const { account } = data
</script>

<div class={`w-full flex flex-row ${message.From == account?.Email ? "justify-end" : "justify-start"}`}>
  <div class={`rounded-md p-4 mb-2 ${message.From == account?.Email ? "bg-red-100" : "bg-gray-100"}`}>
    {message.Body}
  </div>
</div>