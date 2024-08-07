import { PUBLIC_BACKEND_URL } from '$env/static/public';
import type { RouteStep, Stop, Route } from '$lib/types/global';
import { redirect } from '@sveltejs/kit';

export const actions = {
	signOut: async ({ cookies, locals }) => {
		cookies.delete('accessToken', { path: '/' });
		cookies.delete('refreshToken', { path: '/' });
		locals.session = null;
		redirect(301, '/')
	}
}

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
	response = await fetch(`${PUBLIC_BACKEND_URL}:3000/route-step/get-all-stops`);
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

