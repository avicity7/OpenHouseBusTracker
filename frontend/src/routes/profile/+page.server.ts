import fetch from 'node-fetch'
import { PUBLIC_BACKEND_URL } from '$env/static/public'
// import type { Agent } from 'https'
// let httpsAgent: Agent
// if (PUBLIC_ENV !== "DEV") {
//   httpsAgent = new https.Agent({
//     rejectUnauthorized: false,
//   });
// }
export const actions = {
  login: async ({ request, cookies, locals }) => {
    const values = await request.formData()
    const email = values.get('email')
    const password = values.get('password')
    const result = await fetch(`${PUBLIC_BACKEND_URL}:3000/api/auth/signInUser`, {
      method: 'POST',
      body: JSON.stringify({
        email,
        password
      }), 
      headers: {
        'content-type': 'application/json'
      },
    })
    
    try {
      const parsed = await result.json() as {token: string, user: { email:string, username: string }}
      cookies.set('jwt-token', parsed.token, {path: '/'})
      const { email, username } = parsed.user
      locals.session = {user: {email, username }}
    } catch(err) {
      console.log(err)
    }
  },
  signOut: async ({ cookies, locals }) => {
    cookies.delete('jwt-token', {path: '/'})
    locals.session = null
  }
}

// export const load = async (event) => {
//   const session = event.locals.session
//   const backend_uri = PUBLIC_BACKEND_URL 
//   return {
//     backend_uri,
//     session,
//   }
// }