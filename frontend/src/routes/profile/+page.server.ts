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
    const result = await fetch(`${PUBLIC_BACKEND_URL}:3000/auth/login`, {
      method: 'POST',
      body: JSON.stringify({
        "Email": email,
        "Password": password
      }), 
      headers: {
        'content-type': 'application/json'
      },
    })
    
    try {
      const parsed = await result.json() as { User: { Email: string, Role: string }, AccessToken: string, RefreshToken: string }
      const { User, AccessToken, RefreshToken } = parsed
      cookies.set('accessToken', AccessToken, {path: '/'})
      cookies.set('refreshToken', RefreshToken, {path: '/'}) 
      locals.session = User
    } catch(err) {
      console.log(err)
    }
  },
  signOut: async ({ cookies, locals }) => {
    cookies.delete('accessToken', {path: '/'})
    cookies.delete('refreshToken', {path: '/'})
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