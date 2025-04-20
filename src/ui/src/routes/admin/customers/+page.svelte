<script lang="ts">
    import { t } from "$lib/i18n";
    import { onMount } from "svelte";
    import { usersState, fetchUsers } from "$lib/states/users";
    import { Button } from "$lib/components/ui/button/index.js";
    import * as Table from "$lib/components/ui/table/index.js";
    import { HamburgerMenu, EnvelopeClosed, File } from "svelte-radix";
    import { productState, fetchProducts } from "$lib/states/product";
    import { map, firstValueFrom, filter } from "rxjs";
    import { Skeleton } from "$lib/components/ui/skeleton";
    import { writable } from "svelte/store";

    let users = usersState.getAsyncState();
    let loading = writable(true);

    // Pagination
    let currentPage = 1;
    let itemsPerPage = 10;

    $: totalPages = Math.ceil(($users?.length || 0) / itemsPerPage);
    $: paginatedUsers = $users ? 
        $users.slice((currentPage - 1) * itemsPerPage, currentPage * itemsPerPage) 
        : [];
    
    function prevPage() {
        if (currentPage > 1) currentPage--;
    }
    
    function nextPage() {
        if (currentPage < totalPages) currentPage++;
    }

    onMount(async () => {
        try {
            await Promise.allSettled([
                await fetchUsers(),
                await fetchProducts(),
                new Promise((resolve) => setTimeout(resolve, 150)),
            ]);
        } catch {
        } finally {
            loading.set(false);
        }
    });

    const products = productState.observable();

    function productsForCustomer(customerId: string) {
        return firstValueFrom(
            products.pipe(
                filter((productsList) => !!productsList),
                map((productsList) =>
                    productsList.filter(
                        (product) => product.renterInfo?.userId === customerId,
                    ),
                ),
            ),
        );
    }
</script>

<div class="mx-auto">
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

    {#if $loading}
        <div class="flex justify-center items-center h-full">
            <Skeleton class="w-full h-[30rem]" />
        </div>
    {:else if $users && $users.length > 0}
        <div class="rounded-md border">
            <Table.Root>
                <Table.Header>
                    <h1 class="text-2xl font-bold px-4 pt-4">
                        {$t("admin.sidebar.customers")}
                    </h1>
                    <p class="text-sm text-gray-500 px-4 pt-3 pb-4">
                        {$t("admin.customers.description")}
                    </p>
                    <Table.Row>
                        <Table.Head
                            >{$t("admin.customers.table.email")}</Table.Head
                        >
                        <Table.Head
                        >{$t("admin.customers.table.products")}</Table.Head
                        >
                        <Table.Head
                            >{$t("admin.customers.table.role")}</Table.Head
                        >
                        <Table.Head class="text-right"
                            >{$t("admin.customers.table.actions")}</Table.Head
                        >
                    </Table.Row>
                </Table.Header>
                <Table.Body>
                    {#each paginatedUsers as customer}
                        <Table.Row>
                            <Table.Cell>{customer.email || "-"}</Table.Cell>
                            <Table.Cell>
                                <ul>
                                    {#await productsForCustomer(customer.id || "")}
                                        <li>Loading...</li>
                                    {:then products}
                                        {#each products as product}
                                            <li>{product.name}</li>
                                        {/each}
                                    {/await}
                                </ul>
                            </Table.Cell>
                            <Table.Cell>
                                <span
                                    class="inline-flex items-center px-2.5 py-0.5"
                                >
                                    {$t(
                                        `admin.customers.table.role.${customer.role}`,
                                    )}
                                </span>
                            </Table.Cell>
                            <Table.Cell class="flex flex-col items-end">
                                    <EnvelopeClosed class="h-4 w-4 mr-4" />
                            </Table.Cell>
                        </Table.Row>
                    {/each}
                </Table.Body>
            </Table.Root>
            
            <!-- Simple Pagination -->
            <div class="flex items-center justify-between px-4 py-3 border-t">
                <div class="text-sm text-gray-700">
                    {$t("common.pagination.showing")} <span class="font-medium">{((currentPage - 1) * itemsPerPage) + 1}</span> {$t("common.pagination.to")} 
                    <span class="font-medium">{Math.min(currentPage * itemsPerPage, $users?.length || 0)}</span> {$t("common.pagination.of")} 
                    <span class="font-medium">{$users?.length || 0}</span> {$t("common.pagination.entries")}
                </div>
                
                <div class="flex gap-2">
                    <Button variant="outline" size="sm" disabled={currentPage === 1} on:click={prevPage}>
                        {$t("common.pagination.previous")}
                    </Button>
                    
                    <div class="flex items-center">
                        <span class="mx-2 text-sm">{$t("common.pagination.page")} {currentPage} {$t("common.pagination.of")} {totalPages || 1}</span>
                    </div>
                    
                    <Button variant="outline" size="sm" disabled={currentPage === totalPages || totalPages === 0} on:click={nextPage}>
                        {$t("common.pagination.next")}
                    </Button>
                </div>
            </div>
        </div>
    {:else}
        <div
            class="flex flex-col items-center justify-center p-8 border rounded-lg"
        >
            <p class="text-gray-500 mb-4">{$t("admin.customers.empty")}</p>
        </div>
    {/if}
</div>
