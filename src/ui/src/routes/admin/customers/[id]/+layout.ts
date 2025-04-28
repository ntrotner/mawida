import { fetchUserById } from '$lib/states/users';

export const prerender = false;

export const load = async ({ params }) => {
  const { id } = params;
  const customer = await fetchUserById(id);
  return { customer };
};
