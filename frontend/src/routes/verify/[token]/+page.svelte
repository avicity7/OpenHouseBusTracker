<script>
  import { page } from '$app/stores'
	import { onMount } from 'svelte';
  export let data
  const { backend_uri } = data
  const token = $page.params.token

  let loaded = false

  const verifyToken = async() => {
    fetch(`${backend_uri}:3000/auth/verify/${token}`)
  }

  onMount(async() => {
    await verifyToken()
    loaded = true
  })
</script>

{#if loaded}
  <div class="flex flex-col text-center items-center my-auto mx-8">
    <h1 class="mt-16 text-3xl font-semibold">You're verified!</h1>
    <p class="mt-8 text-xl">Close this page and refresh the other window to proceed.</p>
  </div>
{:else}
  <div class="flex flex-col text-center items-center my-auto mx-8">
    <h1 class="mt-16 text-3xl font-semibold">We're verifying you.</h1>
    <p class="mt-8 text-xl">Hold on, this won't take too long.</p>
  </div>
{/if}
  
  