import { derived } from 'svelte/store';
import { DefaultState } from '../common/state';
import type { AuthenticationStatus } from './model';
import { userState } from '../user';

/**
 * AuthenticationState class is used to manage the state of the authentication.
 */
export class AuthenticationState extends DefaultState<AuthenticationStatus> {
  constructor() {
    super();
    this.setState({ authenticated: undefined })
  }

  public setAuthStatus(isAuthenticated: boolean) {
    this.setState({ ...this.getSyncState(), authenticated: isAuthenticated })
  }

  public isAdmin() {
    return derived([this.getAsyncState(), userState.getAsyncState()], ([auth, user]) => auth?.authenticated && user?.role === "admin");
  } 
}
export const authenticationState = new AuthenticationState();

export const AUTHENTICATION_STATE = 'authState';