import { PUBLIC_BACKEND_URL } from '$env/static/public'
import type { RouteStop } from '../types/global.js'

export const load = async () => {
  let stops: Array<Array<RouteStop>> = []

  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event/get-all-stops`, {
    method: 'GET',
    headers: {
      'content-type': 'application/json'
    },
  })
  stops = await response.json() as Array<Array<RouteStop>>

  return {
    stops
  }
}