import { PUBLIC_BACKEND_URL } from "$env/static/public"
import type { Route, RouteStep } from "$lib/types/global"

export const load = async () => {

  let response = await fetch(`${PUBLIC_BACKEND_URL}:3000/route/`)
  let routes = await response.json() as Array<Route>
  if (!routes) {
    routes = []
  }
  response = await fetch(`${PUBLIC_BACKEND_URL}:3000/route-step/`)
  let routeSteps = await response.json() as Array<RouteStep>
  if (!routeSteps) {
    routeSteps = []
  }

  return {
    routes,
    routeSteps
  }
}