<script>
    import { page } from "$app/stores";
</script>

<div>
    <h1>Payment</h1>
    <p>Product ID: {$page.data.productId}</p>
    <p>Success: {$page.data.success}</p>
</div>