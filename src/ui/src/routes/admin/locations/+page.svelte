<script lang="ts">
  import { t } from "$lib/i18n";
  import { onMount } from "svelte";
  import { locationState, fetchLocations } from "$lib/states/location";
  import { Button } from "$lib/components/ui/button/index.js";
  import * as Table from "$lib/components/ui/table/index.js";

  let locations = locationState.getAsyncState();
  
  onMount(async () => {
    await fetchLocations();
  });
</script>

<div class="container mx-auto">
  <div class="flex justify-between items-center mb-6">
    <h1 class="text-2xl font-bold">{$t("sidebar.locations")}</h1>
    
    <Button variant="outline">
      {$t("admin.locations.add", { default: "Add Location" })}
    </Button>
  </div>
  
  {#if $locations && $locations.length > 0}
    <div class="rounded-md border">
      <Table.Root>
        <Table.Header>
          <Table.Row>
            <Table.Head>{$t("admin.locations.table.city", { default: "City" })}</Table.Head>
            <Table.Head>{$t("admin.locations.table.street", { default: "Street" })}</Table.Head>
            <Table.Head>{$t("admin.locations.table.postalCode", { default: "Postal Code" })}</Table.Head>
            <Table.Head>{$t("admin.locations.table.buildingName", { default: "Building" })}</Table.Head>
            <Table.Head class="text-right">{$t("admin.locations.table.actions", { default: "Actions" })}</Table.Head>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {#each $locations as location}
            <Table.Row>
              <Table.Cell>{location.city || "-"}</Table.Cell>
              <Table.Cell>{location.street || "-"}</Table.Cell>
              <Table.Cell>{location.postalCode || "-"}</Table.Cell>
              <Table.Cell>{location.buildingName || "-"}</Table.Cell>
              <Table.Cell class="text-right">
                <div class="flex justify-end gap-2">
                  <Button variant="outline" size="sm">
                    {$t("admin.locations.view", { default: "View" })}
                  </Button>
                  <Button variant="outline" size="sm">
                    {$t("admin.locations.edit", { default: "Edit" })}
                  </Button>
                </div>
              </Table.Cell>
            </Table.Row>
          {/each}
        </Table.Body>
      </Table.Root>
    </div>
  {:else}
    <div class="flex flex-col items-center justify-center p-8 border rounded-lg">
      <p class="text-gray-500 mb-4">{$t("admin.locations.empty", { default: "No locations available yet." })}</p>
      <Button variant="outline">
        {$t("admin.locations.add", { default: "Add Your First Location" })}
      </Button>
    </div>
  {/if}
</div>
