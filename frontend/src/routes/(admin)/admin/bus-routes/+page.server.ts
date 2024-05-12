import { PUBLIC_BACKEND_URL } from "$env/static/public"
import type { Route, RouteStep } from "$lib/types/global"

export const load = async () => {

  let response = await fetch(`${PUBLIC_BACKEND_URL}:3000/route/`)
  const routes = await response.json() as Array<Route>
  response = await fetch(`${PUBLIC_BACKEND_URL}:3000/route-step/`)
  const routeSteps = await response.json() as Array<RouteStep>

  return {
    routes,
    routeSteps
  }
}