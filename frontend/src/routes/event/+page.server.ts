import { PUBLIC_BACKEND_URL } from '$env/static/public'
import type { EventBus, FollowBusEvent, RouteStep } from '$lib/types/global.js'

export const load = async ({ locals, fetch }) => {
  let followBus: FollowBusEvent = {} as FollowBusEvent
  let stops: Array<RouteStep> = []

  if (locals.session?.Email != undefined) {
    let response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event/get-follow-bus/${locals.session?.Email}`, {
      method: 'GET',
      headers: {
        'content-type': 'application/json'
      },
    })
    if (response) {
      followBus = await response.json() as FollowBusEvent
      if (followBus?.Carplate != '') {
        response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event/get-route-steps/${followBus?.RouteName}`, {
          method: 'GET',
          headers: {
            'content-type': 'application/json'
          },
        })
        stops = await response.json() as Array<RouteStep>
      }
    }
  }

  let buses: Array<EventBus> = []
  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event/get-buses`, {
    method: 'GET',
    headers: {
      'content-type': 'application/json'
    },
  })
  if (response) {
    buses = await response.json() as Array<EventBus>
  }

  return {
    followBus,
    buses,
    stops
  }
}