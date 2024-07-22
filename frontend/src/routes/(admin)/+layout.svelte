<script lang="ts">
  import SideBarHeader from '$lib/components/SideBarHeader.svelte';
</script>
<aside class="fixed h-full bg-white flex flex-col block" data-testid="sidebar">
  <SideBarHeader uri="/admin/users" text="Users" />
  <SideBarHeader uri="/admin/schedule" text="Schedules"/>
  <SideBarHeader uri="/admin/drivers" text="Drivers" />
  <SideBarHeader uri="/admin/buses" text="Buses" />
  <SideBarHeader uri="/admin/bus-routes" text="Routes" />
  <SideBarHeader uri="/admin/event-helper" text="Helpers" />
</aside>
<aside class="ml-40 bg-zinc-100 transition-all ease-in-out duration-50 min-h-[94vh]">
  <slot />
</aside>