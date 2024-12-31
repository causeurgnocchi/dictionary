<script lang="ts">
    import Header from "../Header.svelte";
    import SearchBar from "../SearchBar.svelte";
</script>

<div class="container">
    <Header pageTitle="Dictionary"/>
    <SearchBar />
</div>

<style>
    .container {
        height: 100vh;
        display: grid;
        grid-template-rows: 50px 44px;
    }
</style>
