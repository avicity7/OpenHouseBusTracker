<script lang="ts">
  import { PUBLIC_BACKEND_URL } from '$env/static/public';
  import SuccessMessage from '$lib/components/SuccessMessage.svelte';
  import ErrorMessage from '$lib/components/ErrorMessage.svelte';
  import { page } from '$app/stores'
  const token = $page.params.token

  let password: string,
    confirmPassword: string,
    sent = false,
    error = false,
    used = false

  const resetPassword = async () => {
    if (password == confirmPassword) {
		  const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/auth/reset-password`, {
		  	method: 'POST',
		  	body: JSON.stringify({
          Token: token,
          Password: password 
		  	}),
		  	headers: {
		  		'content-type': 'application/json'
		  	}
		  })
      if (response.ok) {
        used = false
        sent = true
        password = ""
        confirmPassword = ""
      } else {
        used = true
      }
    } else {
      error = true
    }
  }
</script>

<div class="w-screen flex flex-col items-center text-zinc-800 pt-16">
  <div class="max-w-xl mt-12">
    <h1 class="font-bold text-3xl w-full">Reset Password</h1>
    <form 
      class="flex flex-col mt-24"
      on:submit|preventDefault={resetPassword}
    >
      <div class="font-medium mb-4">New password</div>
      <input 
        type="password"
        autocomplete="off"
        bind:value={password}
        class="px-3 border-2 border-zinc-200 rounded focus:border-red-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800 mb-8"
      />
      <div class="font-medium mb-4">Confirm new password</div>
      <input 
        type="password"
        autocomplete="off"
        bind:value={confirmPassword}
        class="px-3 border-2 border-zinc-200 rounded focus:border-red-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
      />
      <button class="p-4 py-2 bg-red-800 hover:bg-red-900 font-medium text-white text-sm rounded mt-12 mb-12" type="submit">Update Password</button>
      {#if sent}
        <SuccessMessage message="Password updated." />
      {/if}
      {#if error}
        <ErrorMessage message="Passwords do not match." />
      {/if}
      {#if used}
        <ErrorMessage message="Password has been used in the past." />
      {/if}
    </form>
  </div>
</div>
