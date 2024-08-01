<script lang="ts">
  export let data

  const { swaps } = data
  const formatTime = (timestamp: string | number | Date) => {
    const utcDate = new Date(timestamp);
    return utcDate.toLocaleTimeString([], { day: "numeric", month: "short", hour: '2-digit', minute: '2-digit' });
  }
</script>

<div class="p-8">
  <table class="min-w-full divide-y divide-gray-200">
    <thead class="bg-gray-50">
      <tr>
        <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
          Time
        </th>
        <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
          From
        </th>
        <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
          With
        </th>
        <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
          For
        </th>
      </tr>
    </thead>
    <tbody class="bg-white divide-y divide-gray-200">
      {#if swaps.length != 0}
        {#each swaps as swap}
          <tr class="hover:bg-gray-100">
            <td class="px-6 py-4 whitespace-nowrap">{formatTime(swap.Timestamp)}</td>
            <td class="px-6 py-4 whitespace-nowrap">{swap.FromName}</td>
            <td class="px-6 py-4 whitespace-nowrap">{swap.WithName}</td>
            <td class="px-6 py-4 whitespace-nowrap">{swap.TargetShift ? "PM": "AM"}</td>
          </tr>
        {/each}
      {:else}
        <tr>
          <td colspan="8" class="px-6 py-4 whitespace-nowrap text-center">
            No swaps to display.
          </td>
        </tr>
      {/if}
    </tbody>
  </table>
</div>