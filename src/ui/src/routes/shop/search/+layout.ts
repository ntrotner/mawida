import { browser } from "$app/environment";

export const load = async ({ parent, url }) => {
  if (browser) {
    await parent();

    const searchParams = url.searchParams;
    const category = searchParams.get("category");
    const location = searchParams.get("location");

    return { category, location };
  }
}