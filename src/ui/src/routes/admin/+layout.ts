import { browser } from "$app/environment";
import { isUserAdmin, isUserAuthenticated } from "$lib/routes/guards/authentication";
import { goto } from "$app/navigation";
import { ROUTES } from "$lib/routes";
import { firstValueFrom, map } from "rxjs";
import { configState } from "$lib/states/config";
import { type AppConfig, AppConfigKey } from "$lib/states/status";
import { fetchLocations, locationState } from "$lib/states/location";
import { fetchSales, salesState } from "$lib/states/sales";

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
        }

        const locations = locationState.getSyncState();
        if (!locations || locations.length === 0) {
          fetchLocations();
        }

        const sales = salesState.getSyncState();
        if (!sales || sales.length === 0) {
          fetchSales();
        }
      })
  }
}