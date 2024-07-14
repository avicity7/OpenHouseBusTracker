import { PUBLIC_BACKEND_URL } from "$env/static/public"
import type { EventBus, Route, Schedule } from "$lib/types/global"

export const load = async() => {
  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/bus/get-buses`)
  let buses: Array<EventBus> = []
  const parsed = await response.json() as Array<EventBus>
  if (parsed) {
    buses = parsed
  }

  const routeResponse = await fetch(`${PUBLIC_BACKEND_URL}:3000/route`);
  if (!routeResponse.ok) {
      throw new Error("Failed to fetch routes");
  }
  const routes = await routeResponse.json() as Route[];

  const scheduleBusResponse = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/get-schedules`);
  if (!scheduleBusResponse.ok) {
    throw new Error("Failed to fetch buses that are in schedules");
  }
  
  const scheduleBus = await scheduleBusResponse.json() as Schedule[];
  
  return {
    buses,
    routes,
    scheduleBus
  }
}

export const actions = {
  updateScheduleRoutes: async({ request }) => {
    const formData = await request.formData();
    const assignmentData = JSON.parse(formData.get('assignmentData') as string);
    console.log('Received assignment data:', assignmentData);

    const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/update-schedule-route`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(assignmentData)
    });
  
    if (!response.ok) {
      throw new Error("Failed to update schedule");
    }
  }    
}