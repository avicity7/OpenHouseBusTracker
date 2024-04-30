<div class="p-6 md:p-12">
  <h1 class="text-3xl font-semibold">
    Drivers (Admin)
  </h1>
</div>

<script lang="ts">
  import { onMount } from 'svelte';
  import { writable } from 'svelte/store';
  import type { Driver } from '../../../types/global';
  import { PUBLIC_BACKEND_URL } from '$env/static/public';

  let drivers = writable<Driver[]>([]);
  let newName = '';
  let alert = '';

  // Add Driver
  async function addDriver() {
    if (newName.trim() === '') return;
    const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/driver/add-driver`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ DriverName: newName })
    });
    if (response.ok) {
      getDrivers();
      alert = `${newName} has been added!`;
      newName = '';
      setTimeout(() => { alert = ''; }, 3000);
    }
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
    const updatedName = prompt('Enter the new name:', name);
    if (updatedName === null) return;
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

// Delete Driver
async function deleteDriver(id: number, name: string) {
    const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/driver/delete-driver/${id}`, { method: 'DELETE' });
    if (response.ok) {
        getDrivers(); // Fetch updated list of drivers
        alert = `${name} has been deleted!`;
        setTimeout(() => { alert = ''; }, 3000);
    }
}

</script>

<div class="container px-8 mx-3">
  <div class="flex mb-4">
    <input type="text" class="border border-gray-300 rounded p-2 m-2" bind:value={newName} placeholder="Enter driver name">
    <button class="bg-blue-500 text-white p-3 m-2 rounded" on:click={addDriver}>Add Driver</button>
  </div>
  {#if alert}
    <div class="mb-4 p-2 bg-green-200 text-green-800">{alert}</div>
  {/if}
  <ul>
    {#each $drivers as driver}
      <li class="border border-gray-300 rounded p-2 mb-2 flex justify-between items-center">
        <span>{driver.DriverName}</span>
        <div>
          <button class="bg-yellow-500 text-white px-3 py-1 rounded mr-2" on:click={() => updateDriver(driver.DriverId, driver.DriverName)}>Update</button>
          <button class="bg-red-500 text-white px-3 py-1 rounded" on:click={() => deleteDriver(driver.DriverId, driver.DriverName)}>Delete</button>
        </div>
      </li>
    {/each}
  </ul>
</div>
