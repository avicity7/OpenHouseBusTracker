<script lang="ts">
  import { onMount } from "svelte"
  import type { Session, User, UserRole } from "../../../../lib/types/global";
  export let data
  let { backend_uri, session, users, roles } = data

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
</script>

<div class="p-6 md:p-12">
  <div class="max-w-sm md:max-w-4xl mx-auto bg-white p-2 md:p-8 rounded-lg">
    <table class="w-full">
      <thead>
        <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Email</th>
        <th class="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Role</th>
      </thead>
      <tbody class="divide-y divide-stone-300">
        {#each users as user}
          <tr>
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