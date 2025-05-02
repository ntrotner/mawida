<script lang="ts">
    import { productState } from "$lib/states/product";
    import { t } from "$lib/i18n";
    import * as Table from "$lib/components/ui/table/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import { formatCurrency } from "$lib/helpers/price";
    import { goto } from "$app/navigation";
    import { ROUTES } from "$lib/routes";
    import { configState } from "$lib/states/config/config";
    import {
        ProductConfigKey,
        type ProductConfig,
    } from "$lib/states/config/collection/product";
    import { locationState, fetchLocations } from "$lib/states/location";
    import { onMount } from "svelte";
    import { Star, StarFilled } from "svelte-radix";
    import { StarHalf } from "lucide-svelte";

    export let locationId: string | undefined;
    export let categoryId: string | undefined;

    const productsState = productState.getAsyncState();
    const productConfig =
        configState.getConfig<ProductConfig>(ProductConfigKey);
    const locations = locationState.getAsyncState();

    // Generate random review data for products
    const productReviews = new Map<string, { score: number; count: number }>();

    function getProductReview(productId: string | undefined): {
        score: number;
        count: number;
    } {
        if (!productId) return { score: 0, count: 0 };

        if (!productReviews.has(productId)) {
            // Generate random score between 3.0 and 5.0
            const randomScore = Math.random() * 2 + 3;
            // Generate random review count between 5 and 100
            const randomCount = Math.floor(Math.random() * 95) + 5;

            productReviews.set(productId, {
                score: parseFloat(randomScore.toFixed(1)),
                count: randomCount,
            });
        }

        return productReviews.get(productId) || { score: 0, count: 0 };
    }

    function getProductImage(product: any): string {
        if (product?.images && product.images.length > 0) {
            return `data:image/*;base64,${product.images[0].data}`;
        }
        return "";
    }

    function getLocationCity(locationId: string | undefined): string {
        if (!locationId || !$locations) return "-";

        const location = $locations.find((loc) => loc.id === locationId);
        return location?.city || "-";
    }

    // Pagination
    let currentPage = 1;
    let itemsPerPage = 10;

    $: derivedProducts =
        $productsState?.filter((product) => {
            if (!locationId && !categoryId) {
                return true;
            }

            if (locationId && categoryId) {
                return (
                    product.location === locationId &&
                    product.dynamicAttributes?.["category"] === categoryId
                );
            }

            if (locationId) {
                return product.location === locationId;
            }

            if (categoryId) {
                return product.dynamicAttributes?.["category"] === categoryId;
            }

            return true;
        }) || [];

    $: totalPages = Math.ceil(derivedProducts.length / itemsPerPage);
    $: paginatedProducts = derivedProducts.slice(
        (currentPage - 1) * itemsPerPage,
        currentPage * itemsPerPage,
    );

    function prevPage() {
        if (currentPage > 1) currentPage--;
    }

    function nextPage() {
        if (currentPage < totalPages) currentPage++;
    }

    function getCategoryName(categoryId: string | undefined): string {
        if (!categoryId) return $t("product.category.undefined");

        const category = $productConfig?.categories?.find(
            (c) => c.id === categoryId,
        );
        if (category) {
            return $t(category.translationKey);
        }

        if (categoryId === $productConfig?.categoryFallback.id) {
            return $t($productConfig.categoryFallback.translationKey);
        }

        return $t("product.category.undefined");
    }

    function handleProductClick(productId: string | undefined) {
        if (productId) {
            goto(`${ROUTES.SHOP_PRODUCT.replace("{productId}", productId)}`);
        }
    }

    onMount(async () => {
        // Make sure we have locations data
        if (!$locations || $locations.length === 0) {
            await fetchLocations();
        }
    });
</script>

<div class="w-full">
    {#if derivedProducts.length === 0}
        <div
            class="flex flex-col items-center justify-center p-8 border rounded-lg"
        >
            <p class="text-gray-500 mb-4">{$t("shop.products.empty")}</p>
        </div>
    {:else}
        <div class="rounded-md border">
            <Table.Root>
                <Table.Header>
                    <Table.Row>
                        <Table.Head class="min-w-24"></Table.Head>
                        <Table.Head>{$t("shop.products.table.name")}</Table.Head
                        >
                        <Table.Head
                            >{$t("shop.products.table.reviews")}</Table.Head
                        >
                        <Table.Head
                            >{$t("shop.products.table.status")}</Table.Head
                        >
                        <Table.Head
                            >{$t("shop.products.table.location")}</Table.Head
                        >
                        <Table.Head
                            >{$t("shop.products.table.price")}</Table.Head
                        >
                        <Table.Head
                            >{$t("shop.products.table.deposit")}</Table.Head
                        >
                    </Table.Row>
                </Table.Header>
                <Table.Body>
                    {#each paginatedProducts as product}
                        <Table.Row class="cursor-pointer" on:click={() => handleProductClick(product.id)}>
                            <Table.Cell class="w-16 h-16 p-2">
                                <img
                                    src={getProductImage(product)}
                                    alt={product.name}
                                    class="w-16 h-16 object-cover rounded"
                                    style="cursor: pointer;"
                                />
                            </Table.Cell>
                            <Table.Cell>
                                <div class="flex flex-col">
                                    <span
                                        class="text-md cursor-pointer hover:underline"
                                    >
                                        {product.name || "-"}
                                    </span>
                                    <span class="text-sm text-gray-500 pt-1">
                                        {#if product.dynamicAttributes?.["category"]}
                                            {getCategoryName(
                                                product.dynamicAttributes?.["category"],
                                            )}
                                        {:else}
                                            {$t("product.category.undefined")}
                                        {/if}
                                    </span>
                                </div>
                            </Table.Cell>
                            <Table.Cell>
                                {@const review = getProductReview(product.id)}
                                <div class="flex items-center">
                                    <div class="flex mr-1">
                                        {#each Array(5) as _, i}
                                            {@const starValue = i + 1}
                                            {@const filled =
                                                review.score >=
                                                starValue - 0.25}
                                            {@const halfFilled =
                                                !filled &&
                                                review.score >=
                                                    starValue - 0.75}

                                            {#if filled}
                                                <StarFilled
                                                    class="h-4 w-4 text-yellow-400 fill-yellow-400"
                                                />
                                            {:else if halfFilled}
                                                <StarHalf
                                                    class="h-4 w-4 text-yellow-400 fill-yellow-400"
                                                />
                                            {:else}
                                                <Star
                                                    class="h-4 w-4 text-gray-300"
                                                />
                                            {/if}
                                        {/each}
                                    </div>
                                    <span class="text-sm text-gray-500"
                                        >({review.score}) • {review.count}</span
                                    >
                                </div>
                            </Table.Cell>
                            <Table.Cell>
                                <span
                                    class={`px-2 py-1 rounded-md text-xs font-medium border border-gray-150`}
                                >
                                    {product.isRented
                                        ? $t("shop.products.table.rented")
                                        : $t("shop.products.table.available")}
                                </span>
                            </Table.Cell>
                            <Table.Cell>
                                {getLocationCity(product.location)}
                            </Table.Cell>
                            <Table.Cell>
                                <div class="flex flex-col">
                                    <span
                                        >{formatCurrency(
                                            product.pricing?.price || 0,
                                            "€",
                                        )}</span
                                    >
                                </div>
                            </Table.Cell>
                            <Table.Cell>
                                {formatCurrency(
                                    product.pricing?.deposit || 0,
                                    "€",
                                )}
                            </Table.Cell>
                        </Table.Row>
                    {/each}
                </Table.Body>
            </Table.Root>

            <!-- Pagination -->
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
                            derivedProducts.length,
                        )}</span
                    >
                    {$t("common.pagination.of")}
                    <span class="font-medium">{derivedProducts.length}</span>
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
    {/if}
</div>
