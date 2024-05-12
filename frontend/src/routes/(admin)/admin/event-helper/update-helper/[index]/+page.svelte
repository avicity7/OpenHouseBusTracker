<script lang="ts">
    import { onMount } from "svelte";
    import type { EventHelper } from "../../../../../../types/global";
	import { page } from "$app/stores";

    export let data;

    let helper = {
        Carplate: '',
        Email: '',
        StartTime: '',
        EndTime: ''
    };

    let carplates: string[] = [];
    let emails: string[] = [];
    let selectedCarplate: string | null = null;
    let selectedEmail: string | null = null;
    let selectedStartTime = "";
    let selectedEndTime = "";

    function setEventHelperDropdownOptions(data: EventHelper[] | undefined) {
        if (!data) return;

        const uniqueCarplates = new Set<string>();
        const uniqueEmails = new Set<string>();

        data.forEach(({ Carplate, Email }: EventHelper) => {

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
    // updated information is not shown when update is successful because of url not changing
    onMount(() => {
        const { index } = $page.params; 
        helper = JSON.parse(decodeURIComponent(index));
        selectedCarplate = helper.Carplate;
        selectedEmail = helper.Email;
        selectedStartTime = helper.StartTime.split('T')[0] + 'T' + helper.StartTime.split('T')[1].slice(0, 5);
        selectedEndTime = helper.EndTime.split('T')[0] + 'T' + helper.EndTime.split('T')[1].slice(0, 5);

        if (data && data.data) {
            setEventHelperDropdownOptions(data.data)
        }
        
    });
    
</script>

<div class="flex justify-center items-center h-full mt-24">
    <div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3">
        <h1 class="text-2xl font-semibold mb-4">Update Selected Event Helper</h1>
        <form method="POST" action="?/updateEventHelper">
            <input type="hidden" name="old_carplate" value={helper.Carplate}>
            <input type="hidden" name="old_email" value={helper.Email}>
            <input type="hidden" name="old_start_time" value={helper.StartTime}>
            <input type="hidden" name="old_end_time" value={helper.EndTime}>

            <div class="mb-4">
                <label for="carplate" class="block text-sm font-medium mb-1">Carplate:</label>
                <select id="carplate" name="carplate" bind:value={selectedCarplate} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500">
                    {#each carplates as carplate}
                        <option value={carplate}>{carplate}</option>
                    {/each}
                </select>
            </div>
            
            <div class="mb-4">
                <label for="email" class="block text-sm font-medium mb-1">Email:</label>
                <select id="email" name="email" bind:value={selectedEmail} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500">
                    {#each emails as email}
                        <option value={email}>{email}</option>
                    {/each}
                </select>
            </div>

            <div class="mb-4">
                <label for="startTime" class="block text-sm font-medium mb-1">Start Time:</label>
                <input type="datetime-local" id="start_time" name="start_time" required bind:value={selectedStartTime} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500" />
            </div>

            <div class="mb-4">
                <label for="endTime" class="block text-sm font-medium mb-1">End Time:</label>
                <input type="datetime-local" id="end_time" name="end_time" required bind:value={selectedEndTime} class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500" />
            </div>

            <div class="mt-4 flex justify-center">
                <button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800">Update Event Helper</button>
            </div>
        </form>
    </div>
</div>

