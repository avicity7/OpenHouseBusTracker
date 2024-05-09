import { PUBLIC_BACKEND_URL } from "$env/static/public"
import type { Bus } from "../../../../types/global"

export const load = async() => {
  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/bus/get-buses`)
  const buses = await response.json() as Array<Bus>
  
  return {
    buses
  }
}