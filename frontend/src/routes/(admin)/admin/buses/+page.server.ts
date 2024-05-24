import { PUBLIC_BACKEND_URL } from "$env/static/public"
import type { Bus } from "$lib/types/global"

export const load = async() => {
  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/bus/get-buses`)
  let buses: Array<Bus> = []
  if (response) {
    buses = await response.json() as Array<Bus>
  }
  
  return {
    buses
  }
}