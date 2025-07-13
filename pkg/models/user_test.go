package models

import (
	"testing"
)

func TestUserValidate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name:    "valid user",
			user:    User{PhoneNumber: "919876543210", Name: "Alice"},
			wantErr: false,
		},
		{
			name:    "invalid phone number (too short)",
			user:    User{PhoneNumber: "9198765432", Name: "Bob"},
			wantErr: true,
		},
		{
			name:    "invalid phone number (wrong prefix)",
			user:    User{PhoneNumber: "929876543210", Name: "Carol"},
			wantErr: true,
		},
		{
			name:    "invalid phone number (non-numeric)",
			user:    User{PhoneNumber: "91abcdefghij", Name: "Dan"},
			wantErr: true,
		},
		{
			name:    "empty name",
			user:    User{PhoneNumber: "919876543210", Name: ""},
			wantErr: true,
		},
		{
			name:    "name too long",
			user:    User{PhoneNumber: "919876543210", Name: string(make([]byte, 101))},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.user.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error: %v, got: %v", tc.wantErr, err)
			}
		})
	}
}
