<script lang="ts">
    import * as Breadcrumb from "$lib/components/ui/breadcrumb";
    import { ChevronRight } from "lucide-svelte";
    import { t } from "$lib/i18n";
    import { page } from "$app/stores";
    import { ROUTES } from "$lib/routes/routes";
    import { goto } from "$app/navigation";
    type SegmentType = { name: string; href: string };

    $: segments = $page.url.pathname
        .split("/")
        .slice(2)
        .reduce<[string, (SegmentType | undefined)[]]>(
            (acc, segment) => {
                if (!segment) return acc;
                const [accPath, items] = acc;
                const newPath = `${accPath}${accPath.endsWith("/") ? "" : "/"}${segment}`;

                if (
                    acc[1].at(-1)?.name === "locations" &&
                    segment !== "create"
                ) {
                    return [newPath, [...items, undefined]];
                }

                return [newPath, [...items, { name: segment, href: newPath }]];
            },
            [ROUTES.SHOP, []],
        );
</script>

<Breadcrumb.Root>
    <Breadcrumb.List>
        <Breadcrumb.Item>
            <Breadcrumb.Link
                ><span class="cursor-pointer" on:click={() => goto(ROUTES.SHOP)}
                    >{$t("shop.home")}</span
                ></Breadcrumb.Link
            >
        </Breadcrumb.Item>

        {#each segments[1] as segment, i}
            {#if segment}
                <Breadcrumb.Separator>
                    <ChevronRight class="h-4 w-4" />
                </Breadcrumb.Separator>
                <Breadcrumb.Item>
                    {#if i === segments[1].length - 1}
                        <Breadcrumb.Page
                            >{$t(
                                `shop.breadcrumb.${segment.name}`,
                            )}</Breadcrumb.Page
                        >
                    {:else}
                        <Breadcrumb.Link>
                            <span
                                class="cursor-pointer"
                                on:click={() => goto(segment.href)}
                                >{$t(`shop.breadcrumb.${segment.name}`)}</span
                            ></Breadcrumb.Link
                        >
                    {/if}
                </Breadcrumb.Item>
            {/if}
        {/each}
    </Breadcrumb.List>
</Breadcrumb.Root>
