package service

import (
	"context"
	"errors"
	"testing"

	"github.com/yourusername/truecaller-lite/pkg/dao/daoerrors"
	"github.com/yourusername/truecaller-lite/pkg/dao/mock"
	"github.com/yourusername/truecaller-lite/pkg/models"
)

// Test cases for UserService.UploadContacts
func TestUserService_UploadContacts(t *testing.T) {
	tests := []struct {
		name      string
		ctx       context.Context
		owner     string
		contacts  []models.Contact
		mockSetup func(m *mock.UserDAOMock, pb *mock.PhoneBookDAOMock)
		wantErr   bool
	}{
		{
			name:     "valid upload",
			ctx:      context.Background(),
			owner:    "919876543210",
			contacts: []models.Contact{{PhoneNumber: "919123456789", Name: "Bob"}},
			mockSetup: func(m *mock.UserDAOMock, pb *mock.PhoneBookDAOMock) {
				pb.OnCreateOrUpdatePhoneBook = func(ctx context.Context, phoneBook *models.PhoneBook) error { return nil }
			},
			wantErr: false,
		},
		{
			name:      "invalid owner phone number",
			ctx:       context.Background(),
			owner:     "123",
			contacts:  []models.Contact{{PhoneNumber: "919123456789", Name: "Bob"}},
			mockSetup: func(m *mock.UserDAOMock, pb *mock.PhoneBookDAOMock) {},
			wantErr:   true,
		},
		{
			name:      "invalid contact phone number",
			ctx:       context.Background(),
			owner:     "919876543210",
			contacts:  []models.Contact{{PhoneNumber: "123", Name: "Bob"}},
			mockSetup: func(m *mock.UserDAOMock, pb *mock.PhoneBookDAOMock) {},
			wantErr:   true,
		},
		{
			name:      "context canceled",
			ctx:       func() context.Context { c, cancel := context.WithCancel(context.Background()); cancel(); return c }(),
			owner:     "919876543210",
			contacts:  []models.Contact{{PhoneNumber: "919123456789", Name: "Bob"}},
			mockSetup: func(m *mock.UserDAOMock, pb *mock.PhoneBookDAOMock) {},
			wantErr:   true,
		},
		{
			name:     "DAO returns error",
			ctx:      context.Background(),
			owner:    "919876543210",
			contacts: []models.Contact{{PhoneNumber: "919123456789", Name: "Bob"}},
			mockSetup: func(m *mock.UserDAOMock, pb *mock.PhoneBookDAOMock) {
				pb.OnCreateOrUpdatePhoneBook = func(ctx context.Context, phoneBook *models.PhoneBook) error { return errors.New("dao error") }
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			userDAO := &mock.UserDAOMock{}
			phoneBookDAO := &mock.PhoneBookDAOMock{}
			if tc.mockSetup != nil {
				tc.mockSetup(userDAO, phoneBookDAO)
			}
			svc := NewUserService(userDAO, phoneBookDAO)
			err := svc.UploadContacts(tc.ctx, tc.owner, tc.contacts)
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error: %v, got: %v", tc.wantErr, err)
			}
		})
	}
}

// Test cases for UserService.LookupUser
func TestUserService_LookupUser(t *testing.T) {
	tests := []struct {
		name      string
		ctx       context.Context
		phone     string
		mockSetup func(m *mock.UserDAOMock)
		wantName  string
		wantSpam  bool
		wantErr   bool
	}{
		{
			name:  "valid lookup",
			ctx:   context.Background(),
			phone: "919876543210",
			mockSetup: func(m *mock.UserDAOMock) {
				m.OnGetUserByPhoneNumber = func(ctx context.Context, phone string) (*models.User, error) {
					return &models.User{PhoneNumber: phone, Name: "Alice", IsSpam: true}, nil
				}
			},
			wantName: "Alice",
			wantSpam: true,
			wantErr:  false,
		},
		{
			name:  "user not found",
			ctx:   context.Background(),
			phone: "919876543210",
			mockSetup: func(m *mock.UserDAOMock) {
				m.OnGetUserByPhoneNumber = func(ctx context.Context, phone string) (*models.User, error) {
					return nil, daoerrors.ErrUserNotFound
				}
			},
			wantName: "",
			wantSpam: false,
			wantErr:  true,
		},
		{
			name:      "invalid phone number",
			ctx:       context.Background(),
			phone:     "123",
			mockSetup: func(m *mock.UserDAOMock) {},
			wantName:  "",
			wantSpam:  false,
			wantErr:   true,
		},
		{
			name:      "context canceled",
			ctx:       func() context.Context { c, cancel := context.WithCancel(context.Background()); cancel(); return c }(),
			phone:     "919876543210",
			mockSetup: func(m *mock.UserDAOMock) {},
			wantName:  "",
			wantSpam:  false,
			wantErr:   true,
		},
		{
			name:  "DAO returns error",
			ctx:   context.Background(),
			phone: "919876543210",
			mockSetup: func(m *mock.UserDAOMock) {
				m.OnGetUserByPhoneNumber = func(ctx context.Context, phone string) (*models.User, error) {
					return nil, errors.New("dao error")
				}
			},
			wantName: "",
			wantSpam: false,
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			userDAO := &mock.UserDAOMock{}
			if tc.mockSetup != nil {
				tc.mockSetup(userDAO)
			}
			svc := NewUserService(userDAO, nil)
			name, spam, err := svc.LookupUser(tc.ctx, tc.phone)
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error: %v, got: %v", tc.wantErr, err)
			}
			if name != tc.wantName {
				t.Errorf("expected name: %s, got: %s", tc.wantName, name)
			}
			if spam != tc.wantSpam {
				t.Errorf("expected spam: %v, got: %v", tc.wantSpam, spam)
			}
		})
	}
}
