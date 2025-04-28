import { AdministrationApi, type UserProfile, type ModelError, type Success, ResponseError, type ChangeRoleRoleEnum } from "$lib/open-api";
import { usersState } from "./users";

/**
 * Fetch all users.
 * @returns {Promise<UserProfile[]>} - A promise that resolves to an array of users.
 */
export async function fetchUsers(): Promise<UserProfile[]> {
  try {
    const adminApi = new AdministrationApi();
    const users = await adminApi.administrationUsersGet();
    usersState.setState(users);
    return users;
  } catch (e) {
    console.error('Error fetching users:', e);
    return [];
  }
}

/**
 * Fetch a specific user by ID.
 * @param {string} userId - The ID of the user to fetch.
 * @returns {Promise<UserProfile | undefined>} - A promise that resolves to a user or undefined.
 */
export async function fetchUserById(userId: string): Promise<UserProfile | undefined> {
  try {
    // Simulate fetching a specific user
    // Replace with actual API call when available
    const users = await fetchUsers();
    return users.find(user => user.email === userId);
  } catch (e) {
    console.error(`Error fetching user with ID ${userId}:`, e);
    return undefined;
  }
}

/**
 * Change the role of a user.
 * @param {string} userId - The ID of the user.
 * @param {ChangeRoleRoleEnum} role - The new role for the user.
 * @returns {Promise<Success & ModelError | undefined>} - A promise that resolves to a Success object or undefined.
 */
export async function changeUserRole(userId: string, role: ChangeRoleRoleEnum): Promise<Success & ModelError | undefined> {
  const adminApi = new AdministrationApi();
  try {
    const response = await adminApi.administrationChangeRoleUserIdPost({
      userId,
      changeRole: { role }
    });
    // Refresh the users list after changing a user's role
    await fetchUsers();
    return response;
  } catch (e: unknown) {
    let errorResponse: ModelError | undefined = undefined;
    if (e instanceof ResponseError) {
      errorResponse = await e.response.json() as ModelError;
    }
    return errorResponse;
  }
} 