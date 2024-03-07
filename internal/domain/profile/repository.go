package profile

import "context"

type Repository interface {
	CreateProfile(ctx context.Context, userID string, params CreateProfileParams) error
	// GetProfileByID(ctx context.Context, profileID int) Profile
	// GetProfileByUserID(ctx context.Context, userID int) Profile
	// UpdateProfile(ctx context.Context, params UpdateProfileParams) error
	// DeleteProfile(ctx context.Context, profileID int) error
}
