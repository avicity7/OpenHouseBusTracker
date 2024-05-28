import { PUBLIC_BACKEND_URL } from '$env/static/public';
import type { AuthedResponse, UserRole } from '$lib/types/global';
// import fs from 'fs';
// import path from 'path'

export const load = async ({ fetch, cookies }) => {
	const getUsers = new Promise<AuthedResponse>((resolve) => {
		fetch(`${PUBLIC_BACKEND_URL}:3000/users/get-users`, {
			method: 'GET',
			headers: {
				'content-type': 'application/json'
			},
			credentials: 'include'
		}).then(async (response) => {
			const data = await response.json();
			resolve(data);
		});
	});

	const getRoles = new Promise<Array<UserRole>>((resolve) => {
		fetch(`${PUBLIC_BACKEND_URL}:3000/users/get-roles`).then(async (response) => {
			const data = (await response.json()) as Array<UserRole>;
			resolve(data);
		});
	});

	const [userResponse, roles] = await Promise.all([
		Promise.resolve(getUsers),
		Promise.resolve(getRoles)
	]);

	const users = userResponse;

	return {
		users,
		roles
	};
};

// export const actions = {
//   uploadCSV: async ({ request }) => {
//     const formData = await request.formData();
//     const file = formData.get('file') as File;

//     if (!file) {
//       console.error('No file uploaded');
//       return { success: false, error: 'No file uploaded' };
//     }

//     console.log('File object:', file);

//     const fileName = file.name || 'uploaded-file.csv';
//     const tempFilePath = path.join('/tmp', fileName);
//     const fileBuffer = Buffer.from(await file.arrayBuffer());

//     try {
//       await fs.promises.writeFile(tempFilePath, fileBuffer);

//       const backendFormData = new FormData();
//       backendFormData.append('file', new Blob([fileBuffer], { type: file.type }), fileName);

//       const result = await fetch(`${PUBLIC_BACKEND_URL}:3000/auth/bulk-create-users`, {
//         method: 'POST',
//         body: backendFormData,
//       });

//       if (result.ok) {
//         return { success: true };
//       } else {
//         const error = await result.text();
//         return { success: false, error };
//       }
//     } catch (err) {
//       console.error('Error processing file:', err);
//       return { success: false, error: 'Error processing file' };
//     }
//   }
// }

