<script lang="ts">
    import { onMount } from "svelte";
    import { fetchSales } from "$lib/states/sales";
    import { salesState } from "$lib/states/sales";
    import { t } from "$lib/i18n";
    import { Button } from "$lib/components/ui/button/index.js";
    import * as Table from "$lib/components/ui/table/index.js";
    import * as Card from "$lib/components/ui/card";
    import { HamburgerMenu, File } from "svelte-radix";
    import { formatCurrency } from "$lib/helpers/price";
    import { writable } from "svelte/store";
    import { Skeleton } from "$lib/components/ui/skeleton";
    import * as ToggleGroup from "$lib/components/ui/toggle-group";

    let sales = salesState.getAsyncState();
    let loading = writable(true);
    let filteredSales = writable<typeof $sales>([]);
    
    // Time period filter
    let timePeriod = "all";

    // Pagination
    let currentPage = 1;
    let itemsPerPage = 15;

    $: totalPages = Math.ceil(($filteredSales?.length || 0) / itemsPerPage);
    $: paginatedSales = $filteredSales
        ? $filteredSales.slice(
              (currentPage - 1) * itemsPerPage,
              currentPage * itemsPerPage,
          )
        : [];

    // Calculate statistics for the summary cards
    $: totalPaid = $filteredSales
        ? $filteredSales.reduce(
              (sum, sale) => sum + (sale?.paid ? sale.price || 0 : 0),
              0,
          )
        : 0;

    $: allSales = $filteredSales
        ? $filteredSales.reduce((sum, sale) => sum + (sale.price || 0), 0)
        : 0;

    $: totalPaidPercentage = totalPaid > 0 ? (totalPaid / allSales) * 100 : 0;

    $: totalDepositsWithRefunds = $filteredSales
        ? $filteredSales.reduce((sum, sale) => {
              if (sale.depositRefundedAt) {
                  return (
                      sum +
                      (sale.deposit || 0) -
                      (sale.depositRefundedAmount || 0)
                  );
              }
              return sum;
          }, 0)
        : 0;

    $: pendingDeposits = $filteredSales
        ? $filteredSales.reduce((sum, sale) => {
              if (!sale?.paid || sale.depositRefundedAt) {
                  return sum;
              }
              // Only count deposits that haven't been fully refunded
              const pendingAmount = sale.deposit || 0;
              return sum + (pendingAmount > 0 ? pendingAmount : 0);
          }, 0)
        : 0;

    $: pendingDepositsPercentage =
        totalPaid > 0 ? (totalDepositsWithRefunds / pendingDeposits) * 100 : 0;
    $: pendingDepositeColor =
        pendingDepositsPercentage > 0 ? "bg-green-500" : "bg-red-500";

    function prevPage() {
        if (currentPage > 1) currentPage--;
    }

    function nextPage() {
        if (currentPage < totalPages) currentPage++;
    }

    // Format timestamp to readable date
    function formatDate(timestamp: number | undefined) {
        if (!timestamp) return "-";
        return new Date(timestamp * 1000).toLocaleDateString();
    }

    // Filter sales by time period
    function filterSalesByTimePeriod(period: string, salesData: any[]): any[] {
        if (!salesData || period === "all") return salesData;
        
        const now = new Date();
        // Calculate start timestamps for different periods
        const startOfWeek = new Date(now.getFullYear(), now.getMonth(), now.getDate() - Math.max(now.getDay() - 1, 0)).getTime();
        const startOfMonth = new Date(now.getFullYear(), now.getMonth(), 1).getTime();
        const startOfYear = new Date(now.getFullYear(), 0, 1).getTime();
        
        let startTimestamp: number;
        
        switch(period) {
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
        
        return salesData.filter(sale => {
            const timestamp = sale.createdAt || 0;
            return (timestamp * 1000) >= startTimestamp;
        });
    }
    
    // Handle time period change
    function handleTimePeriodChange(value: string) {
        timePeriod = value;
        currentPage = 1; // Reset to first page when changing filter
        filteredSales.set(filterSalesByTimePeriod(value, $sales || []));
    }

    onMount(async () => {
        try {
            await Promise.allSettled([
                await fetchSales(),
                new Promise((resolve) => setTimeout(resolve, 150)),
            ]);
            // Initialize filtered sales with all sales
            filteredSales.set($sales || []);
        } catch {
        } finally {
            loading.set(false);
        }
    });
</script>

<div class="mx-auto space-y-6">
    <!-- Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <!-- Total Sales Card -->
        <Card.Root>
            <Card.Header>
                <Card.Title>{$t("admin.statistics.totalSales")}</Card.Title>
                <Card.Description
                    >{$t(
                        "admin.statistics.totalSalesDescription",
                    )}</Card.Description
                >
            </Card.Header>
            <Card.Content>
                {#if $loading}
                    <Skeleton class="w-full h-[50px]" />
                {:else}
                     <!-- Simple bar visualization -->
                     <div class="space-y-2">
                        <div class="text-3xl font-bold">
                            {formatCurrency(totalPaid, "€")}
                        </div>
                         <div class="w-full bg-muted rounded-full h-4">
                            <div
                                class={`h-4 rounded-full bg-primary`}
                                style={`width: ${Math.abs(totalPaidPercentage)}%`}
                            ></div>
                        </div>
                     </div>
                {/if}
            </Card.Content>
        </Card.Root>
        
        <!-- Pending Deposits Card -->
        <Card.Root>
            <Card.Header>
                <Card.Title>{$t("admin.statistics.pendingDeposits")}</Card.Title
                >
                <Card.Description
                    >{$t(
                        "admin.statistics.pendingDepositsDescription",
                    )}</Card.Description
                >
            </Card.Header>
            <Card.Content>
                {#if $loading}
                    <Skeleton class="w-full h-[50px]" />
                {:else}
                     <!-- Simple bar visualization -->
                     <div class="space-y-2">
                        <div class="text-3xl font-bold">
                            {formatCurrency(pendingDeposits, "€")}
                        </div>
                         <div class="w-full bg-muted rounded-full h-4">
                             <!-- Make width relative to total paid -->
                            <div
                                class={`h-4 rounded-full ${pendingDepositeColor}`}
                                style={`width: ${Math.abs(pendingDepositsPercentage)}%`}
                            ></div>
                         </div>
                     </div>
                {/if}
            </Card.Content>
        </Card.Root>
    </div>

    <!-- Sales Table Section -->
    <div class="flex justify-between items-center mb-3 gap-3">
        <!-- Time Period Toggle Group -->
        <div class="flex items-center space-x-2">
            <ToggleGroup.Root 
                type="single" 
                value={timePeriod} 
                onValueChange={(value) => value && handleTimePeriodChange(value)}
                variant="outline"
                class="rounded-md bg-gray-100 p-1"
            >
                <ToggleGroup.Item class="border-none rounded-md data-[state=on]:bg-white" value="week">{$t("admin.statistics.week")}</ToggleGroup.Item>
                <ToggleGroup.Item class="border-none rounded-md data-[state=on]:bg-white" value="month">{$t("admin.statistics.month")}</ToggleGroup.Item>
                <ToggleGroup.Item class="border-none rounded-md data-[state=on]:bg-white" value="year">{$t("admin.statistics.year")}</ToggleGroup.Item>
                <ToggleGroup.Item class="border-none rounded-md data-[state=on]:bg-white" value="all">{$t("admin.statistics.all")}</ToggleGroup.Item>
            </ToggleGroup.Root>
        </div>
        
        <div class="flex gap-3">
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
    {:else if $filteredSales && $filteredSales.length > 0}
        <div class="rounded-md border">
            <Table.Root>
                <Table.Header>
                    <h2 class="text-xl font-bold px-4 pt-4">
                        {$t("admin.statistics.salesOverview")}
                    </h2>
                    <p class="text-sm text-gray-500 px-4 pt-3 pb-4">
                        {$t("admin.statistics.salesOverviewDescription")}
                    </p>
                    <Table.Row>
                        <Table.Head
                            >{$t("admin.statistics.table.id")}</Table.Head
                        >
                        <Table.Head
                            >{$t("admin.statistics.table.product")}</Table.Head
                        >
                        <Table.Head
                            >{$t("admin.statistics.table.user")}</Table.Head
                        >
                        <Table.Head
                            >{$t("admin.statistics.table.price")}</Table.Head
                        >
                        <Table.Head
                            >{$t("admin.statistics.table.deposit")}</Table.Head
                        >
                        <Table.Head
                            >{$t("admin.statistics.table.status")}</Table.Head
                        >
                        <Table.Head
                            >{$t(
                                "admin.statistics.table.createdDate",
                            )}</Table.Head
                        >
                        <Table.Head
                            >{$t("admin.statistics.table.paidDate")}</Table.Head
                        >
                    </Table.Row>
                </Table.Header>
                <Table.Body>
                    {#each paginatedSales as sale}
                        <Table.Row>
                            <Table.Cell class="font-medium"
                                >{sale.id?.substring(0, 8) || "-"}</Table.Cell
                            >
                            <Table.Cell>{sale.productId || "-"}</Table.Cell>
                            <Table.Cell>{sale.userId || "-"}</Table.Cell>
                            <Table.Cell
                                >{formatCurrency(
                                    sale.price || 0,
                                    "€",
                                )}</Table.Cell
                            >
                            <Table.Cell
                                >{formatCurrency(
                                    sale.deposit || 0,
                                    "€",
                                )}</Table.Cell
                            >
                            <Table.Cell>
                                {#if sale.paid}
                                    <span
                                        class="px-2 py-1 rounded-md border"
                                    >
                                        {$t("admin.statistics.paid")}
                                    </span>
                                {:else}
                                    <span
                                        class="px-2 py-1 rounded-md border"
                                        >{$t("admin.statistics.pending")}
                                    </span>
                                {/if}
                            </Table.Cell>
                            <Table.Cell>{formatDate(sale.createdAt)}</Table.Cell
                            >
                            <Table.Cell
                                >{sale.paidAt
                                    ? formatDate(sale.paidAt)
                                    : "-"}</Table.Cell
                            >
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
                            $filteredSales?.length || 0,
                        )}</span
                    >
                    {$t("common.pagination.of")}
                    <span class="font-medium">{$filteredSales?.length || 0}</span>
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
                            {totalPages}</span
                        >
                    </div>
                    
                    <Button
                        variant="outline"
                        size="sm"
                        disabled={currentPage === totalPages}
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
                {$t("admin.statistics.noSalesDataAvailable")}
            </p>
        </div>
    {/if}
</div>
