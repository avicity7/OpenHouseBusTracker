<script lang="ts">
  import { onMount } from 'svelte';

  export let data;
  const { UserSettings } = data;

  let name = '';
  let contact = '';
  let email = '';
  
  onMount(() => {
    if(UserSettings){
      name = UserSettings.Name || '';
      contact = UserSettings.Contact || '';
      email = UserSettings.Email || '';
    }

    console.log("this is user settitngs data", UserSettings)
  })
</script>

<div class="flex flex-col justify-center items-center mt-28">
  <div class="font-albert w-full max-w-sm text-zinc-800 dark:text-zinc-100 px-6">
    <h1 class="text-3xl font-bold mb-5 text-center">Settings Page</h1>

    <form method="POST" action="?/updateUserSettings" class="w-full">
      <div class="flex flex-col space-y-4">
        <div>
          <label for="name" class="block text-sm font-medium mb-1">Name</label>
          <input
            type="text"
            id="name"
            name="name"
            required
            bind:value={name}
            class="w-full rounded-md border border-gray-300 px-3 py-2 shadow-sm focus:outline-none focus:ring-1 focus:ring-red-500"
          />
        </div>

        <div>
          <label for="contact" class="block text-sm font-medium mb-1">Contact Number</label>
          <input
            type="tel"
            id="contact"
            name="contact"
            required
            bind:value={contact}
            class="w-full rounded-md border mb-4 border-gray-300 px-3 py-2 shadow-sm focus:outline-none focus:ring-1 focus:ring-red-500"
          />
        </div>

        <input type="hidden" name="email" value={email} />

        <button type="submit" class="w-full mt-6 px-3 py-1.5 bg-red-800 hover:bg-red-900 hover:text-zinc-200 rounded-md text-zinc-100 text-sm font-medium">
          Update Settings
        </button>
      </div>
    </form>
  </div>
</div>
