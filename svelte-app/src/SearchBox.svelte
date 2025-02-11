<script lang="ts">
    import searchIcon from '/src/static/search-icon.svg';
</script>

<div>
    <input type="text" />
    <button type="submit">
        <img src={searchIcon} alt="search icon">
    </button>
</div>

<style>
    div {
        display: grid;
        grid-template-rows: 24px;
        grid-template-columns: 200px 24px;
        grid-row: 2;
        align-self: center;
        justify-self: center;
        border: 1px solid black;
    }

    input {
        border: none;
        font-size: 15px;
    }

    button {
        display: flex;
        border: none;
        border-left: 1px solid black;
    }

    button > img {
        background-color: white;
        flex-grow: 1;
    }
</style>