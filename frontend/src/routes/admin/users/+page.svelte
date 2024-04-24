<script lang="ts">
  import { onMount } from "svelte"
  import type { User, UserRole } from "../../../types/global";
  export let data
  const { backend_uri, session } = data
  let users: Array<User> = []
  let roles: Array<UserRole> = []
  let notLoaded = true
  
  const getUsers = () => {
    return new Promise(async(resolve, reject) => {
      let response = await fetch(`${window.location.href.split('/')[0]}/api/users/get-users`, {
        method: 'GET',
        credentials: 'include'
      })
      let parsed = await response.json()
      users = parsed
      resolve(users)
    })
  } 

  const getRoles = () => {
    return new Promise(async(resolve, reject) => {
      let response = await fetch(`${backend_uri}:3000/users/get-roles`)
      roles = await response.json()
      resolve(roles) 
    })
  } 

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

  onMount(async() => {
    await Promise.all([getUsers(), getRoles()])
    notLoaded = false
  })

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