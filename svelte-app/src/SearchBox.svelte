<script lang="ts">
  import SearchIcon from "/src/static/search-icon.svg";
  let search = $state("");
</script>

<div class="container">
  <input class="input" type="text" bind:value={search} /><a
    class="search-link"
    href={`/search/${search}`}
  >
    <img class="search" src={SearchIcon} alt="search icon" />
  </a>
</div>

<style>
  .container {
    margin: var(--margin, 0);
    width: fit-content;
  }

  .input {
    box-sizing: border-box;
    width: 200px;
    height: 25px;
    border: 1px solid black;
    vertical-align: top;
  }

  .search-link {
    display: inline-block;
    width: 25px;
    height: 25px;
  }

  .search {
    box-sizing: border-box;
    width: 25px;
    height: 25px;
    position: relative;
    left: -1px;
    border: 1px solid black;
  }
</style>
