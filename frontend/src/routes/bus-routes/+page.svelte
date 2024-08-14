<script lang="ts">
  import type { Route, RouteStep, Stop } from '$lib/types/global.js';

	export let data;

	const { 
    backend_uri, 
    env, 
    osr_key, 
    mapbox_key, 
    routeSteps, 
    stops, 
    routes 
  }: { 
    backend_uri: string, 
    env: string, 
    osr_key: string, 
    mapbox_key: string,
    routeSteps: Array<RouteStep>,
    stops: Array<Stop>
    routes: Array<Route>
  } = data;
	import { onMount } from 'svelte';
	import type { CurrentBus, FollowBusEvent } from '$lib/types/global.js';
	import type { MapboxOptions, MarkerOptions, Map, Marker } from 'mapbox-gl';
	import pkg from 'mapbox-gl';
	const { Map, Marker, Popup } = pkg;
	import tinycolor from 'tinycolor2';

	let map: Map;
	let mapContainer: MapboxOptions['container'];

	let lng = 103.7772;
	let lat = 1.3093;
	let zoom = 16;

	let initialState = { lng: lng, lat: lat, zoom: zoom };

	let busList: Array<CurrentBus> = [];
	let ws: WebSocket;

	let loaded = false;

	const coordCollection: GeoJSON.LineString[][] = [];
	let markers: Array<Marker> = [];

	onMount(async () => {
		if (routes && routeSteps) {
			for (const route of routes) {
				const coords: GeoJSON.LineString[] = [];
				const inputCoords: [number, number][] = [];
				let positions = '';
				routeSteps.forEach((stop) => {
					if (stop.RouteName == route.RouteName) {
						positions = `${stop.Lng},${stop.Lat};` + positions;
						inputCoords.push([stop.Lng, stop.Lat]);
					}
				});
				positions = positions.slice(0, -1);

				let response = await fetch(
					`https://api.openrouteservice.org/v2/directions/driving-car/geojson`,
					{
						method: 'POST',
						headers: {
							'Content-Type': 'application/json',
							Authorization: osr_key
						},
						body: JSON.stringify({
							alternative_routes: { target_count: 2, weight_factor: 10 },
							coordinates: [inputCoords[inputCoords.length - 1], inputCoords[0]]
						})
					}
				);
				let parsed = await response.json();
				coords.push(parsed.features[parsed.features.length - 1].geometry);


				if (inputCoords.length > 2) {
					response = await fetch(
						`https://api.mapbox.com/matching/v5/mapbox/driving/${positions}?geometries=geojson&access_token=${mapbox_key}`
					);
					parsed = await response.json();
					coords.push(parsed.matchings[0].geometry);
				} else {
					response = await fetch(
						`https://api.mapbox.com/matching/v5/mapbox/driving/${positions}?geometries=geojson&access_token=${mapbox_key}`
					);
					parsed = await response.json();
					coords.push(parsed.matchings[0].geometry);
				}

				coordCollection.push(coords);
			}
		}

		map = new Map({
			container: mapContainer,
			accessToken: mapbox_key,
			style: `mapbox://styles/mapbox/streets-v12`,
			center: [initialState.lng, initialState.lat],
			zoom: initialState.zoom
		});
		map.on('load', () => {
			stops.forEach((stop) => {
				const waypoint = new Marker({ color: '#b91c1c' })
					.setLngLat([stop.Lng, stop.Lat])
					.addTo(map)
					.on('dragend', () => {
						console.log(waypoint.getLngLat());
					});
			});

			for (let j = 0; j < routes.length; j++) {
				const route = routes[j];
				const coords = coordCollection[j];
				for (let i = 0; i < coords.length; i++) {
					const geo = coords[i];
					const color = tinycolor(route.Color)
						.darken(i * 20)
						.toString();
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
		});
	});
</script>

<div class="map-wrap h-[100%]">
	<div class="map w-full h-[100%]" bind:this={mapContainer} />
	<div class="absolute top-0 mt-24 ml-8 bg-white rounded-md text-xl px-8 py-3 font-semibold">
		Bus Routes
	</div>
</div>
{#if !loaded}
	<div class="flex p-24">
		<p class="m-auto text-stone-500 font-semibold">Loading</p>
	</div>
{/if}

<style>
	.map {
		position: absolute;
	}
</style>
