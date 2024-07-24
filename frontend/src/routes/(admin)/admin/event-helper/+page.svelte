<script lang="ts">
    import { onMount } from 'svelte'
    import type { EventHelper } from "$lib/types/global";
    import ToolTip from '$lib/components/ToolTip.svelte';

    export let data;
    const { backend_uri, helperData, busData } = data
    let eventHelper: EventHelper[] = [];
    let filteredHelpers: EventHelper[] = [];
    let carplates: string[] = [];
    let selectedCarplate: string = '';
    let selectedShift: string = '';
    let searchQuery: string = '';

    let csvFile: HTMLInputElement | null = null;
    let isLoading = false;

    async function deleteEventHelper(eventHelperToDelete: EventHelper) {
        try {
            const response = await fetch(`${backend_uri}:3000/event-helpers/delete-helper`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(eventHelperToDelete)
            });

            if (!response.ok) {
                throw new Error('Failed to delete event helper');
            }

            eventHelper = eventHelper.filter(item => item !== eventHelperToDelete);
            applyFilters();
        } catch (error) {
            console.error('Error deleting event helper:', error);
        }
    }

    function applyFilters() {
        filteredHelpers = eventHelper.filter(helper => {
            const matchesCarplate = selectedCarplate === '' || helper.Carplate === selectedCarplate;
            const matchesShift = selectedShift === '' || (selectedShift === 'AM' ? helper.Shift : !helper.Shift);
            const matchesSearchQuery = helper.Name.toLowerCase().includes(searchQuery.toLowerCase());
            return matchesCarplate && matchesShift && matchesSearchQuery;
        });
    } 

    const uploadCSV = async () => {
        if (!csvFile || !csvFile.files || csvFile.files.length === 0) {
          alert('Please select a file.');
          return;
        }

        isLoading = true;

        const formData = new FormData();
        formData.append('file', csvFile.files[0]);

        try {
            const response = await fetch(`${backend_uri}:3000/event-helpers/bulk-create-shifts`, {
                method: 'POST',
                body: formData
            });

            if (response.ok) {
                console.log("File uploaded successfully");
            } else {
                const error = await response.text();
                console.error(`Error: ${error}`);
            }
        } catch (error) {
            console.error("Fetch error:", error);
        }

        isLoading = false;
        location.reload();
    };

    onMount(() => {
        if (helperData){
            eventHelper = helperData  
            applyFilters();  
        }

        if (Array.isArray(busData)) {
          carplates = busData.map(bus => bus.Carplate);
        }
    })
</script>

<svelte:head>
	<title>Manage - Event Helpers | SPOH Bus Tracker</title>  
</svelte:head>

<div class="p-6 md:p-12">
    <div class="flex items-center space-x-4 mb-6 justify-center">
      <label class="relative cursor-pointer text-black font-semibold py-2 rounded-md text-sm">
        <input type="file" bind:this={csvFile}/>
      </label>
      <button 
        on:click={uploadCSV} 
        disabled={isLoading} 
        class="bg-red-700 text-white font-semibold py-2 px-4 rounded-md text-sm enabled:hover:bg-red-800 disabled:opacity-75 disabled:cursor-not-allowed"
      >
        {#if isLoading}
          <svg aria-hidden="true" role="status" class="inline w-4 h-4 me-3 text-white animate-spin" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="#E5E7EB"/>
            <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentColor"/>
          </svg>
        {/if}
        Upload Shifts CSV
      </button>  
    </div> 

    <div class="flex items-center justify-between mb-4">
        <a href="/admin/event-helper/add-helper" class="border-black text-white font-semibold text-md px-6 py-2 rounded-xl bg-red-700 hover:bg-red-800 mr-2">
            Add Event Helper
        </a>

        <div class="ml-auto">
          <select bind:value={selectedCarplate} on:change={applyFilters} class="border border-gray-300 text-sm rounded-xl px-3 py-2" data-testid="search-route">
              <option value="">All Carplates</option>
              {#each carplates as carplate}
                  <option value={carplate}>{carplate}</option>
              {/each}
          </select>
          <select bind:value={selectedShift} on:change={applyFilters} class="border border-gray-300 text-sm rounded-xl px-3 py-2" data-testid="search-carplate">
              <option value="">All Shifts</option>
              <option value="AM">AM</option>
              <option value="PM">PM</option>
          </select>
        </div>
      
      <div class="ml-auto">
        <input bind:value={searchQuery} on:input={applyFilters} type="text" placeholder="Search Students..." class="border border-gray-300 rounded-md px-3 py-2 w-60" data-testid="search-input">
      </div>
    </div>

    <div class="mt-8">
        <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
                <tr>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Bus Carplate</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Student Name</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Shift</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
                </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
                  {#if filteredHelpers.length === 0} 
                      <tr>
                          <td colspan="8" class="px-6 py-4 whitespace-nowrap text-center">No matching event helper found</td>
                      </tr>
                  {:else}
                    {#each filteredHelpers as helper}
                        <tr class="hover:bg-gray-100">
                            <td class="px-6 py-4 whitespace-nowrap">{helper.Carplate}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{helper.Name}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{helper.Shift ? 'AM' : 'PM'}</td>
                            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                                <div class="flex items-center">
                                    <a href={`event-helper/update-helper/${encodeURIComponent(JSON.stringify(helper))}`} class="text-stone-500 hover:text-green-500 mr-8">
                                        <ToolTip text="Update Event Helper"> 
                                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                            <path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/>
                                        </svg>
                                        </ToolTip>
                                    </a>
                                    <button class="text-stone-500 hover:text-red-600 text-2xl" on:click={() => deleteEventHelper(helper)}>
                                        <ToolTip text="Delete Event Helper"> 
                                        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" {...$$props}>
                                            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24" {...$$props}>
                                                <g fill="none">
                                                    <path d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022m-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z" />
                                                    <path fill="currentColor" d="M14.28 2a2 2 0 0 1 1.897 1.368L16.72 5H20a1 1 0 1 1 0 2l-.003.071l-.867 12.143A3 3 0 0 1 16.138 22H7.862a3 3 0 0 1-2.992-2.786L4.003 7.07A1.01 1.01 0 0 1 4 7a1 1 0 0 1 0-2h3.28l.543-1.632A2 2 0 0 1 9.721 2zm3.717 5H6.003l.862 12.071a1 1 0 0 0 .997.929h8.276a1 1 0 0 0 .997-.929zM10 10a1 1 0 0 1 .993.883L11 11v5a1 1 0 0 1-1.993.117L9 16v-5a1 1 0 0 1 1-1m4 0a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1m.28-6H9.72l-.333 1h5.226z" />
                                                </g>
                                            </svg>
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
</div>
