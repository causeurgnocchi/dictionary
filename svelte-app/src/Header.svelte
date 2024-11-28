<script>
    const {pageTitle} = $props();
</script>

<header>
    <h1 class="page-title">{pageTitle}</h1>
</header>

<style>
    *, *::before, *::after {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }
    
    header {
        display: flex;
    }

    .page-title {
        font-weight: normal;
        font-size: 24px;
        margin-left: 20px;
        align-self: center;
    }
</style>