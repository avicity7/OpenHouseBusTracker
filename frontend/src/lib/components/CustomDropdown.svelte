<script lang="ts">
    import { onMount } from 'svelte';
    export let options: any[] = [];
    export let selected: any = null;
    export let label: string = "";
    export let required: boolean = false;
    export let name: string = "";
    export let searchable: boolean = false; 

    let isOpen = false;
    let error = "";
    let selectedDisplay: any;
    let searchQuery: string = '';
    let filteredOptions: any[] = [];

    let dropdownId = `dropdown-${Math.random().toString(36).substring(2,9)}`; // either this or increment by 1

    function toggleDropdown() {
        isOpen = !isOpen;
    }

    function selectOption(option: any) {
        selected = typeof(option) == "object" ? JSON.stringify(option) : option;
        selectedDisplay = option;
        isOpen = false;
        validateSelection();
    }

    function validateSelection() {
        if (required && !selected) {
            error = "Selection is required.";
        } else {
            error = "";
        }
    }

    function filterOptions() {
        filteredOptions = options.filter(option => {
            const displayText = typeof(option) === 'object' ?
                (option.DriverName || option.RouteName || option.Carplate) :
                option;

            return displayText.toLowerCase().includes(searchQuery.toLowerCase());
        });
    }

    onMount(() => {
        let mutated = typeof(selected) == "object";
        selected = typeof(selected) == "object" ? JSON.stringify(selected) : selected;
        selectedDisplay = mutated ? JSON.parse(selected) : selected;

        const handleClickOutside = (event: MouseEvent) => {
            const dropdown = document.getElementById(dropdownId);
            if (isOpen && dropdown && !dropdown.contains(event.target as Node)) {
                isOpen = false;
            }
        };

        document.addEventListener('click', handleClickOutside);

        return () => {
            document.removeEventListener('click', handleClickOutside);
        };
    });

    $: searchQuery, filterOptions();
    $: options, filterOptions();
</script>

<div class="mb-4">
    <label for={dropdownId} class="block text-sm font-medium mb-1">{label}</label>
    <div
        id={dropdownId}
        role="combobox"
        aria-haspopup="listbox"
        aria-expanded={isOpen ? 'true' : 'false'}
        aria-controls={isOpen ? `${dropdownId}-list` : undefined}
        aria-invalid={!!error ? 'true' : 'false'}
        aria-describedby={!!error ? `${dropdownId}-error` : undefined}
        class="relative"
        data-testid="{name}"
    >
        <button
            type="button"
            class="w-full appearance-none rounded-md border border-gray-300 bg-white text-gray-700 px-3 py-2 shadow-sm focus:outline-none focus:ring-1 focus:ring-red-500 cursor-pointer"
            on:click={toggleDropdown}
            aria-haspopup="listbox"
            aria-expanded={isOpen}
            data-testid="dropdown-input"
        >
            {typeof(selectedDisplay) == "boolean" ? selectedDisplay ? "AM" : "PM" : selectedDisplay ? typeof(selectedDisplay) == "object" ? "DriverName" in selectedDisplay ? selectedDisplay.DriverName : "RouteName" in selectedDisplay ? selectedDisplay.RouteName : selectedDisplay.Carplate : selectedDisplay : "Select an option"}
            <svg class="w-5 h-5 text-gray-400 absolute right-3 top-1/2 transform -translate-y-1/2 pointer-events-none" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
        </button>
        {#if isOpen}
            <ul id={`${dropdownId}-list`} role="listbox" class="absolute z-10 w-full bg-white border border-gray-300 rounded-md shadow-lg mt-1 max-h-96 overflow-y-auto">
                {#if searchable}
                    <div class="px-3 py-2">
                        <input
                            type="text"
                            placeholder="Search..."
                            bind:value={searchQuery}
                            class="block w-full px-3 py-2 border rounded-md text-sm focus:border-red-500 focus:outline-none"
                        />
                    </div>
                {/if}
                {#if filteredOptions.length > 0}
                    {#each filteredOptions as option (option)}
                        <li>
                            <button
                                type="button"
                                role="option"
                                id={`option-${option.DriverId}`}
                                aria-selected={option === selected}
                                class="block w-full px-4 py-2 text-gray-700 hover:bg-red-100 focus:bg-red-100 focus:outline-none cursor-pointer"
                                on:click={() => selectOption(option)}
                            >
                                {typeof(option) == "boolean" ? option ? "AM" : "PM" : option && typeof(option) == "object" ? "DriverName" in option ? option.DriverName : "RouteName" in option ? option.RouteName : option.Carplate : option}
                            </button>
                        </li>
                    {/each}
                {:else}
                    <li>
                        <button
                            type="button"
                            role="option"
                            aria-selected="false"
                            aria-disabled="true"
                            class="block w-full px-4 py-2 text-gray-700 cursor-not-allowed"
                        >
                            No options available
                        </button>
                    </li>
                {/if}
            </ul>
        {/if}
    </div>
    {#if error}
        <p id={`${dropdownId}-error`} class="text-red-500 text-sm mt-1">{error}</p>
    {/if}
</div>

<input type="hidden" id="selectedValue" name={name} value={selected ? selected.DriverId || selected.value || selected : ''}/>
