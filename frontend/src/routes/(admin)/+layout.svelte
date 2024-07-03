<script lang="ts">
  import SideBarHeader from '$lib/components/SideBarHeader.svelte';
  import { fade } from 'svelte/transition';

  export let data;
</script>
<aside class="fixed h-full bg-white flex flex-col block">
  <SideBarHeader uri="/admin/users" text="Users" />
  <SideBarHeader uri="/admin/schedule" text="Schedules" />
  <SideBarHeader uri="/admin/drivers" text="Drivers" />
  <SideBarHeader uri="/admin/buses" text="Buses" />
  <SideBarHeader uri="/admin/bus-routes" text="Routes" />
  <SideBarHeader uri="/admin/event-helper" text="Helpers" />
</aside>
<aside class="ml-40 bg-zinc-100 transition-all ease-in-out duration-50">
  {#key data.pathname}
    <div in:fade={{ duration: 100, delay: 150 }} out:fade={{ duration: 100 }}>
      <slot />
    </div>
  {/key}
</aside>