<script lang="ts">
  import { page } from '$app/stores';  
</script>
<aside class="fixed h-full bg-white flex flex-col block">
  <a href="/admin/users" class={"mt-10 px-8 py-2 mb-4 rounded-r-full font-medium "+($page.url.pathname == '/admin/users' ? "text-white bg-red-800" : "hover:text-red-700")}>Users</a>
  <a href="/admin/schedule" class={"px-8 py-2 mb-4 rounded-r-full font-medium "+($page.url.pathname == '/admin/schedule' ? "text-white bg-red-800" : "hover:text-red-700")}>Schedule</a>
  <a href="/admin/drivers" class={"px-8 py-2 mb-4 rounded-r-full font-medium "+($page.url.pathname == '/admin/drivers' ? "text-white bg-red-800" : "hover:text-red-700")}>Drivers</a>
  <a href="/admin/buses" class={"px-8 py-2 mb-4 rounded-r-full font-medium "+($page.url.pathname == '/admin/buses' ? "text-white bg-red-800" : "hover:text-red-700")}>Buses</a>
  <a href="/admin/bus-routes" class={"px-8 py-2 mb-4 rounded-r-full font-medium "+($page.url.pathname == '/admin/bus-routes' ? "text-white bg-red-800" : "hover:text-red-700")}>Routes</a>
  <a href="/admin/event-helper" class={"px-8 py-2 mb-4 rounded-r-full font-medium "+($page.url.pathname == '/admin/event-helper' ? "text-white bg-red-800" : "hover:text-red-700")}>Event Helper</a>
</aside>
<aside class="ml-40 bg-zinc-100">
  <slot />
</aside>