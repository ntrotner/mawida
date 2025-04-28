import { AdministrationApi, type Sale, type ModelError, type Success, ResponseError } from "$lib/open-api";
import { salesState } from "./sales";

/**
 * Fetch all sales.
 * @returns {Promise<Sale[]>} - A promise that resolves to an array of sales.
 */
export async function fetchSales(): Promise<Sale[]> {
  try {
    const adminApi = new AdministrationApi();
    const sales = await adminApi.administrationSalesGet();
    salesState.setState(sales);
    return sales;
  } catch (e) {
    console.error('Error fetching sales:', e);
    return [];
  }
}

/**
 * Fetch a specific sale by ID.
 * @param {string} saleId - The ID of the sale to fetch.
 * @returns {Promise<Sale | undefined>} - A promise that resolves to a sale or undefined.
 */
export async function fetchSaleById(saleId: string): Promise<Sale | undefined> {
  try {
    const sales = await fetchSales();
    return sales.find(sale => sale.id === saleId);
  } catch (e) {
    console.error(`Error fetching sale with ID ${saleId}:`, e);
    return undefined;
  }
} 