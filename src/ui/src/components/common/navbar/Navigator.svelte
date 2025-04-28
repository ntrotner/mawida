<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import { userState } from "$lib/states/user";
    import { authenticationState } from "$lib/states/authentication";
    import { t } from "$lib/i18n";
    import { ROUTES } from "$lib/routes";
    import { Avatar, AvatarFallback } from "$lib/components/ui/avatar";
    import { goto } from "$app/navigation";
    export let navigationItems: { key: string; action: () => void }[] = [];

    const isAdmin = userState.isAdmin();
    const authStatus = authenticationState.getAsyncState();
</script>

<nav class="flex items-center justify-between px-10 py-4 bg-white shadow-sm">
    <!-- Company Logo -->
    <div class="flex-shrink-0">
        <Avatar>
            <AvatarFallback>
                MWD
            </AvatarFallback>
        </Avatar>
    </div>

    <!-- Right Side Group -->
    <div class="flex items-center space-x-8">
        <!-- Navigation Items -->
        <div class="flex items-center space-x-2">
            {#each navigationItems as item}
                <Button on:click={() => item.action()} variant="link">
                    {$t(item.key)}
                </Button>
            {/each}
        </div>

        <!-- Shopping Cart Button -->
        <Button on:click={() => goto(ROUTES.SHOP)} variant="outline">
            {$t('common.navigation.explore-products')}
        </Button>

        <!-- Login/Admin Button -->
        {#if !$authStatus?.authenticated}
            <Button on:click={() => goto(ROUTES.LOGIN)} variant="default">{$t('common.nav-menu.login')}</Button>
        {:else if $isAdmin}
            <Button on:click={() => goto(ROUTES.ADMIN)} variant="default">{$t('common.navigation.admin')}</Button>
        {/if}
    </div>
</nav>

<style>
    nav {
        width: 100%;
    }
</style>
