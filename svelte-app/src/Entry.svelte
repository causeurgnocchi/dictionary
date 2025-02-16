<script lang="ts">
    import Word from './Word.svelte';

    let { vocabulary, onclick } = $props();
</script>

<button class="container" onclick={onclick}>
    <Word {vocabulary} />
    <ul class="meanings">
        {#each vocabulary.meanings as meaning}
        <li class="meaning">{meaning}</li>
        {/each}
    </ul>
</button>

<style>
    .container {
        width: 500px;
        padding: 10px;
        border: 1px solid black;
        box-shadow: 5px 5px 5px rgb(128, 128, 128, 0.5);
    }
</style>