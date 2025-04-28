import { getLocale, loadTranslations, setLocale, setLocaleInStorage, t } from "$lib/i18n";
import { logger } from "$lib/analytics";
import { take } from "rxjs";
import { browser } from "$app/environment";
import { fetchConfigurations, configState } from "../lib/states/config";
import { type AppConfig, AppConfigKey, checkStatus } from "../lib/states/status";
import { toast } from "svelte-sonner";
import { TOKEN_REFRESH_IN_MS } from "../lib/open-api/helpers";
import { appState } from "../lib/states/app";
import { refreshToken } from "../lib/states/authentication";

export const prerender = false;
export const ssr = false;

function setupLocalization() {
  let locale = getLocale();
  if (!locale) {
    locale = 'en';
    setLocaleInStorage(locale);
  }
  setLocale(locale);
  document.documentElement.lang = locale;
  loadTranslations(locale);
}

async function loadConfig() {
  await fetchConfigurations(import.meta.env.VITE_configUrl || window.configUrl);
}

async function checkBackendStatus() {
  let healthStatus = false;
  let firstTry = false;

  while (!healthStatus) {
    try {
      healthStatus = await checkStatus();
      if (!healthStatus && !firstTry) {
        toast.error(t.get('common.backend-down'), { duration: 10000000 })
      } else if (healthStatus) {
        toast.dismiss()
      }
    } catch {
      if (!firstTry) {
        toast.error(t.get('common.backend-down'), { duration: 10000000 })
      }
    } finally {
      firstTry = true;
    }
    if (!healthStatus) {
      await new Promise(resolve => setTimeout(resolve, 5000));
    }
  }
}

async function setupToken() {
  setInterval(() => refreshToken(), TOKEN_REFRESH_IN_MS);
  refreshToken();
}

export const load = async () => {
  if (browser) {
    setupLocalization();
    await loadConfig();

    configState.getConfig<AppConfig>(AppConfigKey).pipe(
      take(1)
    ).subscribe(async appConfig => {
      logger.info("page init")
      if (appConfig?.isBackendAware) {
        checkBackendStatus();
      }

      if (appConfig?.user) {
        setupToken();
      } else {
        appState.setLoaded(true)
      }
    })
  }
  return {};
}