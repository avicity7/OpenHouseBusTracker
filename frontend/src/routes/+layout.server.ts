import { PUBLIC_BACKEND_URL } from '$env/static/public'
import { redirect } from '@sveltejs/kit'

export const load = async (event) => {
  const path = event.url.pathname
  const session = event.locals.session
  const backend_uri = PUBLIC_BACKEND_URL 
  
  if (path.includes("admin") && session?.Role != "admin") {
    redirect(301, "/")
  }

  return {
    backend_uri,
    session,
  }
}