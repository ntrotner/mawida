import { browser } from "$app/environment";
import { isUserAdmin, isUserAuthenticated } from "$lib/routes/guards/authentication";
import { goto } from "$app/navigation";
import { ROUTES } from "$lib/routes";
import { firstValueFrom, map } from "rxjs";
import { configState } from "$lib/states/config";
import { type AppConfig, AppConfigKey } from "$lib/states/status";

export const load = async ({ parent, url }) => {
  if (browser) {
    await parent();

    await firstValueFrom(configState.getConfig<AppConfig>(AppConfigKey)
      .pipe(map((config) => config?.user))).then(async isUserEnabled => {
        if (!isUserEnabled) {
          goto(ROUTES.HOME);
        }
        const authenticated = await isUserAuthenticated();
        if (!authenticated) {
          goto(ROUTES.HOME);
        }
        
        const isAdmin = await isUserAdmin();
        if (!isAdmin) {
          goto(ROUTES.HOME);
        } else if (url.pathname.endsWith(ROUTES.ADMIN)) {
          goto(ROUTES.ADMIN_STATISTICS);
        }
      })
  }
}