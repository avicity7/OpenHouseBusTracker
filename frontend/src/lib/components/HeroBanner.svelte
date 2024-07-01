<script>
  export let bold
  export let text
  export let source
</script>


<div class="max-w-3xl">
  <h2 class="w-full flex flex-row justify-center mt-20 mb-8 font-semibold text-2xl text-stone-700 whitespace-pre">{bold} 
    <span class="font-light">{text}</span>
  </h2>
  <img 
    src={source} 
    alt={`${bold} ${text}`}
    class="rounded-3xl aspect-video object-cover"
  />
</div>