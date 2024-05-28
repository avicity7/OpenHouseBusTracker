import { PUBLIC_BACKEND_URL } from '$env/static/public';

/** @type {import('./$types').RequestHandler} */
export async function GET({ fetch, cookies, params }) {
	const uri = params.uri;
	console.log(uri);
	const result = await fetch(`${PUBLIC_BACKEND_URL}:3000/${uri}`, {
		method: 'GET',
		headers: {
			'content-type': 'application/json'
		},
		credentials: 'include'
	});
	const parsed = await result.json();
	const AccessToken = parsed.Tokens.AccessToken;
	const RefreshToken = parsed.Tokens.RefreshToken;
	if (AccessToken != '') {
		cookies.set('accessToken', AccessToken, { path: '/' });
		cookies.set('refreshToken', RefreshToken, { path: '/' });
	}
	return new Response(JSON.stringify(parsed.Output));
}
