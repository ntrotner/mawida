import { DefaultState } from '../common/state';
import type { RentContract } from '$lib/open-api';

/**
 * RentContractsState class is used to manage the state of rent contracts.
 */
export class RentContractsState extends DefaultState<RentContract[]> {

    addRentContract(rentContract: RentContract) {
        const currentState = this.getSyncState();
        const nextState = [...(currentState?.filter(contract => contract.id !== rentContract.id) || []), rentContract];
        this.setState(nextState);
    }

    getRentContractById(id: string) {
        const currentState = this.getSyncState();
        return currentState?.find(contract => contract.id === id);
    }
}

export const rentContractsState = new RentContractsState();

export const RENT_CONTRACTS_STATE = 'rentContractsState'; 