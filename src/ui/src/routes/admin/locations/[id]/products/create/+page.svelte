<script lang="ts">
    import { t } from "$lib/i18n";
    import * as Card from "$lib/components/ui/card/index.js";
    import * as Form from "$lib/components/ui/form/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { createProduct } from "$lib/states/product/effects";
    import { goto } from "$app/navigation";
    import { ROUTES } from "$lib/routes";
    import { superForm } from "sveltekit-superforms";
    import { writable } from "svelte/store";
    import Alert from "../../../../../../components/alert/Alert.svelte";
    import { page } from "$app/stores";
    import { formatStringToCurrency } from "$lib/helpers/price";
    // Get the location ID from the URL
    const locationId = $page.params.id;

    // Product form data
    const productData = {
        id: "",
        name: "",
        description: "",
        location: locationId,
        pricing: {
            price: 0,
            deposit: 0,
        },
        images: [] as Array<{ id: string; name: string; data: string }>,
        documents: [] as Array<{ id: string; name: string; data: string }>,
        dynamicAttributes: {},
    };

    // Initialize superForm
    const productForm = superForm(productData, {
        dataType: "json",
    });
    const { form } = productForm;

    // State management
    const errors = writable<string[]>([]);
    const isSubmitting = writable(false);

    // Handle file uploads

    async function handleImageUpload(event: Event) {
        const input = event.target as HTMLInputElement;
        if (!input.files || input.files.length === 0) return;

        const images = $form.images || [];

        for (let i = 0; i < input.files.length; i++) {
            const file = input.files[i];
            const reader = new FileReader();

            reader.onload = (e) => {
                const result = e.target?.result as string;
                if (result) {
                    const base64Data = result.split(",")[1]; // Remove data URL prefix
                    images.push({
                        id: crypto.randomUUID(),
                        name: file.name,
                        data: base64Data,
                    });
                    $form.images = [...images];
                }
            };

            reader.readAsDataURL(file);
        }

        // Reset input
        input.value = "";
    }

    async function handleDocumentUpload(event: Event) {
        const input = event.target as HTMLInputElement;
        if (!input.files || input.files.length === 0) return;

        const documents = $form.documents || [];

        for (let i = 0; i < input.files.length; i++) {
            const file = input.files[i];
            const reader = new FileReader();

            reader.onload = (e) => {
                const result = e.target?.result as string;
                if (result) {
                    const base64Data = result.split(",")[1]; // Remove data URL prefix
                    documents.push({
                        id: crypto.randomUUID(),
                        name: file.name,
                        data: base64Data,
                    });
                    $form.documents = [...documents];
                }
            };

            reader.readAsDataURL(file);
        }

        // Reset input
        input.value = "";
    }

    function removeImage(id: string) {
        const images = $form.images || [];
        $form.images = images.filter((img) => img.id !== id);
    }

    function removeDocument(id: string) {
        const documents = $form.documents || [];
        $form.documents = documents.filter((doc) => doc.id !== id);
    }

    // Handle form submission
    async function handleSubmit() {
        $isSubmitting = true;
        errors.set([]);

        try {
            const result = await createProduct({
                id: $form.id,
                name: $form.name,
                description: $form.description,
                location: $form.location,
                pricing: {
                    price: formatStringToCurrency(
                        $form.pricing.price.toString(),
                    ),
                    deposit: formatStringToCurrency(
                        $form.pricing.deposit.toString(),
                    ),
                },
                images: $form.images || [],
                documents: $form.documents || [],
                dynamicAttributes: $form.dynamicAttributes || {},
            });

            if (result) {
                goto(`${ROUTES.ADMIN_LOCATIONS}/${locationId}/products`);
            }
        } catch (error: any) {
            console.error("Error creating product:", error);
            if (error?.errorMessages) {
                errors.set(
                    error.errorMessages?.flatMap((msg: any) =>
                        msg.message ? [msg.message] : [],
                    ),
                );
            } else {
                errors.set(["An unexpected error occurred"]);
            }
        } finally {
            $isSubmitting = false;
        }
    }
</script>

<div class="flex flex-col gap-4">
    <Alert
        messages={$errors}
        title={$t("admin.products.create.form-error")}
        variant="destructive"
    />

    <Card.Root>
        <form on:submit|preventDefault={handleSubmit}>
            <Card.Header>
                <Card.Title>{$t("admin.products.create.form-title")}</Card.Title
                >
            </Card.Header>

            <Card.Content class="space-y-4">
                <Form.Field form={productForm} name="name">
                    <Form.Control let:attrs>
                        <Form.Label>
                            <span class="font-bold"
                                >{$t("admin.products.create.form-name")}</span
                            >
                        </Form.Label>
                        <Input
                            {...attrs}
                            bind:value={$form.name}
                            required
                            disabled={$isSubmitting}
                            placeholder={$t("admin.products.table.name")}
                        />
                        <Form.FieldErrors
                            class="text-xs text-destructive mt-1"
                        />
                    </Form.Control>
                </Form.Field>

                <Form.Field form={productForm} name="description">
                    <Form.Control let:attrs>
                        <Form.Label
                            ><span class="font-bold"
                                >{$t(
                                    "admin.products.create.form-description",
                                )}</span
                            ></Form.Label
                        >
                        <textarea
                            class="flex w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 min-h-[100px]"
                            {...attrs}
                            bind:value={$form.description}
                            required
                            disabled={$isSubmitting}
                            placeholder={$t("admin.products.table.description")}
                        />
                        <Form.FieldErrors
                            class="text-xs text-destructive mt-1"
                        />
                    </Form.Control>
                </Form.Field>

                <Form.Field form={productForm} name="pricing.price">
                    <Form.Control let:attrs>
                        <Form.Label
                            ><span class="font-bold"
                                >{$t(
                                    "admin.products.create.form-price",
                                )}</span
                            ></Form.Label
                        >
                        <Input
                            {...attrs}
                            type="number"
                            step="0.01"
                            min="0"
                            bind:value={$form.pricing.price}
                            required
                            disabled={$isSubmitting}
                            placeholder={$t("admin.products.table.price")}
                        />
                        <Form.FieldErrors
                            class="text-xs text-destructive mt-1"
                        />
                        <Form.Description
                            class="text-xs text-muted-foreground mt-1"
                        >
                            {$t("admin.products.create.form-price-description")}
                        </Form.Description>
                    </Form.Control>
                </Form.Field>

                <Form.Field form={productForm} name="pricing.deposit">
                    <Form.Control let:attrs>
                        <Form.Label
                            ><span class="font-bold"
                                >{$t(
                                    "admin.products.create.form-deposit",
                                )}</span
                            ></Form.Label
                        >
                        <Input
                            {...attrs}
                            type="number"
                            step="0.01"
                            min="0"
                            bind:value={$form.pricing.deposit}
                            required
                            disabled={$isSubmitting}
                            placeholder={$t("admin.products.table.deposit")}
                        />
                        <Form.FieldErrors
                            class="text-xs text-destructive mt-1"
                        />
                        <Form.Description
                            class="text-xs text-muted-foreground mt-1"
                        >
                            {$t(
                                "admin.products.create.form-deposit-description",
                            )}
                        </Form.Description>
                    </Form.Control>
                </Form.Field>

                <div class="border rounded-md p-4">
                    <h3 class="text-lg font-medium mb-2">
                        {$t("admin.products.create.form-images")}
                    </h3>
                    <div class="mb-4">
                        <Input
                            type="file"
                            multiple
                            accept="image/*"
                            on:change={handleImageUpload}
                            disabled={$isSubmitting}
                        />
                        <p class="text-xs text-muted-foreground mt-1">
                            {$t(
                                "admin.products.create.form-images-description",
                            )}
                        </p>
                    </div>

                    {#if $form.images && $form.images.length > 0}
                        <div
                            class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-2"
                        >
                            {#each $form.images as image}
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

                <div class="border rounded-md p-4">
                    <h3 class="text-lg font-medium mb-2">
                        {$t("admin.products.create.form-documents")}
                    </h3>
                    <div class="mb-4">
                        <Input
                            type="file"
                            multiple
                            accept=".pdf,.doc,.docx,.txt"
                            on:change={handleDocumentUpload}
                            disabled={$isSubmitting}
                        />
                        <p class="text-xs text-muted-foreground mt-1">
                            {$t(
                                "admin.products.create.form-documents-description",
                            )}
                        </p>
                    </div>

                    {#if $form.documents && $form.documents.length > 0}
                        <div class="grid grid-cols-1 gap-2">
                            {#each $form.documents as document}
                                <div
                                    class="relative border rounded-md p-3 flex items-center"
                                >
                                    <div class="mr-2">
                                        <svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            width="24"
                                            height="24"
                                            viewBox="0 0 24 24"
                                            fill="none"
                                            stroke="currentColor"
                                            stroke-width="2"
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            class="text-muted-foreground"
                                        >
                                            <path
                                                d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"
                                            />
                                            <polyline points="14 2 14 8 20 8" />
                                        </svg>
                                    </div>
                                    <div class="flex-1">
                                        <p class="text-sm font-medium truncate">
                                            {document.name}
                                        </p>
                                    </div>
                                    <button
                                        type="button"
                                        class="ml-2 bg-destructive text-destructive-foreground rounded-full w-5 h-5 flex items-center justify-center"
                                        on:click={() =>
                                            removeDocument(document.id)}
                                        >×</button
                                    >
                                </div>
                            {/each}
                        </div>
                    {/if}
                </div>
            </Card.Content>

            <Card.Footer class="flex justify-start gap-2">
                <Button type="submit" disabled={$isSubmitting}>
                    {$t("admin.products.create.form-submit")}
                </Button>
                <Button
                    variant="outline"
                    on:click={() =>
                        goto(
                            `${ROUTES.ADMIN_LOCATIONS}/${locationId}/products`,
                        )}
                >
                    {$t("admin.products.create.form-cancel")}
                </Button>
            </Card.Footer>
        </form>
    </Card.Root>
</div>
