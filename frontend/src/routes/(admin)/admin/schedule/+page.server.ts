import type { Load } from '@sveltejs/kit';
import type { EventBus, Route, Schedule } from '$lib/types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';


export const load: Load = async ({ fetch }) => {
    try {
        const getSchedules = new Promise<Schedule[]>((resolve) => {
            fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/get-schedules`)
            .then(async (response) => {
                const data = await response.json() as Schedule[];
                resolve(data);
            })
            .catch((error) => {
                console.error("Failed to fetch bus schedules:", error);
                resolve([]);
            });
        });

        const getRoutes = new Promise<Route>((resolve) => {
            fetch(`${PUBLIC_BACKEND_URL}:3000/route/`)
            .then(async (response) => {
                const data = await response.json();
                resolve(data);
            })
            .catch((error) => {
                console.error("Failed to fetch routes data:", error);
            });
        });

        const getCarplates = new Promise<EventBus>((resolve) => {
            fetch(`${PUBLIC_BACKEND_URL}:3000/bus/get-buses`)
            .then(async (response) => {
                const data = await response.json();
                resolve(data);
            })
            .catch((error) => {
                console.error("Failed to fetch buses data:", error);
            });
        });


        const [schedules, routes, carplates] = await Promise.all([getSchedules, getRoutes, getCarplates]);

        return {
            schedules,
            routes,
            carplates
        };
    } catch (error) {
        console.error("Internal Server Error:", error);
        return {
            status: 500,
            body: { error: 'Internal Server Error' }
        };
    }
};


// export const load: Load = async ({ fetch }) => {
//     try {
//         const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/get-schedules`);
//         if (!response.ok) {
//             throw new Error("Failed to fetch bus schedules");
//         }
        
//         const data = await response.json() as Schedule[];
//         return {
//             data
//         };

//     } catch (error) {
//         console.error(error);
//         return {
//             status: 500,
//             body: { error: 'Internal Server Error' }
//         };
//     }
// };

// export const actions = {
//     deleteBusSchedule: async (scheduleId) => {
//       const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/delete-schedule/${scheduleId}`, {
//         method: 'DELETE'
//       });
  
//       if (!response.ok) {
//         throw new Error("Failed to delete bus schedule");
//       }
//     }
//   };
  