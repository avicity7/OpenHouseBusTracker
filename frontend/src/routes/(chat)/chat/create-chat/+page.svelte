<script lang="ts">
	export let data;
  import type { User } from '$lib/types/global.js';

	const { users, account, backend_uri } = data;

	let selectedUser: User = {} as User

  const createChatRoom = async () => {
		await fetch(`${backend_uri}:3000/chat/create-chat-room`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({SelectedUsers: [selectedUser.Email, account?.Email], Groupname: ""})
		});

		location.replace("/chat")
  }
</script>

<div class="flex justify-center items-center h-full">
	<div class="bg-white shadow-md rounded-lg p-8 w-full md:w-3/4 lg:w-2/3 xl:w-1/3 mt-24">
		<h1 class="text-2xl font-semibold mb-4">Start Chat</h1>
		<form>
			<div class="mb-4">
				<label for="user" class="block text-sm font-medium mb-1">User:</label>
				<select
          name="user"
					required
					bind:value={selectedUser}
					class="w-full rounded-md border border-gray-300 px-3 py-2 focus:outline-none focus:ring-1 focus:ring-blue-500"
				>
					{#each users as user}
						{#if user.Email != account?.Email}
							<option value={user}>{user.Name}</option>
						{/if}
					{/each}
				</select>
			</div>

			<div class="mt-4 flex justify-center">
				<button on:click={createChatRoom} class="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-800"
					>Start Chat</button
				>
			</div>
		</form>
	</div>
</div>
