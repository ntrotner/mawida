import { RentalApi, type RentContract, type ModelError, type Success, ResponseError, type PickupConfirmation, type ReturnProduct } from "$lib/open-api";
import { rentContractsState } from "./rent_contracts";

/**
 * Fetch all rent contracts.
 * @returns {Promise<RentContract[]>} - A promise that resolves to an array of rent contracts.
 */
export async function fetchRentContracts(): Promise<RentContract[]> {
  try {
    const rentalApi = new RentalApi();
    const rentContracts = await rentalApi.rentalsGet();
    rentContractsState.setState(rentContracts);
    return rentContracts;
  } catch (e) {
    console.error('Error fetching rent contracts:', e);
    return [];
  }
}

/**
 * Fetch a specific rent contract by ID.
 * @param {string} rentContractId - The ID of the rent contract to fetch.
 * @returns {Promise<RentContract | undefined>} - A promise that resolves to a rent contract or undefined.
 */
export async function fetchRentContractById(rentContractId: string): Promise<RentContract | undefined> {
  try {
    const rentalApi = new RentalApi();
    const rentContract = await rentalApi.rentalsRentContractIdGet({ rentContractId });
    rentContractsState.addRentContract(rentContract);
    return rentContract;
  } catch (e) {
    console.error(`Error fetching rent contract with ID ${rentContractId}:`, e);
    return undefined;
  }
}

/**
 * Cancel a rent contract.
 * @param {string} rentContractId - The ID of the rent contract to cancel.
 * @returns {Promise<Success & ModelError | undefined>} - A promise that resolves to a Success object or undefined.
 */
export async function cancelRentContract(rentContractId: string): Promise<Success & ModelError | undefined> {
  const rentalApi = new RentalApi();
  try {
    const response = await rentalApi.rentalsRentContractIdCancelPost({ rentContractId });
    // Refresh the rent contracts list after cancelling
    await fetchRentContracts();
    return response;
  } catch (e: unknown) {
    let errorResponse: ModelError | undefined = undefined;
    if (e instanceof ResponseError) {
      errorResponse = await e.response.json() as ModelError;
    }
    return errorResponse;
  }
} 

/**
 * Pick up a product from a rent contract.
 * @param {string} rentContractId - The ID of the rent contract to pick up.
 * @param {PickupConfirmation} pickupConfirmation - The pickup confirmation object.
 * @returns {Promise<Success & ModelError | undefined>} - A promise that resolves to a Success object or undefined.
 */
export async function pickUpProduct(rentContractId: string, pickupConfirmation: PickupConfirmation): Promise<Success & ModelError | undefined> {
  const rentalApi = new RentalApi();
  try {
    const response = await rentalApi.rentalsRentContractIdPickupPost({ rentContractId, pickupConfirmation });
    return response;
  } catch (e: unknown) {
    return undefined;
  }
}

/**
 * Return a product to a rent contract.
 * @param {string} rentContractId - The ID of the rent contract to return the product to.
 * @param {ReturnProduct} returnProduct - The return product object.
 * @returns {Promise<Success & ModelError | undefined>} - A promise that resolves to a Success object or undefined.
 */
export async function returnProduct(rentContractId: string, returnProduct: ReturnProduct): Promise<Success & ModelError | undefined> {
  const rentalApi = new RentalApi();
  try {
    const response = await rentalApi.rentalsRentContractIdReturnPost({ rentContractId, returnProduct });
    return response;
  } catch (e: unknown) {
    return undefined;
  }
}