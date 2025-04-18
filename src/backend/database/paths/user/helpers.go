package database_user

// RoleMapping maps a string role to a UserRole
func RoleMapping(role string) UserRole {
	switch role {
	case "admin":
		return AdminUser
	case "user":
		return ConfirmedUser
	case "temporary":
		return TemporaryUser
	default:
		return TemporaryUser
	}
}

// RoleMappingReverse UserRole to string
func RoleMappingReverse(role UserRole) string {
	switch role {
	case AdminUser:
		return "admin"
	case ConfirmedUser:
		return "user"
	case TemporaryUser:
		return "temporary"
	default:
		return "temporary"
	}
}
