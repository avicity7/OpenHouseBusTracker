<script lang="ts">
	import type { MapboxOptions, MarkerOptions } from "mapbox-gl";
  import { onMount, onDestroy } from "svelte";
  import { Map, Marker } from "mapbox-gl";
  import tinycolor from "tinycolor2"
  export let data
  const { mapbox_key, coordCollection, routes, stops } = data
  let map:Map
  let mapContainer: MapboxOptions["container"]

  let lng = 103.7772;
  let lat = 1.3093;
  let zoom = 16;

  let initialState = { lng: lng, lat: lat, zoom: zoom };

  onMount(() => {
    map = new Map({
      container: mapContainer,
      accessToken: mapbox_key,
      style: `mapbox://styles/mapbox/streets-v12`,
      center: [initialState.lng, initialState.lat],
      zoom: initialState.zoom
    })
    map.on('load', () => {
      stops.forEach(stop => {
        const waypoint = new Marker({color: '#b91c1c', draggable: true}).setLngLat([stop.Lng, stop.Lat]).addTo(map).on('dragend', () => { console.log(waypoint.getLngLat())})
      })
      for (let j = 0; j < routes.length; j++) {
        const route = routes[j]
        const coords = coordCollection[j]
        for (let i = 0; i < coords.length; i++) {
          const geo = coords[i]
          const color = tinycolor(route.Color).darken(i * 20).toString()
          const waypoint = new Marker({color: '#b91c1c', draggable: true}).setLngLat([geo.coordinates[Math.floor(geo.coordinates.length / 2)][0], geo.coordinates[Math.floor(geo.coordinates.length / 2)][1]]).addTo(map).on('dragend', () => { console.log(waypoint.getLngLat())})
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

  onDestroy(() => {
    map.remove();
  })
</script>
<style>
  .map {
  position: absolute;
}
</style>
<div class="p-6 md:p-12 flex flex-col">
  <h1 class="text-3xl font-semibold mb-2">
    Routes
  </h1>
  <div class="map-wrap">
    <div class="map absolute w-[87.5%] h-[80%]" bind:this="{mapContainer}" />
  </div>
</div>