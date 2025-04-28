<script lang="ts">
  import ShopSidebar from "../../components/shop/sidebar/Sidebar.svelte";
  import ShopBreadcrumb from "../../components/shop/breadcrumb/Breadcrumb.svelte";
  import ShopProfile from "../../components/shop/profile/Profile.svelte";
  import { onMount } from "svelte";
  import { fetchLocations } from "$lib/states/location";
  import { fetchProducts } from "$lib/states/product";
  import { locationState } from "$lib/states/location";
  import { productState } from "$lib/states/product";

  const location = locationState.getAsyncState();
  const product = productState.getAsyncState();

  onMount(() => {
    if (!$location || $location.length === 0) {
      fetchLocations();
    }
    if (!$product || $product.length === 0) {
      fetchProducts();
    }
  });
</script>

<div class="flex h-screen w-full relative">
  <!-- Sidebar -->
  <ShopSidebar />

  <!-- Main content area -->
  <main class="flex-1 ml-[3.5rem] p-6 overflow-auto">
    <div class="flex flex-row gap-4 items-center justify-between pb-6">
      <div class="flex flex-row gap-4">
        <ShopBreadcrumb />
      </div>
      <div class="flex flex-row gap-4">
        <ShopProfile />
      </div>
    </div>
    <slot />
  </main>
</div>
