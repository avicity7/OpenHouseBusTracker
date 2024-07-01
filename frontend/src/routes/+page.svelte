<script lang="ts">

	import BusTrackerMap from "$lib/components/BusTrackerMap.svelte";
	import LandingPage from "$lib/components/LandingPage.svelte";
	
	export let data;
	const { session } = data;

</script>

{#if session?.Role == "admin"}
	<BusTrackerMap {data}/>
{:else}
	<LandingPage {data}/>
{/if}
