<script lang="ts">
    import { onMount } from "svelte";
    import type { EventHelper } from "$lib/types/global";
	import { page } from "$app/stores";
	import CustomDropdown from "$lib/components/CustomDropdown.svelte";

    export let data;
    const { dropdownData } = data;

    let helper: EventHelper = {BusId: '', Carplate: '', Name: '', Shift: false};
    let carplates: string[] = [];
    let names: string[] = [];
    let selectedCarplate: string | null = null;
    let selectedName: string | null = null;
    let selectedShift: string

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
        names = Array.from(uniqueNames)
    }

    onMount(() => {
        const { index } = $page.params; 
        helper = JSON.parse(decodeURIComponent(index));
        selectedCarplate = helper.Carplate;
        selectedName = helper.Name;
        selectedShift = helper.Shift ? "AM" : "PM"

        if (dropdownData) {
            setEventHelperDropdownOptions()
        }
        
    });

    const shiftOptions = [true, false]
</script>

<div class="flex justify-center items-center h-full">
    <div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-24">
        <h1 class="text-2xl font-semibold mb-4">Update Selected Event Helper</h1>
        <form method="POST" action="?/updateEventHelper">
            <input type="hidden" name="old_carplate" value={helper.Carplate}>
            <input type="hidden" name="old_name" value={helper.Name}>
            <input type="hidden" name="old_shift" value={helper.Shift}>

            <!-- <div class="mb-4">
                <label for="carplate" class="block text-sm font-medium mb-1">Carplate:</label>
                <select id="carplate" name="carplate" required bind:value={selectedCarplate} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500">
                    {#each carplates as carplate}
                        <option value={carplate}>{carplate}</option>
                    {/each}
                </select>
            </div>
            
            <div class="mb-4">
                <label for="name" class="block text-sm font-medium mb-1">Name:</label>
                <select id="name" name="name" required bind:value={selectedName} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500">
                    {#each names as name}
                        <option value={name}>{name}</option>
                    {/each}
                </select>
            </div> -->
            
            <div class="mb-4">
                <CustomDropdown
                    label="Carplate"
                    name="carplate"
                    options={carplates}
                    required
                    bind:selected={selectedCarplate}
                />
            </div>
            
            <!-- this should be a fieldset no? -->
            <div class="mb-4">
                <CustomDropdown
                    label="Name"
                    name="name"
                    options={names}
                    required
                    bind:selected={selectedName}
                />
            </div>

            <!-- same issue, getting from backend in boolean, unable to translate into string in displayOption -->
            <CustomDropdown
				label="Shift"
				name="shift"
				options={shiftOptions}
				bind:selected={selectedShift}
			/>

            <!-- <div class="mb-4">
                <label for="shift" class="block text-sm font-medium mb-1">Shift:</label>
                <select id="shift" name="shift" required bind:value={selectedShift} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500" >
                    <option value={true}>AM</option>
                    <option value={false}>PM</option>
                </select>
            </div> -->
<!-- 
            <div class="mb-4">
                <label for="endTime" class="block text-sm font-medium mb-1">End Time:</label>
                <input type="datetime-local" id="end_time" name="end_time" required bind:value={selectedEndTime} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500" />
            </div> -->

            <div class="mt-4 flex justify-center">
                <button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800">Update Event Helper</button>
            </div>
        </form>
    </div>
</div>

