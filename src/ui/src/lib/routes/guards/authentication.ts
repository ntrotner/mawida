import { filter, firstValueFrom, take } from "rxjs";
import { authenticationState } from "../../states/authentication";
import { userState } from "../../states/user";

export async function isUserAuthenticated() {
  const userAuthenticationState = await firstValueFrom(authenticationState.observable().pipe(
    filter(state => typeof state.authenticated !== 'undefined'),
    take(1)
  ))
  const user = userState.getSyncState();
  return !!userAuthenticationState.authenticated || !!user?.email;
}

export async function isUserAdmin() {
  const userAuthenticationState = await firstValueFrom(authenticationState.observable().pipe(
    filter(state => typeof state.authenticated !== 'undefined'),
    take(1)
  ))
  const user = userState.getSyncState();
  return userAuthenticationState.authenticated && user?.role === "admin";
}
