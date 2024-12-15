<script lang="ts">
    import magnifyingGlass from '/src/static/magnifying-glass.svg';
</script>

<form action="/results">
    <input type="text" name="search" id="search">
    <button type="submit" class="search">
        <img src={magnifyingGlass} alt="search">
    </button>
</form>

<style>
    *, *::before, *::after {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }
    
    form {
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