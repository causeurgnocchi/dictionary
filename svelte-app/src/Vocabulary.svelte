<script lang="ts">
    import FuriganaWriting from "./FuriganaWriting.svelte";
    import Meanings from "./Meanings.svelte";
    
    const {vocabulary} : {vocabulary: Vocabulary} = $props();
</script>

<div class="container">
    <FuriganaWriting {vocabulary} />
    <div class="meanings-wrapper">
        <Meanings meanings={vocabulary.meanings} />
    </div>
</div>

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
    }

    .meanings-wrapper {
        margin-top: 15px;
    }
</style>