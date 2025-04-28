<script>
    import { page } from "$app/stores";
    import Products from "../../../components/shop/products/Products.svelte";

    $: console.log($page.data);
</script>


<Products locationId={$page.data.location} categoryId={$page.data.category} />