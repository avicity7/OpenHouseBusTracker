<script lang="ts">
    import { onMount } from 'svelte';
  
    export let options: string[] = [];
    export let selected: string | null = null;
    export let label: string = "";
    export let required: boolean = false;
  
    let isOpen = false;
    let error = "";
  
    let dropdownId = `dropdown-${Math.random().toString(36).substring(2,9)}`; // either this or just increment numbers
  
    function toggleDropdown() {
      isOpen = !isOpen;
    }
  
    function selectOption(option: string) {
      selected = option;
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
  
    onMount(() => {
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
    >
      <button
        type="button"
        class="w-full appearance-none rounded-md border border-gray-300 bg-white text-gray-700 px-3 py-2 shadow-sm focus:outline-none focus:ring-1 focus:ring-red-500 cursor-pointer"
        on:click={toggleDropdown}
        aria-haspopup="listbox"
        aria-expanded={isOpen}
      >
        {selected ? selected : 'Select an option'}
        <svg class="w-5 h-5 text-gray-400 absolute right-3 top-1/2 transform -translate-y-1/2 pointer-events-none" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
        </svg>
      </button>
      {#if isOpen && options.length > 0}
        <ul id={`${dropdownId}-list`} role="listbox" class="absolute z-10 w-full bg-white border border-gray-300 rounded-md shadow-lg mt-1">
          {#each options as option (option)}
            <li>
              <button
                type="button"
                role="option"
                id={`option-${option}`}
                aria-selected={option === selected}
                class="block w-full px-4 py-2 text-gray-700 hover:bg-red-100 focus:bg-red-100 focus:outline-none cursor-pointer"
                on:click={() => selectOption(option)}
              >
                {option}
              </button>
            </li>
          {/each}
        </ul>
      {:else if options.length === 0}
        <p class="absolute z-10 w-full bg-white border border-gray-300 rounded-md shadow-lg mt-1 px-4 py-2 text-gray-700">
          No options available
        </p>
      {/if}
    </div>
    {#if error}
      <p id={`${dropdownId}-error`} class="text-red-500 text-sm mt-1">{error}</p>
    {/if}
  </div>
  