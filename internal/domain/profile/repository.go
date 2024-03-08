package profile

import "context"

type Repository interface {
	CreateProfile(ctx context.Context, userID string, params CreateProfileParams) error
	GetProfileByID(ctx context.Context, profileID int) (Profile, error)
	GetProfileByUserID(ctx context.Context, userID string) (Profile, error)
	UpdateProfile(ctx context.Context, userID string, params UpdateProfileParams) error
}
