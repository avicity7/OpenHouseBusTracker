import type { Load } from '@sveltejs/kit';
import type { Schedule } from '../../../../types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

// For initial load using SSR
export const load: Load = async ({ fetch }) => {
    try {
        const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/get-dropdown-data`);
        if (!response.ok) {
            throw new Error("Failed to fetch dropdown data");
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

// For form actions, submissions    
  export const actions = {
    createBusSchedule: async({ request}): Promise<void> =>{
      const form =await request.formData()

        const Carplate = form.get('carplate');
        const RouteName = form.get('route_name');
        const DriverIdString = form.get('driver_id');
        const DriverId = DriverIdString ? +DriverIdString : null;
         
        const StartTime = form.get('start_time') + ":00Z";
        const EndTime = form.get('end_time') + ":00Z";

        // console.log('Retrieved start_time:', StartTime);
        // console.log('Retrieved end_time:', EndTime);       
        const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/add-schedule`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
                  Carplate,
                  RouteName,
                  DriverId,
                  StartTime,
                  EndTime
              })
        });
      
        if (!response.ok) {
          throw new Error("Failed to create bus schedule");
        }
      }    
  }

// export const actions = {
//   createSchedule: async (formData: FormData): Promise<void> => {
//       try {
//      
//           const carplate = formData.get('carplate');
//           const routeName = formData.get('routeName');
//           const driverId = formData.get('driverId');
//           const startTime = formData.get('startTime');
//           const endTime = formData.get('endTime');
  
//        
//           const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/add-schedule`, {
//               method: 'POST',
//               headers: {
//                   'Content-Type': 'application/json'
//               },
//               body: JSON.stringify({
//                   carplate,
//                   routeName,    
//                   driverId,
//                   startTime,
//                   endTime
//               })
//           });
  
//   
//           if (!response.ok) {
//               throw new Error('Failed to create schedule');
//           }
  
//           console.log('Schedule created successfully');
//       } catch (error) {
//           console.error('Error creating schedule:', error);
//           throw error;
//       }
//   }
// };