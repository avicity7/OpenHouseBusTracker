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
    updateEventHelper: async({ request}): Promise<void> =>{
      const form = await request.formData()
  
      const Carplate = form.get('carplate');
      const Name = form.get('name');
      const ShiftString = form.get('shift');

      const Shift = ShiftString === 'true';

      const OldCarplate = form.get('old_carplate');
      const OldName = form.get('old_name');
      const OldShiftString =  form.get('old_shift');

      const OldShift = OldShiftString === 'true';

      const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event-helpers/update-helper`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            NewCarplate: Carplate,
            NewName: Name,
            NewShift: Shift,
            OldCarplate,
            OldName,
            OldShift,
            })
      });
    
      if (!response.ok) {
        throw new Error("Failed to update event helper");
      }

      redirect(301, '/admin/event-helper')
    }    
  }
  
