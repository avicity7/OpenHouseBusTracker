import { PUBLIC_BACKEND_URL, PUBLIC_OSR_KEY, PUBLIC_MAPBOX_KEY } from '$env/static/public'
import type { Route, RouteStep, Stop } from '$lib/types/global.js'

export const load = async ({ fetch }) => {
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
  response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event/get-all-stops`)
  let stops = await response.json() as Array<Stop>
  if (!stops) {
    stops = []
  }
  
  const coordCollection = []
  if (routes && routeSteps) {
    for (const route of routes) {
      const coords = []
      const inputCoords:[number, number][] = []
      let positions = ''
      routeSteps.forEach(stop => {
        if (stop.RouteName == route.RouteName) {
          positions = `${stop.Lng},${stop.Lat};` + positions
          inputCoords.push([stop.Lng, stop.Lat])
        } 
      })
      positions = positions.slice(0, -1)

      let response = await fetch (`https://api.openrouteservice.org/v2/directions/driving-car/geojson`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": PUBLIC_OSR_KEY 
        },
        body: JSON.stringify({
          "alternative_routes": {"target_count": 2, "weight_factor": 10},
          "coordinates": [inputCoords[inputCoords.length - 1], inputCoords[0]]
        })
      })
      let parsed = await response.json()
      coords.push(parsed.features[parsed.features.length - 1].geometry)
      
      if (parsed.features.length > 1) {
        response = await fetch (`https://api.openrouteservice.org/v2/directions/driving-car/geojson`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "Authorization": PUBLIC_OSR_KEY 
          },
          body: JSON.stringify({
            "continue_straight": false,
            "coordinates": inputCoords
          })
        })
        parsed = await response.json()
        coords.push(parsed.features[0].geometry)
      } else {
        response = await fetch(`https://api.mapbox.com/matching/v5/mapbox/driving/${positions}?geometries=geojson&access_token=${PUBLIC_MAPBOX_KEY}`)
        parsed = await response.json()
        coords.push(parsed.matchings[0].geometry)
      }
      
      coordCollection.push(coords)
    }
  }

  return {
    coordCollection,
    routes,
    stops
  }
}
