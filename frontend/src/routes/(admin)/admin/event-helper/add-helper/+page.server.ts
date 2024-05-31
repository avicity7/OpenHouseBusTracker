import type { EventHelper } from '$lib/types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';
import { redirect } from '@sveltejs/kit';

export const load = async ({ fetch }) => {
    try {
        const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event-helpers/get-event-dropdown`);
        if (!response.ok) {
            throw new Error("Failed to fetch event helpers");
        }
        
        const dropdownData = await response.json() as EventHelper[];
        return {
            dropdownData
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
    createEventHelper: async({ request}): Promise<void> =>{
      const form = await request.formData()
  
      const Carplate = form.get('carplate');
      const Name = form.get('Name');
      const ShiftString = form.get('shift');
      
      const Shift = ShiftString === 'true';

      const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event-helpers/create-helper`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
                Carplate,
                Name,
                Shift,
            })
      });
    
      if (!response.ok) {
        throw new Error("Failed to create event helper");
      }

      redirect(301, '/admin/event-helper')
    }    
  }