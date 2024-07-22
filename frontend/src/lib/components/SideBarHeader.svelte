<script lang="ts">
  import { page } from '$app/stores';  
  import User from './User.svelte';
  import Schedule from './Schedule.svelte';
  import Driver from './Driver.svelte';
  import Bus from './Bus.svelte';
  import Route from './Route.svelte';
  import Helper from './Helper.svelte';

  export let uri
  export let text
</script>

<a href={uri} class={"text-sm flex flex-row items-center mt-10 px-6 py-2 mb-4 mx-2 rounded-lg font-medium "+($page.url.pathname == uri ? "text-white bg-red-800" : "hover:bg-red-50")}>
  {#if text == "Users"}
    <User />
  {:else if text == "Schedules"}
    <Schedule />
  {:else if text == "Drivers"}
    <Driver />
  {:else if text == "Buses"}
    <Bus />
  {:else if text == "Routes"}
    <Route />  
  {:else if text == "Helpers"}
    <Helper />
  {/if}
  <span class="ml-3">
    {text}
  </span>
</a>