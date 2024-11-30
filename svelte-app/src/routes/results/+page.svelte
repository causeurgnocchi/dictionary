<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import Vocabulary from '../../Vocabulary.svelte';
    import StagingArea from '../../StagingArea.svelte';
    import Header from '../../Header.svelte';
    import SearchBar from '../../SearchBar.svelte';

    const search = $page.url.searchParams.get('search');
    let vocabularies = $state([]);

    onMount(async () => {
        fetch(`http://[::1]:8080/api/vocabularies/${search}`)
        .then(response => response.json())
        .then(data => {
            vocabularies = data;
        })
    });
</script>

<div class="container">
    <Header pageTitle="Dictionary"/>
    <SearchBar />
    <div class="vocabularies">
        {#each vocabularies as vocabulary}
            <StagingArea {vocabulary} />
            <Vocabulary {vocabulary} />
        {/each}
    </div>
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