<script lang="ts">
    import { onMount } from 'svelte';
    import * as Counter from '../wailsjs/go/main/Counter.js'
    import { main as models } from '../wailsjs/go/models.js';
    import CounterRow from './CounterRow.svelte';

    let counter = $state<models.Counter | null>(null)

    interface CounterRow {
        label: string
        actionIncrement: () => Promise<models.Counter>
        actionDecrement: () => Promise<models.Counter>
        actionAuto: (() => Promise<models.Counter>) | null
        property: keyof models.Counter
    }

    const counterRows:CounterRow[] = [
        {
            label: "Attempts",
            actionIncrement: Counter.IncrementAttempts,
            actionDecrement: Counter.DecrementAttempts,
            actionAuto: Counter.AutoFail,
            property: "attempts"
        },
        {
            label: "Success",
            actionIncrement: Counter.IncrementSuccess,
            actionDecrement: Counter.DecrementSuccess,
            actionAuto: Counter.AutoSuccess,
            property: "success"
        },
        {
            label: "Success (Perfect)",
            actionIncrement: Counter.IncrementSuccessPerfect,
            actionDecrement: Counter.DecrementSuccessPerfect,
            actionAuto: Counter.AutoSuccessPerfect,
            property: "successPerfect"
        },
        {
            label: "Current Streak",
            actionIncrement: Counter.IncrementCurrentStreak,
            actionDecrement: Counter.DecrementCurrentStreak,
            actionAuto: null,
            property: "currentStreak"
        },
        {
            label: "Best Streak",
            actionIncrement: Counter.IncrementBestStreak,
            actionDecrement: Counter.DecrementBestStreak,
            actionAuto: null,
            property: "bestStreak"
        },
    ]

    onMount(async () => {
        try{
            counter = await Counter.GetCounter()
        }catch {
            alert("Une erreur est survenue lors de l'initialisation de l'application")
        }

    })

    async function handleAction(action: () => Promise<models.Counter>) {
        try{
            counter = await action()
        }catch{
            alert('Une erreur est survenue')
        }
    }
</script>

<main>
    {#if counter}
        {#each counterRows as row }
            <CounterRow label={row.label}
                actionIncrement={() => handleAction(row.actionIncrement)}
                actionDecrement={() => handleAction(row.actionDecrement)}
                actionAuto={row.actionAuto ? () => handleAction(row.actionAuto!) : null}
                count={counter[row.property]}
                />
        {/each}
    {/if}
</main>
