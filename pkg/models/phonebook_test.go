package models

import (
	"testing"
)

func TestContactValidate(t *testing.T) {
	tests := []struct {
		name    string
		contact Contact
		wantErr bool
	}{
		{
			name:    "valid contact",
			contact: Contact{PhoneNumber: "919876543210", Name: "Bob"},
			wantErr: false,
		},
		{
			name:    "invalid phone number",
			contact: Contact{PhoneNumber: "9187654321", Name: "Bob"},
			wantErr: true,
		},
		{
			name:    "empty name",
			contact: Contact{PhoneNumber: "919876543210", Name: ""},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.contact.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error: %v, got: %v", tc.wantErr, err)
			}
		})
	}
}

func TestPhoneBookValidate(t *testing.T) {
	tests := []struct {
		name      string
		phoneBook PhoneBook
		wantErr   bool
	}{
		{
			name:      "valid phonebook",
			phoneBook: PhoneBook{PhoneNumber: "919876543210", Contacts: []Contact{{PhoneNumber: "919123456789", Name: "Alice"}}},
			wantErr:   false,
		},
		{
			name:      "invalid owner phone number",
			phoneBook: PhoneBook{PhoneNumber: "9187654321", Contacts: []Contact{{PhoneNumber: "919123456789", Name: "Alice"}}},
			wantErr:   true,
		},
		{
			name:      "invalid contact in phonebook",
			phoneBook: PhoneBook{PhoneNumber: "919876543210", Contacts: []Contact{{PhoneNumber: "9181234567", Name: "Bob"}}},
			wantErr:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.phoneBook.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error: %v, got: %v", tc.wantErr, err)
			}
		})
	}
}
