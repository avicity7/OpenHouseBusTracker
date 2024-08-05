<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import { onMount } from 'svelte';
  import ErrorMessage from '$lib/components/ErrorMessage.svelte';

  export let data
  export let form
  const { captcha_key } = data
  
  const getCaptchaRes = () => {
    grecaptcha.ready(function() {
      grecaptcha.render("g-recaptcha", {
        "sitekey": captcha_key
      });
    });
  }

  onMount(() => {
    if (data.session) {
      window.location.replace('/profile')
    }
    getCaptchaRes()
  })

  afterNavigate(getCaptchaRes)
</script>

<svelte:head>
  <script src="https://www.google.com/recaptcha/api.js?render=explicit" async defer></script>
	<title>Sign up | SPOH Bus Tracker</title>
</svelte:head>

<div class="w-screen flex flex-col items-center text-zinc-800 pt-16">
  <div class="max-w-xl mt-12">
    <h1 class="font-bold text-3xl w-full mb-6">Sign Up</h1>
    {#if form?.success === false}
      <ErrorMessage message="Ensure all fields are filled correctly."/>
    {/if}
    <form 
      class="flex flex-col"
      method='POST'
      action="?/signUp"
    >
      <div class="font-medium mb-4 mt-2">Name</div>
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
        class="px-3 border-2 border-zinc-200 rounded-full focus:border-red-400 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800 mb-8"
      />
      <div id="g-recaptcha">
      </div>
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
