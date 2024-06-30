<script lang="ts">
  import { page } from '$app/stores';  

  export let uri
  export let text
</script>

<a href={uri} class={"transition-color ease-in-out duration-100 mt-10 px-8 py-2 mb-4 mx-2 rounded-lg font-medium "+($page.url.pathname == uri ? "text-white bg-red-800" : "hover:bg-stone-200")}>{text}</a>