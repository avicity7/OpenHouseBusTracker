import {
	PUBLIC_BACKEND_URL,
	PUBLIC_MAPBOX_KEY,
	PUBLIC_ENV,
	PUBLIC_OSR_KEY
} from '$env/static/public';
import { redirect } from '@sveltejs/kit';
import type { User } from '$lib/types/global.ts';

export const load = async ({ url, locals, fetch }) => {
	const path = url.pathname;
	const session = locals.session;
	const backend_uri = PUBLIC_BACKEND_URL;
	const mapbox_key = PUBLIC_MAPBOX_KEY;
	const osr_key = PUBLIC_OSR_KEY;
	const env = PUBLIC_ENV;

	if (path.includes('admin') && session?.Role != 'admin') {
		redirect(301, '/');
	}

	const getUser = async () => {
		const response = await fetch(`${backend_uri}:3000/auth/get-user/${session?.Email}`, {
			method: 'GET',
			headers: {
				'content-type': 'application/json'
			},
			credentials: 'include'
		});
		let parsed: User = session;
		console.log(parsed);
		if (response) {
			parsed = (await response.json()) as User;
		}
		return parsed;
	};

	const account = session;

	return {
		backend_uri,
		session,
		osr_key,
		mapbox_key,
		env,
		account
	};
};
