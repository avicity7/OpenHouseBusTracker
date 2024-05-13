import { PUBLIC_BACKEND_URL, PUBLIC_MAPBOX_KEY } from '$env/static/public'
import type { Stop } from '$lib/types/global.js'

export const load = async ({ fetch }) => {
  let stops:Array<Stop> = []
  const coords = []

  let response = await fetch(`${PUBLIC_BACKEND_URL}:3000/route-step/get-all-stops`, {
    method: 'GET',
    headers: {
      'content-type': 'application/json'
    },
  })
  if (response) {
    stops = await response.json() as Array<Stop>
    let string = ''
    stops.forEach(stop => {
      string += `${stop.Lng},${stop.Lat};`
    })
    string = string.slice(0, -1)
    response = await fetch(`https://api.mapbox.com/directions/v5/mapbox/walking/${string}?geometries=geojson&access_token=${PUBLIC_MAPBOX_KEY}`)
    let geo = await response.json()
    coords.push(geo.routes[0].geometry)

    response = await fetch(`https://api.mapbox.com/directions/v5/mapbox/driving/${string}?geometries=geojson&access_token=${PUBLIC_MAPBOX_KEY}`)
    geo = await response.json()
    coords.push(geo.routes[0].geometry)
  }

  return {
    stops,
    coords
  }
}