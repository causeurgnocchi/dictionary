<script lang="ts">
    import FuriganaWriting from "./FuriganaWriting.svelte";
    import Meanings from "./Meanings.svelte";
    
    const {vocabulary} : {vocabulary: Vocabulary} = $props();
</script>

<div>
    <FuriganaWriting {vocabulary} />
    <Meanings meanings={vocabulary.meanings} />
</div>

<style>
    *, *::before, *::after {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }

    div {
        border: 1px solid black;
        width: 500px;
        padding: 10px;
    }
</style>