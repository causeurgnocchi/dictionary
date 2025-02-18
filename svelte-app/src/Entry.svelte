<script lang="ts">
  import Word from "./Word.svelte";

  let { vocabulary, onclick } = $props();
</script>

<button class="entry" {onclick}>
  <Word {vocabulary} />
  <ul class="meanings">
    {#each vocabulary.meanings as meaning}
      <li class="meaning">{meaning}</li>
    {/each}
  </ul>
</button>

<style>
  .entry {
    width: 500px;
    padding: 10px;
    border: 1px solid black;
    border-radius: 5px;
    box-shadow: 5px 5px 5px rgb(128, 128, 128, 0.5);
  }

  .meaning {
    text-align: start;
  }
</style>
