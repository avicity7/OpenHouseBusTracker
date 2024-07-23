import { PUBLIC_BACKEND_URL } from "$env/static/public";
import type { Driver } from "$lib/types/global";

export const load = async () => {
	try {
		let response = await fetch(`${PUBLIC_BACKEND_URL}:3000/driver/get-driver-hours`);
		if (!response.ok) {
			throw new Error(`Failed to fetch: ${response.statusText}`);
		}
		const ScheduleTimeDiff = await response.json();

		response = await fetch(`${PUBLIC_BACKEND_URL}:3000/driver/get-drivers`);
		if (!response.ok) {
				throw new Error(`Failed to fetch: ${response.statusText}`);
		}
		const drivers = await response.json() as Driver[];

		return {
			ScheduleTimeDiff,
			drivers
		};
	} catch (error) {
		console.error("Error fetching driver hours:", error);
        return {
            status: 500,
            body: { error: 'Internal Server Error' }
        };
	}
};