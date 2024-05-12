import type { EventHelper } from '../../../../../../types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

export const load = async ({ fetch }) => {
    try {
        const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event-helpers/get-event-dropdown`);
        if (!response.ok) {
            throw new Error("Failed to fetch event helpers");
        }
        
        const data = await response.json() as EventHelper[];
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


export const actions = {
    updateEventHelper: async({ request}): Promise<void> =>{
      const form =await request.formData()
  
      const Carplate = form.get('carplate');
      const Email = form.get('email');
      const StartTime = form.get('start_time') + ":00+08:00";
      const EndTime = form.get('end_time') + ":00+08:00";
  
      const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event-helpers/update-helper`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            NewCarplate: Carplate,
            NewEmail: Email,
            NewStartTime: StartTime,
            NewEndTime: EndTime,
            OldCarplate: form.get('old_carplate'),
            OldEmail: form.get('old_email'),
            OldStartTime: form.get('old_start_time'),
            OldEndTime: form.get('old_end_time'),
            })
      });
    
      // whats the correct way of navigating in server side svelte? goto and redirect doesnt work
      if (!response.ok) {
        throw new Error("Failed to update event helper");
      }
    }    
  }
  
