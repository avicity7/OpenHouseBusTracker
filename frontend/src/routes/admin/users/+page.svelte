<script lang="ts">
  import { onMount } from "svelte"
  import type { Session, User, UserRole } from "../../../types/global";
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
  <h1 class="text-3xl font-semibold">
    Users
  </h1>
  <div class="mt-12">
    {#each users as user}
      <div class="flex flex-row justify-between max-w-3xl mt-3">
        <h1>{user.Email}</h1>
        <select bind:value={user.Role} on:change={() => {updateUserRole(user)}}>
          {#each roles as role}
            <option value={role.Description} selected={user.Role == role.Description ?? role.Description}>{role.Description}</option>
          {/each}
        </select>
      </div>
    {/each}
  </div>
</div>