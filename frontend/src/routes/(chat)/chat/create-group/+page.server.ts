import { PUBLIC_BACKEND_URL } from '$env/static/public';
import type { User } from '$lib/types/global.js';
export const actions = {
  createGroupChat: async({ request, fetch, locals}) => {
    const form = await request.formData()

    const SelectedUsers = form.getAll('name');
    const Groupname = form.get("groupname")
    SelectedUsers.push(locals.session?.Email as FormDataEntryValue)
    
		await fetch(`${PUBLIC_BACKEND_URL}:3000/chat/create-chat-room`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ SelectedUsers, Groupname })
		});
  }
}
export const load = async ({ fetch }) => {
  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/users/get-users`)
  const users = await response.json() as Array<User> 

  return {
    users
  }
}
