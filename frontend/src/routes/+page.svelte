<script lang="ts">

	import BusTrackerMap from "$lib/components/BusTrackerMap.svelte";
	import LandingPage from "$lib/components/LandingPage.svelte";
	import { onMount } from "svelte";
	
	export let data;
	const { session } = data;

</script>

<svelte:head>
	<title>SPOH Bus Tracker</title>
</svelte:head>

{#if session?.Role == "admin"}
	<BusTrackerMap {data}/>
{:else}
	<LandingPage {data}/>
{/if}
