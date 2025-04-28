import { DefaultState } from '../common/state';
import type { Location } from '$lib/open-api';

/**
 * LocationState class is used to manage the state of locations.
 */
export class LocationState extends DefaultState<Location[]> {
    public reset() {
        this.setState([]);
    }
}

export const locationState = new LocationState();

export const LOCATION_STATE = 'locationState'; 