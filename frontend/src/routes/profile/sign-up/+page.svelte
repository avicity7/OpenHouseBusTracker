<script lang="ts">
	import { onMount } from 'svelte';

  export let data
  onMount(() => {
    if (data.session) {
      window.location.replace('/profile')
    }
  })
</script>

<svelte:head>
	<title>Sign up | SPOH Bus Tracker</title>
</svelte:head>

<div class="w-screen flex flex-col items-center text-zinc-800 pt-16">
  <div class="max-w-xl mt-12">
    <h1 class="font-bold text-3xl w-full">Sign Up</h1>
    <form 
      class="flex flex-col"
      method='POST'
      action="?/signUp"
    >
      <div class="font-medium mb-4 mt-8">Name</div>
      <input 
        type="text"
        autocomplete="off"
        name="name"
        class="px-3 border-2 border-zinc-200 rounded-full focus:border-red-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
      />
      <div class="font-medium mb-4 mt-8">Email</div>
      <input 
        type="text"
        autocomplete="off"
        name="email"
        class="px-3 border-2 border-zinc-200 rounded-full focus:border-red-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
      />
      <div class="font-medium mb-4 mt-8">Password</div>
      <input 
        type="password"
        autocomplete="off"
        name="password"
        class="px-3 border-2 border-zinc-200 rounded-full focus:border-red-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
      />
      <button class="py-1 bg-red-800 hover:bg-red-900 font-medium text-white rounded-full mt-6" type="submit">Sign up</button>
      <div class="text-zinc-400 font-medium text-sm mx-6 md:mx-24 mt-8 text-center">
        Already have an account?
        <a href="/profile" class="text-blue-400 hover:text-blue-500">
          Sign in
        </a>
      </div>
    </form>
  </div>
</div>
