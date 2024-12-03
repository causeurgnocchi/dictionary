<script lang="ts">
    const {vocabulary} : {vocabulary: Vocabulary} = $props();
    const kanjiSize = 30;
    const hasFurigana = vocabulary.characters.some((c: Character) => c.furigana !== '');
</script>

<div class={hasFurigana ? "has-furigana" : "not-has-furigana"} style:--kanji-size={`${kanjiSize}px`}>
    {#if hasFurigana}
        <p class="furigana">
            {#each vocabulary.characters as character}
                <span style={`font-size: ${kanjiSize / (character.furigana.length === 1 ? 2 : character.furigana.length)}px`}>{character.furigana}</span>
            {/each}
        </p>
    {/if}
    <p class="writing">
        {vocabulary.writing}
    </p>
</div>

<style>
    *, *::before, *::after {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }
    
    .has-furigana {
        display: grid;
        grid-template-rows: 18px 35px;
    }

    .not-has-furigana {
        display: grid;
        grid-template-rows: 35px;
    }

    .furigana {
        display: flex;
        align-items: center;
    }

    .furigana > span {
        width: var(--kanji-size);
        text-align: center;
        max-height: 18px;
    }

    .writing {
        font-size: var(--kanji-size);
    }
</style>