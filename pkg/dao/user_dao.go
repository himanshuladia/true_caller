package dao

import (
	"context"

	"github.com/yourusername/truecaller-lite/pkg/models"
)

// UserDAO defines the data access contract for user-related operations.
// All methods accept a context for timeouts and cancellations, and return errors for data access or validation failures.
type UserDAO interface {
	// CreateOrUpdateUser creates a new user or updates an existing user by phone number.
	// Params:
	//   ctx: context for timeout/cancellation
	//   user: the user object to create or update
	// Returns:
	//   error: if validation fails or storage error occurs
	// Example:
	//   err := dao.CreateOrUpdateUser(ctx, &models.User{PhoneNumber: "919876543210", Name: "Alice"})
	CreateOrUpdateUser(ctx context.Context, user *models.User) error

	// GetUserByPhoneNumber retrieves a user by their phone number.
	// Params:
	//   ctx: context for timeout/cancellation
	//   phoneNumber: the user's phone number (must be 12 digits, starts with 91)
	// Returns:
	//   user: the user object if found, or nil
	//   error: if not found or storage error occurs
	// Example:
	//   user, err := dao.GetUserByPhoneNumber(ctx, "919876543210")
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error)

	// GetAllUsers returns all users in the system.
	// Params:
	//   ctx: context for timeout/cancellation
	// Returns:
	//   users: slice of all users
	//   error: if storage error occurs
	// Example:
	//   users, err := dao.GetAllUsers(ctx)
	GetAllUsers(ctx context.Context) ([]*models.User, error)

	// UpdateSpamStatus updates the spam status for a user.
	// Params:
	//   ctx: context for timeout/cancellation
	//   phoneNumber: the user's phone number
	//   isSpam: new spam status
	// Returns:
	//   error: if user not found or storage error occurs
	// Example:
	//   err := dao.UpdateSpamStatus(ctx, "919876543210", true)
	UpdateSpamStatus(ctx context.Context, phoneNumber string, isSpam bool) error
}

// Error handling pattern: All methods return error for not found, validation, or storage errors. Use errors.Is for type checks.
