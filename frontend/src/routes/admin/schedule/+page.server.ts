import type { Load } from '@sveltejs/kit';
import type { Schedule } from '../../../types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

export const load: Load = async ({ fetch }) => {
    try {
        const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/get-schedule`);
        if (!response.ok) {
            throw new Error("Failed to fetch bus schedules");
        }
        
        const data = await response.json() as Schedule[];
        return {
            data
        };

    } catch (error) {
        console.error(error);
        return {
            status: 500,
            body: { error: 'Internal Server Error' }
        };
    }
};
