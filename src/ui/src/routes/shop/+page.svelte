<script lang="ts">
    import { configState } from "$lib/states/config/config";
    import { userState } from "$lib/states/user";
    import {
        ProductConfigKey,
        type ProductConfig,
    } from "$lib/states/config/collection/product";
    import { t } from "$lib/i18n";
    import * as Card from "$lib/components/ui/card/index.js";
    import * as Root from "$lib/components/ui/carousel/carousel.svelte";
    import * as Content from "$lib/components/ui/carousel/carousel-content.svelte";
    import * as Item from "$lib/components/ui/carousel/carousel-item.svelte";
    import * as Previous from "$lib/components/ui/carousel/carousel-previous.svelte";
    import * as Next from "$lib/components/ui/carousel/carousel-next.svelte";
    import { Button } from "$lib/components/ui/button/index.js";
    import { goto } from "$app/navigation";
    import * as ToggleGroup from "$lib/components/ui/toggle-group";
    import { HamburgerMenu } from "svelte-radix";
    import { ROUTES } from "$lib/routes";
    import Locations from "../../components/shop/locations/Locations.svelte";
    import { Microwave, Truck, Wrench } from "lucide-svelte";

    // Get product categories from config
    const productConfig =
        configState.getConfig<ProductConfig>(ProductConfigKey);

    const user = userState.getAsyncState();

    function handleCategoryClick(category: string) {
        const searchParams = new URLSearchParams();
        searchParams.set("category", category);
        goto(`${ROUTES.SHOP_SEARCH}?${searchParams.toString()}`);
    }
</script>

<div class="container mx-auto pb-8 px-4">
    <Card.Root class="mb-8">
        <Card.Header>
            <Card.Title
                >{$t("shop.welcome-message.title")}
                {$user?.email ?? ""}</Card.Title
            >
        </Card.Header>
        <Card.Content class="flex flex-col gap-4 items-start">
            {$t("shop.welcome-message.description")}
            <Button variant="default">
                {$t("shop.welcome-message.button")}
            </Button>
        </Card.Content>
    </Card.Root>

    <!-- Sales Table Section -->
    <div class="flex justify-between items-center mb-3 gap-3">
        <!-- Time Period Toggle Group -->
        <div class="flex items-center space-x-2">
            <ToggleGroup.Root
                type="single"
                value="all"
                variant="outline"
                class="rounded-md bg-gray-100 p-1"
            >
                <ToggleGroup.Item
                    class="border-none rounded-md data-[state=on]:bg-white"
                    value="all"
                    >{$t("shop.categories-selector.all")}</ToggleGroup.Item
                >
                <ToggleGroup.Item
                    class="border-none rounded-md data-[state=on]:bg-white"
                    value="tools"
                    >
                    <Wrench class="w-4 h-4 mr-2" />
                    {$t("shop.categories-selector.tools")}</ToggleGroup.Item
                >
                <ToggleGroup.Item
                    class="border-none rounded-md data-[state=on]:bg-white"
                    value="machines"
                    >
                    <Microwave class="w-4 h-4 mr-2" />
                    {$t("shop.categories-selector.machines")}</ToggleGroup.Item
                >
                <ToggleGroup.Item
                    class="border-none rounded-md data-[state=on]:bg-white"
                    value="vehicles"
                    >
                    <Truck class="w-4 h-4 mr-2" />
                    {$t("shop.categories-selector.vehicles")}</ToggleGroup.Item
                >
            </ToggleGroup.Root>
        </div>

        <div class="flex gap-3">
            <Button variant="outline">
                <HamburgerMenu class="h-4 w-4 mr-2" />
                {$t("common.filter")}
            </Button>
        </div>
    </div>
    <!-- Categories Section -->
    <section class="mb-12">
        <Card.Root>
            <Card.Header>
                <Card.Title>{$t("shop.popular-categories")}</Card.Title>
                <Card.Description
                    >{$t(
                        "shop.popular-categories.description",
                    )}</Card.Description
                >
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
                            {#each $productConfig?.categories || [] as category}
                                <Item.default
                                    class="md:basis-1/2 lg:basis-1/3 xl:basis-1/4"
                                >
                                    <!-- svelte-ignore a11y-click-events-have-key-events -->
                                    <div
                                        on:click={() => handleCategoryClick(category.id)}
                                        class="p-1"
                                    >
                                        <Card.Root
                                            class="h-full overflow-hidden hover:bg-accent/50 transition-colors cursor-pointer"
                                        >
                                            <div
                                                class="w-[325px] h-full overflow-hidden"
                                            >
                                                <img
                                                    src={category.image}
                                                    alt={$t(
                                                        category.translationKey,
                                                    )}
                                                    class="object-cover w-full h-full"
                                                />
                                            </div>
                                            <div class="p-4 text-left">
                                                <h3 class="text-md mb-1">
                                                    {$t(
                                                        category.translationKey,
                                                    )}
                                                </h3>
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
    </section>
    <section>
        <Locations />
    </section>
</div>
