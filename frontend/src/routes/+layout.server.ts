import { PUBLIC_BACKEND_URL, PUBLIC_MAPBOX_KEY } from '$env/static/public'
import { redirect } from '@sveltejs/kit'

export const load = async (event) => {
  const path = event.url.pathname
  const session = event.locals.session
  const backend_uri = PUBLIC_BACKEND_URL 
  const mapbox_key = PUBLIC_MAPBOX_KEY
  
  if (path.includes("admin") && session?.Role != "admin") {
    redirect(301, "/")
  }

  return {
    backend_uri,
    session,
    mapbox_key
  }
}