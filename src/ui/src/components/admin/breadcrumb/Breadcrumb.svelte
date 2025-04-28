<script lang="ts">
    import * as Breadcrumb from "$lib/components/ui/breadcrumb";
    import { ChevronRight } from "lucide-svelte";
    import { t } from "$lib/i18n";
    import { page } from "$app/stores";
    import { ROUTES } from "$lib/routes/routes";
    import { goto } from "$app/navigation";
    import { usersState } from "$lib/states/users";
    import { locationState } from "$lib/states/location";
    import { derived } from "svelte/store";

    type SegmentType = { name: string; href: string };

    function getEmailFromId(id: string) {
        const user = usersState.getSyncState()?.find((user) => user.id === id);
        return user?.email;
    }

    function getLocationNameFromId(id: string) {
        const location = locationState
            .getSyncState()
            ?.find((location) => location.id === id);
        return location?.city;
    }

    $: segments = derived(
        [page, locationState.getAsyncState()],
        ([path, locations]) =>
            path.url.pathname
                .split("/")
                .slice(2)
                .reduce<[string, (SegmentType | undefined)[]]>(
                    (acc, segment) => {
                        if (!segment) return acc;
                        const [accPath, items] = acc;
                        const newPath = `${accPath}${accPath.endsWith("/") ? "" : "/"}${segment}`;

                        if (
                            acc[1].at(-1)?.name ===
                                "admin.breadcrumb.locations" &&
                            segment !== "create"
                        ) {
                            return [
                                newPath,
                                [
                                    ...items,
                                    {
                                        name:
                                            getLocationNameFromId(segment) ??
                                            "",
                                        href: acc[1].at(-1)?.href || "",
                                    },
                                ],
                            ];
                        }

                        if (
                            acc[1].at(-1)?.name === "admin.breadcrumb.customers"
                        ) {
                            return [
                                newPath,
                                [
                                    ...items,
                                    {
                                        name: getEmailFromId(segment) ?? "",
                                        href: newPath,
                                    },
                                ],
                            ];
                        }

                        return [
                            newPath,
                            [
                                ...items,
                                {
                                    name: `admin.breadcrumb.${segment}`,
                                    href: newPath,
                                },
                            ],
                        ];
                    },
                    [ROUTES.ADMIN, []],
                ) || [],
    );
</script>

<Breadcrumb.Root>
    <Breadcrumb.List>
        <Breadcrumb.Item>
            <Breadcrumb.Link
                ><span
                    class="cursor-pointer"
                    on:click={() => goto(ROUTES.ADMIN)}>{$t("admin.home")}</span
                ></Breadcrumb.Link
            >
        </Breadcrumb.Item>

        {#each $segments[1] as segment, i}
            {#if segment}
                <Breadcrumb.Separator>
                    <ChevronRight class="h-4 w-4" />
                </Breadcrumb.Separator>
                <Breadcrumb.Item>
                    {#if i === $segments[1].length - 1}
                        <Breadcrumb.Page>
                            <span
                                class="cursor-pointer"
                                on:click={() => goto(segment.href)}
                                >{$t(segment.name)}</span
                            ></Breadcrumb.Page
                        >
                    {:else}
                        <Breadcrumb.Link
                            ><span
                                class="cursor-pointer"
                                on:click={() => goto(segment.href)}
                                >{$t(segment.name)}</span
                            ></Breadcrumb.Link
                        >
                    {/if}
                </Breadcrumb.Item>
            {/if}
        {/each}
    </Breadcrumb.List>
</Breadcrumb.Root>
