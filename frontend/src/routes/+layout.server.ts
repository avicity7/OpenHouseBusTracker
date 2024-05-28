import {
	PUBLIC_BACKEND_URL,
	PUBLIC_MAPBOX_KEY,
	PUBLIC_ENV,
	PUBLIC_OSR_KEY
} from '$env/static/public';
import { redirect } from '@sveltejs/kit';
import type { User } from '$lib/types/global.ts';

export const load = async (event) => {
	const path = event.url.pathname;
	const session = event.locals.session;
	const backend_uri = PUBLIC_BACKEND_URL;
	const mapbox_key = PUBLIC_MAPBOX_KEY;
	const osr_key = PUBLIC_OSR_KEY;
	const env = PUBLIC_ENV;

	if (path.includes('admin') && session?.Role != 'admin') {
		redirect(301, '/');
	}

	const getProfile = async () => {
		const response = await fetch(`${backend_uri}:3000/auth/get-user/${session?.Email}`);
		const parsed = (await response.json()) as User;
		return parsed;
	};
	const account = await getProfile();

	return {
		backend_uri,
		session,
		osr_key,
		mapbox_key,
		env,
		account
	};
};
