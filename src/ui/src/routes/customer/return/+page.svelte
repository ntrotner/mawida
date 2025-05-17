<script lang="ts">
	import { onMount } from "svelte";
	import { page } from "$app/stores";
	import { goto } from "$app/navigation";
	import { Skeleton } from "$lib/components/ui/skeleton";
	import { Button } from "$lib/components/ui/button";
	import {
		fetchRentContractById,
		returnProduct,
	} from "$lib/states/rent_contracts/effects";
	import { rentContractsState } from "$lib/states/rent_contracts";
	import { ROUTES } from "$lib/routes";
	import { fetchProduct, productState } from "$lib/states/product";
	import { Checkbox } from "$lib/components/ui/checkbox";
	import { t } from "$lib/i18n";
	import * as Textarea from "$lib/components/ui/textarea";
	import { writable } from "svelte/store";
	import { Input } from "$lib/components/ui/input";
	import StarRating from "$lib/components/ui/star-rating.svelte";

	const productsState = productState.getAsyncState();
	let rentContract: any = null;
	let product: any = null;
	let loading = true;
	let error: string | null = null;
	let returnImages = [] as Array<{ id: string; name: string; data: string }>;
	let reviewStars: number = 0;
	let reviewDescription: string = "";

	let additionalNotes: string = "";

	let damage = false;
	let damageDescription: string | null = "";
	let inspected = false;

	let isSubmitting = writable(false);

	onMount(async () => {
		const rentContractId = $page.data.rentContractId;
		if (!rentContractId) {
			error = "No rental contract ID specified";
			loading = false;
			return;
		}

		try {
			// Fetch the rental contract
			await Promise.allSettled([
				new Promise((resolve) => setTimeout(resolve, 150)),
				fetchRentContractById(rentContractId),
			]);
			rentContract =
				rentContractsState.getRentContractById(rentContractId);

			const localProduct = $productsState?.find(
				(product) => product.id === rentContract.productId,
			);
			if (localProduct) {
				product = localProduct;
			} else {
				await fetchProduct(rentContract.productId);
				product = $productsState?.find(
					(product) => product.id === rentContract.productId,
				);
			}
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
				returnImages: returnImages.map((image) => image.data),
				additionalNotes: additionalNotes || "no additional notes",
			});

			if (!response?.errorMessages) {
				const productId =
					rentContractsState.getRentContractById(
						rentContractId,
					)?.productId;
				// Success - redirect to dashboard
				if (productId) {
					await fetchProduct(productId);
				}
				goto(ROUTES.CUSTOMER_PROFILE);
			} else if (response?.errorMessages?.length) {
				// Error with messages
				error =
					response.errorMessages[0].message ||
					"Failed to confirm return";
			}
		} catch (err: any) {
			console.error("Return confirmation failed:", err);
			error =
				err?.errorMessages?.[0]?.message || "Failed to confirm return";
		}
	}

	function damageChanged() {
		damage = !damage;
		damageDescription = "";
	}

	async function handleImageUpload(event: Event) {
		const input = event.target as HTMLInputElement;
		if (!input.files || input.files.length === 0) return;

		const images = returnImages || [];

		for (let i = 0; i < input.files.length; i++) {
			const file = input.files[i];
			const reader = new FileReader();

			reader.onload = (e) => {
				const result = e.target?.result as string;
				if (result) {
					const base64Data = result.split(",")[1]; // Remove data URL prefix
					images.push({
						id:
							typeof crypto.randomUUID === "function"
								? crypto.randomUUID()
								: Math.random().toString(36).substring(2, 15),
						name: file.name,
						data: base64Data,
					});
					returnImages = [...images];
				}
			};

			reader.readAsDataURL(file);
		}

		// Reset input
		input.value = "";
	}

	function removeImage(id: string) {
		returnImages = returnImages.filter((image) => image.id !== id);
	}
</script>

<div class="container mx-auto p-4">
	<div class="grid grid-cols-10 gap-4">
		<div class="col-span-7">
			{#if rentContract && product}
				<div class="rounded-md border p-4 mb-4">
					<p class="text-xl font-bold">{$t("shop.return.title")}</p>
					<p class="text-sm text-gray-500 mt-2">
						{$t("shop.return.description")}
					</p>
				</div>

				<div class="rounded-md border p-6 my-4">
					<p class="text-xl font-bold">
						{$t("shop.return.damage.title")}
					</p>
					<p class="text-sm text-gray-500 mt-2">
						{$t("shop.return.damage.description")}
					</p>

					<div class="flex flex-row gap-2 my-6">
						<span
							class="flex items-center justify-start"
							on:click={damageChanged}
						>
							<Checkbox checked={damage} />
							<p class="text-sm text-gray-500 mx-2 cursor-pointer">
								{$t("shop.return.damage.checkbox")}
							</p>
						</span>
					</div>

					<div class="flex flex-col gap-2">
						<p class="text-sm font-bold">
							{$t("shop.return.damage.description.title")}
						</p>
						<Textarea.Root
							bind:value={damageDescription}
							disabled={!damage}
							placeholder={$t(
								"shop.return.damage.description.placeholder",
							)}
						/>
					</div>
				</div>

				<div class="rounded-md border p-6 my-4">
					<p class="text-xl font-bold">
						{$t("shop.return.images.title")}
					</p>
					<p class="text-sm text-gray-500 mt-2">
						{$t("shop.return.images.description")}
					</p>

					<div class="pt-4">
						<Input
							type="file"
							multiple
							accept="image/*"
							on:change={handleImageUpload}
							disabled={$isSubmitting}
						/>
					</div>

					{#if returnImages && returnImages.length > 0}
						<div
							class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-2"
						>
							{#each returnImages as image}
								<div class="relative border rounded-md p-2">
									<img
										src={`data:image/*;base64,${image.data}`}
										alt={image.name}
										class="w-full h-24 object-cover rounded"
									/>
									<p class="text-xs truncate mt-1">
										{image.name}
									</p>
									<button
										type="button"
										class="absolute top-1 right-1 bg-destructive text-destructive-foreground rounded-full w-5 h-5 flex items-center justify-center"
										on:click={() => removeImage(image.id)}
										>×</button
									>
								</div>
							{/each}
						</div>
					{/if}
				</div>
				<div class="rounded-md border p-6 my-4">
					<p class="text-xl font-bold">
						{$t("shop.return.review.title")}
					</p>
					<p class="text-sm text-gray-500 mt-2">
						{$t("shop.return.review.description")}
					</p>

					<div class="flex flex-col gap-4 mt-4">
						<div class="flex flex-row items-center justify-center gap-2">
							<span class="text-sm text-gray-500">{$t("shop.return.review.rating.very-bad")}</span>
							<StarRating
								bind:value={reviewStars}
								on:change={(e) =>
									(reviewStars = e.detail.value)}
							/>
							<span class="text-sm text-gray-500">{$t("shop.return.review.rating.very-good")}</span>
						</div>

						<div class="flex flex-col gap-2">
							<p class="text-sm font-bold">
								{$t("shop.return.review.description.title")}
							</p>
							<Textarea.Root
								bind:value={reviewDescription}
								placeholder={$t(
									"shop.return.review.description.placeholder",
								)}
							/>
						</div>
					</div>
				</div>
			{:else}
				<div class="flex justify-center items-center h-full">
					<Skeleton class="w-full h-full" />
				</div>
			{/if}
		</div>

		<div class="col-span-3">
			{#if rentContract && product}
				<div class="rounded-md border p-6 mb-4">
					<img
						src={"data:image/*;base64," +
							product?.images?.[0]?.data}
						alt={product?.name}
						class="h-full h-full object-cover rounded-md"
					/>

					<div class="flex flex-col gap-4">
						<p class="text-xl font-bold mt-4">{product?.name}</p>
						<p
							class="text-sm text-gray-500 cursor-pointer"
							on:click={() => (inspected = !inspected)}
						>
							<Checkbox checked={inspected} />
							{$t("shop.return.damage.inspected")}
							<span class="text-red-600">*</span>
						</p>
						<Button
							variant="default"
							type="submit"
							on:click={handleReturn}
							disabled={$isSubmitting || !inspected}
						>
							{$t("shop.return.confirm")}
						</Button>
						<Button
							variant="destructive"
							type="submit"
							disabled={$isSubmitting}
						>
							{$t("shop.return.cancel")}
						</Button>
						<Button
							variant="outline"
							type="submit"
							disabled={$isSubmitting}
						>
							{$t("shop.return.contact")}
						</Button>
					</div>
				</div>
			{:else}
				<div class="flex flex-col gap-4 justify-center items-center">
					<Skeleton class="w-full h-32" />
				</div>
			{/if}
		</div>
	</div>
</div>
