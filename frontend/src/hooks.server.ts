// import {
//  NEXTAUTH_SECRET,
// } from '$env/static/private';
import jwt, { type JwtPayload } from 'jsonwebtoken'
import type { Handle } from '@sveltejs/kit';

const emailPassword: Handle = async ({ event, resolve }) => {
  const cookie = event.request.headers.get('cookie')
  if (cookie) {
    const tokens = cookie.split('; ')
    const accessToken = tokens[0].split("=")[1]
    const decode = await jwt.decode(accessToken as string) as JwtPayload
    event.locals.session = { Email: decode["Email"], Role: decode["Role"] }
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
  if (event.request.headers.has('cookie')) {  
    const tokens = String(event.request.headers.get('cookie')).split(" ")
    request.headers.set('access', tokens[0].split("=")[1].replace(";",""))
    request.headers.set('refresh', tokens[1].split("=")[1])
  
    return fetch(request);
  }

}
