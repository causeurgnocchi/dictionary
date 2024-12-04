<script lang="ts">
    import FuriganaWriting from "./FuriganaWriting.svelte";
    import Meanings from "./Meanings.svelte";
    
    const {vocabulary, onclick} : {vocabulary: Vocabulary, onclick: () => void} = $props();
</script>

<button class="container" onclick={onclick}>
    <FuriganaWriting {vocabulary} />
    <Meanings meanings={vocabulary.meanings} --margin-top="15px" />
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
        height: calc-size(auto);
        transition: height 1s;
        user-select: text;
    }
</style>