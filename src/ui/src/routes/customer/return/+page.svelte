<script lang="ts">
	import { onMount } from "svelte";
	import { page } from "$app/stores";
	import { goto } from "$app/navigation";
	import { Skeleton } from "$lib/components/ui/skeleton";
	import { Button } from "$lib/components/ui/button";
	import { fetchRentContractById, returnProduct } from "$lib/states/rent_contracts/effects";
	import { rentContractsState } from "$lib/states/rent_contracts";
    import { ROUTES } from "$lib/routes";
    import { fetchProduct } from "$lib/states/product";

	let rentContract: any = null;
	let loading = true;
	let error: string | null = null;
	let returnImages: string[] = [];
	let additionalNotes: string = "";
	
	onMount(async () => {
		const rentContractId = $page.data.rentContractId;
		if (!rentContractId) {
			error = "No rental contract ID specified";
			loading = false;
			return;
		}

		try {
			// Fetch the rental contract
			await fetchRentContractById(rentContractId);
			rentContract = rentContractsState.getRentContractById(rentContractId);
			
			if (!rentContract) {
				error = "Rental contract not found";
			}
			
			loading = false;
		} catch (err) {
			console.error("Failed to load rental contract:", err);
			error = "Failed to load rental contract details";
			loading = false;
		}
	});

	async function handleReturn() {
		if (!rentContract) return;
		
		const rentContractId = rentContract.id;
		
		try {
			error = null;
			const response = await returnProduct(rentContractId, {
				returnImages: returnImages,
				additionalNotes: additionalNotes || "no additional notes"
			});
			
			if (!response?.errorMessages) {
				const productId = rentContractsState.getRentContractById(rentContractId)?.productId;
				// Success - redirect to dashboard
				if (productId) {
					await fetchProduct(productId);
				}
				goto(ROUTES.CUSTOMER_PROFILE);
			} else if (response?.errorMessages?.length) {
				// Error with messages
				error = response.errorMessages[0].message || "Failed to confirm return";
			}
		} catch (err: any) {
			console.error("Return confirmation failed:", err);
			error = err?.errorMessages?.[0]?.message || "Failed to confirm return";
		}
	}
</script>

<div class="container mx-auto p-4">
	{#if loading || error}
		<Skeleton class="w-full h-32" />
	{:else if rentContract}
		<h1>Confirm Return</h1>
		
		<div class="product-summary">
			<h2>Rental Contract #{rentContract.id}</h2>
			
			<div class="rental-details">
				<div class="detail-item">
					<span class="label">Start Date:</span>
					<span class="value">{new Date(rentContract.rentalStartDate * 1000).toLocaleString()}</span>
				</div>
				<div class="detail-item">
					<span class="label">End Date:</span>
					<span class="value">{new Date(rentContract.rentalEndDate * 1000).toLocaleString()}</span>
				</div>
				<div class="detail-item">
					<span class="label">Total Amount:</span>
					<span class="value">{(rentContract.totalAmount / 100).toFixed(2)} €</span>
				</div>
			</div>
		</div>

		<form on:submit|preventDefault={handleReturn} class="return-form">
			<Button variant="default" type="submit">Confirm Return</Button>
		</form>
	{:else}
		<p>Rental contract not found</p>
	{/if}
</div>
