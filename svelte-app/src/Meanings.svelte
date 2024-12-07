<script lang="ts">
    let {meanings} : {meanings: string[]} = $props();
</script>

<ul>
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

    ul {
        padding-inline-start: 30px;
        margin-top: var(--margin-top, 0);
    }
</style>