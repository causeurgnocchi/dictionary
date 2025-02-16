<script lang="ts">
    import Word from "./Word.svelte";

    let {vocabulary} : {vocabulary: Vocabulary} = $props();
    let chosenMeaning = $state(vocabulary.meanings[0]);
    let unchosenMeanings = $derived(vocabulary.meanings.filter((meaning) => meaning !== chosenMeaning))
</script>


<div class="chosen-entry">
    <Word {vocabulary} />
    <span class="chosen-meaning">{chosenMeaning}</span>
    {#if vocabulary.meanings.length > 1}
    <div class="meanings">
        {#each unchosenMeanings as meaning}
        <button class="meaning" onclick={() => { chosenMeaning = meaning }}>{meaning}</button>
        {/each}
    </div>
    {/if}
</div>


<style>
    .chosen-entry {
        width: 350px;
    }

    .chosen-meaning {
        text-align: center;
        margin-top: 10px;
        margin-bottom: 25px;
    }

    .meanings {
        display: flex;
        flex-direction: column;
        row-gap: 13px;
        margin-top: var(--margin-top, 0);
        list-style: none;
    }

    .meaning {
        border: 1px black solid;
        border-radius: 5px;
        /* padding: 5px 10px; */
        text-align: center;
        /* cursor: pointer; */
    }
</style>