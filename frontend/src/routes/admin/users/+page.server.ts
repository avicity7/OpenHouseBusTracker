import { PUBLIC_BACKEND_URL } from '$env/static/public';
import type { AuthedResponse, UserRole } from '../../../types/global';
export const load = async ({fetch, cookies}) => {
  const getUsers = new Promise<AuthedResponse>((resolve) => {
    fetch(`${PUBLIC_BACKEND_URL}:3000/users/get-users`, {
      method: 'GET',
      headers: {
        'content-type': 'application/json'
      },
      credentials: 'include'
    })
    .then(async (response) => {
      const data = await response.json()
      resolve(data)
    })
  })

  const getRoles = new Promise<Array<UserRole>>((resolve) => {
    fetch(`${PUBLIC_BACKEND_URL}:3000/users/get-roles`)
    .then(async (response) => {
      const data = await response.json() as Array<UserRole>
      resolve(data)
    })
  })

  const [userResponse, roles] = await Promise.all([Promise.resolve(getUsers), Promise.resolve(getRoles)])

  const AccessToken = userResponse.Tokens.AccessToken
  const RefreshToken = userResponse.Tokens.RefreshToken
  if (AccessToken != '') {
    cookies.set("accessToken", AccessToken, { path: "/" })
    cookies.set("refreshToken", RefreshToken, { path: "/" })
  }

  const users = userResponse.Output

  return {
    users,
    roles
  }
}