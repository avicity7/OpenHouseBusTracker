<script lang="ts">
  let email: string
  import { PUBLIC_BACKEND_URL } from "$env/static/public";
	import SuccessMessage from "$lib/components/SuccessMessage.svelte";

  let sent = false;

  const startResetPassword = async () => {
		const response = await fetch(`${PUBLIC_BACKEND_URL}:3000/auth/start-reset-password`, {
			method: 'POST',
			body: JSON.stringify({
				Email: email,
			}),
			headers: {
				'content-type': 'application/json'
			}
		})

    if (response.ok) {
      sent = true
      email = ""
    }
  }
</script>

<div class="w-screen flex flex-col items-center text-zinc-800 pt-16">
  <div class="max-w-xl mt-12">
    <h1 class="font-bold text-3xl w-full">Reset Password</h1>
    <form 
      class="flex flex-col mt-24"
      on:submit={startResetPassword}
    >
      <div class="font-medium mb-4">Email</div>
      <input 
        type="text"
        autocomplete="off"
        name="email"
        bind:value={email}
        class="px-3 border-2 border-zinc-200 rounded focus:border-red-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
      />
      <button class="p-4 py-2 bg-red-800 hover:bg-red-900 font-medium text-white text-sm rounded mt-12 mb-12" type="submit">Send password reset email</button>
      {#if sent}
        <SuccessMessage message="Password reset email sent." />
      {/if}
    </form>
  </div>
</div>
