<script lang="ts">
  import * as Tabs from "$lib/components/ui/tabs/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import * as Form from "$lib/components/ui/form/index.js";
  import { Button } from "$lib/components/ui/button/index.js";
  import { Input } from "$lib/components/ui/input/index.js";
  import { t } from "$lib/i18n";
  import { appState } from "$lib/states/app";
  import { login, register } from "$lib/states/authentication";
  import { type ModelError, type Success } from "$lib/open-api";
  import { goto } from "$app/navigation";
  import { ROUTES } from "$lib/routes";
  import { writable, type Unsubscriber } from "svelte/store";
  import { onDestroy, onMount, afterUpdate } from "svelte";
  import Alert from "../../components/alert/Alert.svelte";
  import { map } from "rxjs";
  import { superForm } from "sveltekit-superforms";
  import { GithubLogo } from "svelte-radix";
  import SlimNavigator from "../../components/navigator/SlimNavigator.svelte";
  import loginBackground from '$lib/assets/login-background.png';
    import { userState } from "$lib/states/user";

  // Reactive state management
  const errors = writable<string[]>([]);
  const requestInProcess = writable(false);

  // Form data stores
  const loginData = { username: "", password: "" };
  const registerData = { username: "", password: "", repeatPassword: "" };
  $: currentTab = "login";

  // Initialize superForms
  const loginForm = superForm(loginData);
  const registerForm = superForm(registerData);
  const { form: loginFormData } = loginForm;
  const { form: registerFormData } = registerForm;

  // Responsive state
  const mobile = appState
    .observable()
    .pipe(map((state) => (state?.width || 0) <= 640));

  function isValid(input: string) {
    return input.trim().length !== 0;
  }

  function fillErrors(errorInput: (Success & ModelError) | undefined) {
    $errors = [];
    if ((errorInput as ModelError)?.errorMessages) {
      $errors =
        ((errorInput as ModelError).errorMessages
          ?.map(({ message }) => message)
          .filter((message) => typeof message === "string") as string[]) || [];
    }
  }

  function isError(input: (Success & ModelError) | undefined) {
    return !!(input as ModelError)?.errorMessages;
  }

  async function submitLogin() {
    requestInProcess.set(true);
    try {
      const { username, password } = $loginFormData;

      if (!isValid(username) || !isValid(password)) {
        fillErrors({
          errorMessages: [{ code: "200", message: "User or Password Invalid" }],
        });
        throw new Error("");
      }

      const response = await login(username, password);

      if (!response) {
        fillErrors({
          errorMessages: [{ code: "100", message: "Server Failure" }],
        });
      } else if (isError(response)) {
        fillErrors(response);
      } else {
        const user = await userState.getSyncState();
        if (user?.role === "admin") {
          goto(ROUTES.ADMIN);
        } else if (user?.role === "user") {
          goto(ROUTES.CUSTOMER_PROFILE);
        } else {
          goto(ROUTES.HOME);
        }
      }
    } finally {
      requestInProcess.set(false);
    }
  }

  async function submitRegister() {
    requestInProcess.set(true);
    try {
      const { username, password, repeatPassword } = $registerFormData;

      if (!isValid(username) || !isValid(password)) {
        fillErrors({
          errorMessages: [{ code: "200", message: "User or Password Invalid" }],
        });
        throw new Error("");
      }

      if (password !== repeatPassword) {
        fillErrors({
          errorMessages: [{ code: "200", message: "Passwords don't match" }],
        });
        throw new Error("");
      }

      const response = await register(username, password);

      if (!response) {
        fillErrors({
          errorMessages: [{ code: "100", message: "Server Failure" }],
        });
      } else if (isError(response)) {
        fillErrors(response);
      } else {
        goto(ROUTES.HOME);
      }
    } finally {
      requestInProcess.set(false);
    }
  }

  function switchToRegister() {
    currentTab = 'register';
  }

  function switchToLogin() {
    currentTab = 'login';
  }

  function clear() {
    $errors = [];
    $loginFormData = { username: "", password: "" };
    $registerFormData = { username: "", password: "", repeatPassword: "" };
  }

  onMount(() => {
    // Background setup
    // @ts-ignore
    document.firstElementChild.style.backgroundImage = `url(${loginBackground})`;
    // @ts-ignore
    document.firstElementChild.style.backgroundSize = 'cover';
    // @ts-ignore
    document.firstElementChild.style.backgroundPosition = 'center';
    // @ts-ignore
    document.firstElementChild.style.backgroundRepeat = 'no-repeat';
    // @ts-ignore
    document.firstElementChild.style.backgroundAttachment = 'fixed';
  });


  onDestroy(() => {
    // @ts-ignore
    document.firstElementChild.style.backgroundImage = '';
    // @ts-ignore
    document.firstElementChild.style.backgroundSize = '';
    // @ts-ignore
    document.firstElementChild.style.backgroundPosition = '';
    // @ts-ignore
    document.firstElementChild.style.backgroundRepeat = '';
    // @ts-ignore
    document.firstElementChild.style.backgroundAttachment = '';
  });
</script>

<div class="mt-[12lvh] flex flex-row justify-center items-start">
  <div>
    <Tabs.Root onValueChange={clear} value="{currentTab}" class="w-[450px]">
      <div class="hidden">
        <Tabs.List class="grid w-full grid-cols-2 mb-4">
          <Tabs.Trigger value="login" disabled={$requestInProcess}>
            {$t("login.login-title")}
          </Tabs.Trigger>
          <Tabs.Trigger value="register" disabled={$requestInProcess}>
            {$t("login.register-title")}
          </Tabs.Trigger>
        </Tabs.List>
      </div>

      <Tabs.Content class="m-0" value="login">
        <Card.Root>
          <Alert
            messages={$errors}
            title={$t("login.input-error-title")}
            variant="destructive"
          />
          <form class="px-6" on:submit|preventDefault={submitLogin}>
            <Card.Content class="space-y-2 mt-5">
              <div
                class="flex flex-col justify-center items-center flex-nowrap pt-4 pb-6"
              >
                <p class="font-bold pb-4 text-2xl">
                  {$t("login.welcome-login")}
                </p>
                <p>
                  {$t("login.welcome-login.description")}
                </p>
              </div>

              <Form.Field form={loginForm} name="username">
                <Form.Control let:attrs>
                  <Form.Label>
                    <p class="font-semibold text-sm">
                      {$t("login.username")}
                    </p></Form.Label
                  >
                  <Input
                    {...attrs}
                    type="email"
                    bind:value={$loginFormData.username}
                    disabled={$requestInProcess}
                  />
                </Form.Control>
              </Form.Field>

              <Form.Field form={loginForm} name="password">
                <Form.Control let:attrs>
                  <Form.Label
                    class="mt-6 flex flex-row flex-nowrap justify-between align-end"
                  >
                    <p class="font-semibold text-sm">
                      {$t("login.password")}
                    </p>
                    <div class="text-md">
                      {$t("login.reset-password")}
                    </div>
                  </Form.Label>
                  <Input
                    {...attrs}
                    type="password"
                    bind:value={$loginFormData.password}
                    disabled={$requestInProcess}
                  />
                </Form.Control>
              </Form.Field>
            </Card.Content>

            <Card.Footer>
              <Button type="submit" class="w-full" disabled={$requestInProcess}>
                {$t("login.login-button")}
              </Button>
            </Card.Footer>
          </form>
          <div class="relative w-[80%] justify-self-center">
            <div class="absolute inset-0 flex items-center">
              <span class="w-full border-t" />
            </div>
            <div class="relative flex justify-center text-xs uppercase">
              <span class="bg-background text-muted-foreground px-2">
                {$t("login.login-with")}
              </span>
            </div>
          </div>

          <div class="py-6 px-8 flex flex-row justify-between">
            <Button class="py-6 px-10" variant="outline" type="button" disabled={$requestInProcess}>
              <svg
                class="h-6 w-6"
                fill="#000000"
                version="1.1"
                viewBox="0 0 22.773 22.773"
                xml:space="preserve"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="m15.769 0h0.162c0.13 1.606-0.483 2.806-1.228 3.675-0.731 0.863-1.732 1.7-3.351 1.573-0.108-1.583 0.506-2.694 1.25-3.561 0.69-0.808 1.955-1.527 3.167-1.687z"
                />
                <path
                  d="m20.67 16.716v0.045c-0.455 1.378-1.104 2.559-1.896 3.655-0.723 0.995-1.609 2.334-3.191 2.334-1.367 0-2.275-0.879-3.676-0.903-1.482-0.024-2.297 0.735-3.652 0.926h-0.462c-0.995-0.144-1.798-0.932-2.383-1.642-1.725-2.098-3.058-4.808-3.306-8.276v-1.019c0.105-2.482 1.311-4.5 2.914-5.478 0.846-0.52 2.009-0.963 3.304-0.765 0.555 0.086 1.122 0.276 1.619 0.464 0.471 0.181 1.06 0.502 1.618 0.485 0.378-0.011 0.754-0.208 1.135-0.347 1.116-0.403 2.21-0.865 3.652-0.648 1.733 0.262 2.963 1.032 3.723 2.22-1.466 0.933-2.625 2.339-2.427 4.74 0.176 2.181 1.444 3.457 3.028 4.209z"
                />
              </svg>
            </Button>
            <Button class="py-6 px-10" variant="outline" type="button" disabled={$requestInProcess}>
              <svg class="h-6 w-6" version="1.1" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <title>google [#178]</title>
                <g fill="none" fill-rule="evenodd">
                <g transform="translate(-300 -7399)" fill="#000">
                <g transform="translate(56 160)">
                <path d="m263.82 7247h-9.6102c0 0.99944 0 2.9983-0.006126 3.9978h5.5689c-0.2134 0.99944-0.97001 2.3986-2.0391 3.1033-0.001021-1e-3 -0.002042 0.00599-0.004084 0.00499-1.4213 0.93848-3.297 1.1514-4.6897 0.87151-2.183-0.43375-3.9107-2.0169-4.6121-4.0277 0.004084-3e-3 0.007147-0.03098 0.01021-0.03298-0.43906-1.2473-0.43906-2.9174 0-3.9168h-0.001021c0.56567-1.837 2.3454-3.513 4.5315-3.9718 1.7583-0.37279 3.7422 0.03098 5.2013 1.3962 0.194-0.18989 2.6854-2.6225 2.8722-2.8204-4.9848-4.5145-12.966-2.9264-15.953 2.9034h-0.001021s0.001021 0-0.005105 0.01099c-1.4775 2.8634-1.4162 6.2375 0.01021 8.964-0.004084 0.00299-0.007147 0.00499-0.01021 0.00799 1.2927 2.5086 3.6452 4.4325 6.4797 5.1651 3.0111 0.78956 6.8432 0.24986 9.4101-2.0718l0.003063 3e-3c2.1749-1.9589 3.5288-5.296 2.8447-9.5866"></path>
                </g>
                </g>
                </g>
              </svg>
            </Button>
            <Button class="py-6 px-10" variant="outline" type="button" disabled={$requestInProcess}>
              <GithubLogo class="h-6 w-6" />
            </Button>
          </div>

          <div class="pb-4 text-sm justify-self-center">
            <p>
              {$t('login.prompt-to-register')}
              <Button class="underline" on:click="{() => switchToRegister()}" variant="link">{$t('login.redirect-to-register')}</Button>
            </p>
          </div>
        </Card.Root>
      </Tabs.Content>

      <Tabs.Content class="m-0" value="register">
        <Card.Root>
          <Alert
            messages={$errors}
            title={$t("login.input-error-title")}
            variant="destructive"
          />
          <form class="px-6" on:submit|preventDefault={submitRegister}>
            <Card.Content class="space-y-2 mt-5">
              <div
                class="flex flex-col justify-center items-center flex-nowrap pt-4 pb-6"
              >
                <p class="font-bold pb-4 text-2xl">
                  {$t("login.welcome-register")}
                </p>
                <p>
                  {$t("login.welcome-register.description")}
                </p>
              </div>

              <Form.Field form={registerForm} name="username">
                <Form.Control let:attrs>
                  <Form.Label>
                    <p class="font-semibold text-sm">
                      {$t("login.username")}
                    </p></Form.Label
                  >
                  <Input
                    {...attrs}
                    type="email"
                    bind:value={$registerFormData.username}
                    disabled={$requestInProcess}
                  />
                </Form.Control>
              </Form.Field>

              <Form.Field form={registerForm} name="password">
                <Form.Control let:attrs>
                  <Form.Label>
                    <p class="pt-4 font-semibold text-sm">
                      {$t("login.password")}
                    </p></Form.Label
                  >
                  <Input
                    {...attrs}
                    type="password"
                    bind:value={$registerFormData.password}
                    disabled={$requestInProcess}
                  />
                </Form.Control>
              </Form.Field>

              <Form.Field form={registerForm} name="repeatPassword">
                <Form.Control let:attrs>
                  <Form.Label>
                    <p class="pt-4 font-semibold text-sm">
                      {$t("login.password-repeat")}
                    </p></Form.Label
                  >
                  <Input
                    {...attrs}
                    type="password"
                    bind:value={$registerFormData.repeatPassword}
                    disabled={$requestInProcess}
                  />
                </Form.Control>
              </Form.Field>
            </Card.Content>

            <Card.Footer>
              <Button type="submit" class="w-full" disabled={$requestInProcess}>
                {$t("login.register-button")}
              </Button>
            </Card.Footer>
          </form>

          <div class="relative w-[80%] justify-self-center">
            <div class="absolute inset-0 flex items-center">
              <span class="w-full border-t" />
            </div>
            <div class="relative flex justify-center text-xs uppercase">
              <span class="bg-background text-muted-foreground px-2">
                {$t("login.register-with")}
              </span>
            </div>
          </div>

          <div class="py-6 px-8 flex flex-row justify-between">
            <Button class="py-6 px-10" variant="outline" type="button" disabled={$requestInProcess}>
              <svg
                class="h-6 w-6"
                fill="#000000"
                version="1.1"
                viewBox="0 0 22.773 22.773"
                xml:space="preserve"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="m15.769 0h0.162c0.13 1.606-0.483 2.806-1.228 3.675-0.731 0.863-1.732 1.7-3.351 1.573-0.108-1.583 0.506-2.694 1.25-3.561 0.69-0.808 1.955-1.527 3.167-1.687z"
                />
                <path
                  d="m20.67 16.716v0.045c-0.455 1.378-1.104 2.559-1.896 3.655-0.723 0.995-1.609 2.334-3.191 2.334-1.367 0-2.275-0.879-3.676-0.903-1.482-0.024-2.297 0.735-3.652 0.926h-0.462c-0.995-0.144-1.798-0.932-2.383-1.642-1.725-2.098-3.058-4.808-3.306-8.276v-1.019c0.105-2.482 1.311-4.5 2.914-5.478 0.846-0.52 2.009-0.963 3.304-0.765 0.555 0.086 1.122 0.276 1.619 0.464 0.471 0.181 1.06 0.502 1.618 0.485 0.378-0.011 0.754-0.208 1.135-0.347 1.116-0.403 2.21-0.865 3.652-0.648 1.733 0.262 2.963 1.032 3.723 2.22-1.466 0.933-2.625 2.339-2.427 4.74 0.176 2.181 1.444 3.457 3.028 4.209z"
                />
              </svg>
            </Button>
            <Button class="py-6 px-10" variant="outline" type="button" disabled={$requestInProcess}>
              <svg class="h-6 w-6" version="1.1" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <g fill="none" fill-rule="evenodd">
                <g transform="translate(-300 -7399)" fill="#000">
                <g transform="translate(56 160)">
                <path d="m263.82 7247h-9.6102c0 0.99944 0 2.9983-0.006126 3.9978h5.5689c-0.2134 0.99944-0.97001 2.3986-2.0391 3.1033-0.001021-1e-3 -0.002042 0.00599-0.004084 0.00499-1.4213 0.93848-3.297 1.1514-4.6897 0.87151-2.183-0.43375-3.9107-2.0169-4.6121-4.0277 0.004084-3e-3 0.007147-0.03098 0.01021-0.03298-0.43906-1.2473-0.43906-2.9174 0-3.9168h-0.001021c0.56567-1.837 2.3454-3.513 4.5315-3.9718 1.7583-0.37279 3.7422 0.03098 5.2013 1.3962 0.194-0.18989 2.6854-2.6225 2.8722-2.8204-4.9848-4.5145-12.966-2.9264-15.953 2.9034h-0.001021s0.001021 0-0.005105 0.01099c-1.4775 2.8634-1.4162 6.2375 0.01021 8.964-0.004084 0.00299-0.007147 0.00499-0.01021 0.00799 1.2927 2.5086 3.6452 4.4325 6.4797 5.1651 3.0111 0.78956 6.8432 0.24986 9.4101-2.0718l0.003063 3e-3c2.1749-1.9589 3.5288-5.296 2.8447-9.5866"></path>
                </g>
                </g>
                </g>
              </svg>
            </Button>
            <Button class="py-6 px-10" variant="outline" type="button" disabled={$requestInProcess}>
              <GithubLogo class="h-6 w-6" />
            </Button>
          </div>

          <div class="pb-4 text-sm justify-self-center">
            <p>
              {$t('login.prompt-to-login')}
              <Button on:click="{() => switchToLogin()}" class="underline" variant="link">{$t('login.redirect-to-login')}</Button>
            </p>
          </div>
        </Card.Root>
      </Tabs.Content>
    </Tabs.Root>
  </div>
</div>

<style>
  :global(html) {
    min-height: 100vh;
  }
  :global(body) {
    background-color: transparent;
  }
</style>
