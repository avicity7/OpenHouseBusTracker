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
	let filteredNames: string[] = [];
	let searchQuery: string = '';

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
		filterNames();
	}

	function filterNames() {
		filteredNames = names.filter(name =>
			name.toLowerCase().includes(searchQuery.toLowerCase())
		);
	}

	function toggleNameSelection(name: string, checked: boolean) {
		if (checked) {
		selectedNames.add(name);
		} else {
		selectedNames.delete(name);
		}
	}

	function handleCheckboxChange(event: Event, name: string) {
		const target = event.target as HTMLInputElement;
		toggleNameSelection(name, target.checked);
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
	
	$: searchQuery, filterNames();

	const shiftOptions = [
		{ label: "AM", value: true },
		{ label: "PM", value: false }
	];
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
					searchable
					bind:selected={selectedCarplate}
				/>
			</div>

			<div class="mb-4">
				<fieldset>
					<legend class="block text-sm font-medium mb-1">Names:</legend>
					<div class="mb-4 ">
						<input
							type="text"
							placeholder="Search names..."
							bind:value={searchQuery}
							class="block w-full px-3 py-2 border rounded-md text-sm focus:border-red-600 focus:outline-none"
							/>
					</div>
					{#if filteredNames.length === 0 || filteredNames === null}
						<p class="text-sm text-gray-500">No options available</p>
					{:else}
						{#each filteredNames as name}
							<div class="flex items-center mb-2">
								<input
									type="checkbox"
									name="name"
									id={name}
									value={name}
									on:change={(e) => handleCheckboxChange(e, name)}
									class="mr-2"
									checked={selectedNames.has(name)}
								/>
								<label for={name} class="text-sm">{name}</label>
							</div>
						{/each}
					{/if}
				</fieldset>
			</div>

			<div class="mb-4">
				<fieldset>
					<legend class="block text-sm font-medium mb-1">Shift:</legend>
					<div class="flex items-center space-x-4">
						{#each shiftOptions as { label, value }}
							<div class="flex items-center">
								<input
									type="radio"
									name="shift"
									id={label}
									value={value}
									bind:group={selectedShift}
									class="mr-2"
									required
								/>
								<label for={label} class="text-sm">{label}</label>
							</div>
						{/each}
					</div>
				</fieldset>
			</div>

			<div class="mt-4 flex justify-center">
				<button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800">
					Add Event Helper
				</button>
			</div>
		</form>
	</div>
</div>
