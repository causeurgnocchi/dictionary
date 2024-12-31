<script lang="ts">
    import FuriganaWriting from "./FuriganaWriting.svelte";
    
    let {vocabulary, onclick} : {vocabulary: Vocabulary, onclick: () => void} = $props();
</script>

<button class="container" onclick={onclick}>
    <FuriganaWriting {vocabulary} />
    <ul class="meanings">
        {#each vocabulary.meanings as meaning}
            <li class="meaning">{meaning}</li>
        {/each}
    </ul>
</button>

<style>
    *, *::before, *::after {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }

    .container {
        border: 1px solid black;
        width: 500px;
        padding: 10px;
        box-shadow: 5px 5px 5px rgb(128, 128, 128, 0.5);
        text-align: start;
        display: flex;
        flex-direction: column;
        font-family: 'Times New Roman';
        font-size: 16px;
    }

    .meanings {
        margin-top: 15px;
        padding-inline-start: 30px;
    }
</style>