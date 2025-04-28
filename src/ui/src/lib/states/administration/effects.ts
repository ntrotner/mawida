import { AdministrationApi } from "$lib/open-api";
import { fetchUsers } from "$lib/states/users";
/**
 * Promote a user to a customer.
 * @param {string} userId - The ID of the user to promote.
 * @returns {Promise<void>} - A promise that resolves when the user is promoted.
 */
export async function promotoToCustomer(userId: string) {
    try {
        const administrationApi = new AdministrationApi();
        const response = await administrationApi.administrationChangeRoleUserIdPost({
            userId: userId,
        changeRole: {
            role: "user"
        }
        });

        await fetchUsers();
        return response;
    } catch (error) {
        console.error(error);
        return undefined;
    }
}