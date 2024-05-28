import { PUBLIC_BACKEND_URL } from '$env/static/public';
import type { RouteStep, Stop, Route } from '$lib/types/global';

export const load = async () => {
	let response = await fetch(`${PUBLIC_BACKEND_URL}:3000/route/`);
	let routes = (await response.json()) as Array<Route>;
	if (!routes) {
		routes = [];
	}
	response = await fetch(`${PUBLIC_BACKEND_URL}:3000/route-step/`);
	let routeSteps = (await response.json()) as Array<RouteStep>;
	if (!routeSteps) {
		routeSteps = [];
	}
	response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event/get-all-stops`);
	let stops = (await response.json()) as Array<Stop>;
	if (!stops) {
		stops = [];
	}

	return {
		stops,
		routes,
		routeSteps
	};
};

