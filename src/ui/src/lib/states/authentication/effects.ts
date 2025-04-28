import { AuthenticationApi, ResponseError, type ModelError, type Success } from "$lib/open-api";
import { fetchUserProfile, userState } from "../user";
import { authenticationState } from ".";
import { clearToken, existsToken, isTokenTimeValid } from "../../open-api/helpers";
import { appState } from "../app";
import { productState } from "../product";
import { locationState } from "../location";
import { rentContractsState } from "../rent_contracts";

/**
 * Login the user.
 * @param {string} email - The email of the user.
 * @param {string} password - The password of the user.
 * @returns {Promise} - A promise that resolves success or error object.
 */
export async function login(email: string, password: string): Promise<Success & ModelError | undefined> {
  const authApi = new AuthenticationApi();

  try {
    const response = await authApi.loginPost({
      userLogin: {
        email,
        password
      }
    });
    await fetchUserProfile();
    productState.reset();
    locationState.reset();
    rentContractsState.reset();
    authenticationState.setAuthStatus(true)
    return response;
  } catch (e: unknown) {
    let errorResponse: ModelError | undefined = undefined;
    if (e instanceof ResponseError) {
      errorResponse = await e.response.json() as ModelError;
    }
    userState.setState(undefined);
    authenticationState.setAuthStatus(false);
    return errorResponse;
  }
}

/**
 * Register the user.
 * @param {string} email - The email of the user.
 * @param {string} password - The password of the user.
 * @returns {Promise} - A promise that resolves success or error object.
 */
export async function register(email: string, password: string): Promise<Success & ModelError | undefined> {
  const authApi = new AuthenticationApi();

  try {
    const response = await authApi.registerPost({
      userRegistration: {
        email,
        password
      }
    });
    await fetchUserProfile();
    authenticationState.setAuthStatus(true);
    return response;
  } catch (e: unknown) {
    let errorResponse: ModelError | undefined = undefined;
    if (e instanceof ResponseError) {
      errorResponse = await e.response.json() as ModelError;
    }
    userState.setState(undefined);
    authenticationState.setAuthStatus(false);
    return errorResponse;
  }
}

/**
 * Logout the user.
 * @returns {Promise<boolean>} - A promise that resolves to a boolean indicating the success of the logout.
 */
export async function logout() {
  const authApi = new AuthenticationApi();
  let requestFailed = false;

  try {
    await authApi.logoutPost();
  } catch {
    requestFailed = true;
  }

  userState.setState(undefined);
  authenticationState.setAuthStatus(false);
  return !requestFailed;
}

/**
 * Refresh token for authentication.
 * @returns {Promise<boolean>} - A promise that resolves to a boolean indicating the success of the token refresh.
 */
export async function refreshToken() {
  if (!existsToken() || !isTokenTimeValid()) {
    clearToken();
    authenticationState.setAuthStatus(false);
    appState.setLoaded(true);
    return false;
  }
  const authApi = new AuthenticationApi();

  try {
    await authApi.refreshTokenPost();
    await fetchUserProfile();
    authenticationState.setAuthStatus(true);
    appState.setLoaded(true);
    return true;
  } catch {
    userState.setState(undefined);
    authenticationState.setAuthStatus(false);
    appState.setLoaded(true);
    clearToken();
    return false;
  }
}