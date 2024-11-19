<script>
    import Header from "../Header.svelte";
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

<div class="container">
    <Header />
</div>