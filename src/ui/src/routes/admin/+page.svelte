<script lang="ts">
    import { goto } from "$app/navigation";
    import { Button } from "$lib/components/ui/button";
    import * as Card from "$lib/components/ui/card";
    import { t } from "$lib/i18n";
    import { ROUTES } from "$lib/routes";
    import { userState } from "$lib/states/user";
    import * as Table from "$lib/components/ui/table";
    import { locationState } from "$lib/states/location";
    import * as Menubar from "$lib/components/ui/menubar";
    import { DotsHorizontal } from "svelte-radix";
    import { salesState } from "$lib/states/sales";
    import { formatCurrency } from "$lib/helpers/price";

    const user = userState.getAsyncState();
    const locations = locationState.getAsyncState();
    const sales = salesState.getAsyncState();

    // Filter sales by time period
    function filterSalesByTimePeriod(
        period: string,
        salesData: any[] | undefined,
    ): any[] {
        if (!salesData || period === "all") return salesData || [];

        const now = new Date();
        // Calculate start timestamps for different periods
        const startOfDay = new Date(
            now.getFullYear(),
            now.getMonth(),
            now.getDate(),
        ).getTime();
        const startOfWeek = new Date(
            now.getFullYear(),
            now.getMonth(),
            now.getDate() - Math.max(now.getDay() - 1, 0),
        ).getTime();
        const startOfMonth = new Date(
            now.getFullYear(),
            now.getMonth(),
            1,
        ).getTime();
        const startOfYear = new Date(now.getFullYear(), 0, 1).getTime();

        let startTimestamp: number;

        switch (period) {
            case "day":
                startTimestamp = startOfDay;
                break;
            case "week":
                startTimestamp = startOfWeek;
                break;
            case "month":
                startTimestamp = startOfMonth;
                break;
            case "year":
                startTimestamp = startOfYear;
                break;
            default:
                return salesData;
        }

        return salesData.filter((sale) => {
            const timestamp = sale.createdAt || 0;
            return timestamp * 1000 >= startTimestamp && sale.paid;
        });
    }

    $: dailySales = filterSalesByTimePeriod("day", $sales);
    $: weeklySales = filterSalesByTimePeriod("week", $sales);
    $: monthlySales = filterSalesByTimePeriod("month", $sales);
    $: dailySalesAmount =
        dailySales?.reduce((acc, sale) => acc + (sale.price || 0), 0) || 0;
    $: weeklySalesAmount =
        weeklySales?.reduce((acc, sale) => acc + (sale.price || 0), 0) || 0;
    $: monthlySalesAmount =
        monthlySales?.reduce((acc, sale) => acc + (sale.price || 0), 0) || 0;
    $: dailySalesAmountPercentage =
        (dailySalesAmount / weeklySalesAmount) * 100;

    $: weeklySalesAmountPercentage =
        (weeklySalesAmount / monthlySalesAmount) * 100;
</script>

<div class="flex flex-col gap-6">
    <Card.Root>
        <Card.Header>
            <Card.Title>
                {$t("admin.home.hello").replace(
                    "{name}",
                    $user?.email ?? "Max",
                )}
            </Card.Title>
            <Card.Description>
                <p class="text-md text-black pt-2">
                    {$t("admin.home.description")}
                </p>
            </Card.Description>
        </Card.Header>
    </Card.Root>

    <div class="grid grid-cols-2 gap-4">
        <div class="h-full">
            <Card.Root class="h-full">
                <Card.Header>
                    <Card.Title>
                        <div class="flex justify-between items-center">
                            {$t("admin.home.locations.title")}
                            <Button
                                on:click={() => goto(ROUTES.ADMIN_LOCATIONS)}
                            >
                                {$t("admin.home.locations.overview")}
                            </Button>
                        </div>
                    </Card.Title>
                    <Card.Description>
                        <p>{$t("admin.home.locations.description")}</p>
                    </Card.Description>
                </Card.Header>
                <Card.Content>
                    <Table.Root>
                        <Table.Header>
                            <Table.Row>
                                <Table.Head
                                    >{$t(
                                        "admin.home.locations.city",
                                    )}</Table.Head
                                >
                                <Table.Head
                                    >{$t(
                                        "admin.home.locations.postalCode",
                                    )}</Table.Head
                                >
                                <Table.Head
                                    >{$t(
                                        "admin.home.locations.actions",
                                    )}</Table.Head
                                >
                            </Table.Row>
                        </Table.Header>
                        <Table.Body>
                            {#if $locations}
                                {#each $locations.slice(0, 2) as location}
                                    <Table.Row>
                                        <Table.Cell>{location.city}</Table.Cell>
                                        <Table.Cell
                                            >{location.postalCode}</Table.Cell
                                        >
                                        <Table.Cell>
                                            <div class="flex justify-end gap-2">
                                                <Menubar.Root
                                                    class="p-0 border-none bg-transparent"
                                                >
                                                    <Menubar.Menu>
                                                        <Menubar.Trigger>
                                                            <DotsHorizontal
                                                                class="w-4 h-4"
                                                            />
                                                        </Menubar.Trigger>
                                                        <Menubar.Content>
                                                            <Menubar.Item
                                                                on:click={() => {
                                                                    goto(
                                                                        ROUTES.ADMIN_PRODUCTS.replace(
                                                                            "{locationId}",
                                                                            location.id ||
                                                                                "",
                                                                        ),
                                                                    );
                                                                }}
                                                            >
                                                                {$t(
                                                                    "admin.locations.products.view",
                                                                )}
                                                            </Menubar.Item>
                                                        </Menubar.Content>
                                                    </Menubar.Menu>
                                                </Menubar.Root>
                                            </div>
                                        </Table.Cell>
                                    </Table.Row>
                                {/each}
                            {/if}
                        </Table.Body>
                    </Table.Root>
                </Card.Content>
            </Card.Root>
        </div>
        <div>
            <div class="flex flex-col gap-4 h-full">
                <Card.Root class="h-full">
                    <Card.Header>
                        <Card.Description>
                            {$t("admin.home.statistics.day.title")}
                        </Card.Description>
                        <Card.Content class="p-0 pt-1">
                            <div class="space-y-2">
                                <div class="text-3xl font-bold">
                                    {formatCurrency(dailySalesAmount, "€")}
                                </div>
                                <div class="w-full bg-muted rounded-full h-4">
                                    <div
                                        class={`h-4 rounded-full bg-primary`}
                                        style={`width: ${Math.abs(dailySalesAmountPercentage)}%`}
                                    ></div>
                                </div>
                            </div>
                        </Card.Content>
                    </Card.Header>
                </Card.Root>
                <Card.Root class="h-full">
                    <Card.Header>
                        <Card.Description>
                            {$t("admin.home.statistics.week.title")}
                        </Card.Description>
                        <Card.Content class="p-0 pt-1">
                            <div class="space-y-2">
                                <div class="text-3xl font-bold">
                                    {formatCurrency(weeklySalesAmount, "€")}
                                </div>
                                <div class="w-full bg-muted rounded-full h-4">
                                    <div
                                        class={`h-4 rounded-full bg-primary`}
                                        style={`width: ${Math.abs(weeklySalesAmountPercentage)}%`}
                                    ></div>
                                </div>
                            </div>
                        </Card.Content>
                    </Card.Header>
                </Card.Root>
            </div>
        </div>
    </div>
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
                                >{$t("customer.dashboard.name")}</Table.Cell
                            >
                            <Table.Cell>Bjorn Baumann</Table.Cell>
                        </Table.Row>
                        <Table.Row>
                            <Table.Cell
                                >{$t("customer.dashboard.street")}</Table.Cell
                            >
                            <Table.Cell>Street 123</Table.Cell>
                        </Table.Row>
                        <Table.Row>
                            <Table.Cell
                                >{$t("customer.dashboard.email")}</Table.Cell
                            >
                            <Table.Cell>{$user?.email}</Table.Cell>
                        </Table.Row>
                        <Table.Row>
                            <Table.Cell
                                >{$t("customer.dashboard.phone")}</Table.Cell
                            >
                            <Table.Cell>1234567890</Table.Cell>
                        </Table.Row>
                        <Table.Row>
                            <Table.Cell
                                >{$t("customer.dashboard.about")}</Table.Cell
                            >
                            <Table.Cell>
                            </Table.Cell>
                        </Table.Row>
                    </Table.Body>
                </Table.Root>
            </Card.Content>
        </Card.Root>
    </section>
</div>
