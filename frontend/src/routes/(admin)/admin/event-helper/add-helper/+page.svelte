<script lang="ts">
	import { onMount } from 'svelte';
	import type { EventHelper } from '$lib/types/global';

	export let data
  	const { dropdownData } = data

	let carplates: string[] = [];
	let emails: string[] = [];

	let selectedCarplate: string | null = null;
	// let selectedEmail: string | null = null;
	let selectedEmails: Set<string> = new Set();
	let selectedShift: boolean | null = null;

	function setEventHelperDropdownOptions() {
		if (!dropdownData) return;
		
		const uniqueCarplates = new Set<string>();
		const uniqueEmails = new Set<string>();

		dropdownData.forEach(({ Carplate, Email }: EventHelper) => {
			if (Carplate) {
				uniqueCarplates.add(Carplate);
			}
			if (Email) {
				uniqueEmails.add(Email);
			}
		});

		carplates = Array.from(uniqueCarplates);
		emails = Array.from(uniqueEmails);
	}

	function toggleEmailSelection(email: string) {
		if (selectedEmails.has(email)) {
			selectedEmails.delete(email);
		} else {
			selectedEmails.add(email);
		}
	}


	onMount(() => {
		if (dropdownData) {
			setEventHelperDropdownOptions()
		}
	});
</script>

<div class="flex justify-center items-center h-full">
	<div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-20">
		<h1 class="text-2xl font-semibold mb-4">Add New Event Helper</h1>
		<form method="POST" action="?/createEventHelper">
			<div class="mb-4">
				<label for="carplate" class="block text-sm font-medium mb-1">Carplate:</label>
				<select
					id="carplate"
					name="carplate"
					bind:value={selectedCarplate}
					required
					class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500"
				>
					{#each carplates as carplate}
						<option value={carplate}>{carplate}</option>
					{/each}
				</select>
			</div>

			<!-- <div class="mb-4">
				<label for="email" class="block text-sm font-medium mb-1">Email:</label>
				<select
					id="email"
					name="email"
					bind:value={selectedEmail}
					class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500"
				>
					{#each emails as email}
						<option value={email}>{email}</option>
					{/each}
				</select>
			</div> -->

			<div class="mb-4">
				<fieldset>
					<legend class="block text-sm font-medium mb-1">Emails:</legend>
					{#each emails as email}
						<div class="flex items-center mb-2">
							<input
								type="checkbox"
								name="email"
								id={email}
								value={email}
								on:change={() => toggleEmailSelection(email)}
								class="mr-2"
							/>
							<label for={email} class="text-sm">{email}</label>
						</div>
					{/each}
				</fieldset>
			</div>

			<div class="mb-4">
				<label for="shift" class="block text-sm font-medium mb-1">Shift:</label>
				<select
					id="shift"
					name="shift"
					bind:value={selectedShift}
					required
					class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500"
				>
					<option value="true">AM</option>
					<option value="false">PM</option>
				</select>
			</div>

			<div class="mt-4 flex justify-center">
				<button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800">
					Add Event Helper
				</button>
			</div>
		</form>
	</div>
</div>
