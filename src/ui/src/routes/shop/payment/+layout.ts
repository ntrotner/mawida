import { browser } from "$app/environment";
import { ROUTES } from "$lib/routes";
import { goto } from "$app/navigation";

export const load = async ({ parent, url }) => {
  if (browser) {
    await parent();

    const searchParams = url.searchParams;
    const productId = searchParams.get("productId");
    const success = searchParams.get("success");

    goto(ROUTES.CUSTOMER_PROFILE);
  }
}