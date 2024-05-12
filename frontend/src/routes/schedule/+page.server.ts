import type { Schedule } from '$lib/types/global.js';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

export const load = async ({ fetch, locals }) => {
    try {
        const email = locals?.session?.Email;

        if (!email) {
            throw new Error("Email not found in session");
        }

        const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/get-user-schedule/${email}`);
        if (!response.ok) {
            throw new Error("Failed to fetch current schedules");
        }
        
        const data = await response.json() as Schedule[];
        
        return {
          data
        };

    } catch (error) {
        console.error(error);
        return {
            status: 500,
            error: 'Internal Server Error'
        };
    }
};