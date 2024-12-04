<script lang="ts">
    import { page } from '$app/stores';
    import Vocabulary from '../../Vocabulary.svelte';
    import StagingArea from '../../StagingArea.svelte';
    import Header from '../../Header.svelte';
    import SearchBar from '../../SearchBar.svelte';
    import type { PageData } from './$types';

    const search = $page.url.searchParams.get('search') || '';
    let { data }: { data : PageData } = $props();
    let chosenVocabulary = $state.raw(data.vocabularies.length > 0 ? data.vocabularies[0] : []);
    let remainingVocabulary = $derived(data.vocabularies.filter((v: Vocabulary) => v !== chosenVocabulary));
    const vocabularyOnClick = (vocabulary: Vocabulary) => { chosenVocabulary = vocabulary; };
</script>

<div class="container">
    <Header pageTitle="Dictionary"/>
    <SearchBar lastSearch={search}/>
    {#if data.vocabularies.length > 0}
        <StagingArea vocabulary={chosenVocabulary} />
        {#if remainingVocabulary.length > 0}
            <div class="vocabularies">
                {#each remainingVocabulary as remaining}
                    <Vocabulary vocabulary={remaining} onclick={()=>vocabularyOnClick(remaining)} />
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
        display: flex;
        flex-direction: column;
        row-gap: 30px;
        padding: 20px;
        transition: 1s;
    }
</style>