<script lang="ts">
  import type { Session, User, UserRole } from "$lib/types/global";
  export let data
  let { backend_uri, session, users, roles } = data

  let csvFile: HTMLInputElement | null = null; 
  let isLoading = false;

  const updateUserRole = async(user: User) => {
    let roleInt = roles.find((role) => role.Description == user.Role)?.RoleId
    await fetch(`${backend_uri}:3000/users/update-user`, {
      method: 'PUT',
      body: JSON.stringify({

        Email: user.Email,
        Role: roleInt
      }), 
      headers: {
        'content-type': 'application/json'
      },
      credentials: 'include'
    })
  }

  const uploadCSV = async () => {
    if (!csvFile || !csvFile.files || csvFile.files.length === 0) {
      alert('Please select a file.');
      return;
    }
    
    isLoading = true;

    const formData = new FormData();
    formData.append('file', csvFile.files[0]);

    const response = await fetch(`${backend_uri}:3000/auth/bulk-create-users`, {
      method: 'POST',
      body: formData,
    });

    if (response.ok) {
      console.log('Users imported successfully');
      users = users
    } else {
      const error = await response.text();
      console.log('Error importing users: ', error);
    }

    isLoading = false;
  };

</script>

<div class="p-6 md:p-12">
  <div class="flex items-center space-x-4 pb-2 justify-center">
    <label class="relative cursor-pointer text-black font-semibold py-2 rounded-md text-sm">
      <input type="file" bind:this={csvFile} class=""/>
    </label>
    <button on:click={uploadCSV} class="bg-red-700 text-white font-semibold py-2 px-4 rounded-md text-sm hover:bg-red-800">
      {#if isLoading}
        <svg aria-hidden="true" role="status" class="inline w-4 h-4 me-3 text-white animate-spin" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="#E5E7EB"/>
          <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentColor"/>
        </svg>
      {/if}
      Upload CSV
    </button>  
  </div>  

  <div class="max-w-sm md:max-w-4xl mx-auto bg-white p-2 md:p-8 rounded-lg">
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
  </div>
</div>