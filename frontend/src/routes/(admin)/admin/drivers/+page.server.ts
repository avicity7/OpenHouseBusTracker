import { PUBLIC_BACKEND_URL } from "$env/static/public";

export const load = async () => {
	try {
		const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/driver/get-driver-hours`);
		if (!response.ok) {
			throw new Error(`Failed to fetch: ${response.statusText}`);
		}
		const ScheduleTimeDiff = await response.json();

		return {
			ScheduleTimeDiff
		};
	} catch (error) {
		console.error("Error fetching driver hours:", error);
        return {
            status: 500,
            body: { error: 'Internal Server Error' }
        };
	}
};