import { LocationApi, type Success, type ModelError, ResponseError, type Location } from "../../../lib/open-api";
import { locationState } from "./location";

/**
 * Fetch all locations.
 * @returns {Promise<Location[]>} - A promise that resolves to an array of locations.
 */
export async function fetchLocations(): Promise<Location[]> {
  try {
    const locationApi = new LocationApi();
    const locations = await locationApi.locationsGet();
    locationState.setState(locations);
    return locations;
  } catch (e) {
    console.error('Error fetching locations:', e);
    return [];
  }
}

/**
 * Fetch a specific location by ID.
 * @param {string} locationId - The ID of the location to fetch.
 * @returns {Promise<Location | undefined>} - A promise that resolves to a location or undefined.
 */
export async function fetchLocationById(locationId: string): Promise<Location | undefined> {
  try {
    const locationApi = new LocationApi();
    return await locationApi.locationLocationIdGet({ locationId });
  } catch (e) {
    console.error(`Error fetching location with ID ${locationId}:`, e);
    return undefined;
  }
}

/**
 * Create a new location.
 * @param {Location} location - The location to create.
 * @returns {Promise<Success & ModelError | undefined>} - A promise that resolves to a Success object or undefined.
 */
export async function createLocation(location: Location): Promise<Success & ModelError | undefined> {
  const locationApi = new LocationApi();
  try {
    const response = await locationApi.locationsPost({ location });
    // Refresh the locations list after creating a new one
    await fetchLocations();
    return response;
  } catch (e: unknown) {
    let errorResponse: ModelError | undefined = undefined;
    if (e instanceof ResponseError) {
      errorResponse = await e.response.json() as ModelError;
    }
    return errorResponse;
  }
} 