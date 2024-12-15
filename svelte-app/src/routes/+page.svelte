<script>
    import Header from "../Header.svelte";
    import SearchBar from "../SearchBar.svelte";
</script>

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
</style>

<svelte:head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dictionary</title>
</svelte:head>
<div class="container">
    <Header pageTitle="Dictionary"/>
    <SearchBar />
</div>