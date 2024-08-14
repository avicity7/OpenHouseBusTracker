<script lang="ts">
    export let show: boolean;
    export let message: string;
    export let confirmText: string;
    export let cancelText: string;
    export let onConfirm: () => void;
    export let onCancel: () => void;
  
    if (show && typeof window !== "undefined") {
      window.scrollTo(0, 0);
      document.body.style.overflow = 'hidden';
    } else if (!show && typeof window !== "undefined") {
      document.body.style.overflow = 'auto';
    }
  </script>
  
  {#if show}
    <div
      class="fixed top-0 left-0 w-full h-full flex items-center justify-center bg-gray-900 bg-opacity-50"
    >
      <div class="bg-white p-8 rounded-lg shadow-lg max-w-md w-full">
        <p class="text-lg mb-4">{message}</p>
        <div class="flex justify-end">
          <button class="px-4 py-2 mr-4 text-gray-600" on:click={onCancel}>{cancelText}</button>
          <button class="px-4 py-2 bg-red-700 text-white rounded" on:click={onConfirm}>{confirmText}</button>
        </div>
      </div>
    </div>
  {/if}
  