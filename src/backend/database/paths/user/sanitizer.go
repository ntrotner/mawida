package database_user

// SanitizeUserProfile sanitizes a user profile
func SanitizeUserProfile(user *UserProfile) *PublicUserProfile {
	sanitized := PublicUserProfile{}
	sanitized.Email = user.Email
	sanitized.Role = user.Roles
	return &sanitized
}
