<script lang="ts">
  import type { Session, User, UserRole } from "$lib/types/global";
  export let data
  let { backend_uri, session, users, roles } = data

  let csvFile: HTMLInputElement | null = null; 

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
    
    const formData = new FormData();
    formData.append('file', csvFile.files[0]);

    const response = await fetch(`${backend_uri}:3000/auth/bulk-create-users`, {
      method: 'POST',
      body: formData,
    });

    if (response.ok) {
      console.log('Users imported successfully');
    } else {
      const error = await response.text();
      console.log('Error importing users: ', error);
    }
  };

</script>

<div class="p-6 md:p-12">
  <input type="file" bind:this={csvFile} />
  <button on:click={uploadCSV}>Upload CSV</button>
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