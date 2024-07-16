<script lang="ts">
	import { onMount } from 'svelte';
	import type { EventHelper } from '$lib/types/global';
	import { page } from '$app/stores';
	import ErrorMessage from '$lib/components/ErrorMessage.svelte';
	import CustomDropdown from '$lib/components/CustomDropdown.svelte';

	export let data;
	const { dropdownData, buses } = data;

	let carplates: string[] = [];
	let names: string[] = [];

	let selectedCarplate: string;
	let selectedNames: Set<string> = new Set();
	let selectedShift: boolean;
	let errorMessage: string;

	function setEventHelperDropdownOptions() {
		if (!dropdownData) return;

		const uniqueCarplates = new Set<string>();
		const uniqueNames = new Set<string>();
			
		dropdownData.forEach(({ Carplate, Name }: EventHelper) => {
			if (Carplate) {
				uniqueCarplates.add(Carplate);
			}
			if (Name) {
				uniqueNames.add(Name);
			}
		});

		carplates = Array.from(uniqueCarplates);
		names = Array.from(uniqueNames);
	}

	function toggleNameSelection(name: string) {
		if (selectedNames.has(name)) {
			selectedNames.delete(name);
		} else {
			selectedNames.add(name);
		}
	}

	if ($page.status === 409) {
		errorMessage =
			$page.error?.message || 'This event helper(s) has already been assigned to a bus.';
	}

	onMount(() => {
		if (dropdownData) {
			setEventHelperDropdownOptions();
		}
	});

	const shiftOptions = [true, false];
</script>

<div class="flex justify-center items-center h-full">
	<div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-20">
		<h1 class="text-2xl font-semibold mb-4">Add New Event Helper</h1>
		<ErrorMessage message={errorMessage} />
		<form method="POST" action="?/createEventHelper">
			<div class="mb-4">
				<CustomDropdown
					label="Bus"
					name="bus"
					options={buses}
					required
					bind:selected={selectedCarplate}
				/>
			</div>

			<div class="mb-4">
				<fieldset>
					<legend class="block text-sm font-medium mb-1">Names:</legend>
					{#if names.length === 0 || names === null}
						<p class="text-sm text-gray-500">No options available</p>
					{:else}
						{#each names as name}
							<div class="flex items-center mb-2">
								<input
									type="checkbox"
									name="name"
									id={name}
									value={name}
									on:change={() => toggleNameSelection(name)}
									class="mr-2"
								/>
								<label for={name} class="text-sm">{name}</label>
							</div>
						{/each}
					{/if}
				</fieldset>
			</div>

			<!-- <div class="mb-4">
				<label for="shift" class="block text-sm font-medium mb-1">Shift:</label>
				<select
					id="shift"
					name="shift"
					bind:value={selectedShift}
					required
					class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500"
				>
					<option value="true">AM</option>
					<option value="false">PM</option>
				</select>
			</div> -->

			<CustomDropdown
				label="Shift"
				name="shift"
				options={shiftOptions}
				bind:selected={selectedShift}
			/>

			<div class="mt-4 flex justify-center">
				<button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800">
					Add Event Helper
				</button>
			</div>
		</form>
	</div>
</div>
