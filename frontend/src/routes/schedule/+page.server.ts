import type { Schedule, SwapRequest } from '$lib/types/global.ts';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

export const load = async ({ fetch, locals }) => {
    try {
        const email = locals?.session?.Email;

        if (!email) {
            throw new Error("Email not found in session");
        }

        const currentResponse = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/get-user-schedule/${email}`);
        if (!currentResponse.ok) {
            throw new Error("Failed to fetch current schedules");
        }
        
        const futureResponse = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/get-future-user-schedule/${email}`);
        if (!futureResponse.ok) {
            throw new Error("Failed to fetch future schedules");
        }

        const swapRequestResponse = await fetch(`${PUBLIC_BACKEND_URL}:3000/event-helpers/get-swap-requests/${email}`)

        const currentSchedules = await currentResponse.json() as Schedule[];
        const futureSchedules = await futureResponse.json() as Schedule[];
        const swapRequests = await swapRequestResponse.json() as SwapRequest[]

        return {
            currentSchedules,
            futureSchedules,
            swapRequests
        };

    } catch (error) {
        console.error(error);
        return {
            status: 500,
            error: 'Internal Server Error'
        };
    }
};
