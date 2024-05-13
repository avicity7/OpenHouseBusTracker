<script lang="ts">
	import type { MapboxOptions, MarkerOptions } from "mapbox-gl";
  import { onMount, onDestroy } from "svelte";
  import { Map, Marker } from "mapbox-gl";
  export let data
  const { mapbox_key, stops, coords } = data
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
      style: `mapbox://styles/mapbox/outdoors-v11`,
      center: [initialState.lng, initialState.lat],
      zoom: initialState.zoom
    })
    stops.forEach(stop => {
      const marker = new Marker({color: '#b91c1c', draggable: true}).setLngLat([stop.Lng, stop.Lat]).addTo(map).on('dragend', () => { console.log(marker.getLngLat())})
    })
    map.on('load', () => {
      let id = 0
      coords.forEach(geo => {
        map.addLayer({
          id: `route${id}`,
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
            'line-color': '#FFFF00',
            'line-width': 8,
            'line-opacity': 0.8
          }
        })
        id += 1
      })
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