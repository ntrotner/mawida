<script lang="ts">
  import * as Menubar from "$lib/components/ui/menubar";
  import * as Sheet from "$lib/components/ui/sheet/index.js";
  import Globe from "svelte-radix/Globe.svelte";
  import HamburgerMenu from "svelte-radix/HamburgerMenu.svelte";
  import { Separator } from "$lib/components/ui/separator/index.js";
  import { Button } from "$lib/components/ui/button/index.js";
  import { languages, setLocale, setLocaleInStorage, t } from "$lib/i18n";
  import { goto } from "$app/navigation";
  import { browser } from "$app/environment";
  import { ROUTES } from "$lib/routes";
  import { appState } from "$lib/states/app";
  import { userState } from "$lib/states/user";
  import { onDestroy } from "svelte";
  import type { Unsubscriber } from "svelte/store";
  import { map } from "rxjs";
  import * as Avatar from "../../lib/components/ui/avatar";

  const subscriptions: Unsubscriber[] = [];
  const user = userState.getAsyncState();
  const isAdmin = userState.isAdmin();
  const mobile = appState
    .observable()
    .pipe(map((state) => (state?.width || 0) <= 640));
  let menuBarOptions: Menubar.Menubar;
  let profileBarOptions: Menubar.Menubar;

  function changeMenuState(nextState: boolean) {
    if (menuBarOptions) {
      // @ts-ignore
      menuBarOptions.$set({ open: nextState });
    }
  }

  function changeProfileState(nextState: boolean) {
    if (profileBarOptions) {
      // @ts-ignore
      profileBarOptions.$set({ open: nextState });
    }
  }

  function changeLanguage(locale: string) {
    setLocale(locale);
    setLocaleInStorage(locale);
    if (browser) {
      document.documentElement.lang = locale;
    }
  }

  function redirect(page: string) {
    goto(page);
  }

  onDestroy(() => subscriptions.forEach((s) => s()));
</script>

<Sheet.Root>
  <Sheet.Trigger asChild let:builder>
    <div
      class="container px-2 pt-0 sm:pt-2 mt-3 h-12 grid grid-flow-col grid-cols-8 justify-stretch"
    >
      <div class="nav-left col-span-2">
        <div class="title">
          <Button
            on:click={() => redirect(ROUTES.HOME)}
            variant="link"
            class={($mobile ? "text-xl" : "text-lg") + " font-semibold"}
            >{$t("common.nav-title")}</Button
          >
        </div>
      </div>

      <div class="nav-middle col-span-4"></div>

      <div class="nav-right flex col-span-2 justify-end">
        <div class="language mr-2">
          <Button builders={[builder]} variant="ghost" size="icon">
            <Globe class="h-5 w-5" />
          </Button>
        </div>
        {#if $user?.email}
          <div>
            <div>
              <Menubar.Root class="border-0">
                <Menubar.Menu
                  onOutsideClick={() => {
                    changeProfileState(false);
                    changeMenuState(false);
                  }}
                  bind:this={profileBarOptions}
                >
                  <Menubar.Trigger>
                    <Avatar.Root>
                      <Avatar.Fallback
                        >{$user.email.slice(0, 2)}</Avatar.Fallback
                      >
                    </Avatar.Root>
                  </Menubar.Trigger>
                  <Menubar.Content>
                    <Menubar.Item on:click={() => redirect(ROUTES.PROFILE)}
                      >{$t("common.nav-menu.profile")}</Menubar.Item
                    >
                    <Menubar.Separator />
                    <Menubar.Item on:click={() => redirect(ROUTES.LOGOUT)}
                      >{$t("common.nav-menu.logout")}</Menubar.Item
                    >
                  </Menubar.Content>
                </Menubar.Menu>
              </Menubar.Root>
            </div>
          </div>
        {/if}
        <div>
          <Menubar.Root class="border-0">
            <Menubar.Menu
              onOutsideClick={() => {
                changeProfileState(false);
                changeMenuState(false);
              }}
              bind:this={menuBarOptions}
            >
              <Menubar.Trigger
                ><HamburgerMenu class="h-5 w-5" /></Menubar.Trigger
              >
              <Menubar.Content>
                {#if $user?.email}
                  <Menubar.Item on:click={() => redirect(ROUTES.HOME)}
                    >{$t("common.nav-links.home")}</Menubar.Item
                  >
                {:else}
                  <Menubar.Item on:click={() => redirect(ROUTES.LOGIN)}
                    >{$t("common.nav-menu.login")}</Menubar.Item
                  >
                {/if}
                {#if $isAdmin}
                  <Menubar.Item on:click={() => redirect(ROUTES.ADMIN)}
                    >{$t("common.nav-menu.admin")}</Menubar.Item
                  >
                {/if}
              </Menubar.Content>
            </Menubar.Menu>
          </Menubar.Root>
        </div>
      </div>
    </div>
  </Sheet.Trigger>
  <Sheet.Content side="right">
    <Sheet.Header>
      <Sheet.Title>{$t("common.nav-language.language")}</Sheet.Title>
      <Sheet.Description>
        {$t("common.nav-language.language-description")}
      </Sheet.Description>
    </Sheet.Header>
    <Sheet.Close asChild let:builder>
      <div class="space-y-4 py-4">
        <div class="space-y-1">
          {#each languages as language}
            <Button
              builders={[builder]}
              on:click={() => changeLanguage(language.locale)}
              variant="link"
              class="{$mobile
                ? 'justify-center'
                : 'justify-start'} w-full language-select"
            >
              {$t(language.key)}
            </Button>
            <Separator class="my-4" />
          {/each}
        </div>
      </div>
    </Sheet.Close>
  </Sheet.Content>
</Sheet.Root>
