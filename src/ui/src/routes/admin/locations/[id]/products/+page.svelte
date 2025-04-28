<script lang="ts">
    import { t } from "$lib/i18n";
    import { onMount } from "svelte";
    import { productState, fetchProducts } from "$lib/states/product";
    import { Button } from "$lib/components/ui/button/index.js";
    import * as Table from "$lib/components/ui/table/index.js";
    import { page } from "$app/stores";
    import { PlusCircled, HamburgerMenu, File } from "svelte-radix";
    import { formatCurrency } from "$lib/helpers/price";
    import { ROUTES } from "$lib/routes";
    import { goto } from "$app/navigation";
    import { writable } from "svelte/store";
    import { Skeleton } from "$lib/components/ui/skeleton";
    import * as ToggleGroup from "$lib/components/ui/toggle-group/index.js";

    const locationId = $page.params.id;
    let filterStatus = writable("all");
    let products = productState.getAsyncState();
    let filteredProducts = writable<any[]>([]);
    let loading = writable(true);
    // Pagination
    let currentPage = 1;
    let itemsPerPage = 10;

    $: totalPages = Math.ceil($filteredProducts.length / itemsPerPage);
    $: paginatedProducts = $filteredProducts.slice(
        (currentPage - 1) * itemsPerPage,
        currentPage * itemsPerPage,
    );

    function prevPage() {
        if (currentPage > 1) currentPage--;
    }

    function nextPage() {
        if (currentPage < totalPages) currentPage++;
    }

    function handleFilterStatus(value: string) {
        if (value === "available") {
            filteredProducts.set(
                $products?.filter(
                    (product) =>
                        product.location === locationId && !product.isRented,
                ) || [],
            );
        } else if (value === "rented") {
            filteredProducts.set(
                $products?.filter(
                    (product) =>
                        product.location === locationId && !!product.isRented,
                ) || [],
            );
        } else {
            filteredProducts.set(
                $products?.filter(
                    (product) => product.location === locationId,
                ) || [],
            );
        }

        filterStatus.set(value);
        currentPage = 1;
    }

    $: if ($products) {
        handleFilterStatus("all");
    }

    onMount(async () => {
        try {
            await Promise.allSettled([
                await fetchProducts(),
                new Promise((resolve) => setTimeout(resolve, 150)),
            ]);
            handleFilterStatus("all");
        } catch {
        } finally {
            loading.set(false);
        }
    });
</script>

<div class="mx-auto">
    <div class="flex justify-between mb-3 gap-3">
        <div class="flex items-start space-x-2">
            <ToggleGroup.Root
                type="single"
                value={$filterStatus}
                onValueChange={(value) => value && handleFilterStatus(value)}
                variant="outline"
                class="rounded-md bg-gray-100 p-1"
            >
                <ToggleGroup.Item
                    class="border-none rounded-md data-[state=on]:bg-white"
                    value="all">{$t("admin.statistics.all")}</ToggleGroup.Item
                >
                <ToggleGroup.Item
                    class="border-none rounded-md data-[state=on]:bg-white"
                    value="available"
                    >{$t("admin.statistics.available")}</ToggleGroup.Item
                >
                <ToggleGroup.Item
                    class="border-none rounded-md data-[state=on]:bg-white"
                    value="rented"
                    >{$t("admin.statistics.rented")}</ToggleGroup.Item
                >
            </ToggleGroup.Root>
        </div>

        <div class="flex justify-end mb-3 gap-3">
            <Button variant="outline">
                <HamburgerMenu class="h-4 w-4 mr-2" />
                {$t("common.filter")}
            </Button>
            <Button variant="outline">
                <File class="h-4 w-4 mr-2" />
                {$t("common.export")}
            </Button>
            <Button
                variant="default"
                on:click={() => {
                    goto(
                        ROUTES.ADMIN_PRODUCT_CREATE.replace(
                            "{locationId}",
                            locationId,
                        ),
                    );
                }}
            >
                <PlusCircled class="h-4 w-4 mr-2" />
                {$t("admin.products.add")}
            </Button>
        </div>
    </div>

    {#if $loading}
        <div class="flex justify-center items-center h-full">
            <Skeleton class="w-full h-[30rem]" />
        </div>
    {:else if $filteredProducts && $filteredProducts.length > 0}
        <div class="rounded-md border">
            <Table.Root>
                <Table.Header>
                    <h1 class="text-2xl font-bold px-4 pt-4">
                        {$t("admin.sidebar.products")}
                    </h1>
                    <p class="text-sm text-gray-500 px-4 pt-3 pb-4">
                        {$t("admin.products.table.description")}
                    </p>
                    <Table.Row>
                        <Table.Head
                            >{$t("admin.products.table.name")}</Table.Head
                        >
                        <Table.Head
                            >{$t("admin.products.table.status")}</Table.Head
                        >
                        <Table.Head
                            >{$t("admin.products.table.price")}</Table.Head
                        >
                        <Table.Head
                            >{$t("admin.products.table.deposit")}</Table.Head
                        >
                    </Table.Row>
                </Table.Header>
                <Table.Body>
                    {#each paginatedProducts as product}
                        <Table.Row>
                            <Table.Cell>{product.name || "-"}</Table.Cell>
                            <Table.Cell>
                                <span class="px-2 py-1 rounded-md border">
                                    {#if product.renterInfo?.userId && product.renterInfo.rentalEndDate < (Date.now() / 1000)}
                                        {$t("admin.products.table.late")}
                                    {:else if product.renterInfo?.userId}
                                        {$t("admin.products.table.rented")}
                                    {:else}
                                        {$t("admin.products.table.available")}
                                    {/if}
                                </span>
                            </Table.Cell>
                            <Table.Cell
                                >{formatCurrency(
                                    product.pricing?.price || 0,
                                    "€",
                                )}</Table.Cell
                            >
                            <Table.Cell
                                >{formatCurrency(
                                    product.pricing?.deposit || 0,
                                    "€",
                                )}</Table.Cell
                            >
                        </Table.Row>
                    {/each}
                </Table.Body>
            </Table.Root>

            <!-- Simple Pagination -->
            <div class="flex items-center justify-between px-4 py-3 border-t">
                <div class="text-sm text-gray-700">
                    {$t("common.pagination.showing")}
                    <span class="font-medium"
                        >{(currentPage - 1) * itemsPerPage + 1}</span
                    >
                    {$t("common.pagination.to")}
                    <span class="font-medium"
                        >{Math.min(
                            currentPage * itemsPerPage,
                            $filteredProducts.length,
                        )}</span
                    >
                    {$t("common.pagination.of")}
                    <span class="font-medium">{$filteredProducts.length}</span>
                    {$t("common.pagination.entries")}
                </div>

                <div class="flex gap-2">
                    <Button
                        variant="outline"
                        size="sm"
                        disabled={currentPage === 1}
                        on:click={prevPage}
                    >
                        {$t("common.pagination.previous")}
                    </Button>

                    <div class="flex items-center">
                        <span class="mx-2 text-sm"
                            >{$t("common.pagination.page")}
                            {currentPage}
                            {$t("common.pagination.of")}
                            {totalPages || 1}</span
                        >
                    </div>

                    <Button
                        variant="outline"
                        size="sm"
                        disabled={currentPage === totalPages ||
                            totalPages === 0}
                        on:click={nextPage}
                    >
                        {$t("common.pagination.next")}
                    </Button>
                </div>
            </div>
        </div>
    {:else}
        <div
            class="flex flex-col items-center justify-center p-8 border rounded-lg"
        >
            <p class="text-gray-500 mb-4">{$t("admin.products.empty")}</p>
        </div>
    {/if}
</div>
