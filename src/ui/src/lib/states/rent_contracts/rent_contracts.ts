import { DefaultState } from '../common/state';
import type { RentContract } from '$lib/open-api';

/**
 * RentContractsState class is used to manage the state of rent contracts.
 */
export class RentContractsState extends DefaultState<RentContract[]> {
}

export const rentContractsState = new RentContractsState();

export const RENT_CONTRACTS_STATE = 'rentContractsState'; 