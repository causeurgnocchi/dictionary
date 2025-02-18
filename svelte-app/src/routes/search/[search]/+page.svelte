<script lang="ts">
  import ChosenEntry from "../../../ChosenEntry.svelte";
  import Entry from "../../../Entry.svelte";

  let { data } = $props();
  let { results } = $derived(data);
  let clickedEntry = $state(results[0]);
  let chosenEntry = $derived(clickedEntry || results[0]);
  let unchosenEntries = $derived(
    results.filter((r: Entry) => r != chosenEntry)
  );

  $effect(() => {
    clickedEntry = results[0];
  });
</script>

<div class="chosen-entry-container">
  <ChosenEntry
    vocabulary={chosenEntry}
    --margin="auto"
    --space-between="20px"
  />
</div>
<div class="entries">
  {#each results.slice(1) as v}
    <Entry
      vocabulary={v}
      onclick={() => {
        clickedEntry = v;
      }}
    />
  {/each}
</div>

<style>
  .chosen-entry-container {
    box-shadow: 0 5px 5px rgb(128, 128, 128, 0.5);
    padding-bottom: 20px;
  }

  .entries {
    width: fit-content;
    margin: 20px auto 0;
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
</style>
