package models

import (
	"errors"
)

// Contact represents a single contact entry in a user's phone book.
// Business rules:
// - PhoneNumber must be a 10-digit string starting with "91".
// - Name must be non-empty and at most 100 characters.
type Contact struct {
	// PhoneNumber is the contact's unique identifier. Must be a 10-digit string starting with "91".
	PhoneNumber string `json:"phone_number" validate:"required,len=12,startswith=91,numeric"`
	// Name is the contact's name as uploaded from a phone book.
	Name string `json:"name" validate:"required,min=1,max=100"`
}

// Validate checks the Contact fields for business rule compliance.
func (c *Contact) Validate() error {
	user := User{PhoneNumber: c.PhoneNumber, Name: c.Name}
	return user.Validate()
}

// GetPhoneNumber returns the contact's phone number. Returns empty string if receiver is nil.
func (c *Contact) GetPhoneNumber() string {
	if c == nil {
		return ""
	}
	return c.PhoneNumber
}

// GetName returns the contact's name. Returns empty string if receiver is nil.
func (c *Contact) GetName() string {
	if c == nil {
		return ""
	}
	return c.Name
}

// PhoneBook represents a user's phone book (list of contacts).
// Business rules:
// - PhoneNumber is the owner of the phone book (the uploader's phone number).
// - Contacts is a list of contacts uploaded by this user.
type PhoneBook struct {
	// PhoneNumber is the owner's unique identifier. Must be a 10-digit string starting with "91".
	PhoneNumber string `json:"phone_number" validate:"required,len=12,startswith=91,numeric"`
	// Contacts is the list of contacts uploaded by this user.
	Contacts []Contact `json:"contacts" validate:"dive"`
}

// Validate checks the PhoneBook fields and all contained contacts for business rule compliance.
func (pb *PhoneBook) Validate() error {
	if err := (&User{PhoneNumber: pb.PhoneNumber, Name: "dummy"}).Validate(); err != nil {
		return errors.New("invalid phone book owner: " + err.Error())
	}
	for i, c := range pb.Contacts {
		if err := c.Validate(); err != nil {
			return errors.New("invalid contact at index " + string(rune(i)) + ": " + err.Error())
		}
	}
	return nil
}

// GetPhoneNumber returns the phone book owner's phone number. Returns empty string if receiver is nil.
func (pb *PhoneBook) GetPhoneNumber() string {
	if pb == nil {
		return ""
	}
	return pb.PhoneNumber
}

// GetContacts returns the contacts in the phone book. Returns nil if receiver is nil.
func (pb *PhoneBook) GetContacts() []Contact {
	if pb == nil {
		return nil
	}
	return pb.Contacts
}
