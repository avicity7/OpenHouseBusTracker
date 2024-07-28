<script lang="ts">
	import type { EventHelper } from '$lib/types/global';

  export let data
  const { backend_uri, helpers, account } = data
  
	let selectedShift: EventHelper

  const createSwapShift = async () => {
    await fetch(`${backend_uri}:3000/event-helpers/create-swap-request`, {
      method: 'POST',
			body: JSON.stringify({
				From: account?.Email,
				With: selectedShift.Email
			})
    })
    location.replace('/schedule')
  }
</script>
<div class="flex justify-center items-center h-full">
	<div class="bg-white shadow rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-20">
		<h1 class="text-2xl font-semibold mb-4">Swap shift</h1>
		<form on:submit|preventDefault={createSwapShift}>
			<div class="mb-4">
				<label for="startTime" class="block text-sm font-medium mb-1">Student Helper</label>
				<select bind:value={selectedShift}>
					{#each helpers as helper}
						<option value={helper}>
							{helper.Name} {helper.Shift ? "(AM)" : "(PM)"}
						</option>
					{/each}
				</select>
			</div>

			<div class="mt-4 flex justify-center">
				<button type="submit" class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800">
					Request to swap shift
				</button>
			</div>
		</form>
	</div>
</div>
