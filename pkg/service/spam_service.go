package service

import "context"

// SpamService defines the business logic contract for spam status updates.
// All methods accept a context for timeouts and cancellations, and return errors for validation or business rule violations.
type SpamService interface {
	// UpdateSpamStatus updates the spam status for all users based on internal logic or a data science model.
	// This method is intended to be called by a nightly job.
	// Params:
	//   ctx: context for timeout/cancellation
	// Returns:
	//   error: if update fails due to storage or business rule errors
	// Example:
	//   err := service.UpdateSpamStatus(ctx)
	UpdateSpamStatus(ctx context.Context) error
}

// Error handling pattern: All methods return error for storage or business rule errors. Use errors.Is for type checks.
