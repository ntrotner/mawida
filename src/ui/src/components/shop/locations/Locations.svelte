<script lang="ts">
    import * as ToggleGroup from "$lib/components/ui/toggle-group";
    import { Button } from "$lib/components/ui/button";
    import * as Card from "$lib/components/ui/card";
    import { t } from "$lib/i18n";
    import { locationState } from "$lib/states/location";

    import Loc1 from "$lib/assets/locations/loc-1.png";
    import Loc2 from "$lib/assets/locations/loc-2.png";
    import Loc3 from "$lib/assets/locations/loc-3.png";
    import Loc4 from "$lib/assets/locations/loc-4.png";
    import Loc5 from "$lib/assets/locations/loc-5.png";
    import Loc6 from "$lib/assets/locations/loc-6.png";
    import { HamburgerMenu, SewingPin } from "svelte-radix";
    import { goto } from "$app/navigation";
    import { ROUTES } from "$lib/routes";

    const locations = locationState.getAsyncState();

    function getRandomLocationImage() {
        const images = [Loc1, Loc2, Loc3, Loc4, Loc5, Loc6];
        return images[Math.floor(Math.random() * images.length)];
    }

    function getRandomLocationType() {
        const types = ["crafts", "electronics"];
        return types[Math.floor(Math.random() * types.length)];
    }

    function handleLocationClick(location: string | undefined) {
        if (location) {
            goto(ROUTES.SHOP_LOCATION.replace('{locationId}', location));
        }
    }

    function getRandomDistance() {
        return Math.floor(Math.random() * 100) + 1;
    }
</script>

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
                value="crafts"
                >{$t("shop.categories-selector.crafts")}</ToggleGroup.Item
            >
            <ToggleGroup.Item
                class="border-none rounded-md data-[state=on]:bg-white"
                value="electronics"
                >{$t("shop.categories-selector.electronics")}</ToggleGroup.Item
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
<Card.Root>
    <Card.Header>
        <Card.Title>{$t("shop.available-locations")}</Card.Title>
        <Card.Description
            >{$t("shop.available-locations.description")}</Card.Description
        >
    </Card.Header>
    <Card.Content>
        {#if $locations && $locations.length > 0}
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {#each $locations as location}
                    <Card.Root
                        class="w-[250px] h-full overflow-hidden hover:bg-accent/50 transition-colors cursor-pointer"
                    >
                        <!-- svelte-ignore a11y-click-events-have-key-events -->
                        <!-- svelte-ignore a11y-no-static-element-interactions -->
                        <div
                            class="aspect-square overflow-hidden"
                            on:click={() => handleLocationClick(location.id)}
                        >
                            <img
                                src={getRandomLocationImage()}
                                alt={location.buildingName ||
                                    "Location " + location.id}
                                class="object-cover w-full h-full"
                            />
                        </div>
                        <Card.Content class="p-4">
                            <!-- svelte-ignore a11y-click-events-have-key-events -->
                            <!-- svelte-ignore a11y-no-static-element-interactions -->
                            <div
                                on:click={() =>
                                    handleLocationClick(location.id)}
                            >
                                <h3 class="text-md mb-1">
                                    {location.buildingName ||
                                        "Location " + location.id}
                                </h3>
                                <div
                                    class="flex flex-col justify-between text-sm text-muted-foreground mb-2"
                                >
                                    <div>
                                        {$t(
                                            `shop.location-type.${getRandomLocationType()}`,
                                        )}
                                    </div>
                                    <div class="flex items-center mt-1">
                                        <SewingPin class="h-4 w-4 mr-2" />
                                        {location.city}
                                        ({getRandomDistance() +
                                            "" +
                                            $t("common.distance.km")})
                                    </div>
                                </div>
                            </div>
                        </Card.Content>
                    </Card.Root>
                {/each}
            </div>
        {:else}
            <div class="text-center py-8">
                <p class="text-muted-foreground">
                    {$t("shop.no-locations")}
                </p>
            </div>
        {/if}
    </Card.Content>
</Card.Root>
