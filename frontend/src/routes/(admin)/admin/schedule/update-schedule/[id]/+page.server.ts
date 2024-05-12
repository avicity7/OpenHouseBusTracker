import { error, type Load } from '@sveltejs/kit';
import type { Schedule } from '../../../../../../lib/types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

// used page.server.ts instead of page.js for the need of updated information and fast ssr
export const load: Load = async (ctx) => {
    try {
        const dropdownDataPromise = new Promise((resolve, reject) => {
            const fetchData = async () => {
                try {
                    const dropdownUrl = `${PUBLIC_BACKEND_URL}:3000/schedules/get-dropdown-data`;

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
                    const scheduleUrl = `${PUBLIC_BACKEND_URL}:3000/schedules/get-schedule/${id}`;
                    const response = await ctx.fetch(scheduleUrl);
                    if (!response.ok) {
                        throw new Error(`Failed to fetch schedule data from id: ${id}: ${response.statusText}`);
                    }

                    const schedule = await response.json();
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

export const actions = {
    updateBusSchedule: async({ request}): Promise<void> =>{
      const form =await request.formData()
        
        const url = new URL(request.url);
        const id = url.pathname.split('/').pop();

        const BusScheduleId = id ? +id : undefined;
        const Carplate = form.get('carplate');
        const RouteName = form.get('route_name');
        const DriverIdString = form.get('driver_id');
        const DriverId = DriverIdString ? +DriverIdString : null;
        const StartTime = form.get('start_time') + ":00Z";
        const EndTime = form.get('end_time') + ":00Z";
        
        const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/update-schedule`, {
          method: "PUT",
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
                  Carplate, 
                  RouteName,
                  DriverId,
                  StartTime,
                  EndTime,
                  BusScheduleId
              })
        });
        
        console.log(`update bus schedule ${id} successful`)
        if (!response.ok) {
            console.log(error)
          throw new Error("Failed to create bus schedule");
        }
      }
  }


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
