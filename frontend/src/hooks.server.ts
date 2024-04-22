// import {
//  NEXTAUTH_SECRET,
// } from '$env/static/private';
import jwt, { type JwtPayload } from 'jsonwebtoken'
import type { Handle } from '@sveltejs/kit';

const emailPassword: Handle = async ({ event, resolve }) => {
  const cookie = event.request.headers.get('cookie')
  if (cookie && cookie.split('-')[0] != 'next' && cookie.split('-')[0] != '__Secure' && cookie.split('-')[0] != '__Host') {
    const token = cookie.split('=')[1]
    const decode = jwt.decode(token as string) as JwtPayload
    const { email, username } = decode
    event.locals.session = {user: { email, username }}
  }
  if (event.url.search == '?/signOut') {
    event.locals.session = null
  }
  const response = await resolve(event)
  return response
}

export const handle = emailPassword

/** @type {import('@sveltejs/kit').HandleFetch} */
export async function handleFetch({ event, request, fetch }) {
	request.headers.set('cookie', String(event.request.headers.get('cookie')));

	return fetch(request);

}
