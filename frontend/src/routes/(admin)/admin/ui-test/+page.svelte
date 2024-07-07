<script>
    let selectedDashboard = null;
  
    function openDashboard(dashboard) {
      selectedDashboard = dashboard;
    }
  
    function closeDashboard() {
      selectedDashboard = null;
    }
  </script>
  
  {#if selectedDashboard}
    <div class="p-4 min-h-screen">
      <button class="mb-4 p-2 bg-blue-500 text-white rounded" on:click={closeDashboard}>Back to Grid</button>
      {#if selectedDashboard === 'Dashboard1'}
        <p class="bg-red-500">DB1</p>
      {:else if selectedDashboard === 'Dashboard2'}
        <p class="bg-red-500">DB2</p>
      {:else if selectedDashboard === 'Dashboard3'}
        <p class="bg-red-500">DB3</p>
      {:else if selectedDashboard === 'Dashboard4'}
        <p class="bg-red-500">DB4</p>
      {/if}
    </div>
  {:else}
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-2 gap-4 p-4 min-h-screen">
      <button type="button" class="bg-gray-200 p-16 rounded-lg shadow-md hover:bg-gray-300 cursor-pointer" on:click={() => openDashboard('Dashboard1')}>
        <h2 class="text-2xl font-semibold">Dashboard 1</h2>
      </button>
      <button type="button" class="bg-gray-200 p-16 rounded-lg shadow-md hover:bg-gray-300 cursor-pointer" on:click={() => openDashboard('Dashboard2')}>
        <h2 class="text-2xl font-semibold">Dashboard 2</h2>
      </button>
      <button type="button" class="bg-gray-200 p-16 rounded-lg shadow-md hover:bg-gray-300 cursor-pointer" on:click={() => openDashboard('Dashboard3')}>
        <h2 class="text-2xl font-semibold">Dashboard 3</h2>
      </button>
      <button type="button" class="bg-gray-200 p-16 rounded-lg shadow-md hover:bg-gray-300 cursor-pointer" on:click={() => openDashboard('Dashboard4')}>
        <h2 class="text-2xl font-semibold">Dashboard 4</h2>
      </button>
    </div>
  {/if}
  