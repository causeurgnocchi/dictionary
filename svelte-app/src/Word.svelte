<script lang="ts">
  let { vocabulary }: { vocabulary: Vocabulary } = $props();
</script>

<div class="word">
  {#if vocabulary.writing !== vocabulary.reading}
    <span class="reading">
      {#each vocabulary.characters as character}
        <span
          class="character-reading"
          style={`font-size: ${character.furigana.length === 1 ? 30 / 2 : 30 / character.furigana.length}px`}
          >{character.furigana}</span
        >
      {/each}
    </span>
  {/if}
  <span class="writing">
    {vocabulary.writing}
  </span>
</div>

<style>
  .word {
    width: fit-content;
    margin: var(--margin, 0);
  }

  .reading {
    display: block;
  }

  .character-reading {
    display: inline-block;
    width: 30px;
    text-align: center;
  }

  .writing {
    font-size: 30px;
  }
</style>
