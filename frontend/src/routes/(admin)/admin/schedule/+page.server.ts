import type { Load } from '@sveltejs/kit';
import type { Schedule } from '../../../../types/global';
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

        const getDropdownData = new Promise<Schedule[]>((resolve) => {
            fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/get-dropdown-data`)
            .then(async (response) => {
                const data = await response.json();
                resolve(data);
            })
            .catch((error) => {
                console.error("Failed to fetch dropdown data:", error);
                resolve([]);
            });
        });

        const [schedules, dropdownData] = await Promise.all([getSchedules, getDropdownData]);

        return {
            schedules,
            dropdownData
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
  