import type { EventHelper, EventBus } from '$lib/types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';
import { fail, redirect } from '@sveltejs/kit';

export const load = async ({ fetch }) => {
    try {
        let response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event-helpers/get-event-dropdown`);
        if (!response.ok) {
            throw new Error("Failed to fetch event helpers");
        }
        const dropdownData = await response.json() as EventHelper[];

        response = await fetch(`${PUBLIC_BACKEND_URL}:3000/bus/get-buses`);
        if (!response.ok) {
            throw new Error("Failed to fetch dropdown data");
        }
        const buses = await response.json() as EventBus[]
        return {
            dropdownData,
            buses
        }
    } catch (error) {
        console.error(error);
        return {
            status: 500,
            body: { error: 'Internal Server Error' }
        };
    }
};

export const actions = {
    createEventHelper: async({ request }) =>{
      const form = await request.formData()
  
      const Bus = form.get('bus')
      const BusId = JSON.parse(Bus!.toString()).BusId
      const ShiftString = form.get('shift');
      const Shift = ShiftString === 'true';
      const selectedNames = form.getAll('name');
      const EventHelpers = selectedNames.map(name => ({
        BusId,
        Name: name,
        Shift
      }));

      const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event-helpers/create-helpers`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ EventHelpers })
      });
    
      if (!response.ok) {
        const errorText = await response.text();
        console.error('Server Error in eventhelper:', errorText);
        console.error('Server Code:', response.status);
        return fail(response.status, { error: errorText });
      }

      throw redirect(301, '/admin/event-helper')
    }    
  }