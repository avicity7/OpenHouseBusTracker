  <script lang="ts">
    export let message: string | null = null;
  </script>
  
  {#if message}
    <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-4" role="alert">
      <span class="block sm:inline">{message}</span>
    </div>
  {/if}