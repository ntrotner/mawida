<script lang="ts">
  import AdminSidebar from "../../components/sidebar/AdminSidebar.svelte";
</script>

<div class="flex h-screen w-full relative">
  <!-- Sidebar -->
  <AdminSidebar />

  <!-- Main content area -->
  <main class="flex-1 ml-[3.5rem] p-6 overflow-auto">
    <slot />
  </main>
</div>
