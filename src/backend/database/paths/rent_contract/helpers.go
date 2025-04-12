package database_rent_contract

func (s RentContractStatus) CanTransitionTo(newStatus RentContractStatus) bool {
	switch s {
	case RentContractStatusConfirmed:
		return newStatus == RentContractStatusPickupPending
	case RentContractStatusPickupPending:
		return newStatus == RentContractStatusActive
	case RentContractStatusActive:
		return newStatus == RentContractStatusCompleted
	default:
		return false
	}
}

func (s RentContractStatus) IsValid() bool {
	switch s {
	case RentContractStatusPickupPending,
		RentContractStatusConfirmed,
		RentContractStatusActive,
		RentContractStatusCompleted:
		return true
	default:
		return false
	}
}
