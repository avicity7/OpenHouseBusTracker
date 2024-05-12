import { PUBLIC_BACKEND_URL } from "$env/static/public"
import type { Route } from "../../../types/global"

export const load = async () => {

  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/routes/`)
  const routes = await response.json() as Array<Route>

  return {
    routes
  }
}