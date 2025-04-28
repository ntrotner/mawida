<script lang="ts">
    import { onMount } from "svelte";
    import { t } from "$lib/i18n";
    import * as Card from "$lib/components/ui/card/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import * as ToggleGroup from "$lib/components/ui/toggle-group";
    import { Download, File, DotsHorizontal } from "svelte-radix";
    import { userState } from "$lib/states/user";
    import { formatCurrency } from "$lib/helpers/price";
    import { goto } from "$app/navigation";
    import { ROUTES } from "$lib/routes";
    import { rentContractsState } from "$lib/states/rent_contracts";
    import {
        fetchRentContracts,
        cancelRentContract,
    } from "$lib/states/rent_contracts/effects";
    import { fetchProduct, productState } from "$lib/states/product";
    import { locationState } from "$lib/states/location";
    import type { RentContract, RentContractStatus } from "$lib/open-api";
    import * as Table from "$lib/components/ui/table";
    import * as Menubar from "$lib/components/ui/menubar";

    // Get user data
    const user = userState.getAsyncState();
    const rentContracts = rentContractsState.getAsyncState();
    const products = productState.getAsyncState();
    const locations = locationState.getAsyncState();

    // Load rental contracts on mount
    onMount(async () => {
        await fetchRentContracts();
    });

    // Filter contracts by status
    $: currentRentals =
        $rentContracts?.filter(
            (contract) =>
                contract.status === "pickupPending" ||
                contract.status === "confirmed" ||
                contract.status === "active",
        ) || [];

    $: pastRentals =
        $rentContracts?.filter(
            (contract) =>
                contract.status === "completed" ||
                contract.status === "cancelled",
        ) || [];

    // Format date to local format
    function formatDate(timestamp: number | undefined): string {
        if (!timestamp) return "-";
        return new Date(timestamp * 1000).toLocaleDateString();
    }

    // Calculate rental duration in days
    function getRentalDuration(
        startDate: number | undefined,
        endDate: number | undefined,
    ): number {
        if (!startDate || !endDate) return 0;
        const diffTime = Math.abs((endDate - startDate) * 1000);
        return Math.ceil(diffTime / (1000 * 60 * 60 * 24));
    }

    // Get product name from product ID
    function getProductName(productId: string | undefined): string {
        if (!productId || !$products) return "-";
        const product = $products.find((p) => p.id === productId);
        return product?.name || "-";
    }

    // Get product image from product ID
    function getProductImage(productId: string | undefined): string {
        if (!productId || !$products) return "https://via.placeholder.com/150";
        const product = $products.find((p) => p.id === productId);
        if (product?.images && product.images.length > 0) {
            return `data:image/*;base64,${product.images[0].data}`;
        }
        return "https://via.placeholder.com/150";
    }

    // Get location name from location ID
    function getLocationName(locationId: string | undefined): string {
        if (!locationId || !$locations) return "-";
        const location = $locations.find((l) => l.id === locationId);
        return location?.buildingName || location?.city || "-";
    }

    // Active tab state
    let activeRentalTab = "current";

    async function handleCancelRental(rentalId: string): Promise<void> {
        const result = await cancelRentContract(rentalId);
        if (result && !result.errorMessages) {
            // Success
            console.log("Rental cancelled successfully");
        } else {
            // Error
            console.error("Error cancelling rental", result?.errorMessages);
        }

        const rental = $rentContracts?.find(
            (contract) => contract.id === rentalId,
        );
        if (rental?.productId) {
            await fetchProduct(rental.productId);
        }
    }

    // Get status display text and color
    function getStatusInfo(rental: RentContract | undefined): {
        text: string;
        bgClass: string;
    } {
        if (!rental) return { text: "-", bgClass: "bg-gray-100 text-gray-800" };

        switch (rental?.status) {
            case "pickupPending":
                return {
                    text: $t("customer.dashboard.pickup_pending"),
                    bgClass: "bg-blue-100 text-blue-800",
                };
            case "confirmed":
                return {
                    text: $t("customer.dashboard.confirmed"),
                    bgClass: "bg-purple-100 text-purple-800",
                };
            case "active":
                if (rental?.rentalEndDate && rental.rentalEndDate < Date.now() / 1000) {
                    return {
                        text: $t("customer.dashboard.late"),
                        bgClass: "bg-red-100 text-red-800",
                    };
                } else {
                    return {
                        text: $t("customer.dashboard.active"),
                        bgClass: "bg-green-100 text-green-800",
                    };
                }
            case "completed":
                return {
                    text: $t("customer.dashboard.completed"),
                    bgClass: "bg-gray-100 text-gray-800",
                };
            case "cancelled":
                return {
                    text: $t("customer.dashboard.cancelled"),
                    bgClass: "bg-red-100 text-red-800",
                };
            default:
                return { text: "-", bgClass: "bg-gray-100 text-gray-800" };
        }
    }

    // Add pagination state variables and logic after the contract filtering
    // Pagination for current rentals
    let currentRentalsPage = 1;
    let currentRentalsPerPage = 5;

    // Pagination for past rentals
    let pastRentalsPage = 1;
    let pastRentalsPerPage = 5;

    // Calculate paginated data
    $: paginatedCurrentRentals = currentRentals.slice(
        (currentRentalsPage - 1) * currentRentalsPerPage,
        currentRentalsPage * currentRentalsPerPage,
    );

    $: paginatedPastRentals = pastRentals.slice(
        (pastRentalsPage - 1) * pastRentalsPerPage,
        pastRentalsPage * pastRentalsPerPage,
    );

    // Calculate total pages
    $: totalCurrentRentalsPages =
        Math.ceil(currentRentals.length / currentRentalsPerPage) || 1;
    $: totalPastRentalsPages =
        Math.ceil(pastRentals.length / pastRentalsPerPage) || 1;

    // Pagination functions
    function prevCurrentRentalsPage() {
        if (currentRentalsPage > 1) currentRentalsPage--;
    }

    function nextCurrentRentalsPage() {
        if (currentRentalsPage < totalCurrentRentalsPages) currentRentalsPage++;
    }

    function prevPastRentalsPage() {
        if (pastRentalsPage > 1) pastRentalsPage--;
    }

    function nextPastRentalsPage() {
        if (pastRentalsPage < totalPastRentalsPages) pastRentalsPage++;
    }

    function handlePickup(rentContractId: string): void {
        goto(
            ROUTES.CUSTOMER_PICKUP.replace("{rentContractId}", rentContractId),
        );
    }

    function handleReturn(rentContractId: string): void {
        goto(
            ROUTES.CUSTOMER_RETURN.replace("{rentContractId}", rentContractId),
        );
    }

    function redirectToPayment(rentContractId: string): void {
        const rentContract = $rentContracts?.find(
            (contract) => contract.id === rentContractId,
        );
        if (rentContract?.paymentInstructions?.dynamicAttributes?.['url']) {
            document.location.href = rentContract.paymentInstructions.dynamicAttributes['url'];
        }
    }
</script>

<div class="container mx-auto px-4 py-6">
    <!-- Rentals Section -->
    <section class="mb-10">
        <div class="flex justify-between items-center mb-4">
            <ToggleGroup.Root
                type="single"
                value={activeRentalTab}
                onValueChange={(val) => val && (activeRentalTab = val)}
                variant="outline"
                class="rounded-md bg-gray-100 p-1"
            >
                <ToggleGroup.Item
                    class="border-none rounded-md data-[state=on]:bg-white"
                    value="current"
                    >{$t(
                        "customer.dashboard.current_rentals",
                    )}</ToggleGroup.Item
                >
                <ToggleGroup.Item
                    class="border-none rounded-md data-[state=on]:bg-white"
                    value="past"
                    >{$t("customer.dashboard.past_rentals")}</ToggleGroup.Item
                >
            </ToggleGroup.Root>
        </div>

        <Card.Root>
            <Card.Header>
                <h2 class="text-2xl font-semibold">
                    {$t("customer.dashboard.rentals")}
                </h2>
                <p class="text-md text-gray-500">
                    {$t("customer.dashboard.rentals_description")}
                </p>
            </Card.Header>
            <Card.Content>
                <!-- Current Rentals -->
                {#if activeRentalTab === "current"}
                    <div class="grid grid-cols-1 gap-4">
                        {#if currentRentals.length > 0}
                            <Card.Root>
                                <Card.Content class="p-0">
                                    <div class="overflow-x-auto">
                                        <Table.Root>
                                            <Table.Header>
                                                <Table.Row>
                                                    <Table.Head
                                                        >{$t(
                                                            "customer.dashboard.product",
                                                        )}</Table.Head
                                                    >
                                                    <Table.Head
                                                        >{$t(
                                                            "customer.dashboard.status",
                                                        )}</Table.Head
                                                    >
                                                    <Table.Head
                                                        >{$t(
                                                            "customer.dashboard.price",
                                                        )}</Table.Head
                                                    >
                                                    <Table.Head
                                                        >{$t(
                                                            "customer.dashboard.days",
                                                        )}</Table.Head
                                                    >
                                                    <Table.Head
                                                        >{$t(
                                                            "customer.dashboard.rental-date",
                                                        )}</Table.Head
                                                    >
                                                    <Table.Head
                                                        class="text-center"
                                                    ></Table.Head>
                                                </Table.Row>
                                            </Table.Header>
                                            <Table.Body>
                                                {#if $products?.length && $products.length > 0}
                                                    {#each paginatedCurrentRentals as rental}
                                                        <Table.Row
                                                            class="hover:bg-gray-50"
                                                        >
                                                            <Table.Cell>
                                                                <div
                                                                    class="flex items-center gap-3"
                                                                >
                                                                    <img
                                                                        src={getProductImage(
                                                                            rental.productId,
                                                                        )}
                                                                        alt={getProductName(
                                                                            rental.productId,
                                                                        )}
                                                                        class="w-10 h-10 rounded object-cover"
                                                                    />
                                                                    <span
                                                                        class="font-medium"
                                                                        >{getProductName(
                                                                            rental.productId,
                                                                        )}</span
                                                                    >
                                                                </div>
                                                            </Table.Cell>
                                                            <Table.Cell>
                                                                {@const statusInfo =
                                                                    getStatusInfo(
                                                                        rental,
                                                                    )}
                                                                <span
                                                                    class={`px-2 py-1 rounded-full border border-gray-300 text-xs font-medium`}
                                                                >
                                                                    {statusInfo.text}
                                                                </span>
                                                            </Table.Cell>
                                                            <Table.Cell
                                                                >{formatCurrency(
                                                                    rental.totalAmount ||
                                                                        0,
                                                                    "€",
                                                                )}</Table.Cell
                                                            >
                                                            <Table.Cell>
                                                                {getRentalDuration(
                                                                    rental.rentalStartDate,
                                                                    rental.rentalEndDate,
                                                                )}
                                                            </Table.Cell>
                                                            <Table.Cell>
                                                                {formatDate(
                                                                    rental.rentalStartDate,
                                                                )} - {formatDate(
                                                                    rental.rentalEndDate,
                                                                )}
                                                            </Table.Cell>
                                                            <Table.Cell
                                                                class="text-center"
                                                            >
                                                                <Menubar.Root
                                                                    class="p-0 border-none bg-transparent"
                                                                >
                                                                    <Menubar.Menu
                                                                    >
                                                                        <Menubar.Trigger
                                                                        >
                                                                            <DotsHorizontal
                                                                                class="w-4 h-4"
                                                                            />
                                                                        </Menubar.Trigger>
                                                                        <Menubar.Content
                                                                        >
                                                                            {#if rental.status === "confirmed"} 
                                                                                <Menubar.Item
                                                                                    on:click={() =>
                                                                                        handleCancelRental(
                                                                                            rental.id ||
                                                                                                "",
                                                                                        )}
                                                                                >
                                                                                    {$t(
                                                                                        "customer.dashboard.cancel",
                                                                                    )}
                                                                                </Menubar.Item>
                                                                            {/if}
                                                                            {#if rental.status === "pickupPending"}
                                                                                <Menubar.Item
                                                                                    on:click={() =>
                                                                                        handlePickup(
                                                                                            rental.id ||
                                                                                                "",
                                                                                        )}
                                                                                >
                                                                                    {$t(
                                                                                        "customer.dashboard.pickup",
                                                                                    )}
                                                                                </Menubar.Item>
                                                                            {/if}
                                                                            {#if rental.status === "active"}
                                                                                <Menubar.Item
                                                                                    on:click={() =>
                                                                                        handleReturn(
                                                                                            rental.id ||
                                                                                                "",
                                                                                        )}
                                                                                >
                                                                                    {$t(
                                                                                        "customer.dashboard.return",
                                                                                    )}
                                                                                </Menubar.Item>
                                                                            {/if}
                                                                            {#if rental.status === "confirmed"}
                                                                                <Menubar.Item
                                                                                    on:click={() =>
                                                                                        redirectToPayment(
                                                                                            rental.id ||
                                                                                                "",
                                                                                        )}
                                                                                >
                                                                                    {$t(
                                                                                        "customer.dashboard.go-to-payment",
                                                                                    )}
                                                                                </Menubar.Item>
                                                                            {/if}
                                                                        </Menubar.Content>
                                                                    </Menubar.Menu>
                                                                </Menubar.Root>
                                                            </Table.Cell>
                                                        </Table.Row>
                                                    {/each}
                                                {/if}
                                            </Table.Body>
                                        </Table.Root>
                                    </div>
                                </Card.Content>
                                <!-- Add pagination controls for current rentals -->
                                {#if currentRentals.length > currentRentalsPerPage}
                                    <div
                                        class="flex items-center justify-between px-4 py-3 border-t"
                                    >
                                        <div class="text-sm text-gray-700">
                                            {$t("common.pagination.showing")}
                                            <span class="font-medium">
                                                {(currentRentalsPage - 1) *
                                                    currentRentalsPerPage +
                                                    1}
                                            </span>
                                            {$t("common.pagination.to")}
                                            <span class="font-medium">
                                                {Math.min(
                                                    currentRentalsPage *
                                                        currentRentalsPerPage,
                                                    currentRentals.length,
                                                )}
                                            </span>
                                            {$t("common.pagination.of")}
                                            <span class="font-medium"
                                                >{currentRentals.length}</span
                                            >
                                            {$t("common.pagination.entries")}
                                        </div>

                                        <div class="flex gap-2">
                                            <Button
                                                variant="outline"
                                                size="sm"
                                                disabled={currentRentalsPage ===
                                                    1}
                                                on:click={prevCurrentRentalsPage}
                                            >
                                                {$t(
                                                    "common.pagination.previous",
                                                )}
                                            </Button>

                                            <div class="flex items-center">
                                                <span class="mx-2 text-sm">
                                                    {$t(
                                                        "common.pagination.page",
                                                    )}
                                                    {currentRentalsPage}
                                                    {$t("common.pagination.of")}
                                                    {totalCurrentRentalsPages}
                                                </span>
                                            </div>

                                            <Button
                                                variant="outline"
                                                size="sm"
                                                disabled={currentRentalsPage ===
                                                    totalCurrentRentalsPages}
                                                on:click={nextCurrentRentalsPage}
                                            >
                                                {$t("common.pagination.next")}
                                            </Button>
                                        </div>
                                    </div>
                                {/if}
                            </Card.Root>
                        {:else}
                            <div class="p-10 text-center rounded-lg border">
                                <p class="text-gray-500">
                                    {$t(
                                        "customer.dashboard.no_current_rentals",
                                    )}
                                </p>
                                <Button
                                    variant="outline"
                                    class="mt-4"
                                    on:click={() => goto(ROUTES.SHOP)}
                                >
                                    {$t("customer.dashboard.browse_products")}
                                </Button>
                            </div>
                        {/if}
                    </div>
                {:else}
                <div class="grid grid-cols-1 gap-4">
                    {#if pastRentals.length > 0}
                        <Card.Root>
                            <Card.Content class="p-0">
                                <div class="overflow-x-auto">
                                    <Table.Root>
                                        <Table.Header>
                                            <Table.Row>
                                                <Table.Head
                                                    >{$t(
                                                        "customer.dashboard.product",
                                                    )}</Table.Head
                                                >
                                                <Table.Head
                                                    >{$t(
                                                        "customer.dashboard.status",
                                                    )}</Table.Head
                                                >
                                                <Table.Head
                                                    >{$t(
                                                        "customer.dashboard.price",
                                                    )}</Table.Head
                                                >
                                                <Table.Head
                                                    >{$t(
                                                        "customer.dashboard.days",
                                                    )}</Table.Head
                                                >
                                                <Table.Head
                                                    >{$t(
                                                        "customer.dashboard.rental-date",
                                                    )}</Table.Head
                                                >
                                                <Table.Head
                                                    class="text-center"
                                                ></Table.Head>
                                            </Table.Row>
                                        </Table.Header>
                                        <Table.Body>
                                            {#if $products?.length && $products.length > 0}
                                                {#each paginatedPastRentals as rental}
                                                    <Table.Row
                                                        class="hover:bg-gray-50"
                                                    >
                                                        <Table.Cell>
                                                            <div
                                                                class="flex items-center gap-3"
                                                            >
                                                                <img
                                                                    src={getProductImage(
                                                                        rental.productId,
                                                                    )}
                                                                    alt={getProductName(
                                                                        rental.productId,
                                                                    )}
                                                                    class="w-10 h-10 rounded object-cover"
                                                                />
                                                                <span
                                                                    class="font-medium"
                                                                    >{getProductName(
                                                                        rental.productId,
                                                                    )}</span
                                                                >
                                                            </div>
                                                        </Table.Cell>
                                                        <Table.Cell>
                                                            {@const statusInfo =
                                                                getStatusInfo(
                                                                    rental,
                                                                )}
                                                            <span
                                                                class={`px-2 py-1 rounded-full border border-gray-300 text-xs font-medium`}
                                                            >
                                                                {statusInfo.text}
                                                            </span>
                                                        </Table.Cell>
                                                        <Table.Cell
                                                            >{formatCurrency(
                                                                rental.totalAmount ||
                                                                    0,
                                                                "€",
                                                            )}</Table.Cell
                                                        >
                                                        <Table.Cell>
                                                            {getRentalDuration(
                                                                rental.rentalStartDate,
                                                                rental.rentalEndDate,
                                                            )}
                                                        </Table.Cell>
                                                        <Table.Cell>
                                                            {formatDate(
                                                                rental.rentalStartDate,
                                                            )} - {formatDate(
                                                                rental.rentalEndDate,
                                                            )}
                                                        </Table.Cell>
                                                        <Table.Cell
                                                            class="text-center"
                                                        >
                                                        </Table.Cell>
                                                    </Table.Row>
                                                {/each}
                                            {/if}
                                        </Table.Body>
                                    </Table.Root>
                                </div>
                            </Card.Content>
                            <!-- Add pagination controls for current rentals -->
                            {#if pastRentals.length > pastRentalsPerPage}
                                <div
                                    class="flex items-center justify-between px-4 py-3 border-t"
                                >
                                    <div class="text-sm text-gray-700">
                                        {$t("common.pagination.showing")}
                                        <span class="font-medium">
                                            {(pastRentalsPage - 1) *
                                                pastRentalsPerPage +
                                                1}
                                        </span>
                                        {$t("common.pagination.to")}
                                        <span class="font-medium">
                                            {Math.min(
                                                pastRentalsPage *
                                                    pastRentalsPerPage,
                                                pastRentals.length,
                                            )}
                                        </span>
                                        {$t("common.pagination.of")}
                                        <span class="font-medium"
                                            >{pastRentals.length}</span
                                        >
                                        {$t("common.pagination.entries")}
                                    </div>

                                    <div class="flex gap-2">
                                        <Button
                                            variant="outline"
                                            size="sm"
                                            disabled={pastRentalsPage ===
                                                1}
                                            on:click={prevPastRentalsPage}
                                        >
                                            {$t(
                                                "common.pagination.previous",
                                            )}
                                        </Button>

                                        <div class="flex items-center">
                                            <span class="mx-2 text-sm">
                                                {$t(
                                                    "common.pagination.page",
                                                )}
                                                {pastRentalsPage}
                                                {$t("common.pagination.of")}
                                                {totalPastRentalsPages}
                                            </span>
                                        </div>

                                        <Button
                                            variant="outline"
                                            size="sm"
                                            disabled={pastRentalsPage ===
                                                totalPastRentalsPages}
                                            on:click={nextPastRentalsPage}
                                        >
                                            {$t("common.pagination.next")}
                                        </Button>
                                    </div>
                                </div>
                            {/if}
                        </Card.Root>
                    {:else}
                        <div class="p-10 text-center rounded-lg border">
                            <p class="text-gray-500">
                                {$t(
                                    "customer.dashboard.no_past_rentals",
                                )}
                            </p>
                            <Button
                                variant="outline"
                                class="mt-4"
                                on:click={() => goto(ROUTES.SHOP)}
                            >
                                {$t("customer.dashboard.browse_products")}
                            </Button>
                        </div>
                    {/if}
                </div>
                {/if}
            </Card.Content>
        </Card.Root>
    </section>

    <!-- Profile Section -->
    <section>
        <Card.Root>
            <Card.Header>
                <div class="flex flex-row justify-between items-center gap-2">
                    <div>
                        <p class="text-2xl font-semibold">
                            {$t("customer.dashboard.myprofile")}
                        </p>
                        <p class="text-md">
                            {$t("customer.dashboard.myprofile_description")}
                        </p>
                    </div>
                    <div class="flex flex-row gap-2">
                        <Button variant="outline"
                            >{$t("customer.dashboard.edit")}</Button
                        >
                        <Button variant="destructive"
                            >{$t("customer.dashboard.delete")}</Button
                        >
                    </div>
                </div>
            </Card.Header>

            <Card.Content>
                <Table.Root>
                    <Table.Body>
                        <Table.Row>
                            <Table.Cell
                                >{$t("customer.dashboard.email")}</Table.Cell
                            >
                            <Table.Cell>{$user?.email}</Table.Cell>
                        </Table.Row>
                    </Table.Body>
                </Table.Root>
            </Card.Content>
        </Card.Root>
    </section>
</div>
