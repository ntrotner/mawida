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
  import { onDestroy } from "svelte";
  import Alert from "../../components/alert/Alert.svelte";
  import { map } from "rxjs";
  import { superForm } from "sveltekit-superforms";

  // Reactive state management
  const subscriptions: Unsubscriber[] = [];
  const errors = writable<string[]>([]);
  const requestInProcess = writable(false);
  
  // Form data stores
  const loginData = { username: "", password: "" };
  const registerData = { username: "", password: "", repeatPassword: "" };
  
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
    if (!!(errorInput as ModelError)?.errorMessages) {
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
        goto(ROUTES.HOME);
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

  function clear() {
    $errors = [];
    $loginFormData = { username: "", password: "" };
    $registerFormData = { username: "", password: "", repeatPassword: "" };
  }

  onDestroy(() => subscriptions.forEach((s) => s()));
</script>

<div class="container {$mobile ? 'pt-8' : 'pt-16'}">
  <Tabs.Root onValueChange={clear} value="login" class="w-[475px]">
    <Tabs.List class="grid w-full grid-cols-2 mb-4">
      <Tabs.Trigger value="login" disabled={$requestInProcess}>
        {$t("login.login-title")}
      </Tabs.Trigger>
      <Tabs.Trigger value="register" disabled={$requestInProcess}>
        {$t("login.register-title")}
      </Tabs.Trigger>
    </Tabs.List>
    
    <Tabs.Content value="login">
      <Alert
        messages={$errors}
        title={$t("login.input-error-title")}
        variant="destructive"
      />
      
      <Card.Root>
        <form on:submit|preventDefault={submitLogin}>
          <Card.Content class="space-y-2 mt-5">
            <Form.Field form={loginForm} name="username">
              <Form.Control let:attrs>
                <Form.Label>{$t("login.username")}</Form.Label>
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
                <Form.Label>{$t("login.password")}</Form.Label>
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
            <Button
              type="submit"
              class="w-full"
              disabled={$requestInProcess}
            >
              {$t("login.login-button")}
            </Button>
          </Card.Footer>
        </form>
      </Card.Root>
    </Tabs.Content>
    
    <Tabs.Content value="register">
      <Alert
        messages={$errors}
        title={$t("login.input-error-title")}
        variant="destructive"
      />
      
      <Card.Root>
        <form on:submit|preventDefault={submitRegister}>
          <Card.Content class="space-y-2 mt-5">
            <Form.Field form={registerForm} name="username">
              <Form.Control let:attrs>
                <Form.Label>{$t("login.username")}</Form.Label>
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
                <Form.Label>{$t("login.password")}</Form.Label>
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
                <Form.Label>{$t("login.password-repeat")}</Form.Label>
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
            <Button
              type="submit"
              class="w-full"
              disabled={$requestInProcess}
            >
              {$t("login.register-button")}
            </Button>
          </Card.Footer>
        </form>
      </Card.Root>
    </Tabs.Content>
  </Tabs.Root>
</div>

<style>
  .container {
    display: flex;
    flex-direction: row;
    justify-content: center;
  }
</style>
