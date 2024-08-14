<script lang="ts">
  import UserRow from '$lib/components/UserRow.svelte';
  import FileUpload from '$lib/components/FileUpload.svelte';
  import ErrorMessage from '$lib/components/ErrorMessage.svelte';

	export let data;
	let { backend_uri, session, users, roles } = data;

  let selectedFile: File | null = null;
  let showSuccessMessage = false;
	let isLoading = false;

  let pendingSearch = '';
  let verifiedSearch = '';
  let errorMessage: string | null = null;

  $: filteredPendingUsers = users.filter(user =>
    user.Role === 'public' && user.Name.toLowerCase().includes(pendingSearch.toLowerCase())
  );

  $: filteredVerifiedUsers = users.filter(user =>
    user.Role !== 'public' && user.Name.toLowerCase().includes(verifiedSearch.toLowerCase())
  );

  const handleError = (event: { detail: { message: string | null; }; }) => {
    errorMessage = event.detail.message;
  };

  const clearError = () => {
    errorMessage = null;
  };

  const uploadCSV = async () => {
    if (!selectedFile) {
      alert('Please select a file.');
      return;
    }

    isLoading = true; 

    const formData = new FormData();
    formData.append('file', selectedFile);

    const response = await fetch(`${backend_uri}:3000/auth/bulk-create-users`, {
      method: 'POST',
      body: formData
    });

    if (response.ok) {
      showSuccessMessage = true;
      console.log("File uploaded successfully");
      location.reload();
    } else {
      const error = await response.text();
      showSuccessMessage = false;
      console.error(`Error: ${error}`);
    }

    isLoading = false;
  };

  function handleFileSelected(event: CustomEvent<File>) {
    selectedFile = event.detail;
  }
</script>

<svelte:head>
	<title>Manage - Users | SPOH Bus Tracker</title>
</svelte:head>

<div class="p-6 md:p-12">
  <div class="flex flex-col items-center space-y-4 pb-2">
    <FileUpload on:fileSelected={handleFileSelected} />
    <button 
      on:click={uploadCSV}
      disabled={isLoading} 
      class="bg-red-700 text-white font-semibold py-2 px-4 rounded-md text-sm enabled:hover:bg-red-800 disabled:opacity-75 disabled:cursor-not-allowed"  
      data-testid="submit-bulk-import"
    >
      {#if isLoading}
        <svg aria-hidden="true" role="status" class="inline w-4 h-4 me-3 text-white animate-spin" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="#E5E7EB"/>
          <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentColor"/>
        </svg>
      {/if}
      Upload CSV
    </button>  
    {#if showSuccessMessage}
      <p class="text-green-600 text-sm font-semibold mt-2">File uploaded successfully!</p>
    {/if}
  </div>  
  

	<div class="mt-1 flex justify-between items-center">
    <h1 class="text-xl font-semibold p-2 py-8 text-stone-800">Pending Verification</h1>
    <input
      type="text"
      placeholder="Search Pending Users..."
      bind:value={pendingSearch}
      class="border border-gray-300 p-2 rounded-md"
    />
  </div>
  
	<div class="mt-1">
    <table class="min-w-full divide-y divide-gray-200">
			<thead class="bg-gray-50">
				<tr>
					<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
					<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Email</th>
					<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Contact</th>
					<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Role</th>
					<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
				</tr>
			</thead>
			<tbody class="bg-white divide-y divide-gray-200">
        {#if filteredPendingUsers.length === 0}
          <tr>
            <td colspan="4" class="px-6 py-4 whitespace-nowrap text-center text-gray-400">No matching Pending Users found</td>
          </tr>
        {:else}
          {#each filteredPendingUsers as user}
            <UserRow bind:users {user} {roles} />
          {/each}
        {/if}
      </tbody>
		</table>
	</div>

	<div class="mt-8 flex justify-between items-center">
    <h1 class="text-xl font-semibold p-2 py-8 text-stone-800">Admins & Student Helpers</h1>
    
    <input
      type="text"
      placeholder="Search Admins & Student Helpers..."
      bind:value={verifiedSearch}
      class="border border-gray-300 p-2 rounded-md"
    />
  </div>
  
	<div class="mt-8">
    {#if errorMessage}
      <ErrorMessage message={errorMessage} />
    {/if}
    <table class="min-w-full divide-y divide-gray-200">
			<thead class="bg-gray-50">
				<tr>
					<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
					<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Email</th>
					<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Contact</th>
					<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Role</th>
					<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
				</tr>
			</thead>
			<tbody class="bg-white divide-y divide-gray-200">
        {#if filteredVerifiedUsers.length === 0}
          <tr>
            <td colspan="4" class="px-6 py-4 whitespace-nowrap text-center text-gray-400">No matching Admin or Student Helpers found</td>
          </tr>
        {:else}
          {#each filteredVerifiedUsers as user}
            <UserRow 
              bind:users 
              {user} 
              {roles} 
              on:error={handleError} 
              on:clearError={clearError} />
          {/each}
        {/if}
      </tbody>
		</table>
	</div>
</div>

<!-- <div class="max-w-sm md:max-w-4xl mx-auto bg-white p-2 md:p-8 rounded-lg">
    <table class="w-full">
      <thead>
        <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
        <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Email</th>
        <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Role</th>
      </thead>
      <tbody class="divide-y divide-stone-300">
        {#each users as user}
          <tr>
            <td class="px-4 py-6">{user.Name}</td>
            <td class="px-4 py-6">{user.Email}</td>
            <td class="py-6">
              <select
                bind:value={user.Role}
                on:change={() => {updateUserRole(user)}}
              >
                {#each roles as role}
                  <option 
                    value={role.Description}
                    selected={user.Role == role.Description ?? role.Description}
                  >
                    {role.Description}
                  </option>
                {/each}
              </select>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div> -->
