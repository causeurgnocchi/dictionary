<script lang="ts">
    import FuriganaWriting from "./FuriganaWriting.svelte";

    let {vocabulary} : {vocabulary: Vocabulary} = $props();
    let chosenMeaning = $state.raw(vocabulary.meanings[0]);
</script>

<div class="wrapper">
    <div class="staging-area">
        <div class="furigana-writing">
            <FuriganaWriting {vocabulary} />
        </div>
        <p class="chosen-meaning">{chosenMeaning}</p>
        {#if vocabulary.meanings.length > 1}
            <div class="meanings">
                {#each vocabulary.meanings.filter((meaning) => meaning !== chosenMeaning) as meaning}
                    <button class="meaning" onclick={() => { chosenMeaning = meaning }}>{meaning}</button>
                {/each}
            </div>
        {/if}
    </div>
</div>

<style>
    *, *::before, *::after {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }
    
    .wrapper {
        display: flex;
        flex-direction: column;
    }
    
    .staging-area {
        width: 350px;
        margin: 10px 0;
        padding: 20px 0;
        display: flex;
        flex-direction: column;
        align-self: center;
    }

    .furigana-writing {
        align-self: center;
    }

    .chosen-meaning {
        text-align: center;
        margin-top: 10px;
        margin-bottom: 25px;
    }

    .meanings {
        margin-top: var(--margin-top, 0);
        list-style: none;
        display: flex;
        flex-direction: column;
        row-gap: 13px;
    }

    .meaning {
        border: 1px black solid;
        border-radius: 5px;
        padding: 5px 10px;
        text-align: center;
    }
</style>