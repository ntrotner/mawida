<script lang="ts">
    import { page } from "$app/stores";
    import { onDestroy, onMount } from "svelte";
    import { fetchLocations, locationState } from "$lib/states/location";
    import { productState } from "$lib/states/product";
    import * as Card from "$lib/components/ui/card/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import { t } from "$lib/i18n";
    import { derived, writable, type Unsubscriber } from "svelte/store";
    import { Skeleton } from "$lib/components/ui/skeleton/index.js";
    import {
        SewingPin,
        Link1,
        Share2,
        EnvelopeOpen,
        GithubLogo,
    } from "svelte-radix";
    import loc1 from "$lib/assets/locations/loc-1.png";
    import loc2 from "$lib/assets/locations/loc-2.png";
    import loc3 from "$lib/assets/locations/loc-3.png";
    import loc4 from "$lib/assets/locations/loc-4.png";
    import loc5 from "$lib/assets/locations/loc-5.png";
    import ProductsCarousel from "../../../components/shop/products/ProductsCarousel.svelte";
    import { Avatar, AvatarImage } from "$lib/components/ui/avatar";
    import * as ToggleGroup from "$lib/components/ui/toggle-group";
    import { Microwave, Truck, Wrench } from "lucide-svelte";

    const subscriptions: Unsubscriber[] = [];
    const locationId = $page.data.locationId;
    const loading = writable(true);
    const error = writable<string | null>(null);
    const location = writable<any>(null);
    const locations = locationState.getAsyncState();

    // Mock opening times
    const openingTimes = [
        { day: $t("common.monday"), hours: "09:00 - 18:00" },
        { day: $t("common.tuesday"), hours: "09:00 - 18:00" },
        { day: $t("common.wednesday"), hours: "09:00 - 18:00" },
        { day: $t("common.thursday"), hours: "09:00 - 18:00" },
        { day: $t("common.friday"), hours: "09:00 - 17:00" },
        { day: $t("common.saturday"), hours: "10:00 - 15:00" },
        { day: $t("common.sunday"), hours: $t("common.closed") },
    ];

    // Get current day to highlight it
    const today = new Date().getDay();
    const dayNames = [
        $t("common.sunday"),
        $t("common.monday"),
        $t("common.tuesday"),
        $t("common.wednesday"),
        $t("common.thursday"),
        $t("common.friday"),
        $t("common.saturday"),
    ];
    const currentDay = dayNames[today];

    subscriptions.push(
        page.subscribe(async ({ data }) => {
            const locationId = data.locationId;
            if (!data.locationId || typeof data.locationId !== "string") return;

            try {
                loading.set(true);
                await fetchLocations();

                // Small delay for UX
                await new Promise((resolve) => setTimeout(resolve, 200));

                if (!$locations || $locations.length === 0) {
                    error.set("Failed to load locations");
                    throw new Error("Failed to load locations");
                }

                const locationData = $locations.find(
                    (loc) => loc.id === locationId,
                );
                if (!locationData) {
                    error.set("Location not found");
                    throw new Error("Location not found");
                }

                location.set(locationData);
            } catch (err) {
                console.error("Error fetching location:", err);
                error.set("Failed to load location details");
            } finally {
                loading.set(false);
            }
        }),
    );

    function getLocationImage(locationId: string | undefined): string {
        const locationImages = [loc1, loc2, loc3, loc4, loc5];
        const index = locationId
            ? locationId.charCodeAt(0) % locationImages.length
            : 0;
        return locationImages[index];
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
            {:else if $location}
                <!-- Location Overview Card -->
                <Card.Root class="mb-8">
                    <div class="grid grid-cols-1 md:grid-cols-3 gap-8 p-6">
                        <!-- Left side: Location image and basic info -->
                        <div class="flex flex-row gap-4 col-span-1">
                            <!-- Location image -->
                            <div
                                class="aspect-video rounded-lg overflow-hidden border"
                            >
                                <img
                                    src={getLocationImage($location.id)}
                                    alt={$location.name}
                                    class="w-full h-full object-cover"
                                />
                            </div>
                        </div>

                        <!-- Right side: Opening times -->
                        <div class="col-span-2">
                            <!-- Location basic info -->
                            <div>
                                <h1 class="text-2xl font-bold">
                                    {$location.buildingName}
                                </h1>
                                <div
                                    class="flex items-center gap-1 text-gray-600 mt-2"
                                >
                                    <SewingPin class="h-4 w-4" />
                                    <span
                                        >{$location.street}, {$location.postalCode}
                                        {$location.city}</span
                                    >
                                </div>
                            </div>

                            <h2
                                class="text-md text-gray-500 my-4 font-semibold"
                            >
                                {$t("location.opening_hours")}
                            </h2>
                            <div class="my-4">
                                {#each openingTimes as time}
                                    <div
                                        class="flex justify-between items-center py-2 border-b {time.day ===
                                        currentDay
                                            ? 'bg-primary/5 -mx-2 px-2 rounded'
                                            : ''}"
                                    >
                                        <span
                                            class="text-sm font-semibold {time.day ===
                                            currentDay
                                                ? 'text-primary'
                                                : ''}">{time.day}</span
                                        >
                                        <span class="text-gray-600"
                                            >{time.hours}</span
                                        >
                                    </div>
                                {/each}
                            </div>
                        </div>
                    </div>
                </Card.Root>

                <!-- Products from this location -->
                <div class="mb-8">
                    <!-- Time Period Toggle Group -->
                    <div class="flex items-center space-x-2 mb-2">
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
                    <ProductsCarousel
                        locationId={$location.id}
                        categoryId={undefined}
                        title={$t("location.available_products")}
                        description={$t(
                            "location.available_products_description",
                        )}
                    />
                </div>

                <div class="flex items-center space-x-2 mb-2">
                    <ToggleGroup.Root
                        type="single"
                        value="aboutus"
                        variant="outline"
                        class="rounded-md bg-gray-100 p-1"
                    >
                        <ToggleGroup.Item
                            class="border-none rounded-md data-[state=on]:bg-white"
                            value="aboutus"
                            >{$t(
                                "location.aboutus",
                            )}</ToggleGroup.Item
                        >
                        <ToggleGroup.Item
                            class="border-none rounded-md data-[state=on]:bg-white"
                            value="arrivals"
                            >{$t(
                                "location.arrivals",
                            )}</ToggleGroup.Item
                        >
                    </ToggleGroup.Root>
                </div>

                <!-- Location Description Card -->
                <Card.Root>
                    <Card.Header>
                        <Card.Title>{$t("location.about")}</Card.Title>
                    </Card.Header>
                    <Card.Content>
                        <div class="prose max-w-none">
                            Ihre LOXAM Mietstation bei BAUHAUS Mannheim Mallau bietet eine große Auswahl an Mietgeräten für Bauunternehmen, für die Industrie und für den Garten- und Landschaftsbau.Unsere Fachleute in Mannheim Mallau beraten Sie zur Miete und zur Handhabung von Werkzeug. Egal, ob Sie für Ihre Baustelle einen Presslufthammer, einen Mobilbagger, einen Pritschenkipper oder einen Minibagger mieten wollen – Sie finden alles, was Sie brauchen, in Ihrer LOXAM Mietstation bei BAUHAUS Mannheim Mallau.Die LOXAM Mietstation bei BAUHAUS Mannheim Mallau bietet Ihnen verschiedene Mietoptionen mit kurzer, mittlerer oder langer Laufzeit – je nach Ihren Anforderungen. Besuchen Sie Ihre LOXAM Mietstation bei BAUHAUS, um ein Gerüst, eine Arbeitsbühne oder einen Anhänger in Mannheim Mallau zu mieten.
                        </div>
                    </Card.Content>
                </Card.Root>
            {:else}
                <div class="flex justify-center items-center h-64">
                    <div class="text-xl">{$t("location.not_found")}</div>
                </div>
            {/if}
        </div>

        <!-- Sticky right sidebar -->
        <div class="col-span-3">
            <div class="sticky top-0">
                {#if $loading}
                    <div
                        class="flex flex-col gap-4 justify-center items-center"
                    >
                        <Skeleton class="w-full h-32" />
                    </div>
                {:else if $location}
                    <!-- Social Media Card -->
                    <Card.Root class="sticky top-0">
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
                                        src={getLocationImage($location.id)}
                                    />
                                </Avatar>

                                <div class="flex flex-col gap-2">
                                    <span class="text-md">
                                        {$location.buildingName}
                                    </span>
                                    <span
                                        class="text-sm text-gray-500 flex items-center"
                                    >
                                        <SewingPin class="w-4 h-4" />
                                        {$location.city}
                                    </span>
                                </div>
                            </div>
                            <div class="flex flex-col gap-4 mt-4">
                                <Button variant="default">
                                    {$t("product.shop.contact")}
                                </Button>
                            </div>
                            <div class="pt-6 flex flex-row justify-between">
                                <Button
                                    class="py-4 px-6"
                                    variant="outline"
                                    type="button"
                                >
                                    <svg
                                        class="h-6 w-6"
                                        fill="#000000"
                                        version="1.1"
                                        viewBox="0 0 22.773 22.773"
                                        xml:space="preserve"
                                        xmlns="http://www.w3.org/2000/svg"
                                    >
                                        <path
                                            d="m15.769 0h0.162c0.13 1.606-0.483 2.806-1.228 3.675-0.731 0.863-1.732 1.7-3.351 1.573-0.108-1.583 0.506-2.694 1.25-3.561 0.69-0.808 1.955-1.527 3.167-1.687z"
                                        />
                                        <path
                                            d="m20.67 16.716v0.045c-0.455 1.378-1.104 2.559-1.896 3.655-0.723 0.995-1.609 2.334-3.191 2.334-1.367 0-2.275-0.879-3.676-0.903-1.482-0.024-2.297 0.735-3.652 0.926h-0.462c-0.995-0.144-1.798-0.932-2.383-1.642-1.725-2.098-3.058-4.808-3.306-8.276v-1.019c0.105-2.482 1.311-4.5 2.914-5.478 0.846-0.52 2.009-0.963 3.304-0.765 0.555 0.086 1.122 0.276 1.619 0.464 0.471 0.181 1.06 0.502 1.618 0.485 0.378-0.011 0.754-0.208 1.135-0.347 1.116-0.403 2.21-0.865 3.652-0.648 1.733 0.262 2.963 1.032 3.723 2.22-1.466 0.933-2.625 2.339-2.427 4.74 0.176 2.181 1.444 3.457 3.028 4.209z"
                                        />
                                    </svg>
                                </Button>
                                <Button
                                    class="py-4 px-6"
                                    variant="outline"
                                    type="button"
                                >
                                    <svg
                                        class="h-6 w-6"
                                        version="1.1"
                                        viewBox="0 0 20 20"
                                        xmlns="http://www.w3.org/2000/svg"
                                    >
                                        <title>google [#178]</title>
                                        <desc>Created with Sketch.</desc>
                                        <g fill="none" fill-rule="evenodd">
                                            <g
                                                transform="translate(-300 -7399)"
                                                fill="#000"
                                            >
                                                <g
                                                    transform="translate(56 160)"
                                                >
                                                    <path
                                                        d="m263.82 7247h-9.6102c0 0.99944 0 2.9983-0.006126 3.9978h5.5689c-0.2134 0.99944-0.97001 2.3986-2.0391 3.1033-0.001021-1e-3 -0.002042 0.00599-0.004084 0.00499-1.4213 0.93848-3.297 1.1514-4.6897 0.87151-2.183-0.43375-3.9107-2.0169-4.6121-4.0277 0.004084-3e-3 0.007147-0.03098 0.01021-0.03298-0.43906-1.2473-0.43906-2.9174 0-3.9168h-0.001021c0.56567-1.837 2.3454-3.513 4.5315-3.9718 1.7583-0.37279 3.7422 0.03098 5.2013 1.3962 0.194-0.18989 2.6854-2.6225 2.8722-2.8204-4.9848-4.5145-12.966-2.9264-15.953 2.9034h-0.001021s0.001021 0-0.005105 0.01099c-1.4775 2.8634-1.4162 6.2375 0.01021 8.964-0.004084 0.00299-0.007147 0.00499-0.01021 0.00799 1.2927 2.5086 3.6452 4.4325 6.4797 5.1651 3.0111 0.78956 6.8432 0.24986 9.4101-2.0718l0.003063 3e-3c2.1749-1.9589 3.5288-5.296 2.8447-9.5866"
                                                    ></path>
                                                </g>
                                            </g>
                                        </g>
                                    </svg>
                                </Button>
                                <Button
                                    class="py-4 px-6"
                                    variant="outline"
                                    type="button"
                                >
                                    <GithubLogo class="h-6 w-6" />
                                </Button>
                            </div>
                        </Card.Content>
                    </Card.Root>
                {/if}
            </div>
        </div>
    </div>
</div>
