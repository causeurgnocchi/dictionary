<script lang="ts">
    import { page } from '$app/stores';
    import Vocabulary from '../../Vocabulary.svelte';
    import StagingArea from '../../StagingArea.svelte';
    import Header from '../../Header.svelte';
    import SearchBar from '../../SearchBar.svelte';
    import type { PageData } from './$types';

    const search = $page.url.searchParams.get('search') || '';
    let { data }: { data : PageData } = $props();
    let vocabularies = $derived(data.vocabularies);
    let chosenVocabulary = $derived(data.vocabularies[0]);
    let remainingVocabularies = $derived(vocabularies.filter((v: Vocabulary) => v !== chosenVocabulary));
</script>

<div class="container">
    <Header pageTitle="Dictionary"/>
    <SearchBar lastSearch={search}/>
    {#if vocabularies.length > 0}
        <StagingArea vocabulary={chosenVocabulary} />
        {#if remainingVocabularies.length > 0}
            <div class="vocabularies">
                {#each remainingVocabularies as vocabulary}
                    <Vocabulary {vocabulary} onclick={()=>{}} />
                {/each}
            </div>
        {/if}
    {/if}
</div>

<style>
    *, *::before, *::after {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }

    .container {
        display: grid;
        grid-template-rows: 50px 44px 1fr;
    }

    .vocabularies {
        display: grid;
        grid-template-columns: repeat(auto-fit, 500px);
        justify-content: space-around;
        row-gap: 30px;
        padding: 20px;
    }
</style>