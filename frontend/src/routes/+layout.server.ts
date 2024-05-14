import { PUBLIC_BACKEND_URL, PUBLIC_MAPBOX_KEY, PUBLIC_ENV, PUBLIC_OSR_KEY } from '$env/static/public'
import { redirect } from '@sveltejs/kit'

export const load = async (event) => {
  const path = event.url.pathname
  const session = event.locals.session
  const backend_uri = PUBLIC_BACKEND_URL 
  const mapbox_key = PUBLIC_MAPBOX_KEY
  const osr_key = PUBLIC_OSR_KEY
  const env = PUBLIC_ENV
  
  if (path.includes("admin") && session?.Role != "admin") {
    redirect(301, "/")
  }

  return {
    backend_uri,
    session,
    osr_key,
    mapbox_key,
    env
  }
}