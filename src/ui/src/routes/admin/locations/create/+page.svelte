<script lang="ts">
    import { t } from "$lib/i18n";
    import * as Card from "$lib/components/ui/card/index.js";
    import * as Form from "$lib/components/ui/form/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { createLocation } from "$lib/states/location/effects";
    import { goto } from "$app/navigation";
    import { ROUTES } from "$lib/routes";
    import { superForm } from "sveltekit-superforms";
    import { writable } from "svelte/store";
    import Alert from "../../../../components/alert/Alert.svelte";

    // Location form data
    const locationData = {
        id: "",
        city: "",
        street: "",
        postalCode: "",
        buildingName: "",
        coordinates: {
            latitude: 0,
            longitude: 0,
        },
        notes: "",
    };

    // Initialize superForm
    const locationForm = superForm(locationData, {
        dataType: "json",
    });
    const { form, enhance } = locationForm;

    // State management
    const errors = writable<string[]>([]);
    const isSubmitting = writable(false);

    // Handle form submission
    async function handleSubmit() {
        $isSubmitting = true;
        errors.set([]);

        try {
            const result = await createLocation({
                id: $form.id,
                city: $form.city,
                street: $form.street,
                postalCode: $form.postalCode,
                buildingName: $form.buildingName,
                coordinates: {
                    latitude: parseFloat($form.coordinates.latitude.toString()),
                    longitude: parseFloat(
                        $form.coordinates.longitude.toString(),
                    ),
                },
                notes: $form.notes,
            });

            if (result?.errorMessages) {
                errors.set(
                    result.errorMessages?.flatMap((msg) =>
                        msg.message ? [msg.message] : [],
                    ),
                );
            } else {
                goto(ROUTES.ADMIN_LOCATIONS);
            }
        } catch (error) {
            console.error("Error creating location:", error);
            errors.set(["An unexpected error occurred"]);
        } finally {
            $isSubmitting = false;
        }
    }
</script>

<div class="flex flex-col gap-4">
    <Alert
        messages={$errors}
        title={$t("admin.locations.create.form-error")}
        variant="destructive"
    />

    <Card.Root>
        <form on:submit|preventDefault={handleSubmit}>
            <Card.Header>
                <Card.Title
                    >{$t("admin.locations.create.form-title")}</Card.Title
                >
                <Card.Description
                    >{$t(
                        "admin.locations.create.form-description",
                    )}</Card.Description
                >
            </Card.Header>

            <Card.Content class="space-y-4">
                <Form.Field form={locationForm} name="id">
                    <Form.Control let:attrs>
                        <Form.Label>
                            <span class="font-bold"
                                >{$t("admin.locations.create.form-id")}</span
                            >
                        </Form.Label>
                        <Input
                            {...attrs}
                            bind:value={$form.id}
                            required
                            disabled={$isSubmitting}
                            placeholder={$t("admin.locations.create.form-id")}
                        />
                        <Form.FieldErrors
                            class="text-xs text-destructive mt-1"
                        />
                        <Form.Description
                            class="text-xs text-muted-foreground mt-1"
                        >
                            {$t("admin.locations.create.form-id-description")}
                        </Form.Description>
                    </Form.Control>
                </Form.Field>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <Form.Field form={locationForm} name="city">
                        <Form.Control let:attrs>
                            <Form.Label>
                                <span class="font-bold"
                                    >{$t(
                                        "admin.locations.create.form-city",
                                    )}</span
                                >
                            </Form.Label>
                            <Input
                                {...attrs}
                                bind:value={$form.city}
                                required
                                disabled={$isSubmitting}
                                placeholder={$t("admin.locations.table.city")}
                            />
                            <Form.FieldErrors
                                class="text-xs text-destructive mt-1"
                            />
                        </Form.Control>
                    </Form.Field>

                    <Form.Field form={locationForm} name="street">
                        <Form.Control let:attrs>
                            <Form.Label>
                                <span class="font-bold"
                                    >{$t(
                                        "admin.locations.create.form-street",
                                    )}</span
                                >
                            </Form.Label>
                            <Input
                                {...attrs}
                                bind:value={$form.street}
                                required
                                disabled={$isSubmitting}
                                placeholder={$t("admin.locations.table.street")}
                            />
                            <Form.FieldErrors
                                class="text-xs text-destructive mt-1"
                            />
                        </Form.Control>
                    </Form.Field>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <Form.Field form={locationForm} name="postalCode">
                        <Form.Control let:attrs>
                            <Form.Label>
                                <span class="font-bold"
                                    >{$t(
                                        "admin.locations.create.form-postalCode",
                                    )}</span
                                >
                            </Form.Label>
                            <Input
                                {...attrs}
                                bind:value={$form.postalCode}
                                required
                                disabled={$isSubmitting}
                                placeholder={$t(
                                    "admin.locations.table.postalCode",
                                )}
                            />
                            <Form.FieldErrors
                                class="text-xs text-destructive mt-1"
                            />
                        </Form.Control>
                    </Form.Field>

                    <Form.Field form={locationForm} name="buildingName">
                        <Form.Control let:attrs>
                            <Form.Label>
                                <span class="font-bold"
                                    >{$t(
                                        "admin.locations.create.form-buildingName",
                                    )}</span
                                >
                            </Form.Label>
                            <Input
                                {...attrs}
                                bind:value={$form.buildingName}
                                disabled={$isSubmitting}
                                placeholder={$t(
                                    "admin.locations.table.buildingName",
                                )}
                            />
                            <Form.FieldErrors
                                class="text-xs text-destructive mt-1"
                            />
                        </Form.Control>
                    </Form.Field>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <Form.Field form={locationForm} name="coordinates.latitude">
                        <Form.Control let:attrs>
                            <Form.Label>
                                <span class="font-bold"
                                    >{$t(
                                        "admin.locations.create.form-latitude",
                                    )}</span
                                >
                            </Form.Label>
                            <Input
                                {...attrs}
                                type="number"
                                step="0.000001"
                                min="-90"
                                max="90"
                                bind:value={$form.coordinates.latitude}
                                required
                                disabled={$isSubmitting}
                            />
                            <Form.FieldErrors
                                class="text-xs text-destructive mt-1"
                            />
                        </Form.Control>
                    </Form.Field>

                    <Form.Field
                        form={locationForm}
                        name="coordinates.longitude"
                    >
                        <Form.Control let:attrs>
                            <Form.Label>
                                <span class="font-bold"
                                    >{$t(
                                        "admin.locations.create.form-longitude",
                                    )}</span
                                >
                            </Form.Label>
                            <Input
                                {...attrs}
                                type="number"
                                step="0.000001"
                                min="-180"
                                max="180"
                                bind:value={$form.coordinates.longitude}
                                required
                                disabled={$isSubmitting}
                            />
                            <Form.FieldErrors
                                class="text-xs text-destructive mt-1"
                            />
                        </Form.Control>
                    </Form.Field>
                </div>

                <Form.Field form={locationForm} name="notes">
                    <Form.Control let:attrs>
                        <Form.Label>
                            <span class="font-bold"
                                >{$t("admin.locations.create.form-notes")}</span
                            >
                        </Form.Label>
                        <textarea
                            class="flex w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 min-h-[100px]"
                            {...attrs}
                            bind:value={$form.notes}
                            disabled={$isSubmitting}
                        />
                        <Form.FieldErrors
                            class="text-xs text-destructive mt-1"
                        />
                    </Form.Control>
                </Form.Field>
            </Card.Content>

            <Card.Footer class="flex justify-start gap-2">
                <Button type="submit" disabled={$isSubmitting}>
                    {$t("admin.locations.create.form-submit")}
                </Button>
                <Button
                    variant="outline"
                    on:click={() => goto(ROUTES.ADMIN_LOCATIONS)}
                >
                    {$t("admin.locations.create.form-cancel")}
                </Button>
            </Card.Footer>
        </form>
    </Card.Root>
</div>
