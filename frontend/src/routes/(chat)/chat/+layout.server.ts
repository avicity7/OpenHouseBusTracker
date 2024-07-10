import { PUBLIC_BACKEND_URL } from '$env/static/public';
import type { ChatRoom } from '$lib/types/global.js';

export const load = async ({ fetch, locals }) => {
  const email = locals?.session?.Email
  let chat_rooms: Array<ChatRoom>
  if (email != undefined) {
    const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/chat/get-chat-rooms/${email}`)
    chat_rooms = await response.json() as Array<ChatRoom>
    if (!chat_rooms) {
      chat_rooms = []
    }
  } else {
    chat_rooms = []
  }

  return {
    chat_rooms
  }
}