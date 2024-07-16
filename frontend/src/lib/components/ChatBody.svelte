<script lang="ts">
  import type { Message } from "$lib/types/global";
  import { PUBLIC_BACKEND_URL } from "$env/static/public";
  import Trash from "./Trash.svelte";
  export let data
  export let message: Message
  
  const { account } = data

  const deleteMessage = async () => {
    await fetch(`${PUBLIC_BACKEND_URL}:3000/chat/delete-message`, {
      method: "DELETE",
      body: JSON.stringify(message),
      headers: {
        'content-type': 'application/json'
      },
    })
  }
</script>

<div class={`w-full group flex ${message.From == account?.Email ? "flex-row-reverse" : "flex-row"} items-center`}>
  <p class={`max-w-full rounded-md p-4 mb-2 ${message.From == account?.Email ? "bg-red-100" : "bg-gray-100"} break-words text-wrap overflow-wrap`}>
    {message.Body}
  </p>
  <button class="invisible group-hover:visible mx-2 text-gray-200 hover:text-red-400" on:click={deleteMessage}>
    <Trash /> 
  </button>
</div>