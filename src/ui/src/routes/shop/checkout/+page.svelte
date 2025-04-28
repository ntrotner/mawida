<script lang="ts">
	import { onMount } from "svelte";
	import { page } from "$app/stores";
	import { rentProduct, fetchProduct } from "$lib/states/product/effects";
	import type { RentProductFormular } from "$lib/open-api/models";
	import { PaymentMethod } from "$lib/open-api/models";
	import { goto } from "$app/navigation";
	import { Skeleton } from "$lib/components/ui/skeleton";
	import { formatCurrency } from "$lib/helpers/price";
	import { Button } from "$lib/components/ui/button";
	import { fetchRentContractById } from "$lib/states/rent_contracts/effects";
	import { rentContractsState } from "$lib/states/rent_contracts";
	let product: any = null;
	let loading = true;
	let error: string | null = null;

	// Form values
	let startDate: string = "";
	let endDate: string = "";
	let additionalNotes: string = "";
	let selectedLocationId: string = "";
	let selectedPaymentMethod: PaymentMethod = PaymentMethod.Stripe;

	onMount(async () => {
		const productId = $page.data.productId;
		if (!productId) {
			error = "No product ID specified";
			loading = false;
			return;
		}

		try {
			product = await fetchProduct(productId);
			selectedLocationId = $page.data.locationId || "";
			loading = false;
		} catch (err) {
			console.error("Failed to load product:", err);
			error = "Failed to load product details";
			loading = false;
		}
	});

	function convertToTimestamp(dateString: string): number {
		return dateString ? new Date(dateString).getTime() : 0;
	}

	async function handleCheckout() {
		if (!product) return;

		const startTimestamp = convertToTimestamp(startDate);
		const endTimestamp = convertToTimestamp(endDate);

		if (
			!startTimestamp ||
			!endTimestamp ||
			startTimestamp >= endTimestamp
		) {
			error = "Please select valid start and end dates";
			return;
		}

		if (!selectedLocationId) {
			error = "Please select a location";
			return;
		}

		const rentFormular: RentProductFormular = {
			rentalStartDate: startTimestamp / 1000,
			rentalEndDate: endTimestamp / 1000,
			locationId: selectedLocationId,
			paymentMethodId: selectedPaymentMethod,
			additionalNotes: additionalNotes || "no additional notes",
			dynamicAttributes: {},
		};

		try {
			error = null;
			const response = await rentProduct(product.id, rentFormular);
			if (response?.rentContractId) {
				await fetchRentContractById(response.rentContractId);
				const rentContract = rentContractsState.getRentContractById(
					response.rentContractId,
				);
				if (rentContract) {
					const redirectionUrl = rentContract.paymentInstructions?.dynamicAttributes?.url;
					if (redirectionUrl) {
						window.location.href = redirectionUrl;
					}
				}
			}
		} catch (err: any) {
			console.error("Checkout failed:", err);
			error =
				err?.errorMessages?.[0]?.message ||
				"Failed to process rental request";
		}
	}
</script>

<div class="checkout-container">
	{#if loading || error}
		<Skeleton class="w-full h-32" />
	{:else if product}
		<h1>Checkout</h1>

		<div class="product-summary">
			<h2>{product.name}</h2>
			<p>{product.description}</p>

			<div class="pricing-details">
				<div class="price-item">
					<span class="label">Price:</span>
					<span class="amount"
						>{formatCurrency(product.pricing?.price, "€")}</span
					>
				</div>
				<div class="price-item">
					<span class="label">Deposit:</span>
					<span class="amount"
						>{formatCurrency(product.pricing?.deposit, "€")}</span
					>
				</div>
			</div>
		</div>

		<form on:submit|preventDefault={handleCheckout} class="rental-form">
			<div class="form-group">
				<label for="start-date">Start Date</label>
				<input
					id="start-date"
					type="datetime-local"
					bind:value={startDate}
					required
				/>
			</div>

			<div class="form-group">
				<label for="end-date">End Date</label>
				<input
					id="end-date"
					type="datetime-local"
					bind:value={endDate}
					required
				/>
			</div>

			<div class="form-group">
				<label for="notes">Additional Notes</label>
				<textarea
					id="notes"
					bind:value={additionalNotes}
					rows="3"
					placeholder="Any special instructions or requests?"
				></textarea>
			</div>

			<Button variant="default" type="submit">Proceed to Payment</Button>
		</form>
	{:else}
		<p>Product not found</p>
	{/if}
</div>
