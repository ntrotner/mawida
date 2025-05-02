<script lang="ts">
    import { page } from "$app/stores";
    import { onDestroy } from "svelte";
    import { fetchProduct } from "$lib/states/product/effects";
    import { productState } from "$lib/states/product/product";
    import { configState } from "$lib/states/config/config";
    import {
        ProductConfigKey,
        type ProductConfig,
    } from "$lib/states/config/collection/product";
    import * as Card from "$lib/components/ui/card/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import { t } from "$lib/i18n";
    import { formatCurrency } from "$lib/helpers/price";
    import { Star, StarFilled, File, SewingPin } from "svelte-radix";
    import { StarHalf } from "lucide-svelte";
    import { writable, type Unsubscriber } from "svelte/store";
    import { Skeleton } from "$lib/components/ui/skeleton/index.js";
    import { authenticationState } from "$lib/states/authentication";
    import { Avatar, AvatarImage } from "$lib/components/ui/avatar/index.js";
    import loc1 from "$lib/assets/locations/loc-1.png";
    import loc2 from "$lib/assets/locations/loc-2.png";
    import loc3 from "$lib/assets/locations/loc-3.png";
    import loc4 from "$lib/assets/locations/loc-4.png";
    import loc5 from "$lib/assets/locations/loc-5.png";
    import { locationState } from "$lib/states/location";
    import ProductsCarousel from "../../../components/shop/products/ProductsCarousel.svelte";
    import { ROUTES } from "$lib/routes";
    import { goto } from "$app/navigation";

    const subscriptions: Unsubscriber[] = [];
    const authentication = authenticationState.getAsyncState();
    const isAdmin = authenticationState.isAdmin();
    const locations = locationState.getAsyncState();
    const productId = $page.data.productId;
    const loading = writable(true);
    const error = writable<string | null>(null);
    const product = writable<any>(null);
    const selectedImage = writable<string | null>(null);
    const productConfig =
        configState.getConfig<ProductConfig>(ProductConfigKey);

    // Subscribe to product state
    const products = productState.getAsyncState();

    // Generate random review data
    let reviewScore = parseFloat((3 + Math.random() * 2).toFixed(1));
    const reviewCount = Math.floor(Math.random() * 95) + 5; // Random count between 5 and 100

    // Generate random review distribution
    interface StarRating {
        stars: number;
        count: number;
        percentage: number;
    }

    const reviewDistribution: StarRating[] = [
        {
            stars: 5,
            count: Math.floor(Math.random() * reviewCount * 0.4),
            percentage: 0,
        },
        {
            stars: 4,
            count: Math.floor(Math.random() * reviewCount * 0.3),
            percentage: 0,
        },
        {
            stars: 3,
            count: Math.floor(Math.random() * reviewCount * 0.2),
            percentage: 0,
        },
        {
            stars: 2,
            count: Math.floor(Math.random() * reviewCount * 0.1),
            percentage: 0,
        },
        {
            stars: 1,
            count: Math.floor(Math.random() * reviewCount * 0.05),
            percentage: 0,
        },
    ];

    // Make sure the total matches reviewCount
    const totalReviews = reviewDistribution.reduce(
        (sum, item) => sum + item.count,
        0,
    );
    if (totalReviews < reviewCount) {
        reviewDistribution[0].count += reviewCount - totalReviews;
    } else if (totalReviews > reviewCount) {
        const diff = totalReviews - reviewCount;
        for (let i = 0; i < diff; i++) {
            if (reviewDistribution[4].count > 0) reviewDistribution[4].count--;
            else if (reviewDistribution[3].count > 0)
                reviewDistribution[3].count--;
            else if (reviewDistribution[2].count > 0)
                reviewDistribution[2].count--;
            else if (reviewDistribution[1].count > 0)
                reviewDistribution[1].count--;
            else reviewDistribution[0].count--;
        }
    }

    // Calculate percentages for progress bars
    reviewDistribution.forEach((item) => {
        item.percentage = Math.round((item.count / reviewCount) * 100);
    });

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

    // Watch for changes in the page params and product state
    subscriptions.push(
        page.subscribe(async ({ data }) => {
            const productId = data.productId;
            if (!data.productId || typeof data.productId !== "string") return;

            try {
                loading.set(true);

                // Fetch the product data through our effects function
                // This will update the product state automatically
                await Promise.allSettled([
                    fetchProduct(productId),
                    new Promise((resolve) => setTimeout(resolve, 200)),
                ]);

                // The product state is now updated
                const productData = $products?.find((p) => p.id === productId);

                if (!productData) {
                    error.set("Failed to load product details");
                    throw new Error("Product not found");
                }

                product.set(productData);

                // Set the first image as selected
                if (productData.images && productData.images.length > 0) {
                    selectedImage.set(
                        `data:image/*;base64,${productData.images[0].data}`,
                    );
                }
            } catch (err) {
                console.error("Error fetching product:", err);
                error.set("Failed to load product details");
            } finally {
                loading.set(false);
            }
        }),
    );

    // Also subscribe to the product state to get updates if the product changes
    subscriptions.push(
        products.subscribe((updatedProducts) => {
            if (!productId || !updatedProducts) return;

            const currentProduct = updatedProducts.find(
                (p) => p.id === productId,
            );
            if (currentProduct) {
                product.set(currentProduct);
            }
        }),
    );

    function selectImage(imageData: string) {
        selectedImage.set(`data:image/*;base64,${imageData}`);
    }

    function downloadDocument(productDocument: any) {
        if (!productDocument?.data) return;

        // Create a download link
        const link = document.createElement("a");
        link.href = `data:application/octet-stream;base64,${productDocument.data}`;
        link.download = productDocument.name || "document";
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    }

    // After calculating total, set the percentages for each star rating
    let totalActual = reviewDistribution.reduce(
        (sum, item) => sum + item.count,
        0,
    );
    let adjustmentFactor = reviewCount / totalActual;
    let adjustedTotal = 0;

    // Adjust counts to match reviewCount and calculate percentages
    reviewDistribution.forEach((item, i) => {
        item.count =
            i === reviewDistribution.length - 1
                ? reviewCount - adjustedTotal
                : Math.round(item.count * adjustmentFactor);
        adjustedTotal += item.count;
        item.percentage = (item.count / reviewCount) * 100;
    });

    function getRandomLocationImage() {
        const locationImages = [loc1, loc2, loc3, loc4, loc5];
        return locationImages[
            Math.floor(Math.random() * locationImages.length)
        ];
    }

    function getLocationCity(location: string) {
        return $locations?.find((l) => l.id === location)?.city;
    }

    function getLocationBuilding(location: string) {
        return $locations?.find((l) => l.id === location)?.buildingName;
    }

    onDestroy(() => {
        subscriptions.forEach((unsubscribe) => unsubscribe());
    });
</script>

<div class="container mx-auto p-4">
    <div class="grid grid-cols-10 gap-4">
        <div class="col-span-7">
            {#if $loading || $error}
                <div class="flex flex-col gap-4 justify-center items-center">
                    <Skeleton class="w-full h-32" />
                    <Skeleton class="w-full h-32" />
                    <Skeleton class="w-full h-32" />
                </div>
            {:else if $product}
                <Card.Root class="mb-8">
                    <div class="grid grid-cols-1 md:grid-cols-3 gap-8 p-6">
                        <!-- Left side: Product images -->
                        <div class="flex flex-col gap-4">
                            <!-- Main image -->
                            <div
                                class="aspect-square rounded-lg overflow-hidden border max-w-[300px]"
                            >
                                {#if $selectedImage}
                                    <img
                                        src={$selectedImage}
                                        alt={$product.name}
                                        class="h-full h-full object-cover"
                                    />
                                {:else}
                                    <div
                                        class="w-full h-full bg-gray-100 flex items-center justify-center"
                                    >
                                        <span class="text-gray-400"
                                            >No image available</span
                                        >
                                    </div>
                                {/if}
                            </div>

                            <!-- Thumbnail gallery -->
                            {#if $product.images && $product.images.length > 0}
                                <div class="flex flex-wrap gap-2">
                                    {#each $product.images as image}
                                        <button
                                            class="w-16 h-16 rounded-md overflow-hidden border hover:border-primary"
                                            on:click={() =>
                                                selectImage(image.data)}
                                        >
                                            <img
                                                src={`data:image/*;base64,${image.data}`}
                                                alt={image.name ||
                                                    "Product image"}
                                                class="w-full h-full object-cover"
                                            />
                                        </button>
                                    {/each}
                                </div>
                            {/if}
                        </div>

                        <!-- Right side: Product details -->
                        <div class="flex flex-col gap-1 col-span-2">
                            <!-- Product name -->
                            <h1 class="text-2xl font-bold">{$product.name}</h1>

                            <!-- Category -->
                            <div class="text-md text-gray-600">
                                {#if $product.dynamicAttributes?.category}
                                    {getCategoryName(
                                        $product.dynamicAttributes.category,
                                    )}
                                {:else}
                                    {$t("product.category.undefined")}
                                {/if}
                            </div>

                            <!-- Review stars -->
                            <div class="mt-4 flex items-center gap-2">
                                <span class="text-sm text-gray-500"
                                    >{reviewScore}</span
                                >

                                <div class="flex gap-1">
                                    {#each Array(5) as _, i}
                                        {@const starValue = i + 1}
                                        {@const filled =
                                            reviewScore >= starValue - 0.25}
                                        {@const halfFilled =
                                            !filled &&
                                            reviewScore >= starValue - 0.75}

                                        {#if filled}
                                            <StarFilled
                                                class="h-5 w-5 text-yellow-400 fill-yellow-400"
                                            />
                                        {:else if halfFilled}
                                            <StarHalf
                                                class="h-5 w-5 text-yellow-400 fill-yellow-400"
                                            />
                                        {:else}
                                            <Star
                                                class="h-5 w-5 text-gray-300"
                                            />
                                        {/if}
                                    {/each}
                                </div>
                                <span class="text-sm text-gray-500"
                                    >({reviewCount})</span
                                >
                            </div>

                            <!-- Description -->
                            <div class="mt-6">
                                <h2
                                    class="text-md mb-2 font-semibold text-gray-500"
                                >
                                    {$t("product.shop.description")}
                                </h2>
                                <p class="text-gray-700 whitespace-pre-line">
                                    {$product.description}
                                </p>
                            </div>
                        </div>
                    </div>
                </Card.Root>

                <!-- Documents section -->
                <Card.Root>
                    <Card.Header>
                        <Card.Title>{$t("product.shop.documents")}</Card.Title>
                        <Card.Description
                            >{$t(
                                "product.shop.documents_description",
                            )}</Card.Description
                        >
                    </Card.Header>
                    <Card.Content>
                        {#if $product.documents && $product.documents.length > 0}
                            <div class="flex flex-col gap-2">
                                {#each $product.documents as document}
                                    <Card.Root
                                        class="hover:bg-accent/50 transition-colors cursor-pointer"
                                    >
                                        <div
                                            class="flex items-center p-4"
                                            on:click={() =>
                                                downloadDocument(document)}
                                        >
                                            <!-- Document icon -->
                                            <div
                                                class="w-10 h-10 mr-4 bg-gray-100 rounded-full flex items-center justify-center"
                                            >
                                                <File class="w-5 h-5" />
                                            </div>

                                            <!-- Document name -->
                                            <div class="flex-1">
                                                <div class="font-medium">
                                                    {document.name}
                                                </div>
                                                <div
                                                    class="text-sm text-gray-500"
                                                >
                                                    {$t(
                                                        "product.shop.click_to_download",
                                                    )}
                                                </div>
                                            </div>
                                        </div>
                                    </Card.Root>
                                {/each}
                            </div>
                        {:else}
                            <div
                                class="flex flex-row justify-center items-center h-16"
                            >
                                <div class="text-sm text-gray-500">
                                    {$t(
                                        "product.shop.no_documents_description",
                                    )}
                                </div>
                            </div>
                        {/if}
                    </Card.Content>
                </Card.Root>

                <!-- Reviews section -->
                <Card.Root class="mt-8">
                    <Card.Header>
                        <Card.Title>{$t("product.shop.reviews")}</Card.Title>
                    </Card.Header>
                    <Card.Content>
                        <div class="grid grid-cols-10">
                            <!-- Left side: Review statistics -->
                            <div class="flex flex-col col-span-4">
                                <div class="flex items-center gap-1 mb-2">
                                    {#each Array(5) as _, i}
                                        {@const starValue = i + 1}
                                        {@const filled =
                                            reviewScore >= starValue - 0.25}
                                        {@const halfFilled =
                                            !filled &&
                                            reviewScore >= starValue - 0.75}

                                        {#if filled}
                                            <StarFilled
                                                class="h-5 w-5 text-yellow-400 fill-yellow-400"
                                            />
                                        {:else if halfFilled}
                                            <StarHalf
                                                class="h-5 w-5 text-yellow-400 fill-yellow-400"
                                            />
                                        {:else}
                                            <Star
                                                class="h-5 w-5 text-gray-300"
                                            />
                                        {/if}
                                    {/each}

                                    <div
                                        class="flex items-end gap-2 text-sm text-gray-500"
                                    >
                                        <span>{reviewScore}</span>
                                        <span>{$t("product.shop.of")}</span>
                                        <span> 5</span>
                                    </div>
                                </div>

                                <p
                                    class="text-sm font-semibold text-gray-500 mb-4"
                                >
                                    {reviewCount}
                                    {$t("product.shop.total_reviews")}
                                </p>

                                <!-- Star distribution -->
                                <div class="space-y-2">
                                    {#each reviewDistribution as rating}
                                        <div class="flex items-center gap-2">
                                            <div class="flex items-center p-1">
                                                <span
                                                    class="text-gray-500 text-sm"
                                                    >{rating.stars}
                                                    {$t(
                                                        "product.shop.stars",
                                                    )}</span
                                                >
                                            </div>

                                            <div
                                                class="flex-1 bg-gray-200 rounded-full h-4"
                                            >
                                                <div
                                                    class="bg-primary h-4 rounded-full"
                                                    style="width: {rating.percentage}%"
                                                ></div>
                                            </div>

                                            <div class="w-10 text-right">
                                                <span
                                                    class="text-xs text-gray-500"
                                                    >{(
                                                        (rating.count /
                                                            reviewCount) *
                                                        100
                                                    ).toFixed(0)}%</span
                                                >
                                            </div>
                                        </div>
                                    {/each}
                                </div>
                            </div>

                            <div
                                class="flex flex-row justify-center col-span-1"
                            >
                                <div class="h-full w-[1px] bg-gray-200"></div>
                            </div>

                            <!-- Right side: Sample reviews -->
                            <div class="col-span-5">
                                <div class="border-b py-2">
                                    <div class="flex items-center gap-2 mb-2">
                                        <div
                                            class="w-10 h-10 rounded-full bg-gray-200 flex items-center justify-center"
                                        >
                                            <span
                                                class="font-bold text-gray-500"
                                                >{$t(
                                                    "product.shop.review1.initials",
                                                )}</span
                                            >
                                        </div>
                                        <div>
                                            <p class="font-medium">
                                                {$t(
                                                    "product.shop.review1.name",
                                                )}
                                            </p>
                                        </div>
                                    </div>
                                    <div>
                                        <div class="flex items-center gap-2">
                                            <div class="flex items-center">
                                                {#each Array(5) as _, i}
                                                    <StarFilled
                                                        class={`h-4 w-4 ${i < 5 ? "text-yellow-400 fill-yellow-400" : "text-gray-300"}`}
                                                    />
                                                {/each}
                                            </div>
                                            <span class="text-sm font-bold"
                                                >{$t(
                                                    "product.shop.review1.title",
                                                )}</span
                                            >
                                        </div>
                                    </div>
                                    <div>
                                        <p class="text-xs text-gray-500 py-2">
                                            {$t(
                                                "product.shop.review.rented-on",
                                            )}
                                            {$t("product.shop.review1.date")}
                                        </p>
                                    </div>
                                    <p class="text-gray-700 text-sm">
                                        {$t("product.shop.review1.text")}
                                    </p>
                                </div>

                                <div class="py-2">
                                    <div class="flex items-center gap-2 mb-2">
                                        <div
                                            class="w-10 h-10 rounded-full bg-gray-200 flex items-center justify-center"
                                        >
                                            <span
                                                class="font-bold text-gray-500"
                                                >{$t(
                                                    "product.shop.review2.initials",
                                                )}</span
                                            >
                                        </div>
                                        <div>
                                            <p class="font-medium">
                                                {$t(
                                                    "product.shop.review2.name",
                                                )}
                                            </p>
                                        </div>
                                    </div>
                                    <div>
                                        <div class="flex items-center gap-2">
                                            <div class="flex items-center">
                                                {#each Array(5) as _, i}
                                                    <StarFilled
                                                        class={`h-4 w-4 ${i < 5 ? "text-yellow-400 fill-yellow-400" : "text-gray-300"}`}
                                                    />
                                                {/each}
                                            </div>
                                            <span class="text-sm font-bold"
                                                >{$t(
                                                    "product.shop.review2.title",
                                                )}</span
                                            >
                                        </div>
                                    </div>
                                    <div>
                                        <p class="text-xs text-gray-500 py-2">
                                            {$t(
                                                "product.shop.review.rented-on",
                                            )}
                                            {$t("product.shop.review2.date")}
                                        </p>
                                    </div>
                                    <p class="text-gray-700 text-sm">
                                        {$t("product.shop.review2.text")}
                                    </p>
                                </div>
                            </div>
                        </div>
                    </Card.Content>
                </Card.Root>
                <div class="mt-8">
                    <ProductsCarousel
                        locationId={undefined}
                        categoryId={$product?.dynamicAttributes?.category}
                    />
                </div>
            {:else}
                <div class="flex justify-center items-center h-64">
                    <div class="text-xl">{$t("product.shop.not_found")}</div>
                </div>
            {/if}
        </div>
        <div class="col-span-3">
            <div class="sticky top-0">
                {#if $loading}
                    <div
                        class="flex flex-col gap-4 justify-center items-center"
                    >
                        <Skeleton class="w-full h-32" />
                        <Skeleton class="w-full h-32" />
                    </div>
                {:else if $product.id}
                    <Card.Root>
                        <Card.Header>
                            <Card.Title>
                                <span class="text-md font-semibold">
                                    {$t("product.shop.price")}
                                </span>
                            </Card.Title>
                        </Card.Header>
                        <Card.Content>
                            <span class="text-4xl font-bold">
                                {formatCurrency($product.pricing.price, "€")}
                                <span class="text-sm font-normal text-gray-500"
                                    >/{$t("product.shop.per-day")}</span
                                >
                            </span>

                            {#if $authentication?.authenticated && !$isAdmin && !$product.isRented}
                                <div class="flex flex-col gap-4 mt-4">
                                    <Button variant="default"
                                        on:click={() =>
                                            goto(
                                                ROUTES.SHOP_CHECKOUT.replace(
                                                    "{productId}",
                                                    $product.id,
                                                ).replace(
                                                    "{locationId}",
                                                    $product.location,
                                                ),
                                            )}
                                    >
                                        {$t("product.shop.book-now")}
                                    </Button>
                                    <Button variant="outline">
                                        {$t("product.shop.bookmark")}
                                    </Button>
                                </div>
                            {/if}
                        </Card.Content>
                    </Card.Root>
                    <Card.Root class="mt-4 sticky top-0">
                        <Card.Header>
                            <Card.Title>
                                <span class="text-md font-semibold">
                                    {$t("product.shop.location")}
                                </span>
                            </Card.Title>
                        </Card.Header>
                        <Card.Content>
                            <div class="flex items-center gap-2">
                                <Avatar>
                                    <AvatarImage
                                        src={getRandomLocationImage()}
                                    />
                                </Avatar>

                                <div class="flex flex-col gap-2">
                                    <span class="text-md">
                                        {getLocationBuilding($product.location)}
                                    </span>
                                    <span
                                        class="text-sm text-gray-500 flex items-center"
                                    >
                                        <SewingPin class="w-4 h-4" />
                                        {getLocationCity($product.location)}
                                    </span>
                                </div>
                            </div>
                            <div class="flex flex-col gap-4 mt-4">
                                <Button
                                    variant="default"
                                    on:click={() =>
                                        goto(
                                            ROUTES.SHOP_LOCATION.replace(
                                                "{locationId}",
                                                $product.location,
                                            ),
                                        )}
                                >
                                    {$t("product.shop.go-to-location")}
                                </Button>
                                <Button variant="outline">
                                    {$t("product.shop.contact")}
                                </Button>
                            </div>
                        </Card.Content>
                    </Card.Root>
                {/if}
            </div>
        </div>
    </div>
</div>
