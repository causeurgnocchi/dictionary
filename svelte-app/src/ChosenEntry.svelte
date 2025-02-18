<script lang="ts">
  import Word from "./Word.svelte";

  let { vocabulary }: { vocabulary: Vocabulary } = $props();
  let chosenMeaning = $state(vocabulary.meanings[0]);
  let unchosenMeanings = $derived(
    vocabulary.meanings.filter((meaning) => meaning !== chosenMeaning)
  );
</script>

<div class="chosen-entry">
  <Word {vocabulary} --margin="auto" />
  <span class="chosen-meaning">{chosenMeaning}</span>
  {#if vocabulary.meanings.length > 1}
    <div class="meanings">
      {#each unchosenMeanings as meaning}
        <button
          class="meaning"
          onclick={() => {
            chosenMeaning = meaning;
          }}>{meaning}</button
        >
      {/each}
    </div>
  {/if}
</div>

<style>
  .chosen-entry {
    width: 350px;
    margin: var(--margin, 0);
  }

  .chosen-meaning {
    display: block;
    text-align: center;
  }

  .meanings {
    display: flex;
    flex-direction: column;
    row-gap: 15px;
    margin-top: var(--space-between, 0);
    list-style: none;
  }

  .meaning {
    border: 1px black solid;
    border-radius: 5px;
    text-align: center;
    cursor: pointer;
  }
</style>
