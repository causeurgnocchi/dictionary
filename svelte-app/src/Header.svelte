<script>
    let {pageTitle} = $props();
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
        align-self: center;
        margin-left: 20px;
        font-size: 24px;
        font-weight: normal;
    }
</style>