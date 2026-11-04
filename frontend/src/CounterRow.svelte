<script lang="ts">

interface Props {
    label: string
    actionIncrement: () => void
    actionDecrement: () => void
    actionAuto: (() => void) | null
    count: number
}

    let {
        label,
        actionIncrement,
        actionDecrement,
        actionAuto,
        count,
    }:Props = $props()

</script>

<div>
    <div>{label}</div>
    <div>
        <button onclick={actionDecrement}>-</button>
        <span>{count}</span>
        <button onclick={actionIncrement}>+</button>
    </div>
    {#if actionAuto}
        <button onclick={actionAuto}>Action automatique</button>
    {/if}
</div>

<style>

</style>