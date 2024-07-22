import { PUBLIC_BACKEND_URL, PUBLIC_CAPTCHA } from '$env/static/public'
import fetch from 'node-fetch'
export const actions = {
  signUp: async ({ request, cookies, locals }) => {
    const values = await request.formData()
    const Name = values.get("name")
    const Email = values.get("email")
    const Password = values.get("password")
    const captchaRes = values.get("g-recaptcha-response")

    if (Name == "" || Email == "" || captchaRes == "") return { success: false }

    const result = await fetch(`${PUBLIC_BACKEND_URL}:3000/auth/create-user`, {
      method: 'POST',
      body: JSON.stringify({
        Name,
        Email,
        Password,
        Role: 0
      }), 
      headers: {
        'content-type': 'application/json'
      },
    })

    try {
      const parsed = await result.json() as { User: { Name: string, Email: string, Role: string, VerificationToken: string }, AccessToken: string, RefreshToken: string }
      const { User, AccessToken, RefreshToken } = parsed
      cookies.set('accessToken', AccessToken, {path: '/'})
      cookies.set('refreshToken', RefreshToken, {path: '/'}) 
      locals.session = User
      return { success: true }
    } catch(err) {
      console.log(err)
    }
  }
}

export const load = () => {
  const captcha_key = PUBLIC_CAPTCHA

  return {
    captcha_key
  }
}