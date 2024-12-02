<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import Vocabulary from '../../Vocabulary.svelte';
    import StagingArea from '../../StagingArea.svelte';
    import Header from '../../Header.svelte';
    import SearchBar from '../../SearchBar.svelte';

    const search = $page.url.searchParams.get('search') || '';
    let vocabularies = $state([]);
</script>

<div class="container">
    <Header pageTitle="Dictionary"/>
    <SearchBar lastSearch={search}/>
    {#if vocabularies.length > 0}
        <div class="vocabularies">
            <StagingArea vocabulary={vocabularies[0]} />
            {#if vocabularies.length > 1}
                {#each vocabularies.slice(1) as vocabulary}
                    <Vocabulary {vocabulary} />
                {/each}
            {/if}
        </div>
    {/if}
</div>

<style>
    *, *::before, *::after {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }

    .container {
        height: 100vh;
        display: grid;
        grid-template-rows: 50px 44px 1fr;
    }

    .vocabularies {
        display: flex;
        flex-direction: column;
        row-gap: 20px;
        padding: 0 20px 20px;
    }
</style>