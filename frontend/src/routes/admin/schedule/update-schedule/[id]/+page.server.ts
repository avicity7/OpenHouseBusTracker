import type { Load } from '@sveltejs/kit';
import type { Schedule } from '../../../../../types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

// used page.server.ts instead of page.js for the need of updated information and fast ssr
export const load: Load = async (ctx) => {
    try {
        const dropdownDataPromise = new Promise((resolve, reject) => {
            const fetchData = async () => {
                try {
                    const dropdownUrl = `${PUBLIC_BACKEND_URL}/schedules/get-dropdown-data`;
                    console.log("Fetching dropdown data from:", dropdownUrl);

                    const response = await ctx.fetch(dropdownUrl);
                    if (!response.ok) {
                        throw new Error(`Failed to fetch dropdown data: ${response.statusText}`);
                    }

                    const data = await response.json() as Schedule[];
                    resolve({ data });
                } catch (error) {
                    console.error("Error fetching dropdown data:", error);
                    reject({
                        status: 500,
                        body: { error: 'Internal Server Error' }
                    });
                }
            };

            fetchData();
        });

        const scheduleDataPromise = new Promise((resolve, reject) => {
            const fetchData = async () => {
                try {
                    const { id } = ctx.params;
                    console.log("Received ID:", id);
                    const scheduleUrl = `${PUBLIC_BACKEND_URL}/schedules/get-schedule/${id}`;
                    console.log("Fetching schedule data from:", scheduleUrl);

                    const response = await ctx.fetch(scheduleUrl);
                    if (!response.ok) {
                        throw new Error(`Failed to fetch schedule data from id: ${id}: ${response.statusText}`);
                    }

                    const schedule = await response.json();
                    console.log("Fetched schedule data:", schedule);

                    resolve({ schedule });
                } catch (error) {
                    console.error("Error fetching schedule data:", error);
                    reject({
                        status: 500,
                        error: 'Failed to fetch schedule data'
                    });
                }
            };

            fetchData();
        });

        const [dropdownData, scheduleData] = await Promise.all([
            dropdownDataPromise,
            scheduleDataPromise
        ]);

        console.log("Dropdown data result:", dropdownData);

        return {     
            dropdownData,
            scheduleData
        };
    } catch (error) {
        console.error("Error loading data:", error);
        return {
            status: 500,
            error: 'Failed to load data'
        };
    }
};



// export const load: Load = async (ctx) => {
//     try {
//         const dropdownUrl = `${PUBLIC_BACKEND_URL}/schedules/get-dropdown-data`;
//         const scheduleUrl = `${PUBLIC_BACKEND_URL}/schedules/get-schedule/${ctx.params.id}`;
        
//         console.log("Fetching dropdown data from:", dropdownUrl);
//         console.log("Fetching schedule data from:", scheduleUrl);

//         const [dropdownResponse, scheduleResponse] = await Promise.all([
//             ctx.fetch(dropdownUrl),
//             ctx.fetch(scheduleUrl)
//         ]);

//         console.log("am i here")

//         if (!dropdownResponse.ok) {
//             throw new Error(`Failed to fetch dropdown data: ${dropdownResponse.statusText}`);
//         }

//         if (!scheduleResponse.ok) {
//             throw new Error(`Failed to fetch schedule data: ${scheduleResponse.statusText}`);
//         }

//         const dropdownData = await dropdownResponse.json() as Schedule[];
//         const scheduleData = await scheduleResponse.json();

//         return {
//             dropdownData,
//             scheduleData
//         };
//     } catch (error) {
//         console.error("Error loading data:", error);
//         return {
//             status: 500,
//             error: 'Failed to load data'
//         };
//     }
// };
