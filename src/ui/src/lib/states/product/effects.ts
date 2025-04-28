import { ProductApi, type Success, type ModelError, ResponseError, type Product, type ProductsProductIdGet200Response } from "$lib/open-api";
import { productState } from "./product";
import type { RentProductFormular, RentProductConfirmation } from "$lib/open-api/models";

/**
 * Fetch all products.
 * @returns {Promise<Product[]>} - A promise that resolves to an array of products.
 */
export async function fetchProducts(): Promise<Product[]> {
  const productApi = new ProductApi();
  const products = await productApi.productsGet();
  productState.setState(products);
  return products;
}

/**
 * Fetch a single product.
 * @param {string} productId - The ID of the product to fetch.
 * @returns {Promise<ProductsProductIdGet200Response>} - A promise that resolves to a product.
 */
export async function fetchProduct(productId: string): Promise<ProductsProductIdGet200Response> {
  const productApi = new ProductApi();
  const product = await productApi.productsProductIdGet({ productId });
  productState.addProduct(product);
  return product;
}

/**
 * Create a new product.
 * @param {Product} product - The product to create.
 * @returns {Promise<Product>} - A promise that resolves to the created product.
 */
export async function createProduct(product: Product): Promise<Product | undefined> {
  const productApi = new ProductApi();
  try {
    const response = await productApi.productsPost({ product });
    await fetchProducts();
    return response;
  } catch (e: unknown) {
    let errorResponse: ModelError | undefined = undefined;
    if (e instanceof ResponseError) {
      errorResponse = await e.response.json() as ModelError;
    }
    throw errorResponse;
  }
}

/**
 * Delete a product.
 * @param {string} productId - The ID of the product to delete.
 * @returns {Promise<Success>} - A promise that resolves to a Success object.
 */
export async function deleteProduct(productId: string): Promise<Success & ModelError | undefined> {
  const productApi = new ProductApi();
  try {
    const response = await productApi.productsProductIdDelete({ productId });
    await fetchProducts();
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
 * Rent a product.
 * @param {string} productId - The ID of the product to rent.
 * @param {RentProductFormular} rentFormular - The rental formular with required details.
 * @returns {Promise<RentProductConfirmation>} - A promise that resolves to the rent confirmation.
 */
export async function rentProduct(productId: string, rentFormular: RentProductFormular): Promise<RentProductConfirmation | undefined> {
  const productApi = new ProductApi();
  try {
    const response = await productApi.productsProductIdRentPost({ productId, rentProductFormular: rentFormular });
    return response;
  } catch (e: unknown) {
    let errorResponse: ModelError | undefined = undefined;
    if (e instanceof ResponseError) {
      errorResponse = await e.response.json() as ModelError;
    }
    throw errorResponse;
  }
} 