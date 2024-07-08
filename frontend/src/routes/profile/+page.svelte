<script lang='ts'>
  import UserSettingsOutline from '$lib/components/UserSettingsOutline.svelte';
  export let data
  let { backend_uri, session, account } = data;
	import { onMount } from 'svelte';

  onMount(() => {
    if (session && account.Email == '') {
      location.replace("/")
    }
  })
</script>

<div class="flex flex-col justify-center items-center mt-28">
  <div class="font-albert min-w-2xl text-zinc-800 dark:text-zinc-100">
    {#if session && account.Email != ''}
      <div class="flex flex-col">
        <div class="max-h-fit max-w-fit relative pr-2 pb-2">
          <div class={`text-lg relative inline-flex items-center justify-center w-16 h-16 overflow-hidden ${session?.Role == "admin" ? "bg-red-700" : "bg-orange-800"} rounded-full`}>
              {#if session?.Role == "admin"}
                <span class="font-medium text-white">A</span>
              {:else}
                <span class="font-medium text-white">SH</span>
              {/if}
          </div>
          <a href="/profile/settings" class="absolute bottom-0 right-0">
            <UserSettingsOutline />
          </a>
        </div>
        <h2 class="text-2xl font-semibold mt-6">{account.Name}</h2>
        <h2 class="text-2xl font-semibold mt-6">{account.Email}</h2>
        <h2 class="text-xl font-semibold mt-6">+65 {account.Contact}</h2>
        <h3 class="text-xl mt-2 text-stone-700">{account.Role}</h3>

        <form method="POST" action="?/signOut">
          <button
            class="mt-6 px-3 py-1.5 bg-red-800 hover:bg-red-900 hover:text-zinc-200 rounded-md text-zinc-100 text-sm font-medium"
          >
            <div>Sign Out</div>
          </button>
        </form>
      </div>
    {:else}
      <div class="flex flex-col">
        <h1 class="font-bold text-3xl mb-12 w-full">Sign in</h1>
        <form class="flex flex-col" method="POST" action="?/login">
          <div class="font-medium mb-4">Email</div>
          <input
            data-testid="sign-in-email-input"
            type="text"
            autocomplete="off"
            name="email"
            class="px-3 border-2 border-zinc-200 rounded-full focus:border-red-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
          />
          <div class="font-medium mb-4 mt-8">Password</div>
          <input
            data-testid="sign-in-password-input"
            type="password"
            autocomplete="off"
            name="password"
            class="px-3 border-2 border-zinc-200 rounded-full focus:border-red-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
          />
          <button
            class="py-1 bg-red-800 hover:bg-red-900 font-medium text-white rounded-full mt-6"
            type="submit">Sign in</button
          >
          <div class="text-zinc-400 font-medium text-sm mx-6 md:mx-24 mt-8 text-center">
            Don't have an account yet?
            <a href="/profile/sign-up" class="text-blue-400 hover:text-blue-500"> Sign up </a>
          </div>
        </form>
      </div>
    {/if}
  </div>
</div>