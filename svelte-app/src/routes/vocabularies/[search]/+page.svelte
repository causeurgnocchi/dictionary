<script lang=="ts">
    import ChosenEntry from "../../../ChosenEntry.svelte";
    import Entry from "../../../Entry.svelte"

    let { data } = $props();
    let { vocabularies } = $derived(data);
</script>

<ChosenEntry vocabulary={vocabularies[0]} />
{#each vocabularies.slice(1) as v}
<Entry vocabulary={v} onclick={() => {}} />
{/each}
