<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  let csvFile: HTMLInputElement | null = null;
  let fileName: string = '';
  let dragActive = false;

  function handleFileChange() {
    if (csvFile && csvFile.files) {
      const file = csvFile.files[0];
      fileName = file.name;
      dispatch('fileSelected', file);
    }
  }

  function handleDragOver(event: DragEvent) {
    event.preventDefault();
    dragActive = true;
  }

  function handleDragLeave() {
    dragActive = false;
  }

  function handleDrop(event: DragEvent) {
    event.preventDefault();
    dragActive = false;

    if (event.dataTransfer && event.dataTransfer.files) {
      csvFile!.files = event.dataTransfer.files;
      handleFileChange();
    }
  }
</script>

<div
  role="button"
  tabindex="0"
  aria-label="File upload area, drag and drop or browse files"
  class={`bg-gray-100 border-dashed border-2 rounded-lg p-4 w-full md:w-1/2 text-center ${dragActive ? 'border-red-400' : 'border-gray-300'}`}
  on:dragover={handleDragOver}
  on:dragleave={handleDragLeave}
  on:drop={handleDrop}
>
  <label for="file-upload" class="block cursor-pointer text-gray-700 font-semibold mb-2">
    <div class="flex justify-center mb-2">
      <svg xmlns="http://www.w3.org/2000/svg" class="w-12 h-12 text-red-600" viewBox="0 0 24 24">
        <g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2">
          <path d="M22 12h-6l-2 3h-4l-2-3H2" />
          <path d="M5.45 5.11L2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11" />
        </g>
      </svg>
    </div>
    <span class="block text-md mb-2">Drag & drop your CSV file here</span>
    <span class="block text-sm text-gray-500">or</span>
    <span class="block text-md text-red-600 hover:text-red-800" data-testid="file-input">Browse Files</span>
    <input
      id="file-upload"
      type="file"
      accept=".csv"
      bind:this={csvFile}
      class="sr-only"
      on:change={handleFileChange}
    />
  </label>
  {#if fileName}
    <p class="text-sm text-gray-500 mt-2">Selected file: <strong>{fileName}</strong></p>
  {/if}
  <p class="text-sm text-gray-500">Only CSV files are accepted.</p>
</div>
