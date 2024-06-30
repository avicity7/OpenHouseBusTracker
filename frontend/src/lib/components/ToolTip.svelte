<script>
    export let text = '';
  
    let isVisible = false;
  
    const handleMouseEnter = () => {
      isVisible = true;
    };
  
    const handleMouseLeave = () => {
      isVisible = false;
    };
  </script>
  
  <span
    on:mouseenter={handleMouseEnter}
    on:mouseleave={handleMouseLeave}
    class="relative inline-block"
    role="tooltip"
  >
    <slot />
  
    {#if isVisible}
    <div class="absolute bottom-full left-1/2 -transtone-x-1/2 mb-2 px-4 py-2 bg-gray-800 text-white text-sm rounded-lg">
        {text}
      </div>
    {/if}
  </span>
  