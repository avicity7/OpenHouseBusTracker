import { PUBLIC_BACKEND_URL } from '$env/static/public';
import type { User } from '$lib/types/global.js';
export const load = async ({ fetch }) => {
  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/users/get-users`)
  const users = await response.json() as Array<User> 

  return {
    users
  }
}