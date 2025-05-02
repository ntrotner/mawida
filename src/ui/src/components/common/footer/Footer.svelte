<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import { Separator } from "$lib/components/ui/separator";
    import { t } from "$lib/i18n";
    import { ROUTES } from "$lib/routes";
    import { goto } from "$app/navigation";
    export let footerItems: { key: string; action: () => void }[] = [];

    const commonLinks = [
        { key: "common.footer.imprint", href: ROUTES.PRIVACY },
        { key: "common.footer.guidlines", href: ROUTES.TERMS },
        { key: "common.footer.cookies", href: ROUTES.COOKIES },
        { key: "common.footer.settings", href: ROUTES.SETTINGS },
    ];
</script>

<footer class="w-full">
    <!-- Top Section - Customizable -->
    <div class="bg-muted/50 py-12">
        <div
            class="container flex flex-col md:flex-row justify-between items-center gap-8 mx-auto px-12"
        >
            <div class="flex flex-wrap justify-begin gap-8">
                {#each footerItems as item}
                    <Button on:click={() => item.action()} variant="link">
                        {$t(item.key)}
                    </Button>
                {/each}
            </div>

            <div class="max-w-md w-full">
                <div class="flex flex-col sm:flex-row gap-3">
                    <input
                        type="email"
                        placeholder={$t("common.footer.email-placeholder")}
                        class="px-4 py-2 rounded-md border border-input bg-background flex-grow"
                        aria-label={$t("common.footer.email-label")}
                    />
                    <Button variant="default" type="submit">
                        {$t("common.footer.subscribe")}
                    </Button>
                </div>
            </div>
        </div>
    </div>

    <Separator class="mx-auto" />

    <!-- Bottom Section - Common -->
    <div class="bg-muted/50 py-6">
        <div class="container mx-auto px-12">
            <div
                class="flex flex-col md:flex-row justify-between items-center gap-4"
            >
                <!-- Copyright -->
                <div class="text-sm text-muted-foreground">
                    © {new Date().getFullYear()}
                    {$t("common.footer.copyright")}
                </div>

                <!-- Links -->
                <div class="flex gap-6">
                    {#each commonLinks as link}
                        <Button on:click={() => goto(link.href)} variant="link" size="sm">
                            {$t(link.key)}
                        </Button>
                    {/each}
                </div>
            </div>
        </div>
    </div>
</footer>
