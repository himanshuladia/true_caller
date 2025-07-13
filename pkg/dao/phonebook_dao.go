package dao

import "github.com/yourusername/truecaller-lite/pkg/models"

// PhoneBookDAO defines the data access contract for phone book operations.
// All methods return errors for data access or validation failures.
type PhoneBookDAO interface {
	// CreateOrUpdatePhoneBook creates a new phone book or updates an existing one for a user.
	// Params:
	//   phoneBook: the phone book object to create or update
	// Returns:
	//   error: if validation fails or storage error occurs
	// Example:
	//   err := dao.CreateOrUpdatePhoneBook(&models.PhoneBook{PhoneNumber: "919876543210", Contacts: [...]})
	CreateOrUpdatePhoneBook(phoneBook *models.PhoneBook) error

	// GetPhoneBookByUserPhoneNumber retrieves a phone book by the owner's phone number.
	// Params:
	//   phoneNumber: the owner's phone number (must be 12 digits, starts with 91)
	// Returns:
	//   phoneBook: the phone book object if found, or nil
	//   error: if not found or storage error occurs
	// Example:
	//   pb, err := dao.GetPhoneBookByUserPhoneNumber("919876543210")
	GetPhoneBookByUserPhoneNumber(phoneNumber string) (*models.PhoneBook, error)
}

// Error handling pattern: All methods return error for not found, validation, or storage errors. Use errors.Is for type checks.
