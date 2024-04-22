<script lang="ts">
	import { onMount } from 'svelte';

  export let data
  onMount(() => {
    if (data.session) {
      window.location.replace('/profile')
    }
  })
</script>
<div class="max-w-xl text-zinc-800 pt-16">
  <h1 class="w-screen pt-12 pl-6 md:pl-24 font-bold text-3xl mb-24">Sign Up</h1>
  <form 
    class="flex flex-col pl-6 md:pl-24 mb-20"
    method='POST'
    action="?/signUp"
  >
    <div class="font-medium mb-4">Username</div>
    <input 
      type="text"
      autocomplete="off"
      name="username"
      class="border-2 border-zinc-200 rounded-full focus:border-purple-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
    />
    <div class="font-medium mb-4 mt-8">Email</div>
    <input 
      type="text"
      autocomplete="off"
      name="email"
      class="border-2 border-zinc-200 rounded-full focus:border-purple-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
    />
    <div class="font-medium mb-4 mt-8">Password</div>
    <input 
      type="text"
      autocomplete="off"
      name="password"
      class="border-2 border-zinc-200 rounded-full focus:border-purple-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
    />
    <button class="py-1 bg-purple-400 hover:bg-purple-500 font-medium text-white rounded-full mt-6" type="submit">Sign up</button>
    <div class="text-zinc-400 font-medium text-sm mx-6 md:mx-24 mt-8 text-center">
      Already have an account?
      <a href="/profile" class="text-blue-400 hover:text-blue-500">
        Sign in
      </a>
    </div>
  </form>
</div>
