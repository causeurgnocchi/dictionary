<script lang="ts">
    import FuriganaWriting from "./FuriganaWriting.svelte";
    import Meanings from "./Meanings.svelte";

    const {vocabulary} : {vocabulary: Vocabulary} = $props();
</script>

<div class="wrapper">
    <div class="staging-area">
        <div class="furigana-writing">
            <FuriganaWriting {vocabulary} />
        </div>
        <p class="chosen-meaning">{vocabulary.meanings[0]}</p>
        {#if vocabulary.meanings.length > 1}
            <Meanings meanings={vocabulary.meanings.slice(1)} />
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
        align-self: center;
        display: flex;
        flex-direction: column;
    }

    .furigana-writing {
        align-self: center;
    }

    .chosen-meaning {
        text-align: center;
    }
</style>