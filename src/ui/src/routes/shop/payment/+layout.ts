import { browser } from "$app/environment";

export const load = async ({ parent, url }) => {
  if (browser) {
    await parent();

    const searchParams = url.searchParams;
    const productId = searchParams.get("productId");
    const success = searchParams.get("success");
    
    return { productId, success };
  }
}