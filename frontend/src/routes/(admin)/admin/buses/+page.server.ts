import { PUBLIC_BACKEND_URL } from "$env/static/public"
import type { EventBus, Route } from "$lib/types/global"

export const load = async() => {
  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/bus/get-buses`)
  let buses: Array<EventBus> = []
  const parsed = await response.json() as Array<EventBus>
  if (parsed) {
    buses = parsed
  }

  const routeResponse = await fetch(`${PUBLIC_BACKEND_URL}:3000/route`);
  if (!routeResponse.ok) {
      throw new Error("Failed to fetch routes");
  }

  const routes = await routeResponse.json() as Route[];
  
  return {
    buses,
    routes
  }
}