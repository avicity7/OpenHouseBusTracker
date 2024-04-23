import { PUBLIC_BACKEND_URL } from '$env/static/public'
import fetch from 'node-fetch'
export const actions = {
  signUp: async ({ request, cookies, locals }) => {
    const values = await request.formData()
    const Email = values.get("email")
    const Password = values.get("password")

    const result = await fetch(`${PUBLIC_BACKEND_URL}:3000/auth/create-user`, {
      method: 'POST',
      body: JSON.stringify({
        Email,
        Password,
        Role: 0
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
  }
}