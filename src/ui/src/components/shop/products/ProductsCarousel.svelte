<script lang="ts">
    import { productState } from "$lib/states/product";
    import { t } from "$lib/i18n";
    import { goto, invalidateAll } from "$app/navigation";
    import { ROUTES } from "$lib/routes";
    import { configState } from "$lib/states/config/config";
    import {
        ProductConfigKey,
        type ProductConfig,
    } from "$lib/states/config/collection/product";
    import { locationState, fetchLocations } from "$lib/states/location";
    import { onMount } from "svelte";
    import * as Card from "$lib/components/ui/card/index.js";
    import * as Root from "$lib/components/ui/carousel/carousel.svelte";
    import * as Content from "$lib/components/ui/carousel/carousel-content.svelte";
    import * as Item from "$lib/components/ui/carousel/carousel-item.svelte";
    import * as Previous from "$lib/components/ui/carousel/carousel-previous.svelte";
    import * as Next from "$lib/components/ui/carousel/carousel-next.svelte";
    import { formatCurrency } from "$lib/helpers/price";

    export let locationId: string | undefined;
    export let categoryId: string | undefined;
    export let title: string = $t("shop.featured_products");
    export let description: string = $t("shop.featured_products.description");

    const productsState = productState.getAsyncState();
    const productConfig =
        configState.getConfig<ProductConfig>(ProductConfigKey);
    const locations = locationState.getAsyncState();

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
            goto(`${ROUTES.SHOP_PRODUCT.replace("{productId}", productId)}`, {
                replaceState: true,
            });
        }
    }

    onMount(async () => {
        // Make sure we have locations data
        if (!$locations || $locations.length === 0) {
            await fetchLocations();
        }
    });
</script>

<Card.Root>
    <Card.Header>
        <Card.Title>{title}</Card.Title>
        <Card.Description>{description}</Card.Description>
    </Card.Header>
    <Card.Content>
        <div class="relative">
            <Root.default
                opts={{
                    align: "start",
                    loop: false,
                }}
                class="w-full"
            >
                <Content.default>
                    {#each derivedProducts as product}
                        <Item.default
                            class="md:basis-1/2 lg:basis-1/3 xl:basis-1/4"
                        >
                            <!-- svelte-ignore a11y-click-events-have-key-events -->
                            <div
                                on:click={() => handleProductClick(product.id)}
                                class="p-1"
                            >
                                <Card.Root
                                    class="h-full overflow-hidden hover:bg-accent/50 transition-colors cursor-pointer"
                                >
                                    <div class="relative">
                                        <div
                                            class="w-[325px] h-[200px] overflow-hidden"
                                        >
                                            {#if getProductImage(product)}
                                                <img
                                                    src={getProductImage(
                                                        product,
                                                    )}
                                                    alt={product.name}
                                                    class="object-cover w-full h-full"
                                                />
                                            {:else}
                                                <div
                                                    class="w-full h-full bg-gray-100 flex items-center justify-center"
                                                >
                                                    <span class="text-gray-400"
                                                        >No image</span
                                                    >
                                                </div>
                                            {/if}
                                        </div>
                                    </div>
                                    <div class="p-4 text-left">
                                        <h3 class="text-lg font-medium mb-1">
                                            {product.name}
                                        </h3>
                                        <div
                                            class="flex justify-between items-end"
                                        >
                                            <div>
                                                <p class="text-sm text-gray-700">
                                                    {formatCurrency(
                                                        product.pricing
                                                            ?.price || 0,
                                                        "€",
                                                    )}
                                                    <span
                                                        class="text-sm text-gray-500"
                                                    >
                                                        /{$t(
                                                            "product.shop.per-day",
                                                        )}
                                                    </span>
                                                </p>
                                            </div>
                                        </div>
                                    </div>
                                </Card.Root>
                            </div>
                        </Item.default>
                    {/each}
                </Content.default>
                <Previous.default class="left-0" />
                <Next.default class="right-0" />
            </Root.default>
        </div>
    </Card.Content>
</Card.Root>
