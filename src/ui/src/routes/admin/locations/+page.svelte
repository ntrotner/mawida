<script lang="ts">
  import { t } from "$lib/i18n";
  import { onMount } from "svelte";
  import { locationState, fetchLocations } from "$lib/states/location";
  import { Button } from "$lib/components/ui/button/index.js";
  import * as Table from "$lib/components/ui/table/index.js";
  import { ROUTES } from "$lib/routes/routes";
  import { goto } from "$app/navigation";
  import { PlusCircled, HamburgerMenu, File, DotsHorizontal } from "svelte-radix";
  import { Skeleton } from "$lib/components/ui/skeleton";
  import { writable } from "svelte/store";
  import * as Menubar from "$lib/components/ui/menubar/index.js";

  let locations = locationState.getAsyncState();
  let loading = writable(true);

  // Pagination
  let currentPage = 1;
  let itemsPerPage = 10;

  $: totalPages = Math.ceil(($locations?.length || 0) / itemsPerPage);
  $: paginatedLocations = $locations
    ? $locations.slice(
        (currentPage - 1) * itemsPerPage,
        currentPage * itemsPerPage,
      )
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
        await fetchLocations(),
        new Promise((resolve) => setTimeout(resolve, 150)),
      ]);
    } catch {
    } finally {
      loading.set(false);
    }
  });
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
    <Button
      variant="default"
      on:click={() => {
        goto(ROUTES.ADMIN_LOCATIONS_CREATE);
      }}
    >
      <PlusCircled class="h-4 w-4 mr-2" />
      {$t("admin.locations.add")}
    </Button>
  </div>

  {#if $loading}
    <div class="flex justify-center items-center h-full">
      <Skeleton class="w-full h-[30rem]" />
    </div>
  {:else if $locations && $locations.length > 0}
    <div class="rounded-md border">
      <Table.Root>
        <Table.Header>
          <h1 class="text-2xl font-bold px-4 pt-4">
            {$t("admin.sidebar.locations")}
          </h1>
          <p class="text-sm text-gray-500 px-4 pt-3 pb-4">
            {$t("admin.locations.description")}
          </p>
          <Table.Row>
            <Table.Head>{$t("admin.locations.table.city")}</Table.Head>
            <Table.Head>{$t("admin.locations.table.street")}</Table.Head>
            <Table.Head>{$t("admin.locations.table.postalCode")}</Table.Head>
            <Table.Head>{$t("admin.locations.table.buildingName")}</Table.Head>
            <Table.Head class="text-right"
              >{$t("admin.locations.table.actions")}</Table.Head
            >
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {#each paginatedLocations as location}
            <Table.Row>
              <Table.Cell>{location.city || "-"}</Table.Cell>
              <Table.Cell>{location.street || "-"}</Table.Cell>
              <Table.Cell>{location.postalCode || "-"}</Table.Cell>
              <Table.Cell>{location.buildingName || "-"}</Table.Cell>
              <Table.Cell class="text-right">
                <div class="flex justify-end gap-2">
                  <Menubar.Root class="p-0 border-none bg-transparent">
                    <Menubar.Menu>
                      <Menubar.Trigger>
                        <DotsHorizontal class="w-4 h-4" />
                      </Menubar.Trigger>
                      <Menubar.Content>
                        <Menubar.Item
                          on:click={() => {
                            goto(
                              ROUTES.ADMIN_PRODUCTS.replace(
                                "{locationId}",
                                location.id || "",
                              ),
                            );
                          }}
                        >
                          {$t("admin.locations.products.view")}
                        </Menubar.Item>
                      </Menubar.Content>
                    </Menubar.Menu>
                  </Menubar.Root>
                </div>
              </Table.Cell>
            </Table.Row>
          {/each}
        </Table.Body>
      </Table.Root>

      <!-- Simple Pagination -->
      <div class="flex items-center justify-between px-4 py-3 border-t">
        <div class="text-sm text-gray-700">
          {$t("common.pagination.showing")}
          <span class="font-medium">{(currentPage - 1) * itemsPerPage + 1}</span
          >
          {$t("common.pagination.to")}
          <span class="font-medium"
            >{Math.min(
              currentPage * itemsPerPage,
              $locations?.length || 0,
            )}</span
          >
          {$t("common.pagination.of")}
          <span class="font-medium">{$locations?.length || 0}</span>
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
      <p class="text-gray-500 mb-4">{$t("admin.locations.empty")}</p>
      <Button variant="outline">
        {$t("admin.locations.add")}
      </Button>
    </div>
  {/if}
</div>
