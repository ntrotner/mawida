<script lang="ts">
    import { Button } from "$lib/components/ui/button/index.js";
    import { goto } from "$app/navigation";
    import { t } from "$lib/i18n";
    import * as Tooltip from "$lib/components/ui/tooltip/index.js";

    import Person from "svelte-radix/Person.svelte";
    import { ROUTES } from "$lib/routes/routes";
    import { page } from "$app/state";
    import { ShoppingCart } from "lucide-svelte";
    import { authenticationState } from "$lib/states/authentication";

    const authState = authenticationState.getAsyncState();
    const isAdmin = authenticationState.isAdmin();

    function navigate(route: string) {
        goto(route);
    }
</script>

<aside
    class="fixed left-0 top-0 h-full w-[3.5rem] bg-background border-r border-border flex flex-col items-center py-2 z-50"
>
    <!-- Company Logo at the top -->
    <div class="mb-2 p-2">
        <Button
            variant="ghost"
            size="icon"
            class="h-10 w-10 rounded-md hover:bg-muted"
            on:click={() => navigate(ROUTES.HOME)}
        >
            <div
                class="home-button h-5 w-5 bg-primary rounded-xl flex items-center justify-center text-primary-foreground font-bold"
            >
                M
            </div>
        </Button>
    </div>

    <!-- Navigation Icons -->
    <nav class="flex flex-col gap-1 items-center">
        <Tooltip.Root>
            <Tooltip.Trigger asChild>
                <Button
                    variant="ghost"
                    size="icon"
                    class="h-10 w-10 rounded-md hover:bg-muted"
                    on:click={() => navigate(ROUTES.SHOP)}
                >
                    <ShoppingCart class="h-5 w-5 text-gray-500" />
                    <span class="sr-only"
                        >{$t("common.sidebar.cart")}</span
                    >
                </Button>
            </Tooltip.Trigger>
            <Tooltip.Content side="right">
                <p>{$t("common.sidebar.products")}</p>
            </Tooltip.Content>
        </Tooltip.Root>

        {#if $authState?.authenticated && !$isAdmin}
            <Tooltip.Root>
                <Tooltip.Trigger asChild>
                    <Button
                        variant="ghost"
                        size="icon"
                        class="h-10 w-10 rounded-md hover:bg-muted"
                        on:click={() => navigate(ROUTES.CUSTOMER_PROFILE)}
                >
                    <Person class="h-5 w-5 text-gray-500" />
                    <span class="sr-only">{$t("sidebar.customer-dashboard")}</span>
                </Button>
            </Tooltip.Trigger>
            <Tooltip.Content side="right">
                <p>{$t("sidebar.customer-dashboard")}</p>
            </Tooltip.Content>
        </Tooltip.Root>
        {/if}
    </nav>
</aside>
