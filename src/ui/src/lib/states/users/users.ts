import { DefaultState } from '../common/state';
import type { UserProfile } from '$lib/open-api';

/**
 * UsersState class is used to manage the state of customer users.
 */
export class UsersState extends DefaultState<UserProfile[]> {
}

export const usersState = new UsersState();

export const USERS_STATE = 'usersState'; 