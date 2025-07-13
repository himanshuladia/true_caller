package service

import (
	"context"

	"github.com/yourusername/truecaller-lite/pkg/models"
)

// UserService defines the business logic contract for user-related operations.
// All methods accept a context for timeouts and cancellations, and return errors for validation or business rule violations.
type UserService interface {
	// UploadContacts uploads a list of contacts for a user (by phone number).
	// Params:
	//   ctx: context for timeout/cancellation
	//   ownerPhoneNumber: the uploader's phone number (must be 12 digits, starts with 91)
	//   contacts: slice of contacts to upload
	// Returns:
	//   error: if validation fails or business rule is violated
	// Example:
	//   err := service.UploadContacts(ctx, "919876543210", []models.Contact{...})
	UploadContacts(ctx context.Context, ownerPhoneNumber string, contacts []models.Contact) error

	// LookupUser looks up a user by phone number and returns their name and spam status.
	// Params:
	//   ctx: context for timeout/cancellation
	//   phoneNumber: the phone number to look up (must be 12 digits, starts with 91)
	// Returns:
	//   name: the most recent name for the user
	//   isSpam: the current spam status
	//   error: if not found or validation fails
	// Example:
	//   name, isSpam, err := service.LookupUser(ctx, "919876543210")
	LookupUser(ctx context.Context, phoneNumber string) (name string, isSpam bool, err error)
}

// Error handling pattern: All methods return error for not found, validation, or business rule errors. Use errors.Is for type checks.
