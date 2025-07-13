package mock

import (
	"context"

	"github.com/yourusername/truecaller-lite/pkg/models"
)

// SpamUserDAOMock is a mock implementation of UserDAO for spam service tests.
type SpamUserDAOMock struct {
	OnUpdateSpamStatus func(ctx context.Context, phoneNumber string, isSpam bool) error
	OnGetAllUsers      func(ctx context.Context) ([]*models.User, error)
}

func (m *SpamUserDAOMock) UpdateSpamStatus(ctx context.Context, phoneNumber string, isSpam bool) error {
	if m.OnUpdateSpamStatus != nil {
		return m.OnUpdateSpamStatus(ctx, phoneNumber, isSpam)
	}
	return nil
}

func (m *SpamUserDAOMock) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	if m.OnGetAllUsers != nil {
		return m.OnGetAllUsers(ctx)
	}
	return nil, nil
}

// The following methods are not used in SpamService tests, so provide no-op implementations.
func (m *SpamUserDAOMock) CreateOrUpdateUser(ctx context.Context, user *models.User) error {
	return nil
}
func (m *SpamUserDAOMock) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error) {
	return nil, nil
}
