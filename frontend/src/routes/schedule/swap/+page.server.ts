import { PUBLIC_BACKEND_URL } from "$env/static/public"
import type { EventHelper } from "$lib/types/global"

export const load = async ({locals}) => {
  const backend_uri = PUBLIC_BACKEND_URL

  const response = await fetch(`${backend_uri}:3000/event-helpers/get-available-swaps/${locals.session?.Email}`)
  const helpers = await response.json() as EventHelper[]
  
  return {
    helpers
  }
}