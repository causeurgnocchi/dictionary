<script lang="ts">
    import VocabularyComponent from '../../Vocabulary.svelte';
    import StagingArea from '../../StagingArea.svelte';
    import Header from '../../Header.svelte';
    import SearchBar from '../../SearchBar.svelte';
    import type { PageData } from './$types';

    type IndexedVocabulary = Vocabulary & { id: number; };
    let { data }: { data : PageData } = $props();
    let vocabularies = $derived(data.vocabularies.map((v: Vocabulary, i: number) => ({ id: i, ...v })));
    let chosenVocabulary = $state.raw({ id: 0, ...data.vocabularies[0] });
    let stagingAreaVocabulary = $derived(chosenVocabulary || vocabularies[0]);
    let remainingVocabularies = $derived(vocabularies.filter((v: IndexedVocabulary) => v.id !== chosenVocabulary.id));

    $effect(() => { chosenVocabulary = vocabularies[0] })
</script>

<div class="upper-section">
    <Header pageTitle="Dictionary"/>
    <SearchBar />
    {#if vocabularies.length > 0}
        {#key stagingAreaVocabulary}
            <div class="staging-area-wrapper">
                <StagingArea vocabulary={stagingAreaVocabulary} />
            </div>
        {/key}
    {/if}
</div>
{#if vocabularies.length > 1}
    <div class="vocabularies">
        {#each remainingVocabularies as vocabulary (vocabulary.id)}
            <div class="vocabulary-wrapper-flip">
                <VocabularyComponent {vocabulary} onclick={() => { chosenVocabulary = vocabulary }}/>
            </div>
        {/each}
    </div>
{/if}

<style>
    .upper-section {
        position: sticky;
        top: 0;
        display: grid;
        grid-template-rows: 50px 44px 1fr;
        background-color: white;
        box-shadow: 0 5px 5px rgb(128, 128, 128, 0.5);
        z-index: 1;
    }

    .vocabularies {
        padding: 35px 20px;
        display: flex;
        flex-direction: column;
        row-gap: 30px;
    }
</style>