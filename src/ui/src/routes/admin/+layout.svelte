<script lang="ts">
  import AdminSidebar from "../../components/admin/sidebar/Sidebar.svelte";
  import AdminBreadcrumb from "../../components/admin/breadcrumb/Breadcrumb.svelte";
  import AdminProfile from "../../components/admin/profile/Profile.svelte";
</script>

<div class="flex h-screen w-full relative">
  <!-- Sidebar -->
  <AdminSidebar />

  <!-- Main content area -->
  <main class="flex-1 ml-[3.5rem] p-6 overflow-auto">
    <div class="flex flex-row gap-4 items-center justify-between pb-6">
      <div class="flex flex-row gap-4">
        <AdminBreadcrumb />
      </div>
      <div class="flex flex-row gap-4">
        <AdminProfile />
      </div>
    </div>
    <slot />
  </main>
</div>
