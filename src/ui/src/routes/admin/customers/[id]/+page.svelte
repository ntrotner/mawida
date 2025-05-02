<script lang="ts">
    import { t } from "$lib/i18n";
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import { productState, fetchProducts } from "$lib/states/product";
    import { rentContractsState } from "$lib/states/rent_contracts";
    import { fetchRentContracts } from "$lib/states/rent_contracts/effects";
    import { Button } from "$lib/components/ui/button/index.js";
    import * as Table from "$lib/components/ui/table/index.js";
    import { PlusCircled, HamburgerMenu, File } from "svelte-radix";
    import { formatCurrency } from "$lib/helpers/price";
    import { derived, writable } from "svelte/store";
    import { Skeleton } from "$lib/components/ui/skeleton";
    import * as ToggleGroup from "$lib/components/ui/toggle-group/index.js";
    import { fetchUsers, usersState } from "$lib/states/users";

    const customerId = $page.params.id;
    let filterStatus = writable("all");
    let products = productState.getAsyncState();
    let rentContracts = rentContractsState.getAsyncState();
    let filteredContracts = writable<any[]>([]);
    let customer = derived(usersState.getAsyncState(), ($users) =>
        $users?.find((c) => c.id === customerId),
    );
    let loading = writable(true);

    // Pagination
    let currentPage = 1;
    let itemsPerPage = 10;

    $: totalPages = Math.ceil($filteredContracts.length / itemsPerPage);
    $: paginatedContracts = $filteredContracts.slice(
        (currentPage - 1) * itemsPerPage,
        currentPage * itemsPerPage,
    );

    function prevPage() {
        if (currentPage > 1) currentPage--;
    }

    function nextPage() {
        if (currentPage < totalPages) currentPage++;
    }

    function formatDate(timestamp: number): string {
        if (!timestamp) return "-";
        const date = new Date(timestamp * 1000);
        return date.toLocaleDateString();
    }

    function handleFilterStatus(value: string) {
        if (value === "active") {
            filteredContracts.set(
                $rentContracts?.filter(
                    (contract) =>
                        contract.userId === customerId &&
                        (contract.status === "pickupPending" ||
                            contract.status === "confirmed" ||
                            contract.status === "active"),
                ) || [],
            );
        } else if (value === "completed") {
            filteredContracts.set(
                $rentContracts?.filter(
                    (contract) =>
                        contract.userId === customerId &&
                        contract.status === "completed",
                ) || [],
            );
        } else if (value === "cancelled") {
            filteredContracts.set(
                $rentContracts?.filter(
                    (contract) =>
                        contract.userId === customerId &&
                        contract.status === "cancelled",
                ) || [],
            );
        } else {
            filteredContracts.set(
                $rentContracts?.filter(
                    (contract) => contract.userId === customerId,
                ) || [],
            );
        }

        filterStatus.set(value);
        currentPage = 1;
    }

    // Get product name by its ID
    function getProductName(productId: string): string {
        const product = $products?.find((p) => p.id === productId);
        return product?.name || "-";
    }

    $: if ($rentContracts) {
        handleFilterStatus("all");
    }

    onMount(async () => {
        try {
            await Promise.allSettled([
                await fetchProducts(),
                await fetchRentContracts(),
                await fetchUsers(),
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
                    value="active"
                    >{$t(
                        "admin.rent_contracts.status.active",
                    )}</ToggleGroup.Item
                >
                <ToggleGroup.Item
                    class="border-none rounded-md data-[state=on]:bg-white"
                    value="completed"
                    >{$t(
                        "admin.rent_contracts.status.completed",
                    )}</ToggleGroup.Item
                >
                <ToggleGroup.Item
                    class="border-none rounded-md data-[state=on]:bg-white"
                    value="cancelled"
                    >{$t(
                        "admin.rent_contracts.status.cancelled",
                    )}</ToggleGroup.Item
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
        </div>
    </div>

    {#if $loading}
        <div class="flex justify-center items-center h-full">
            <Skeleton class="w-full h-[30rem]" />
        </div>
    {:else if $filteredContracts && $filteredContracts.length > 0}
        <div class="rounded-md border">
            <Table.Root>
                <Table.Header>
                    <h1 class="text-2xl font-bold px-4 pt-4">
                        {$customer?.email}
                    </h1>
                    <p class="text-sm text-gray-500 px-4 pt-3 pb-4">
                        {$t("admin.customer_products.description")}
                    </p>
                    <Table.Row>
                        <Table.Head
                            >{$t("admin.products.table.name")}</Table.Head
                        >
                        <Table.Head
                            >{$t(
                                "admin.rent_contracts.table.status",
                            )}</Table.Head
                        >
                        <Table.Head
                            >{$t(
                                "admin.rent_contracts.table.price",
                            )}</Table.Head
                        >
                        <Table.Head
                            >{$t(
                                "admin.rent_contracts.table.deposit",
                            )}</Table.Head
                        >
                        <Table.Head
                            >{$t(
                                "admin.rent_contracts.table.endDate",
                            )}</Table.Head
                        >
                    </Table.Row>
                </Table.Header>
                <Table.Body>
                    {#each paginatedContracts as contract}
                        <Table.Row>
                            <Table.Cell
                                >{getProductName(
                                    contract.productId || "",
                                )}</Table.Cell
                            >
                            <Table.Cell>
                                <span class="px-2 py-1 rounded-md border">
                                    {$t(
                                        `admin.rent_contracts.status.${contract.status}`,
                                    )}
                                </span>
                            </Table.Cell>
                            <Table.Cell
                                >{formatCurrency(
                                    contract.price || 0,
                                    "€",
                                )}</Table.Cell
                            >
                            <Table.Cell
                                >{formatCurrency(
                                    contract.deposit || 0,
                                    "€",
                                )}</Table.Cell
                            >
                            <Table.Cell
                                >{formatDate(
                                    contract.rentalEndDate,
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
                            $filteredContracts.length,
                        )}</span
                    >
                    {$t("common.pagination.of")}
                    <span class="font-medium">{$filteredContracts.length}</span>
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
            <p class="text-gray-500 mb-4">
                {$t("admin.customer_products.empty")}
            </p>
        </div>
    {/if}
</div>
