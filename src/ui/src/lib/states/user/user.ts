import { DefaultState } from '../common/state';
import type { UserProfile } from '$lib/open-api';
import { map } from 'rxjs';


/**
 * UserState class is used to manage the state of the user.
 */
export class UserState extends DefaultState<UserProfile> {
    public isAdmin() {
        return this.observable().pipe(
            map((user) => user?.role === 'admin')
        );
    }
}
export const userState = new UserState();

export const USER_STATE = 'userState';