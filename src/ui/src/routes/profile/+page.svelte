<script lang="ts">
  import { superForm } from "sveltekit-superforms";
  import { t } from "$lib/i18n";
  import { Separator } from "$lib/components/ui/separator/index";
  import { Input } from "$lib/components/ui/input/index";
  import * as Form from "$lib/components/ui/form/index";
  import {
    changeEmailOfUser,
    changePasswordOfUser,
    userState,
  } from "$lib/states/user";
  import { map } from "rxjs";
  import Alert from "../../components/alert/Alert.svelte";
  import { writable } from "svelte/store";
  import { toast } from "svelte-sonner";

  // Error stores
  const errorsEmail = writable<string[]>([]);
  const errorsPassword = writable<string[]>([]);
  
  // Form data initial values
  const emailInitialData = { value: "", repeated: "" };
  const passwordInitialData = { old: "", value: "", repeated: "" };

  // Initialize superForms
  const emailForm = superForm(emailInitialData);
  const passwordForm = superForm(passwordInitialData);
  const { form: emailFormData } = emailForm;
  const { form: passwordFormData } = passwordForm;
  
  // User email from state
  const userStateEmail = userState
    .observable()
    .pipe(map((state) => state?.email));

  /**
   * Handle email change submission
   */
  async function submitChangeEmail() {
    // Validate form data
    if (
      $emailFormData.value === "" ||
      $emailFormData.value !== $emailFormData.repeated
    ) {
      $errorsEmail = ["Please check your input"];
      return;
    }

    // Submit change request
    const response = await changeEmailOfUser(
      $userStateEmail,
      $emailFormData.value
    );
    
    // Process any error messages
    $errorsEmail = (response?.errorMessages
      ?.map(({ message }) => message)
      .filter((message) => typeof message === "string" && message.length > 0) ||
      []) as string[];

    if ($errorsEmail.length > 0) {
      return;
    }
    
    // Show success message and reset form
    toast.success($t("profile.email-change-successful"));
    $emailFormData = { value: "", repeated: "" };
    $errorsEmail = [];
  }

  /**
   * Handle password change submission
   */
  async function submitChangePassword() {
    // Validate form data
    if (
      $passwordFormData.value === "" ||
      $passwordFormData.old === "" ||
      $passwordFormData.value !== $passwordFormData.repeated
    ) {
      $errorsPassword = ["Please check your input"];
      return;
    }
    
    // Submit change request
    const response = await changePasswordOfUser(
      $passwordFormData.old,
      $passwordFormData.value
    );
    
    // Process any error messages
    $errorsPassword = (response?.errorMessages
      ?.map(({ message }) => message)
      .filter((message) => typeof message === "string" && message.length > 0) ||
      []) as string[];

    if ($errorsPassword.length > 0) {
      return;
    }
    
    // Show success message and reset form
    toast.success($t("profile.password-change-successful"));
    $passwordFormData = { value: "", old: "", repeated: "" };
    $errorsPassword = [];
  }
</script>

<div class="wrapper">
  <div class="header">
    <h3 class="profile-title text-lg font-medium font-bold">
      {$t("profile.profile")}
    </h3>
    <p class="description text-sm text-muted-foreground">
      {$t("profile.profile-description")}
    </p>
  </div>
  
  <Separator />

  <div class="profile-form-wrapper">
    <h3 class="text-md font-bold">{$t("profile.email")}</h3>
    <form
      on:submit|preventDefault={submitChangeEmail}
      class="space-y-8"
      id="email-form"
    >
      <Form.Field form={emailForm} name="value">
        <Form.Control let:attrs>
          <Form.Label>{$t("profile.email-new")}</Form.Label>
          <Input
            placeholder={$userStateEmail}
            {...attrs}
            class="!mb-2"
            bind:value={$emailFormData.value}
          />
          <Form.Label>{$t("profile.email-repeat")}</Form.Label>
          <Input
            {...attrs}
            class="!mb-2"
            bind:value={$emailFormData.repeated}
          />
        </Form.Control>
        <Alert
          messages={$errorsEmail}
          title={$t("login.input-error-title")}
          variant="destructive"
        ></Alert>
      </Form.Field>
      <Form.Button class="!mt-2 !mb-6"
        >{$t("profile.email-update-button")}</Form.Button
      >
    </form>

    <Separator />

    <h3 class="text-md font-bold !mt-6">{$t("profile.password")}</h3>
    <form
      on:submit|preventDefault={submitChangePassword}
      class="space-y-8"
      id="password-form"
    >
      <Form.Field form={passwordForm} name="value">
        <Form.Control let:attrs>
          <Form.Label>{$t("profile.password-old")}</Form.Label>
          <Input
            {...attrs}
            type="password"
            class="!mb-2"
            bind:value={$passwordFormData.old}
          />
          <Form.Label>{$t("profile.password-new")}</Form.Label>
          <Input
            {...attrs}
            type="password"
            class="!mb-2"
            bind:value={$passwordFormData.value}
          />
          <Form.Label>{$t("profile.password-repeat")}</Form.Label>
          <Input
            {...attrs}
            type="password"
            class="!mb-2"
            bind:value={$passwordFormData.repeated}
          />
        </Form.Control>
        <Alert
          messages={$errorsPassword}
          title={$t("login.input-error-title")}
          variant="destructive"
        ></Alert>
      </Form.Field>
      <Form.Button class="!mt-2 !mb-6"
        >{$t("profile.password-update-button")}</Form.Button
      >
    </form>
  </div>
</div>

<style>
  .wrapper {
    margin: 1.8rem auto;
    padding: 0 2.5rem;
    max-width: 40rem;
  }

  .header .description {
    margin: 0.5rem 0;
  }

  .profile-form-wrapper {
    padding: 1rem 0;
    width: 95%;
    max-width: 70rem;
    margin: 0 auto;
  }

  @media only screen and (max-width: 640px) {
    .wrapper {
      margin: 1rem auto;
      padding: 0 2rem;
    }
  }
</style>
