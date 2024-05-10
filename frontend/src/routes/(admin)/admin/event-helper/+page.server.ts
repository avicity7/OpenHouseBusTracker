import type { Load } from '@sveltejs/kit';
import type { EventHelper } from '../../../../types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

export const load: Load = async ({ fetch }) => {
    try {
        const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event-helpers/get-helpers`);
        if (!response.ok) {
            throw new Error("Failed to fetch event helpers");
        }
        
        const data = await response.json() as EventHelper[];
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