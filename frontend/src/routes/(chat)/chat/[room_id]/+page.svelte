<script lang="ts">
  import { afterNavigate, beforeNavigate } from '$app/navigation';
  import ChatBody from '$lib/components/ChatBody.svelte';
  import type { Message } from '$lib/types/global.js';
  export let data

  const { account, env, backend_uri } = data

  let body: string
  let ws: WebSocket;
  let form: HTMLFormElement;
  let chat: HTMLDivElement;
  let display = false;

  const getMessages = async () => {
    const response = await fetch(`${backend_uri}:3000/chat/get-messages/${data.room_id}`)
    data.messages = await response.json() as Array<Message>
  }

  const createMessage = async () => {
    await fetch(`${backend_uri}:3000/chat/create-message`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ From:  account?.Email, RoomId: data.room_id, Body: body })
    });
    getMessages()
  }

  afterNavigate(() => {
    chat.scrollTop = chat.scrollHeight
    display = true

    ws = new WebSocket(`${env == 'PROD' ? 'wss' : 'ws'}://${backend_uri.split('//')[1]}:3000/ws`);

    ws.onmessage = async (msg) => {
      let parts = msg.data.split(' ');
      if (parts[0] == data.room_id) {
        window.location.reload()
      }
    };
  })
</script>

<div class="w-full h-[90vh] grid grid-rows-10">
  <div class={`overflow-auto snap-end row-span-9 ${display ? "opacity-100" : "opacity-0"}`} bind:this={chat}>
    {#each data.messages as message}
      <ChatBody {data} {message} />
    {/each}
  </div>

  <form class="row-span-1 w-[80vw] pl-4 grid grid-cols-10 gap-4 items-center" on:submit={createMessage} bind:this={form}>
    <div class="col-span-3"></div>
    <textarea 
      class="bg-gray-200 hover:bg-gray-200 p-4 rounded-lg w-full col-span-4 resize-y"
      bind:value={body}
      placeholder="Type a message"
      on:keypress={(e) => {
        if (e.key == 'Enter' && !e.shiftKey) {
          e.preventDefault()
          form.requestSubmit()
        }
      }}
    />
    <div class="col-span-1">
      <button class="w-full bg-red-700 hover:bg-red-800 text-white py-2 px-1 rounded" type="submit">
        Send message
      </button>
    </div>
    <div class="col-span-2"></div>
  </form>
</div>