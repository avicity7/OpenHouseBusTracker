<script lang="ts">
    import type { EventBus, EventHelper } from "$lib/types/global";
	import { page } from "$app/stores";
	import CustomDropdown from "$lib/components/CustomDropdown.svelte";

    export let data;
    const { dropdownData, buses } = data;

    let helper: EventHelper = {BusId: '', Carplate: '', Email: '', Name: '', Shift: false};
    let carplates: string[] = [];
    let names: string[] = [];
    let selectedBus: EventBus
    let selectedName: string
    let selectedShift: boolean

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

    const { index } = $page.params; 
    helper = JSON.parse(decodeURIComponent(index));
    buses?.forEach(bus => {
        if (bus.BusId == helper.BusId) selectedBus = bus
    })
    selectedName = helper.Name;
    selectedShift = helper.Shift

    if (dropdownData) {
        setEventHelperDropdownOptions()
    }

    const shiftOptions = [
		{ label: "AM", value: true },
		{ label: "PM", value: false }
	];
</script>

<div class="flex justify-center items-center h-full">
    <div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-24">
        <h1 class="text-2xl font-semibold mb-4">Update Selected Event Helper</h1>
        <form method="POST" action="?/updateEventHelper">
            <input type="hidden" name="old_bus_id" value={helper.BusId}>
            <input type="hidden" name="old_name" value={helper.Name}>
            <input type="hidden" name="old_shift" value={helper.Shift}>
         
            <div class="mb-4">
                <CustomDropdown
                    label="Carplate"
                    name="carplate"
                    options={buses}
                    required
                    searchable
                    bind:selected={selectedBus}
                />
            </div>
            
            <div class="mb-4">
                <CustomDropdown
                    label="Name"
                    name="name"
                    options={names}
                    required
                    searchable
                    bind:selected={selectedName}
                />
            </div>
            <!-- value is assigned but displayOption is not showing -->

            <!-- <CustomDropdown
				label="Shift"
				name="shift"
				options={shiftOptions}
				bind:selected={selectedShift}
			/> -->
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
                                    data-testid="{value}"
								/>
								<label for={label} class="text-sm">{label}</label>
							</div>
						{/each}
					</div>
				</fieldset>
			</div>         

            <div class="mt-4 flex justify-center">
                <button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800">Update Event Helper</button>
            </div>
        </form>
    </div>
</div>
