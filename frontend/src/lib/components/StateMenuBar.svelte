<!-- StateMenuBar.svelte -->
<script lang="ts">
  export let initialState: string;
  export let onStateChange: (newState: string) => void; 

  let currentState = initialState;

  const states = ["Less Than Half", "Half Full", "Full Bus"];

  const handleStateChange = (newState: string) => {
    currentState = newState; 
    onStateChange(newState); 
  };
</script>

<div class="flex justify-center rounded-lg overflow-hidden border max-w-xs sm:max-w-md">
  {#each states as menuState}
    <button 
      class="flex-1 py-2 text-sm font-semibold focus:outline-none focus:bg-red-700 focus:text-white
             hover:bg-red-300 hover:text-gray-800 transition-colors duration-200
             {menuState === currentState ? 'bg-red-700 text-white' : 'bg-red-200 text-gray-700'}"
      on:click={() => handleStateChange(menuState)}
    >
      {menuState}
    </button>
  {/each}
</div>
