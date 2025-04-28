import { DefaultState } from '../common/state';
import type { Sale } from '$lib/open-api';

/**
 * SalesState class is used to manage the state of sales.
 */
export class SalesState extends DefaultState<Sale[]> {
}

export const salesState = new SalesState();

export const SALES_STATE = 'salesState'; 