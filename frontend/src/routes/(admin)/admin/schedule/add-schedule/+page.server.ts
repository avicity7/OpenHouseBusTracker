import { fail, redirect, type Load } from '@sveltejs/kit';
import type { DropdownData, EventBus } from '$lib/types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

export const load: Load = async ({ fetch }) => {
    try {
        let response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/get-dropdown-data`);
        if (!response.ok) {
            throw new Error("Failed to fetch dropdown data");
        }
        
        let dropdownData = await response.json() as DropdownData;
        if (!dropdownData) {
          dropdownData = {} as DropdownData
        }

        response = await fetch(`${PUBLIC_BACKEND_URL}:3000/bus/get-buses`);
        if (!response.ok) {
            throw new Error("Failed to fetch dropdown data");
        }
        const buses = await response.json() as EventBus[]
        
        return {
          dropdownData,
          buses
        };

    } catch (error) {
        console.error(error);
        return {
            status: 500,
            body: { error: 'Internal Server Error' }
        };
    }
};

export const actions = {
  createBusSchedule: async({ request }) =>{
    const form = await request.formData()

    const Bus = form.get('bus');
    const BusId = JSON.parse(Bus!.toString()).BusId
    const Route = form.get('route_name');
    const RouteName = JSON.parse(Route!.toString()).RouteName
    const DriverIdString = form.get('driver_id');
    const DriverId = JSON.parse(DriverIdString!.toString()).DriverId
      
    const StartTime = form.get('start_time') + ":00+08:00";
    const EndTime = form.get('end_time') + ":00+08:00";

    const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/add-schedule`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        BusId,
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
