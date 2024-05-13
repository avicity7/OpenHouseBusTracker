<script lang="ts">
  import { onMount, createEventDispatcher } from 'svelte';
  import { writable } from 'svelte/store';
  import type { Driver } from '$lib/types/global';
  import { PUBLIC_BACKEND_URL } from '$env/static/public';

  let drivers = writable<Driver[]>([]);
  let newName = '';
  let alert = '';
  let search = '';
  let driverToDelete: Driver | null = null;

  // Add Driver
  async function addDriver() {
    let newDriverName: string | null = '';
    let nameExistsMessage = '';

    do {
      newDriverName = prompt(`${nameExistsMessage}Enter the name of the new driver:`);
      if (!newDriverName) return; 
      const trimmedName = newDriverName.trim();
      if (trimmedName === '') return; 

      // Check if the name already exists
      const existingDriver = $drivers.find(driver => driver.DriverName.toLowerCase() === trimmedName.toLowerCase());
      if (existingDriver) {
        nameExistsMessage = `${trimmedName} already exists in the database. `;
      } else {
        // Proceed to add the driver if the name doesn't exist
        const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/driver/add-driver`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ DriverName: trimmedName })
        });

        if (response.ok) {
          getDrivers();
          alert = `${trimmedName} has been added!`;
          setTimeout(() => { alert = ''; }, 3000);
        }
        break;
      }
    } while (true);
  }


  // Get Driver
  const getDrivers = async () => {
    const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/driver/get-driver`);
    if (response.ok) {
      let data = await response.json() as Driver[];
      data.sort((a, b) => a.DriverId - b.DriverId);
      drivers.set(data);
    }
  };
  onMount(getDrivers);

  // Update Driver
  async function updateDriver(id: number, name: string) {
    let updatedName: string = ''; 
    let nameExistsMessage = ''; 

    do {
      updatedName = prompt(`${nameExistsMessage}Enter the new name:`, name) || ''; 

      if (!updatedName) return; 
      
      // Check if the updated name already exists for another driver
      const existingDriver = $drivers.find(driver => driver.DriverName.toLowerCase() === updatedName.toLowerCase());

      if (existingDriver && updatedName.toLowerCase() !== name.toLowerCase()) {
        nameExistsMessage = `${updatedName} already exists in the database. `;
      } else if (updatedName.toLowerCase() === name.toLowerCase()) {
        nameExistsMessage = 'Name has no changes. ';
      } else {
        nameExistsMessage = ''; 
        break; 
      }
    } while (true); 

    if (updatedName !== name) {
      const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/driver/update-driver/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ DriverName: updatedName })
      });

      if (response.ok) {
        getDrivers(); 
        alert = `${name} has been updated to ${updatedName}!`;
        setTimeout(() => { alert = ''; }, 3000);
      }
    }
  }

  // Delete Driver
  async function deleteDriver(id: number, name: string) {
    driverToDelete = { DriverId: id, DriverName: name };
  }

  // Confirmation modal actions
  const dispatch = createEventDispatcher();
  function confirmDelete() {
    if (driverToDelete) {
      const { DriverId, DriverName } = driverToDelete;
      fetch(`${PUBLIC_BACKEND_URL}:3000/driver/delete-driver/${DriverId}`, { method: 'DELETE' })
        .then(response => {
          if (response.ok) {
            getDrivers(); 
            alert = `${DriverName} has been deleted!`;
            setTimeout(() => { alert = ''; }, 3000);
          }
        });
      driverToDelete = null;
    }
  }

  function cancelDelete() {
    driverToDelete = null;
  }

</script>

<div class="mx-auto px-8 relative">
  <div class="flex items-center">
    <div class="ml-auto">
      <input type="text" class="ml-4 p-6 m-3 rounded-full focus:outline-none w-500" bind:value={search} placeholder="Type to search drivers....">
    </div>
  </div>
  {#if alert}
  <div class="fixed top-0 left-0 w-full h-full flex items-center justify-center">
    <div class="bg-gray-100 p-4 m-5 rounded-lg shadow-2xl border border-gray-300 w-800 h-1000">
      <p class="text-lg mb-2">{alert}</p>
    </div>
  </div>
{/if}
<div class="bg-white rounded m-4">
  {#each $drivers.filter(driver => driver.DriverName.toLowerCase().includes(search.toLowerCase())) as driver}
    <div class="flex items-center justify-between hover:bg-gray-100 transition-colors p-8">
      <div>{driver.DriverName}</div>
      <div>
        <button class="text-blue-500 px-3 py-1 rounded" on:click={() => updateDriver(driver.DriverId, driver.DriverName)}>Edit</button>
        <button class="text-red-500 px-3 py-1 rounded" on:click={() => deleteDriver(driver.DriverId, driver.DriverName)}>Delete</button>
      </div>
    </div>
  {/each}
</div>

  <!-- Delete Confirmation Modal -->
  {#if driverToDelete}
    <div class="fixed top-0 left-0 w-full h-full flex items-center justify-center bg-gray-900 bg-opacity-50">
      <div class="bg-white p-8 rounded-lg shadow-lg">
        <p class="text-lg mb-4">Are you sure you want to delete this driver?</p>
        <div class="flex justify-end">
          <button class="px-4 py-2 mr-4 text-gray-600" on:click={cancelDelete}>Cancel</button>
          <button class="px-4 py-2 bg-red-500 text-white rounded" on:click={confirmDelete}>Delete</button>
        </div>
      </div>
    </div>
  {/if}
  <button class="absolute top-0 mt-4 ml-4 bg-red-700 hover:bg-red-800 text-white text-3xl font-semibold w-16 h-16 rounded-full flex items-center justify-center shadow-2xl" on:click={addDriver}>+</button>
</div>
