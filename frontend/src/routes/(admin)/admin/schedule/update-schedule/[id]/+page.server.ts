import { error, redirect, type Load } from '@sveltejs/kit';
import type { Route, EventBus, Driver, DropdownData, Schedule } from '$lib/types/global';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

export const load: Load = async ({ fetch, params }) => {
	const { id } = params
	let response = await fetch(`${PUBLIC_BACKEND_URL}:3000/bus/get-buses`);
	if (!response.ok) {
		throw new Error("Failed to fetch dropdown data");
	}
	const buses = await response.json() as EventBus[]

	response = await fetch(`${PUBLIC_BACKEND_URL}:3000/driver/get-drivers`);
	if (!response.ok) {
		throw new Error("Failed to fetch dropdown data");
	}
	const drivers = await response.json() as Driver[]

	response = await fetch(`${PUBLIC_BACKEND_URL}:3000/route/`);
	let routes = (await response.json()) as Array<Route>;
	if (!routes) {
		routes = [];
	}

	response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/get-dropdown-data`);
	if (!response.ok) {
		throw new Error("Failed to fetch dropdown data");
	}
	const dropdownData = await response.json() as DropdownData;

	const scheduleUrl = `${PUBLIC_BACKEND_URL}:3000/schedules/get-schedule/${id}`;
	response = await fetch(scheduleUrl);
	if (!response.ok) {
		throw new Error(`Failed to fetch schedule data from id: ${id}: ${response.statusText}`);
	}

	const schedule = await response.json() as Schedule;

	return {
		buses,
		drivers,
		routes,
		dropdownData,
		schedule
	};
};

export const actions = {
	updateBusSchedule: async ({ request }): Promise<void> => {
		const form = await request.formData()

		const url = new URL(request.url);
		const id = url.pathname.split('/').pop();

		const BusScheduleId = id ? +id : undefined;
		const Bus = form.get('bus')
		const BusId = JSON.parse(Bus!.toString()).BusId
		const Route = form.get('route');
		const RouteName = JSON.parse(Route!.toString()).RouteName
		const Driver = form.get('driver')
		const DriverId = JSON.parse(Driver!.toString()).DriverId
		const StartTime = form.get('start_time') + ":00+08:00";
		const EndTime = form.get('end_time') + ":00+08:00";

		const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/schedules/update-schedule`, {
			method: "PUT",
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({
				BusId,
				RouteName,
				DriverId,
				StartTime,
				EndTime,
				BusScheduleId
			})
		});

		if (!response.ok) {
			console.log(error)
			throw new Error("Failed to create bus schedule");
		}

		redirect(301, '/admin/schedule')
	}
}