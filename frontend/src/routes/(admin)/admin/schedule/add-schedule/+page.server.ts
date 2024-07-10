import { fail, redirect, type Load } from '@sveltejs/kit';
import type { Schedule } from '$lib/types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

// For initial load using SSR
export const load: Load = async ({ fetch }) => {
    try {
        const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/get-dropdown-data`);
        if (!response.ok) {
            throw new Error("Failed to fetch dropdown data");
        }
        
        const dropdownData = await response.json() as Schedule[];
        return {
            dropdownData
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
  createBusSchedule: async({ request}) =>{
    const form =await request.formData()

    const Carplate = form.get('carplate');
    const RouteName = form.get('route_name');
    const DriverIdString = form.get('driver_id');
    const DriverId = DriverIdString ? +DriverIdString : null;
      
    const StartTime = form.get('start_time') + ":00+08:00";
    const EndTime = form.get('end_time') + ":00+08:00";
  
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
      const errorText = await response.text();
      console.error('Server Error:', errorText);
      console.error('Server Code:', response.status);
      return fail(response.status, { error: errorText });
    }

    throw redirect(301, '/admin/schedule')
  }    
}
