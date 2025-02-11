<div class="container">
  <input class="input" type="text" /><button
    class="search"
    aria-label="search icon"
  ></button>
</div>

<style>
  .container {
    margin: var(--margin, 0);
    width: fit-content;
  }

  .input {
    width: 200px;
    height: 25px;
    border: 1px solid black;
    vertical-align: top;
  }

  .search {
    width: 25px;
    height: 25px;
    border: 1px solid black;
    position: relative;
    left: -1px;
    background-image: url("src/static/search-icon.svg");
    background-repeat: no-repeat;
    background-position: center;
  }
</style>
