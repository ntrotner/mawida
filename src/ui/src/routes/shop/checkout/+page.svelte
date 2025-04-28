<script lang="ts">
	import { onMount } from "svelte";
	import { page } from "$app/stores";
	import { rentProduct, fetchProduct } from "$lib/states/product/effects";
	import type { RentProductFormular } from "$lib/open-api/models";
	import { PaymentMethod } from "$lib/open-api/models";
	import { formatCurrency } from "$lib/helpers/price";
	import { fetchRentContractById } from "$lib/states/rent_contracts/effects";
	import { rentContractsState } from "$lib/states/rent_contracts";
	import { buttonVariants } from "$lib/components/ui/button";
	import * as Table from "$lib/components/ui/table";
	import * as Card from "$lib/components/ui/card";
	import * as Dialog from "$lib/components/ui/dialog";
	import { t } from "$lib/i18n";
	import { Skeleton } from "$lib/components/ui/skeleton";
	import { locationState } from "$lib/states/location";
	import { configState } from "$lib/states/config";
	import {
		ProductConfigKey,
		type ProductConfig,
	} from "$lib/states/config/collection/product";
	import { Button } from "$lib/components/ui/button";
	import { Trash2 } from "lucide-svelte";
	import { authenticationState } from "$lib/states/authentication";
	import { Input } from "$lib/components/ui/input";
	import { Label } from "$lib/components/ui/label";
	import { Dialog as DialogPrimitive } from "bits-ui";
	import { ROUTES } from "$lib/routes";
	import { goto } from "$app/navigation";
	import { writable } from "svelte/store";
	import { productState } from "$lib/states/product";

	const authentication = authenticationState.getAsyncState();
	const isAdmin = authenticationState.isAdmin();
	const productConfig =
		configState.getConfig<ProductConfig>(ProductConfigKey);
	const locations = locationState.getAsyncState();
	const products = productState.getAsyncState();
	let product: any = null;
	let loading = true;
	let error: string | null = null;
	let loadingRedirection = writable(false);

	// Form values
	let startDate: string = new Date().toISOString().split("T")[0];
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
			const localProduct = $products?.find((p) => p.id === productId);
			if (!localProduct) {
				product = await fetchProduct(productId);
			} else {
				product = localProduct;
			}
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

	function getProductImage(product: any): string {
		if (product?.images && product.images.length > 0) {
			return `data:image/*;base64,${product.images[0].data}`;
		}
		return "";
	}

	function getLocationCity(locationId: string | undefined): string {
		if (!locationId || !$locations) return "-";

		const location = $locations.find((loc) => loc.id === locationId);
		return location?.city || "-";
	}

	function getCategoryName(categoryId: string | undefined): string {
		if (!categoryId) return $t("product.category.undefined");

		const category = $productConfig?.categories?.find(
			(c) => c.id === categoryId,
		);
		if (category) {
			return $t(category.translationKey);
		}

		if (categoryId === $productConfig?.categoryFallback.id) {
			return $t($productConfig.categoryFallback.translationKey);
		}

		return $t("product.category.undefined");
	}

	async function handleCheckout() {
		if (!product) return;
		loadingRedirection.set(true);

		const startTimestamp = convertToTimestamp(startDate);
		const endTimestamp = convertToTimestamp(endDate);

		if (
			!startTimestamp ||
			!endTimestamp ||
			startTimestamp >= endTimestamp
		) {
			error = "Please select valid start and end dates";
			loadingRedirection.set(false);
			return;
		}

		if (!selectedLocationId) {
			error = "Please select a location";
			loadingRedirection.set(false);
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
					const redirectionUrl =
						rentContract.paymentInstructions?.dynamicAttributes
							?.url;
					if (redirectionUrl) {
						window.location.href = redirectionUrl;
					} else {
						loadingRedirection.set(false);
					}
				} else {
					loadingRedirection.set(false);
				}
			} else {
				loadingRedirection.set(false);
			}
		} catch (err: any) {
			console.error("Checkout failed:", err);
			error =
				err?.errorMessages?.[0]?.message ||
				"Failed to process rental request";
			loadingRedirection.set(false);
		}
	}
</script>

<div class="container mx-auto p-4">
	<div class="grid grid-cols-10 gap-4">
		<div class="col-span-7">
			{#if product}
				<div class="rounded-md border">
					<Table.Root>
						<Table.Header>
							<Table.Row>
								<Table.Head class="min-w-24"></Table.Head>
								<Table.Head
									>{$t(
										"shop.products.table.name",
									)}</Table.Head
								>
								<Table.Head
									>{$t(
										"shop.products.table.status",
									)}</Table.Head
								>
								<Table.Head
									>{$t(
										"shop.products.table.location",
									)}</Table.Head
								>
								<Table.Head
									>{$t(
										"shop.products.table.price",
									)}</Table.Head
								>
								<Table.Head
									>{$t(
										"shop.products.table.actions",
									)}</Table.Head
								>
							</Table.Row>
						</Table.Header>
						<Table.Body>
							<Table.Row>
								<Table.Cell class="w-16 h-16 p-2">
									<img
										src={getProductImage(product)}
										alt={product.name}
										class="w-16 h-16 object-cover rounded"
										style="cursor: pointer;"
									/>
								</Table.Cell>
								<Table.Cell>
									<div class="flex flex-col">
										<span
											class="text-md cursor-pointer hover:underline"
										>
											{product.name || "-"}
										</span>
										<span
											class="text-sm text-gray-500 pt-1"
										>
											{#if product.dynamicAttributes?.["category"]}
												{getCategoryName(
													product.dynamicAttributes?.[
														"category"
													],
												)}
											{:else}
												{$t(
													"product.category.undefined",
												)}
											{/if}
										</span>
									</div>
								</Table.Cell>
								<Table.Cell>
									<span
										class={`px-2 py-1 rounded-md text-xs font-medium border border-gray-150`}
									>
										{product.isRented
											? $t("shop.products.table.rented")
											: $t(
													"shop.products.table.available",
												)}
									</span>
								</Table.Cell>
								<Table.Cell>
									{getLocationCity(product.location)}
								</Table.Cell>
								<Table.Cell>
									<div class="flex flex-col">
										<span
											>{formatCurrency(
												product.pricing?.price || 0,
												"€",
											)}</span
										>
									</div>
								</Table.Cell>
								<Table.Cell>
									<Button
										disabled={$loadingRedirection}
										variant="ghost"
										on:click={() => {
											goto(
												ROUTES.SHOP_PRODUCT.replace(
													"{productId}",
													product.id,
												),
											);
										}}
									>
										<Trash2 class="w-4 h-4" />
									</Button>
								</Table.Cell>
							</Table.Row>
						</Table.Body>
					</Table.Root>
				</div>
			{:else}
				<div class="flex justify-center items-center h-full">
					<Skeleton class="w-full h-full" />
				</div>
			{/if}
		</div>

		<div class="col-span-3">
			{#if !product}
				<div class="flex flex-col gap-4 justify-center items-center">
					<Skeleton class="w-full h-32" />
					<Skeleton class="w-full h-32" />
				</div>
			{:else}
				<Card.Root>
					<Card.Header>
						<Card.Title>
							<span class="text-md font-semibold">
								{$t("product.shop.price")}
							</span>
						</Card.Title>
					</Card.Header>
					<Card.Content>
						<span class="text-4xl font-bold">
							{formatCurrency(product.pricing.price, "€")}*
							<span class="text-sm font-normal text-gray-500"
								>/{$t("product.shop.per-day")}</span
							>
							<p class="text-sm font-bold">
								+{formatCurrency(product.pricing.deposit, "€")}
								{$t("product.shop.price.deposit")}
							</p>
							<p class="text-sm font-normal text-gray-500">
								*{$t("product.shop.price.note")}
							</p>
						</span>

						{#if $authentication?.authenticated && !$isAdmin && !product.isRented}
							<div class="flex flex-col gap-4 mt-4">
								<Dialog.Root>
									<Dialog.Trigger
										disabled={$loadingRedirection}
										class={buttonVariants({
											variant: "default",
										})}
									>
										{$t("product.shop.book-now")}
									</Dialog.Trigger>
									<Dialog.Content>
										<Dialog.Header>
											<Dialog.Title
												>{$t(
													"product.checkout.dialog.title",
												)}</Dialog.Title
											>
											<Dialog.Description>
												{$t(
													"product.checkout.dialog.description",
												)}
											</Dialog.Description>
										</Dialog.Header>
										<div class="inline-flex flex-row justify-evenly gap-4 py-4">
											<div>
												<Label>
													{$t(
														"product.checkout.dialog.start-date",
													)}
												</Label>
												<Input
													type="date"
													disabled={true}
													bind:value={startDate}
												/>
											</div>
											<div>
												<Label>
													{$t(
														"product.checkout.dialog.end-date",
													)}
												</Label>
												<Input
													type="date"
													bind:value={endDate}
												/>
											</div>
										</div>
										<Dialog.Footer>
											<DialogPrimitive.Close>
												<Button
													type="submit"
													disabled={$loadingRedirection}
													on:click={() =>
														handleCheckout()}
												>
													{$t(
														"product.checkout.dialog.confirm-button",
													)}
												</Button>
												<Button variant="outline">
													{$t("common.cancel")}
												</Button>
											</DialogPrimitive.Close>
										</Dialog.Footer>
									</Dialog.Content>
								</Dialog.Root>

								<Button
									disabled={$loadingRedirection}
									on:click={() => {
										goto(
											ROUTES.SHOP_PRODUCT.replace(
												"{productId}",
												product.id,
											),
										);
									}}
									variant="destructive"
								>
									{$t("product.shop.clean-checkout")}
								</Button>
							</div>
						{/if}
					</Card.Content>
				</Card.Root>
			{/if}
		</div>
	</div>
</div>
