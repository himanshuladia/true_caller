package service

// SpamService defines the business logic contract for spam status updates.
// All methods return errors for validation or business rule violations.
type SpamService interface {
	// UpdateSpamStatus updates the spam status for all users based on internal logic or a data science model.
	// This method is intended to be called by a nightly job.
	// Params: none
	// Returns:
	//   error: if update fails due to storage or business rule errors
	// Example:
	//   err := service.UpdateSpamStatus()
	UpdateSpamStatus() error
}

// Error handling pattern: All methods return error for storage or business rule errors. Use errors.Is for type checks.
