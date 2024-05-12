<script lang="ts">
    import { onMount } from "svelte"
    import type { EventHelper } from '$lib/types/global';

    export let data;

    let carplates: string[] = [];
    let emails: string[] = [];

    let selectedCarplate: string | null = null;
    let selectedEmail: string | null = null;

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
        console.log("here", carplates, emails)
    }

    onMount(() => {
        setEventHelperDropdownOptions(data.data)
        })

</script>

<div class="flex justify-center items-center h-full mt-20">
    <div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3">
        <h1 class="text-2xl font-semibold mb-4">Add New Event Helper</h1>
        <form method="POST" action="?/createEventHelper">
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

        <!-- step in the input field works but it doesnt change the timepicker options, it just checks if the selected time is divisible by 30 mins and show an error if it is -->
        <div class="mb-4">
            <label for="startTime" class="block text-sm font-medium mb-1">Start Time:</label>
            <input type="datetime-local" id="start_time" name="start_time" required class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500"/>
        </div>

        <div class="mb-4">
            <label for="endTime" class="block text-sm font-medium mb-1">End Time:</label>
            <input type="datetime-local" id="end_time" name="end_time" required class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500"/>
        </div>

        <div class="mt-4 flex justify-center">
            <button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800">Add Event Helper</button>
        </div>
        </form>
    </div>
</div>
