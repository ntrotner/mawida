package database_user

// SanitizeUserProfile sanitizes a user profile
func SanitizeUserProfile(user *UserProfile) *PublicUserProfile {
	sanitized := PublicUserProfile{}
	sanitized.Email = user.Email
	sanitized.Role = RoleMappingReverse(user.Roles)
	return &sanitized
}

// SanitizeUserProfileForAdmin sanitizes a user profile for admin
func SanitizeUserProfileForAdmin(user *UserProfile) *PublicUserProfile {
	sanitized := PublicUserProfile{}
	sanitized.Email = user.Email
	sanitized.Role = RoleMappingReverse(user.Roles)
	sanitized.ID = user.ID
	return &sanitized
}
