import {
  PUBLIC_BACKEND_URL,
  PUBLIC_MAPBOX_KEY,
  PUBLIC_ENV,
  PUBLIC_OSR_KEY
} from '$env/static/public';
import { redirect } from '@sveltejs/kit';
import type { User } from '$lib/types/global.ts';

export const load = async ({ url, locals, fetch, cookies }) => {
  const path = url.pathname;
  let session = locals.session
  const backend_uri = PUBLIC_BACKEND_URL;
  const mapbox_key = PUBLIC_MAPBOX_KEY;
  const osr_key = PUBLIC_OSR_KEY;
  const env = PUBLIC_ENV;
  const pathname = url.pathname;


  const response = await fetch(`${backend_uri}:3000/auth/get-user/${locals.session?.Email}`, {
    method: 'GET',
    headers: {
      'content-type': 'application/json'
    },
    credentials: 'include'
  });

  let account = locals.session

  if (response) {
    const parsed = await response.json()
    account = parsed.Output as User


    if (account.Role != locals.session?.Role) {
      const AccessToken = parsed.Tokens.AccessToken
      const RefreshToken = parsed.Tokens.RefreshToken
      if (AccessToken != '') {
        cookies.set("accessToken", AccessToken, { path: "/" })
        cookies.set("refreshToken", RefreshToken, { path: "/" })
      }
      locals.session = account
      session = account
    }
  }

  if (path.includes('admin') && locals.session?.Role != 'admin') {
    redirect(301, '/');
  }

  return {
    backend_uri,
    session,
    osr_key,
    mapbox_key,
    env,
    account,
    pathname
  };
};
