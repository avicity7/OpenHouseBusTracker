<script lang="ts">
  import { afterNavigate } from '$app/navigation';
	import Chat from '$lib/components/Chat.svelte';
	import Send from '$lib/components/Send.svelte';
  import type { Message } from '$lib/types/global.js';
	import { onMount } from 'svelte';
  export let data

  const { account, env, backend_uri } = data

  let body: string
  let ws: WebSocket;
  let form: HTMLFormElement;

  const getMessages = async () => {
    const response = await fetch(`${backend_uri}:3000/chat/get-messages/${data.room_id}`)
    data.messages = await response.json() as Array<Message>
    if (!data.messages) data.messages = []
  }

  const createMessage = async () => {
    await fetch(`${backend_uri}:3000/chat/create-message`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ From:  account?.Email, RoomId: data.room_id, Body: body })
    });
    body = ""
  }

  onMount(() => {
    getMessages()
  })

  afterNavigate(() => {
    body = ""
    ws = new WebSocket(`${env == 'PROD' ? 'wss' : 'ws'}://${backend_uri.split('//')[1]}:3000/ws`);

    ws.onmessage = async (msg) => {
      let parts = msg.data.split(' ');
      if (parts[0] == data.room_id) {
        getMessages()
      } else if (parts[0] == data.room_id + "del") {
        location.replace("/chat")
      }
    };
  })
</script>

<div class="w-full h-[90vh] grid grid-rows-10 md:pl-[22em]">
  <Chat {data} />

  <form class="row-span-1 pl-4 grid grid-cols-10 gap-4 items-center mt-8" on:submit={createMessage} bind:this={form}>
    <div class="md:col-span-1"></div>
    <textarea 
      class="bg-gray-100 p-4 rounded-lg w-full col-span-7 resize-none select-none focus:outline-none"
      bind:value={body}
      placeholder="Type a message"
      on:keypress={(e) => {
        if (e.key == 'Enter' && !e.shiftKey && body != "") {
          e.preventDefault()
          form.requestSubmit()
        }
      }}
    />
    <div class="col-span-1">
      {#if body != ""}
      <button class="text-xl" type="submit">
        <Send />
      </button>
      {/if}
    </div>
    <div class="md:col-span-1"></div>
  </form>
</div>