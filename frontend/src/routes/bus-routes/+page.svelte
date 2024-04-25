<script lang="ts">
	import type { MapboxOptions } from "mapbox-gl";
  import { onMount, onDestroy } from "svelte";
  import { Map } from "mapbox-gl";
  export let data
  const { mapbox_key } = data
  let map:Map
  let mapContainer: MapboxOptions["container"]

  let lng = 103.7772;
  let lat = 1.3093;
  let zoom = 16;

  onMount(() => {
    const initialState = { lng: lng, lat: lat, zoom: zoom };
    map = new Map({
      container: mapContainer,
      accessToken: mapbox_key,
      style: `mapbox://styles/mapbox/outdoors-v11`,
      center: [initialState.lng, initialState.lat],
      zoom: initialState.zoom
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