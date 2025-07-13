package service

import (
	"context"

	"github.com/yourusername/truecaller-lite/pkg/dao"
)

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

// spamService implements SpamService interface.
type spamService struct {
	userDAO dao.UserDAO
}

// NewSpamService creates a new SpamService instance.
func NewSpamService(userDAO dao.UserDAO) SpamService {
	return &spamService{userDAO: userDAO}
}

// UpdateSpamStatus updates the spam status for all users based on a simple rule (simulate DS model).
func (s *spamService) UpdateSpamStatus(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	users, err := s.userDAO.GetAllUsers(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if !user.GetIsSpam() {
			if err := s.userDAO.UpdateSpamStatus(ctx, user.GetPhoneNumber(), true); err != nil {
				return err
			}
		}
	}
	return nil
}

var _ SpamService = (*spamService)(nil)
