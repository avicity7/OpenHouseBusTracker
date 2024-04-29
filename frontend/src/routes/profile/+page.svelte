<script lang='ts'>
  import { Avatar, Spinner } from 'flowbite-svelte';
  import type { User } from '../../types/global.js';
  export let data
  let { backend_uri, session, account } = data;
	import { onMount } from 'svelte';

  onMount(() => {
    if (session && account.Email == '') {
      location.reload()
    }
  })
</script>

<div class="bg-zinc-100 flex flex-col justify-center dark:bg-zinc-800 pb-32 pt-[8rem]">
  <div class="font-albert max-w-xl text-zinc-800 dark:text-zinc-100">
    {#if session && account.Email != ''}
      <div class="ml-6 mr-6 md:mr-0 md:ml-24 flex flex-col">
        <div class="max-h-fit max-w-fit relative pr-2 pb-2">
          <Avatar size="lg" />
          <a href="/profile/settings" class="absolute bottom-0 right-0">
            <!-- <UserSettingsOutline
              class="text-zinc-400 hover:text-purple-600 hover:dark:text-purple-400"
            /> -->
          </a>
        </div>
        <h2 class="text-2xl font-semibold mt-6">{account.Email}</h2>
        <h3 class="text-xl mt-2 text-stone-700">{account.Role}</h3>
        <!-- <h3 class="text-sm mt-2">{session?.Email}</h3> -->
          <form method="POST" action="?/signOut">
            <button
              class="mt-6 px-1 py-1.5 w-[20%] bg-red-800 hover:bg-red-900 hover:text-zinc-200 rounded-md text-zinc-100 text-sm font-medium"
            >
              <div>Sign Out</div>
            </button>
          </form>
      </div>
    {:else}
      <h1 class="w-full pt-12 pl-6 md:pl-24 font-bold text-3xl mb-24">Sign in</h1>
      <form class="flex flex-col pl-6 md:pl-24 mb-20" method="POST" action="?/login">
        <div class="font-medium mb-4">Email</div>
        <input
          type="text"
          autocomplete="off"
          name="email"
          class="px-3 border-2 border-zinc-200 rounded-full focus:border-purple-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
        />
        <div class="font-medium mb-4 mt-8">Password</div>
        <input
          type="text"
          autocomplete="off"
          name="password"
          class="px-3 border-2 border-zinc-200 rounded-full focus:border-purple-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
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
    {/if}
  </div>
</div>