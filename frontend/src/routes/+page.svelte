<script lang="ts">
  export let data 
  const { backend_uri, stops, env, routes, routeSteps, osr_key, mapbox_key } = data
	import { onMount } from "svelte";
  import EventTracker from "$lib/components/EventTracker.svelte"
  import type { CurrentBus, FollowBusEvent } from "$lib/types/global.js"
	import type { MapboxOptions, MarkerOptions } from "mapbox-gl";
  import { Map, Marker } from "mapbox-gl";
  import tinycolor from "tinycolor2"

  let map:Map
  let mapContainer: MapboxOptions["container"]

  let lng = 103.7772;
  let lat = 1.3093;
  let zoom = 16;

  let initialState = { lng: lng, lat: lat, zoom: zoom };

  let busList: Array<CurrentBus> = []
  let ws: WebSocket
  
  let loaded = false
  
  const coordCollection:GeoJSON.LineString[][] = []
  ws = new WebSocket(`${env == "PROD" ? "wss" : "ws"}://${backend_uri.split("//")[1]}:3000/ws`)
  let markers: Array<Marker> = []
  let offset = 0.0002

    
  
  onMount(async() => {
    if (routes && routeSteps) {
      for (const route of routes) {
        const coords:GeoJSON.LineString[] = []
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
            "Authorization": osr_key 
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
              "Authorization": osr_key
            },
            body: JSON.stringify({
              "continue_straight": false,
              "coordinates": inputCoords
            })
          })
          parsed = await response.json()
          coords.push(parsed.features[0].geometry)
        } else {
          response = await fetch(`https://api.mapbox.com/matching/v5/mapbox/driving/${positions}?geometries=geojson&access_token=${mapbox_key}`)
          parsed = await response.json()
          coords.push(parsed.matchings[0].geometry)
        }
    
        coordCollection.push(coords)
      }
    }

    const getCurrentBuses = async () => {
      const response = await fetch(`${backend_uri}:3000/event/get-current-buses`)
      let parsed = await response.json() as Array<CurrentBus>
      if (parsed != null) {
        busList = parsed
      }
      loaded = true
    }
    
    
    getCurrentBuses()
    
    
    map = new Map({
      container: mapContainer,
      accessToken: mapbox_key,
      style: `mapbox://styles/mapbox/streets-v12`,
      center: [initialState.lng, initialState.lat],
      zoom: initialState.zoom
    })
    map.on('load', () => {
      stops.forEach(stop => {
        const waypoint = new Marker({color: '#b91c1c'}).setLngLat([stop.Lng, stop.Lat]).addTo(map).on('dragend', () => { console.log(waypoint.getLngLat())})
      })
      ws.onmessage = async(msg) => {
        let parts = msg.data.split(" ")
        if (parts[0] == 'refresh') {
          await getCurrentBuses()
          markers.forEach(marker => {
            marker.remove()
          })
          busList.forEach(currentBus => {
            const waypoint = new Marker({color: currentBus.Color}).setLngLat([currentBus.Lng, currentBus.Lat + offset]).addTo(map).on('dragend', () => { console.log(waypoint.getLngLat())})
            markers.push(waypoint)
          })
        }
      }
      busList.forEach(currentBus => {
        const waypoint = new Marker({color: currentBus.Color}).setLngLat([currentBus.Lng, currentBus.Lat + offset]).addTo(map).on('dragend', () => { console.log(waypoint.getLngLat())})
        markers.push(waypoint)
      })
      for (let j = 0; j < routes.length; j++) {
        const route = routes[j]
        const coords = coordCollection[j]
        for (let i = 0; i < coords.length; i++) {
          const geo = coords[i]
          const color = tinycolor(route.Color).darken(i * 20).toString()
          // const waypoint = new Marker({color: '#b91c1c', draggable: true}).setLngLat([geo.coordinates[Math.floor(geo.coordinates.length / 2)][0], geo.coordinates[Math.floor(geo.coordinates.length / 2)][1]]).addTo(map).on('dragend', () => { console.log(waypoint.getLngLat())})
          if (map.getSource(`${route.RouteName}${i}`)) {
            map.removeLayer(`${route.RouteName}${i}`);
            map.removeSource(`${route.RouteName}${i}`);
          } else {
            map.addLayer({
              id: `${route.RouteName}${i}`,
              type: 'line',
              source: {
                type: 'geojson',
                data: {
                  type: 'Feature',
                  properties: {},
                  geometry: geo
                }
              },
              layout: {
                'line-join': 'round',
                'line-cap': 'round'
              },
              paint: {
                'line-color': color,
                'line-width': 8,
                'line-opacity': 0.8
              }
            });
          }
        }
      }
    })
  }) 
</script>

<style>
  .map {
  position: absolute;
}
</style>

<div class="p-6 md:p-12">
  <h1 class="text-3xl font-semibold">
    Bus Tracker
  </h1>
  <div>
    <div class="map-wrap">
      <div class="map absolute w-[87.5%] h-[80%]" bind:this={mapContainer} />
    </div>
    {#if !loaded}
      <div class="w-full flex mt-24">
        <p class="m-auto text-stone-500 font-semibold">Loading</p>
      </div>
    {/if} 
  </div>
</div>