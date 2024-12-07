<script lang="ts">
    import Vocabulary from '../../Vocabulary.svelte';
    import StagingArea from '../../StagingArea.svelte';
    import Header from '../../Header.svelte';
    import SearchBar from '../../SearchBar.svelte';
    import type { PageData } from './$types';

    let { data }: { data : PageData } = $props();
    let chosenVocabulary = $state.raw(data.vocabularies.length > 0 ? data.vocabularies[0] : null);
    let stagingAreaVocabulary = $derived(chosenVocabulary || (data.vocabularies.length > 0 ? data.vocabularies[0] : null));
    let remainingVocabularies = $derived(data.vocabularies.filter((v: Vocabulary) => {
        console.log(chosenVocabulary);
        return v !== chosenVocabulary;
    }));

    function setChosenVocabulary(vocabulary: Vocabulary) {
        chosenVocabulary = vocabulary;
    }

    $effect(() => {
        chosenVocabulary = data.vocabularies.length > 0 ? data.vocabularies[0] : null;
    });
</script>

<div class="container">
    <Header pageTitle="Dictionary"/>
    <SearchBar {onsubmit}/>
    {#if data.vocabularies.length > 0}
        <StagingArea vocabulary={stagingAreaVocabulary} />
        {#if data.vocabularies.length > 1}
            <div class="vocabularies">
                {#each remainingVocabularies as vocabulary}
                    <Vocabulary {vocabulary} onclick={()=>setChosenVocabulary(vocabulary)} />
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