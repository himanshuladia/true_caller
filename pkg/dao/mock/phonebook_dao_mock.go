package mock

import (
	"context"

	"github.com/yourusername/truecaller-lite/pkg/models"
)

// PhoneBookDAOMock is a mock implementation of PhoneBookDAO for testing.
type PhoneBookDAOMock struct {
	OnCreateOrUpdatePhoneBook       func(ctx context.Context, phoneBook *models.PhoneBook) error
	OnGetPhoneBookByUserPhoneNumber func(ctx context.Context, phoneNumber string) (*models.PhoneBook, error)
}

func (m *PhoneBookDAOMock) CreateOrUpdatePhoneBook(ctx context.Context, phoneBook *models.PhoneBook) error {
	if m.OnCreateOrUpdatePhoneBook != nil {
		return m.OnCreateOrUpdatePhoneBook(ctx, phoneBook)
	}
	return nil
}

func (m *PhoneBookDAOMock) GetPhoneBookByUserPhoneNumber(ctx context.Context, phoneNumber string) (*models.PhoneBook, error) {
	if m.OnGetPhoneBookByUserPhoneNumber != nil {
		return m.OnGetPhoneBookByUserPhoneNumber(ctx, phoneNumber)
	}
	return nil, nil
}
