<script lang="ts">
    import { PUBLIC_BACKEND_URL } from "$env/static/public";
	import { onMount } from 'svelte';
    import { page } from "$app/stores";
	import type { EventBus } from "$lib/types/global";
	import { goto } from "$app/navigation";
    
    let carplate: string
    let bus: EventBus = { Carplate: '', Status: false, Hidden: false};


    async function updateBus(event: Event) {

    const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/bus/update-bus/${bus.Carplate}  `, {
    method: 'PUT',
    headers: {
        'Content-Type': 'application/json'
    },
    body: JSON.stringify({ Carplate: carplate })
    });

    if (response.ok) {
    console.log(`Driver has been updated to ${carplate}!`)
    goto('/admin/drivers');
    } else {
    console.log(`Failed to update ${carplate}`)
    }
    }

    onMount(() => {
        const { id } = $page.params
        bus = JSON.parse(decodeURIComponent(id))

        carplate = bus.Carplate
    })
  
  </script>
  <div class="flex justify-center items-center h-full">
      <div class="bg-white shadow rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-20">
          <h1 class="text-2xl font-semibold mb-4">Update Selected Bus</h1>
          <form on:submit={updateBus}>
              <div class="mb-4">
                  <label for="startTime" class="block text-sm font-medium mb-1">Carplate</label>
                  <input
                      type="text"
                      required
                      bind:value={carplate}
                      class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-red-500"
                  />
              </div>
  
              <div class="mt-4 flex justify-center">
                  <button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800"
                      >Update Bus</button
                  >
              </div>
          </form>
      </div>
  </div>
  