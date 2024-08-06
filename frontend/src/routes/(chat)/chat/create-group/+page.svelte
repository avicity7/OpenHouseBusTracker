<script lang="ts">
	import type { User } from '$lib/types/global';
	import { onMount } from 'svelte';

	export let data;
	const { users } = data;

	let filteredNames: User[] = []
	let searchQuery: string = '';

	let selectedNames: Array<User>= []


	function filterNames() {
		let arr: Array<User> = []
		users.forEach((user) => {
			if (user.Name.toLowerCase().includes(searchQuery.toLowerCase())) arr.push(user)
		})
		filteredNames = arr
	}

	function toggleNameSelection(user: User, checked: boolean) {
		if (checked) {
			selectedNames.push(user)
		} else {
			selectedNames.splice(selectedNames.indexOf(user), 1)
		}
		filteredNames = filteredNames
	}

	function handleCheckboxChange(event: Event, user: User) {
		const target = event.target as HTMLInputElement;
		toggleNameSelection(user, target.checked);
	}

	onMount(() => {
		filteredNames = users
	})
</script>

<div class="flex justify-center items-center h-full">
	<div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-20">
		<h1 class="text-2xl font-semibold mb-4">Create Group Chat</h1>
		<form method="POST" action="?/createGroupChat">
			<div class="mb-4">
        <label for="groupname" class="text-sm font-medium mb-1">Group Name:</label>
        <div class="mb-4">
          <input
            type="text"
            name="groupname"
            placeholder="Enter group name"
            class="block w-full px-3 py-2 border rounded-md text-sm focus:border-red-600 focus:outline-none"
          />
        </div>
				<fieldset>
					<legend class="block text-sm font-medium mb-1">Names:</legend>
					<div class="mb-4 ">
						<input
							type="text"
							placeholder="Search names..."
							bind:value={searchQuery}
							on:input={filterNames}
							class="block w-full px-3 py-2 border rounded-md text-sm focus:border-red-600 focus:outline-none"
							/>
					</div>
					{#if filteredNames.length === 0 || filteredNames === null}
						<p class="text-sm text-gray-500">No options available</p>
					{:else}
						{#each filteredNames as user}
							<div class="flex items-center mb-2">
								<input
									type="checkbox"
									name="name"
									id={user.Name}
									value={user.Email}
									on:change={(e) => handleCheckboxChange(e, user)}
									class="mr-2"
									checked={selectedNames.includes(user)}
								/>
								<label for={user.Name} class="text-sm">{user.Name}</label>
							</div>
						{/each}
					{/if}
				</fieldset>
			</div>

			<div class="mt-4 flex justify-center">
				<button type="submit" class="px-4 py-2 bg-red-800 text-white rounded-md hover:bg-red-700">
					Create Chat Group
				</button>
			</div>
		</form>
	</div>
</div>
