package models

import (
	"errors"
	"regexp"
	"strings"
)

// User represents a user in the system.
// Business rules:
// - PhoneNumber is the unique identifier for a user.
// - Name is the most recent name uploaded for this user.
// - IsSpam is set by a nightly job, not by user input.
type User struct {
	// PhoneNumber is the user's unique identifier. Must be a 10-digit string starting with "91".
	PhoneNumber string `json:"phone_number" validate:"required,len=12,startswith=91,numeric"`
	// Name is the user's name as uploaded from a phone book.
	Name string `json:"name" validate:"required,min=1,max=100"`
	// IsSpam indicates if the user is marked as spam (populated by nightly job).
	IsSpam bool `json:"is_spam"`
}

// Validate checks the User fields for business rule compliance.
func (u *User) Validate() error {
	if len(u.PhoneNumber) != 12 || !strings.HasPrefix(u.PhoneNumber, "91") {
		return errors.New("phone number must be 12 digits and start with '91'")
	}
	if !regexp.MustCompile(`^91[0-9]{10}$`).MatchString(u.PhoneNumber) {
		return errors.New("phone number must be numeric and 10 digits after '91'")
	}
	if len(strings.TrimSpace(u.Name)) == 0 {
		return errors.New("name is required")
	}
	if len(u.Name) > 100 {
		return errors.New("name must be at most 100 characters")
	}
	return nil
}

// GetPhoneNumber returns the user's phone number. Returns empty string if receiver is nil.
func (u *User) GetPhoneNumber() string {
	if u == nil {
		return ""
	}
	return u.PhoneNumber
}

// GetName returns the user's name. Returns empty string if receiver is nil.
func (u *User) GetName() string {
	if u == nil {
		return ""
	}
	return u.Name
}

// GetIsSpam returns the user's spam status. Returns false if receiver is nil.
func (u *User) GetIsSpam() bool {
	if u == nil {
		return false
	}
	return u.IsSpam
}
