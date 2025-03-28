package database_rent_contract

func (s RentContractStatus) CanTransitionTo(newStatus RentContractStatus) bool {
	switch s {
	case RentContractStatusPending:
		return newStatus == RentContractStatusConfirmed || newStatus == RentContractStatusCancelled
	case RentContractStatusConfirmed:
		return newStatus == RentContractStatusActive || newStatus == RentContractStatusCancelled
	case RentContractStatusActive:
		return newStatus == RentContractStatusCompleted || newStatus == RentContractStatusOverdue || newStatus == RentContractStatusDisputed
	case RentContractStatusOverdue:
		return newStatus == RentContractStatusCompleted || newStatus == RentContractStatusDisputed
	case RentContractStatusDisputed:
		return newStatus == RentContractStatusCompleted || newStatus == RentContractStatusCancelled
	default:
		return false
	}
}

func (s RentContractStatus) IsValid() bool {
	switch s {
	case RentContractStatusPending,
		RentContractStatusConfirmed,
		RentContractStatusActive,
		RentContractStatusCompleted,
		RentContractStatusCancelled,
		RentContractStatusOverdue,
		RentContractStatusDisputed:
		return true
	default:
		return false
	}
}
