<script lang="ts">
    export let checked: boolean;
    export let onToggle: (newChecked: boolean) => void;
  
    const handleToggle = () => {
      onToggle(!checked);
    };
  </script>
  
  <div class="flex items-center space-x-2">
    <!-- <span class="text-gray-500">Auto Scheduling</span>
    <div class="relative inline-block w-10 h-6 align-middle select-none transition duration-200 ease-in">
      <input 
        type="checkbox" 
        name="toggle" 
        id="toggle" 
        class="toggle-checkbox absolute block w-6 h-6 rounded-full bg-white border-4 appearance-none cursor-pointer transition-transform duration-200 ease-in-out transform translate-x-0 left-0 checked:translate-x-full checked:bg-green-400 checked:border-green-400" 
        checked={checked} 
        on:click={handleToggle} 
      />
      <label for="toggle" class="toggle-label block overflow-hidden h-6 rounded-full bg-gray-300 cursor-pointer transition duration-200 ease-in-out"></label>
    </div> -->

    <!-- <span class="text-gray-500">Auto Scheduling</span> -->
    <label class="inline-flex items-center cursor-pointer">
        <input
            type="checkbox"
            name="toggle" 
            id="toggle" 
            checked={checked}
            on:click={handleToggle}
            class="sr-only peer"
            >
        <div class="relative w-12 h-6 bg-red-300 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-red-300 dark:peer-focus:ring-red-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-red-700" />
    </label>
  </div>
  