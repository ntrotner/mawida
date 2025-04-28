import { browser } from "$app/environment";

export const load = async ({ parent, url }) => {
  if (browser) {
    await parent();

    const searchParams = url.searchParams;
    const rentContractId = searchParams.get("rentContractId");
    
    return { rentContractId };
  }
}