package service

import (
	"context"
	"errors"
	"testing"

	"github.com/yourusername/truecaller-lite/pkg/dao/mock"
	"github.com/yourusername/truecaller-lite/pkg/models"
)

func TestSpamService_UpdateSpamStatus(t *testing.T) {
	tests := []struct {
		name      string
		ctx       context.Context
		users     []*models.User
		mockSetup func(m *mock.SpamUserDAOMock)
		wantErr   bool
	}{
		{
			name:  "happy path - marks spam for some users",
			ctx:   context.Background(),
			users: []*models.User{{PhoneNumber: "919876543210", Name: "Alice", IsSpam: false}},
			mockSetup: func(m *mock.SpamUserDAOMock) {
				m.OnUpdateSpamStatus = func(ctx context.Context, phone string, isSpam bool) error { return nil }
				m.OnGetAllUsers = func(ctx context.Context) ([]*models.User, error) {
					return []*models.User{{PhoneNumber: "919876543210", Name: "Alice", IsSpam: false}}, nil
				}
			},
			wantErr: false,
		},
		{
			name:  "no users",
			ctx:   context.Background(),
			users: []*models.User{},
			mockSetup: func(m *mock.SpamUserDAOMock) {
				m.OnGetAllUsers = func(ctx context.Context) ([]*models.User, error) { return []*models.User{}, nil }
			},
			wantErr: false,
		},
		{
			name:      "context canceled",
			ctx:       func() context.Context { c, cancel := context.WithCancel(context.Background()); cancel(); return c }(),
			users:     []*models.User{{PhoneNumber: "919876543210", Name: "Alice", IsSpam: false}},
			mockSetup: func(m *mock.SpamUserDAOMock) {},
			wantErr:   true,
		},
		{
			name:  "DAO returns error on get all users",
			ctx:   context.Background(),
			users: nil,
			mockSetup: func(m *mock.SpamUserDAOMock) {
				m.OnGetAllUsers = func(ctx context.Context) ([]*models.User, error) { return nil, errors.New("dao error") }
			},
			wantErr: true,
		},
		{
			name:  "DAO returns error on update spam status",
			ctx:   context.Background(),
			users: []*models.User{{PhoneNumber: "919876543210", Name: "Alice", IsSpam: false}},
			mockSetup: func(m *mock.SpamUserDAOMock) {
				m.OnGetAllUsers = func(ctx context.Context) ([]*models.User, error) {
					return []*models.User{{PhoneNumber: "919876543210", Name: "Alice", IsSpam: false}}, nil
				}
				m.OnUpdateSpamStatus = func(ctx context.Context, phone string, isSpam bool) error { return errors.New("dao error") }
			},
			wantErr: true,
		},
		{
			name:  "all users already spam",
			ctx:   context.Background(),
			users: []*models.User{{PhoneNumber: "919876543210", Name: "Alice", IsSpam: true}},
			mockSetup: func(m *mock.SpamUserDAOMock) {
				m.OnGetAllUsers = func(ctx context.Context) ([]*models.User, error) {
					return []*models.User{{PhoneNumber: "919876543210", Name: "Alice", IsSpam: true}}, nil
				}
				m.OnUpdateSpamStatus = func(ctx context.Context, phone string, isSpam bool) error { return nil }
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			userDAO := &mock.SpamUserDAOMock{}
			if tc.mockSetup != nil {
				tc.mockSetup(userDAO)
			}
			svc := NewSpamService(userDAO)
			err := svc.UpdateSpamStatus(tc.ctx)
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error: %v, got: %v", tc.wantErr, err)
			}
		})
	}
}
