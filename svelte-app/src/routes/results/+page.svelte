<script lang="ts">
    import { page } from '$app/stores';
    import Vocabulary from '../../Vocabulary.svelte';
    import StagingArea from '../../StagingArea.svelte';
    import Header from '../../Header.svelte';
    import SearchBar from '../../SearchBar.svelte';
    import type { PageData } from './$types';

    const search = $page.url.searchParams.get('search') || '';
    let { data }: { data : PageData } = $props();
    let stagingArea = $state(data.vocabularies[0]);
</script>

<div class="container">
    <Header pageTitle="Dictionary"/>
    <SearchBar lastSearch={search}/>
    {#if data.vocabularies.length > 0}
        <div class="vocabularies">
            <StagingArea vocabulary={stagingArea} />
            {#if data.vocabularies.length > 1}
                {#each data.vocabularies.splice(data.vocabularies.filter((v: Vocabulary) => v !== stagingArea.Vocabulary)) as vocabulary}
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