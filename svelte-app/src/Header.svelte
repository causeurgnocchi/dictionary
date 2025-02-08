<script>
    let {pageTitle} = $props();
</script>

<h1 class="page-title"><a href="/">{pageTitle}</a></h1>

<style>
    *, *::before, *::after {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }

    .page-title {
        margin: 15px 0 0 50px;
        font-size: 24px;
        font-weight: normal;
    }

    a {
        text-decoration: none;
    }
</style>