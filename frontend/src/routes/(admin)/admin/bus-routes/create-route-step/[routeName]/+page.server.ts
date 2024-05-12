import { PUBLIC_BACKEND_URL } from '$env/static/public'

export const load = async ({ fetch }) => {
  let stops:Array<string> = []

  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/route-step/get-all-stops`, {
    method: 'GET',
    headers: {
      'content-type': 'application/json'
    },
  })
  if (response) {
    stops = await response.json() as Array<string>
  }

  return {
    stops
  }
}