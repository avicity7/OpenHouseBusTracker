import type { EventHelper, EventBus } from '$lib/types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';
import { redirect } from '@sveltejs/kit';

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
    updateEventHelper: async({ request}): Promise<void> =>{
      const form = await request.formData()
  
      const Carplate = form.get('carplate');
      const BusId = JSON.parse(Carplate!.toString()).BusId
      const Name = form.get('name');
      const ShiftString = form.get('shift');

      const Shift = ShiftString === 'true';

      const OldBusId = form.get('old_bus_id');
      const OldName = form.get('old_name');
      const OldShiftString =  form.get('old_shift');

      const OldShift = OldShiftString === 'true';

      const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/event-helpers/update-helper`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            NewBusId: BusId,
            NewName: Name,
            NewShift: Shift,
            OldBusId,
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
  
