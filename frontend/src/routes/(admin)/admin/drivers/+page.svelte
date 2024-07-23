<script lang="ts">
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';
	import { jsPDF } from 'jspdf';
	import ToolTip from '$lib/components/ToolTip.svelte';
	import type { Driver } from '$lib/types/global';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import autoTable from 'jspdf-autotable';

	export let data;
	let { ScheduleTimeDiff, drivers } = data;

	if (!drivers) drivers = [];
	let driverHours = data.ScheduleTimeDiff;
	let search = '';
	let driverToDelete: Driver | null = null;
	let timers = writable<Record<number, { startTime: number | null; elapsedTime: number }>>({});
	let payRate = 0;

	const getDrivers = async () => {
		const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/driver/get-drivers`);
		if (response.ok) {
			let data = (await response.json()) as Driver[];
			drivers = data;
		}
	};

	const convertToHours = (nanoseconds: number): number => {
		return nanoseconds / (1000 * 1000 * 1000 * 60 * 60);
	};

	function generateData() {
		const driversData = drivers!;
		return driversData.map((driver) => {
			const driverHour = ScheduleTimeDiff.find(
				(d: { DriverId: number }) => d.DriverId === driver.DriverId
			);
			const hoursWorked = driverHour ? convertToHours(driverHour.TimeDifference) : 0;

			const pay = Math.abs(payRate * hoursWorked);

			return [driver.DriverName, hoursWorked.toFixed(2), pay.toFixed(2)];
		});
	}

	async function exportAllDriversToPDF() {
		const doc = new jsPDF({ putOnlyUsedFonts: true, orientation: 'landscape' });

		doc.text('Driver Paysheet', 15, 10);

		try {
			autoTable(doc, {
				startY: 20,
				theme: 'striped',
				head: [['Driver', 'Hours Worked', 'Pay']],
				headStyles: { fillColor: [185, 28, 28] },
				body: generateData(),
				didDrawCell: (data) => {
					if (data.column.index === 0) {
						data.cell.styles.cellWidth = 80;
					}
				}
			});
		} catch (error) {
			console.error('Error generating table:', error);
		}

		doc.save('drivers_paysheet.pdf');
	}

	function exportDriverToPDF(driver: Driver) {
		const doc = new jsPDF();

		const logoUrl = '/favicon.png';
		const logoWidth = 15;
		const logoHeight = 15;

		doc.addImage(logoUrl, 'PNG', 10, 10, logoWidth, logoHeight);

		doc.setFont('helvetica', 'bold');
		doc.setFontSize(14);

		doc.text('Driver Pay Statement', 10, 35);

		doc.setFontSize(12);
		doc.setFont('helvetica', 'normal');

		doc.text(`Driver Name: ${driver.DriverName}`, 10, 50);
		doc.text(`Pay Rate: $${payRate.toFixed(2)} per hour`, 10, 60);

		const driverHour = ScheduleTimeDiff.find(
			(d: { DriverId: number }) => d.DriverId === driver.DriverId
		);
		const hoursWorked = driverHour ? convertToHours(driverHour.TimeDifference) : 0;
		const pay = Math.abs(payRate * hoursWorked);

		doc.text(`Hours Worked: ${hoursWorked.toFixed(2)} hours`, 10, 70);
		doc.text(`Total Pay: $${pay.toFixed(2)}`, 10, 80);

		doc.setFont('helvetica', 'italic');
		doc.text(
			`This document serves as a formal statement of hours worked and pay for the period specified.`,
			10,
			100
		);
		doc.text(
			`Please review the information provided and sign below to confirm receipt of payment.`,
			10,
			110
		);

		doc.setDrawColor(0, 0, 0);
		doc.line(10, 140, 80, 140);
		doc.text('Signature', 10, 145);

		doc.setDrawColor(0, 0, 0);
		doc.line(100, 140, 170, 140);
		doc.text('Date', 100, 145);

		doc.save(`${driver.DriverName}_paysheet.pdf`);
	}

	async function deleteDriver(driver: Driver) {
		driverToDelete = driver;
	}

	function confirmDelete() {
		if (driverToDelete) {
			fetch(`${PUBLIC_BACKEND_URL}:3000/driver/delete-driver`, {
				method: 'DELETE',
        body: JSON.stringify(driverToDelete)
			}).then((response) => {
				if (response.ok) {
					getDrivers();
				}
			});
			driverToDelete = null;
		}
	}

	function cancelDelete() {
		driverToDelete = null;
	}

	onMount(() => {
		getDrivers();
	});
</script>

<svelte:head>
	<title>Manage - Drivers | SPOH Bus Tracker</title>
</svelte:head>

<div class="p-6 md:p-12">
	<div class="flex items-center justify-between mb-4">
		<a
			href="/admin/drivers/add-driver"
			class="border-black text-white font-semibold text-md px-6 py-2 rounded-xl bg-red-700 hover:bg-red-800 mr-2"
			data-testid="add-driver-button"
		>
			Add Driver
		</a>

		<div class="ml-auto">
			<input
				type="text"
				class="border border-gray-300 rounded-md px-3 py-2 w-60"
				bind:value={search}
				placeholder="Type to search drivers...."
			/>
		</div>
	</div>

	<div class="mt-8">
		<table class="min-w-full divide-y divide-gray-200">
			<thead class="bg-gray-50">
				<tr>
					<th
						class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider"
						>Driver Name</th
					>
					<th
						class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider"
						>Actions</th
					>
				</tr>
			</thead>
			<tbody class="bg-white divide-y divide-gray-200">
				{#if drivers}
					{#each drivers.filter( (driver) => driver.DriverName.toLowerCase().includes(search.toLowerCase()) ) as driver}
						<tr class="hover:bg-gray-100">
							<td class="px-6 py-4 whitespace-nowrap text-center">{driver.DriverName}</td>
							<td class="px-6 py-2 whitespace-nowrap text-sm font-medium ">
								<div class="flex items-center justify-center ml-2 gap-4">
									<a
										href={`drivers/update-driver/${encodeURIComponent(JSON.stringify(driver))}`}
										class="text-stone-500 hover:text-green-500"
									>
										<ToolTip text="Update Driver">	
											<svg
												xmlns="http://www.w3.org/2000/svg"
												class="h-5 w-5 mr-1"
												viewBox="0 0 24 24"
												fill="none"
												stroke="currentColor"
												stroke-width="2"
												stroke-linecap="round"
												stroke-linejoin="round"
											>
												<path
													d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"
												/>
											</svg>
										</ToolTip>
									</a>
									<button
										class="text-stone-500 hover:text-red-600 text-2xl"
										on:click={() => deleteDriver(driver)}
									>
										<ToolTip text="Delete Driver">
											<svg
												xmlns="http://www.w3.org/2000/svg"
												width="1em"
												height="1em"
												viewBox="0 0 24 24"
												{...$$props}
											>
												<svg
													xmlns="http://www.w3.org/2000/svg"
													width="1em"
													height="1em"
													viewBox="0 0 24 24"
													{...$$props}
												>
													<g fill="none">
														<path
															d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022m-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z"
														/>
														<path
															fill="currentColor"
															d="M14.28 2a2 2 0 0 1 1.897 1.368L16.72 5H20a1 1 0 1 1 0 2l-.003.071l-.867 12.143A3 3 0 0 1 16.138 22H7.862a3 3 0 0 1-2.992-2.786L4.003 7.07A1.01 1.01 0 0 1 4 7a1 1 0 0 1 0-2h3.28l.543-1.632A2 2 0 0 1 9.721 2zm3.717 5H6.003l.862 12.071a1 1 0 0 0 .997.929h8.276a1 1 0 0 0 .997-.929zM10 10a1 1 0 0 1 .993.883L11 11v5a1 1 0 0 1-1.993.117L9 16v-5a1 1 0 0 1 1-1m4 0a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1m.28-6H9.72l-.333 1h5.226z"
														/>
													</g>
												</svg>
											</svg>
										</ToolTip>
									</button>
									<!-- colour undetermined LOL -->
									<button
										class="text-stone-500 hover:text-red-600 text-2xl"
										on:click={() => exportDriverToPDF(driver)}
									>
										<ToolTip text="Export to PDF">
											<svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" {...$$props}>
												<g fill="none" fill-rule="evenodd">
													<path d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035q-.016-.005-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427q-.004-.016-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093q.019.005.029-.008l.004-.014l-.034-.614q-.005-.019-.02-.022m-.715.002a.02.02 0 0 0-.027.006l-.006.014l-.034.614q.001.018.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z" />
													<path fill="currentColor" d="M13.586 2a2 2 0 0 1 1.284.467l.13.119L19.414 7a2 2 0 0 1 .578 1.238l.008.176V20a2 2 0 0 1-1.85 1.995L18 22H6a2 2 0 0 1-1.995-1.85L4 20V4a2 2 0 0 1 1.85-1.995L6 2zM12 4H6v16h12V10h-4.5a1.5 1.5 0 0 1-1.493-1.356L12 8.5zm.988 7.848a6.22 6.22 0 0 0 2.235 3.872c.887.717.076 2.121-.988 1.712a6.22 6.22 0 0 0-4.47 0c-1.065.41-1.876-.995-.989-1.712a6.22 6.22 0 0 0 2.235-3.872c.178-1.127 1.8-1.126 1.977 0m-.99 2.304l-.688 1.196h1.38zM14 4.414V8h3.586z" />
												</g>
											</svg>
										</ToolTip>
									</button>
								</div>
							</td>
						</tr>
					{/each}
				{/if}
			</tbody>
		</table>
	</div>

	<div class="mt-8">
		<div>
			<label for="payRate">Pay Rate:</label>
			<input id="payRate" type="number" bind:value={payRate} class="rounded p-1 w-20 mb-2" />
		</div>
		<button
			on:click={exportAllDriversToPDF}
			class="border-black text-white font-semibold text-md px-6 py-2 rounded-xl bg-red-700 hover:bg-red-800"
		>
			Export Data to PDF
		</button>
	</div>

	{#if driverToDelete}
		<div
			class="fixed top-0 left-0 w-full h-full flex items-center justify-center bg-gray-900 bg-opacity-50"
		>
			<div class="bg-white p-8 rounded-lg shadow-lg">
				<p class="text-lg mb-4">Are you sure you want to delete this driver?</p>
				<div class="flex justify-end">
					<button class="px-4 py-2 mr-4 text-gray-600" on:click={cancelDelete}>Cancel</button>
					<button class="px-4 py-2 bg-red-700 text-white rounded" on:click={confirmDelete}
						>Delete</button
					>
				</div>
			</div>
		</div>
	{/if}
</div>
