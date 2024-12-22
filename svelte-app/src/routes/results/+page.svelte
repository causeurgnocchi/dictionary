<script lang="ts">
    import VocabularyComponent from '../../Vocabulary.svelte';
    import StagingArea from '../../StagingArea.svelte';
    import Header from '../../Header.svelte';
    import SearchBar from '../../SearchBar.svelte';

    import type { PageData } from './$types';
    import { flip } from 'svelte/animate';
    import { fade } from 'svelte/transition';

    type IndexedVocabulary = Vocabulary & { id: number; }; 

    let { data }: { data : PageData } = $props();
    let vocabularies = $derived(data.vocabularies.map((v: Vocabulary, i: number) => ({ id: i, ...v })));
    let chosenVocabulary = $state.raw(data.vocabularies.length > 0 ? { id: 0, ...data.vocabularies[0] } : null);
    let stagingAreaVocabulary = $derived(chosenVocabulary || (vocabularies.length > 0 ? vocabularies[0] : null));
    let remainingVocabularies = $derived(vocabularies.filter((v: IndexedVocabulary) => v.id !== chosenVocabulary.id));

    $effect(() => {
        chosenVocabulary = vocabularies.length > 0 ? vocabularies[0] : null;
    })
</script>

<div class="container">
    <Header pageTitle="Dictionary"/>
    <SearchBar />
    {#if vocabularies.length > 0}
        {#each [stagingAreaVocabulary] as v (v)}
            <div class="staging-area-wrapper" in:fade>
                <StagingArea vocabulary={v} />
            </div>
        {/each}
        {#if vocabularies.length > 1}
            <div class="vocabularies">
                {#each remainingVocabularies as vocabulary (vocabulary.id)}
                    <div class="vocabulary-wrapper" animate:flip>
                        <VocabularyComponent {vocabulary} onclick={() => {chosenVocabulary = vocabulary}} />
                    </div>
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
    }
</style>