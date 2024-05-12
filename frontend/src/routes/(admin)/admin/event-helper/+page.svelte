<script lang="ts">
    import { onMount } from 'svelte'
    import type { EventHelper } from "../../../../types/global";
    import ToolTip from '../../../../components/ToolTip.svelte';

    export let data;
    const { backend_uri} = data
    let eventHelper: EventHelper[] = [];

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

            console.log("Deleted Event Helper:", eventHelperToDelete);
        } catch (error) {
            console.error('Error deleting event helper:', error);
        }
    }

    function formatTimestamp(timestamp: string): string {
        const utcDate = new Date(timestamp);
        const formattedDate = utcDate.toLocaleString();
        return formattedDate;
    }
    onMount(() => {
        if (data && data.data){
            eventHelper = data.data
        }
    })
</script>

<div class="p-6 md:p-12">
    <div class="flex items-center justify-between mb-4">
        <a href="event-helper/add-helper" class="border-black text-white font-semibold text-md px-6 py-2 rounded-xl bg-red-700 hover:bg-red-800 mr-2">
            Add Event Helper
        </a>    
    </div>

    <div class="mt-8">
        <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
                <tr>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Bus Carplate</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Student Email</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Start Time</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">End Time</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
                </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
                    {#each eventHelper as helper}
                        <tr class="hover:bg-gray-100">
                            <td class="px-6 py-4 whitespace-nowrap">{helper.Carplate}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{helper.Email}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{formatTimestamp(helper.StartTime)}</td>
                            <td class="px-6 py-4 whitespace-nowrap">{formatTimestamp(helper.EndTime)}</td>
                            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                                <div class="flex items-center">
                                    <a href={`event-helper/update-helper/${encodeURIComponent(JSON.stringify(helper))}`} class="text-slate-500 hover:text-green-500 mr-8">
                                        <ToolTip text="Update Schedule"> 
                                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                                            <path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/>
                                        </svg>
                                        </ToolTip>
                                    </a>
                                    <button class="text-slate-500 hover:text-red-600 text-2xl" on:click={() => deleteEventHelper(helper)}>
                                        <ToolTip text="Delete Schedule"> 
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
            </tbody>
        </table>
    </div>
</div>
