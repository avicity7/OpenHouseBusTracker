<script lang="ts">

    let bus_schedule: {bus_id: number, route_id: number, assigned_driver: number, start_time: string, end_time: string}[] = [
        {bus_id: 1, route_id: 1, assigned_driver: 1, start_time: "2024-04-22T08:00:00", end_time: "2024-04-22T10:00:00"},
        {bus_id: 2, route_id: 2, assigned_driver: 2, start_time: "2024-04-22T09:00:00", end_time: "2024-04-22T11:00:00"}
    ];

    let showModal = false;

    function deleteSchedule(index: number) {
        bus_schedule = bus_schedule.filter((_schedule, i) => i !== index);
    }

    function openModal() {
    showModal = true;
    }

    function closeModal() {
        console.log("close modal clicked")
        showModal = false;
    }

    function addSchedule() {
        openModal();
        console.log('button to open modal clicked');
    }

    function handleSubmit(event: Event) {
        event.preventDefault();
        closeModal();
        console.log('Form submitted');
    }

    function formatTimestamp(timestamp: string): string {
        const date = new Date(timestamp);
        const year = date.getFullYear();
        const month = (date.getMonth() + 1).toString().padStart(2, '0');
        const day = date.getDate().toString().padStart(2, '0');
        const hours = date.getHours().toString().padStart(2, '0');
        const minutes = date.getMinutes().toString().padStart(2, '0');
        const seconds = date.getSeconds().toString().padStart(2, '0');

        return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    }
</script>

<!-- Modal component -->
<div class="fixed inset-0 flex items-center justify-center" style="background-color: rgba(0, 0, 0, 0.5); display: {showModal ? '' : 'none'};">
    <div class="bg-white p-6 rounded-lg shadow-md">
        <h2 class="text-xl font-semibold mb-4">Add New Bus Schedule</h2>
        <form on:submit={handleSubmit}> 
            <label for="bus_id">Bus ID:</label>
            <input type="text" id="bus_id" name="bus_id" />
            
            <!-- Check on what should be inside the form  MAKE NEW PAGE-->

            <div class="mt-4 flex justify-end">
                <button type="button" class="px-4 py-2 bg-gray-200 text-gray-800 rounded-md mr-2" on:click={closeModal}>Cancel</button>
                <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-md">Add Schedule</button>
            </div>
        </form>
    </div>
</div>

<div class="p-6 md:p-12">
    <div class="flex items-center mb-4">
        <h1 class="text-3xl font-semibold mr-4">Bus Schedules</h1>
        <button class="border-2 border-black text-black text-xl px-4 py-2 rounded-full hover:bg-gray-200" on:click={addSchedule}>+</button>
    </div>


    <div class="mt-8">
        <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
                <tr>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Bus ID</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Route ID</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Assigned Driver</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Start Time</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">End Time</th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
                </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
                {#each bus_schedule as schedule, index (schedule.bus_id)}
                <tr>
                    <td class="px-6 py-4 whitespace-nowrap">{schedule.bus_id}</td>
                    <td class="px-6 py-4 whitespace-nowrap">{schedule.route_id}</td>
                    <td class="px-6 py-4 whitespace-nowrap">{schedule.assigned_driver}</td>
                    <td class="px-6 py-4 whitespace-nowrap">{formatTimestamp(schedule.start_time)}</td>
                    <td class="px-6 py-4 whitespace-nowrap">{formatTimestamp(schedule.end_time)}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                        <button class="text-red-600 hover:text-red-900" on:click={() => deleteSchedule(index)}>Delete</button>
                    </td>
                </tr>
                {/each}
            </tbody>
        </table>
    </div>

    <div class="mt-8">
        <h2 class="text-xl font-semibold">Bus Schedule JOINS Bus, Route and Driver tables</h2>
        <p>dropdown for routes, buses is carplate, driver name </p>
    </div>
</div>


<!-- 
    <div class="mt-8">
        <h2 class="text-xl font-semibold">Add New Bus Schedule</h2>
    </div> -->

<!-- <div class="p-6 md:p-12">
    <h1 class="text-3xl font-semibold">Bus Schedules</h1>
    <h3>Admin Dashboard to display all schedules, addition of new bus schedules, deletion of bus schedules</h3>
    <p>display of bus schedules for user side outside of admin route</p>
    
</div> -->