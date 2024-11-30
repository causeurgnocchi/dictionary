<script lang="ts">
    const {meanings} : {meanings: string[]} = $props();
</script>

<ul class="meanings">
    {#each meanings as meaning}
    <li class="meaning">{meaning}</li>
    {/each}
</ul>

<style>
    *, *::before, *::after {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }

    .meanings {
        padding-inline-start: 30px;
        padding-top: 15px;
    }
</style>