package mem

import (
	"context"
	"errors"
	"sync"

	"github.com/yourusername/truecaller-lite/pkg/dao"
	"github.com/yourusername/truecaller-lite/pkg/dao/daoerrors"
	"github.com/yourusername/truecaller-lite/pkg/models"
)

// UserMemDAO is a thread-safe in-memory implementation of UserDAO.
type UserMemDAO struct {
	mu    sync.RWMutex
	users map[string]*models.User // key: phone number
}

// NewUserMemDAO creates a new UserMemDAO instance.
func NewUserMemDAO() *UserMemDAO {
	return &UserMemDAO{
		users: make(map[string]*models.User),
	}
}

// CreateOrUpdateUser creates or updates a user by phone number.
func (dao *UserMemDAO) CreateOrUpdateUser(ctx context.Context, user *models.User) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if err := user.Validate(); err != nil {
		return err
	}
	phone := user.GetPhoneNumber()
	if phone == "" {
		return errors.New("empty phone number")
	}
	dao.mu.Lock()
	defer dao.mu.Unlock()
	// Copy to avoid external mutation
	copyUser := *user
	dao.users[phone] = &copyUser
	return nil
}

// GetUserByPhoneNumber retrieves a user by phone number.
func (dao *UserMemDAO) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	dao.mu.RLock()
	defer dao.mu.RUnlock()
	user, ok := dao.users[phoneNumber]
	if !ok {
		return nil, daoerrors.ErrUserNotFound
	}
	// Return a copy to avoid external mutation
	copyUser := *user
	return &copyUser, nil
}

// GetAllUsers returns all users.
func (dao *UserMemDAO) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	dao.mu.RLock()
	defer dao.mu.RUnlock()
	var result []*models.User
	for _, user := range dao.users {
		copyUser := *user
		result = append(result, &copyUser)
	}
	return result, nil
}

// UpdateSpamStatus updates the spam status for a user.
func (dao *UserMemDAO) UpdateSpamStatus(ctx context.Context, phoneNumber string, isSpam bool) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	dao.mu.Lock()
	defer dao.mu.Unlock()
	user, ok := dao.users[phoneNumber]
	if !ok {
		return daoerrors.ErrUserNotFound
	}
	user.IsSpam = isSpam
	return nil
}

// Ensure UserMemDAO implements dao.UserDAO
var _ dao.UserDAO = (*UserMemDAO)(nil)
