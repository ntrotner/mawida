import { DefaultState } from '../common/state';
import type { ProductPublic } from '$lib/open-api';

/**
 * ProductState class is used to manage the state of products.
 */
export class ProductState extends DefaultState<ProductPublic[]> {
}

export const productState = new ProductState();

export const PRODUCT_STATE = 'productState'; 