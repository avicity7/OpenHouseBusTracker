import type { Message } from '$lib/types/global.js'
import { PUBLIC_BACKEND_URL } from '$env/static/public'

export const load = async (ctx) => {
  const fetch = ctx.fetch
  const room_id = ctx.params.room_id

  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/chat/get-messages/${room_id}`)

  let messages = await response.json() as Array<Message>
  if (!messages) {
    messages = [] as Array<Message>
  }

  return {
    messages,
    room_id
  }
}