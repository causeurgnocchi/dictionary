<script lang="ts">
  import Header from "../Header.svelte";
  import SearchBox from "../SearchBox.svelte";

  let { children } = $props();
</script>

<svelte:head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Dictionary</title>
</svelte:head>
<div class="container">
  <Header />
  <SearchBox --margin="auto" />
  {@render children()}
</div>

<style>
  .container {
    margin: auto;
    max-width: 1200px;
  }
</style>
