package mock

import (
	"context"

	"github.com/yourusername/truecaller-lite/pkg/models"
)

// UserDAOMock is a mock implementation of UserDAO for testing.
type UserDAOMock struct {
	OnCreateOrUpdateUser   func(ctx context.Context, user *models.User) error
	OnGetUserByPhoneNumber func(ctx context.Context, phoneNumber string) (*models.User, error)
	OnGetAllUsers          func(ctx context.Context) ([]*models.User, error)
	OnUpdateSpamStatus     func(ctx context.Context, phoneNumber string, isSpam bool) error
}

func (m *UserDAOMock) CreateOrUpdateUser(ctx context.Context, user *models.User) error {
	if m.OnCreateOrUpdateUser != nil {
		return m.OnCreateOrUpdateUser(ctx, user)
	}
	return nil
}

func (m *UserDAOMock) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error) {
	if m.OnGetUserByPhoneNumber != nil {
		return m.OnGetUserByPhoneNumber(ctx, phoneNumber)
	}
	return nil, nil
}

func (m *UserDAOMock) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	if m.OnGetAllUsers != nil {
		return m.OnGetAllUsers(ctx)
	}
	return nil, nil
}

func (m *UserDAOMock) UpdateSpamStatus(ctx context.Context, phoneNumber string, isSpam bool) error {
	if m.OnUpdateSpamStatus != nil {
		return m.OnUpdateSpamStatus(ctx, phoneNumber, isSpam)
	}
	return nil
}
