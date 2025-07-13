package service

import (
	"context"
	"errors"

	"github.com/yourusername/truecaller-lite/pkg/dao"
	"github.com/yourusername/truecaller-lite/pkg/dao/daoerrors"
	"github.com/yourusername/truecaller-lite/pkg/models"
)

// UserService defines the business logic contract for user-related operations.
// All methods accept a context for timeouts and cancellations, and return errors for validation or business rule violations.
type UserService interface {
	UploadContacts(ctx context.Context, ownerPhoneNumber string, contacts []models.Contact) error
	LookupUser(ctx context.Context, phoneNumber string) (name string, isSpam bool, err error)
}

// userService implements UserService interface.
type userService struct {
	userDAO      dao.UserDAO
	phoneBookDAO dao.PhoneBookDAO
}

// NewUserService creates a new UserService instance.
func NewUserService(userDAO dao.UserDAO, phoneBookDAO dao.PhoneBookDAO) UserService {
	return &userService{userDAO: userDAO, phoneBookDAO: phoneBookDAO}
}

// UploadContacts uploads a list of contacts for a user (by phone number).
func (s *userService) UploadContacts(ctx context.Context, ownerPhoneNumber string, contacts []models.Contact) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	owner := &models.User{PhoneNumber: ownerPhoneNumber, Name: "dummy"}
	if err := owner.Validate(); err != nil {
		return err
	}
	for _, c := range contacts {
		if err := c.Validate(); err != nil {
			return err
		}
	}
	pb := &models.PhoneBook{PhoneNumber: ownerPhoneNumber, Contacts: contacts}
	if err := s.phoneBookDAO.CreateOrUpdatePhoneBook(ctx, pb); err != nil {
		return err
	}
	return nil
}

// LookupUser looks up a user by phone number and returns their name and spam status.
func (s *userService) LookupUser(ctx context.Context, phoneNumber string) (string, bool, error) {
	if ctx.Err() != nil {
		return "", false, ctx.Err()
	}
	tmp := &models.User{PhoneNumber: phoneNumber, Name: "dummy"}
	if err := tmp.Validate(); err != nil {
		return "", false, err
	}
	user, err := s.userDAO.GetUserByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		if errors.Is(err, daoerrors.ErrUserNotFound) {
			return "", false, err
		}
		return "", false, err
	}
	return user.GetName(), user.GetIsSpam(), nil
}

var _ UserService = (*userService)(nil)
