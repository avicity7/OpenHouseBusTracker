import { PUBLIC_BACKEND_URL } from '$env/static/public';
import type { User, LoginResponse } from '$lib/types/global.ts';

export const actions = {
	login: async ({ request, cookies, locals }) => {
		const values = await request.formData();
		const email = values.get('email');
		const password = values.get('password');
		const result = await fetch(`${PUBLIC_BACKEND_URL}:3000/auth/login`, {
			method: 'POST',
			body: JSON.stringify({
				Email: email,
				Password: password
			}),
			headers: {
				'content-type': 'application/json'
			}
		});

		try {
			const parsed = (await result.json()) as LoginResponse;
			const { User, AccessToken, RefreshToken } = parsed;
			cookies.set('accessToken', AccessToken, { path: '/' });
			cookies.set('refreshToken', RefreshToken, { path: '/' });
			locals.session = User;
			return {
				account: User
			};
		} catch (err) {
			console.log(err);
		}
	},
	signOut: async ({ cookies, locals }) => {
		cookies.delete('accessToken', { path: '/' });
		cookies.delete('refreshToken', { path: '/' });
		locals.session = null;
	}
};

export const load = async ({ locals, fetch }) => {
	let account: User = {
		Name: '',
		Email: '',
		Contact: '',
		Role: '',
		VerificationToken: ''
	};

	if (locals.session?.Email != undefined) {
		const response = await fetch(
			`${PUBLIC_BACKEND_URL}:3000/auth/get-user/${locals.session?.Email}`,
			{
				method: 'GET',
				headers: {
					'content-type': 'application/json'
				},
				credentials: 'include'
			}
		);
		if (response) {
			account = (await response.json()).Output as User;
		}
	}

	return {
		account
	};
};
