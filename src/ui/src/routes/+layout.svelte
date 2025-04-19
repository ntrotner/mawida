<script lang="ts">
  import "../app.pcss";
  import { Toaster } from "$lib/components/ui/sonner";
  import { appState } from "$lib/states/app";
  import { DoubleBounce } from "svelte-loading-spinners";
  import { map } from "rxjs";

  const loaded = appState.observable().pipe(map((state) => state.loaded));
</script>

{#if !$loaded}
  <div class="overlay">
    <div class="loading">
      <DoubleBounce color="#0f172a"></DoubleBounce>
    </div>
  </div>
{/if}
<Toaster></Toaster>
{#if $loaded}
  <slot />
{/if}

<style>
  .loading {
    position: absolute;
    top: 50%;
    left: 50%;
    margin-right: -50%;
    transform: translate(-50%, -50%);
  }
  .overlay {
    display: block;
    position: fixed;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.04);
    z-index: 2;
    cursor: pointer;
  }
</style>
