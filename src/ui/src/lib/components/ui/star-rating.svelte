<script lang="ts">
    import { Star } from "svelte-radix";
    import { createEventDispatcher } from "svelte";

    export let value: number = 0;
    let tempValue: number = 0;
    export let disabled: boolean = false;
    export let size: number = 24;
    export let color: string = "#FFD700"; // Gold color for stars

    const dispatch = createEventDispatcher<{
        change: { value: number };
    }>();

    function handleClick(index: number) {
        if (disabled) return;
        if (index + 1 === value) {
            value = 0;
        } else {
            value = index + 1;
        }
        dispatch("change", { value });
    }

    function handleMouseEnter(index: number) {
        if (disabled) return;
        tempValue = index + 1;
    }
</script>

<div class="flex gap-1">
    {#if value >= 0}
        {#each Array(5) as _, i}
            <button
                type="button"
                class="p-0.5 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary rounded-full"
                on:click={() => handleClick(i)}
                on:mouseenter={() => handleMouseEnter(i)}
                {disabled}
            >
                <Star
                    {size}
                    fill={i < value ? color : "none"}
                    stroke={i < value ? color : "#94a3b8"}
                    class="transition-colors duration-200"
                />
            </button>
        {/each}
    {/if}
</div>
