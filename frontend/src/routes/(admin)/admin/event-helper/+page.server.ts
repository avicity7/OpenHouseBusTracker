import type { Load } from '@sveltejs/kit';
import type { EventBus, EventHelper } from '$lib/types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

export const load: Load = async ({ fetch }) => {
    try {
        let response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event-helpers/get-helpers`);
        if (!response.ok) {
            throw new Error("Failed to fetch event helpers");
        }
        
        const helperData = await response.json() as EventHelper[];

        response = await fetch(`${PUBLIC_BACKEND_URL}:3000/bus/get-buses`);
        if (!response.ok) {
            throw new Error("Failed to fetch buses");
        }
        
        const busData = await response.json() as EventBus[];
        return {
            helperData,
            busData
        };

    } catch (error) {
        console.error(error);
        return {
            status: 500,
            body: { error: 'Internal Server Error' }
        };
    }
};