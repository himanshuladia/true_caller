package mem

import (
	"context"
	"errors"
	"sync"

	"github.com/yourusername/truecaller-lite/pkg/dao"
	"github.com/yourusername/truecaller-lite/pkg/dao/daoerrors"
	"github.com/yourusername/truecaller-lite/pkg/models"
)

// PhoneBookMemDAO is a thread-safe in-memory implementation of PhoneBookDAO.
type PhoneBookMemDAO struct {
	mu        sync.RWMutex
	phonebook map[string]*models.PhoneBook // key: owner phone number
}

// NewPhoneBookMemDAO creates a new PhoneBookMemDAO instance.
func NewPhoneBookMemDAO() *PhoneBookMemDAO {
	return &PhoneBookMemDAO{
		phonebook: make(map[string]*models.PhoneBook),
	}
}

// CreateOrUpdatePhoneBook creates or updates a phone book for a user.
func (dao *PhoneBookMemDAO) CreateOrUpdatePhoneBook(ctx context.Context, pb *models.PhoneBook) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if err := pb.Validate(); err != nil {
		return err
	}
	owner := pb.GetPhoneNumber()
	if owner == "" {
		return errors.New("empty phone number")
	}
	dao.mu.Lock()
	defer dao.mu.Unlock()
	// Copy to avoid external mutation
	copyPB := *pb
	dao.phonebook[owner] = &copyPB
	return nil
}

// GetPhoneBookByUserPhoneNumber retrieves a phone book by owner's phone number.
func (dao *PhoneBookMemDAO) GetPhoneBookByUserPhoneNumber(ctx context.Context, phoneNumber string) (*models.PhoneBook, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	dao.mu.RLock()
	defer dao.mu.RUnlock()
	pb, ok := dao.phonebook[phoneNumber]
	if !ok {
		return nil, daoerrors.ErrPhoneBookNotFound
	}
	// Return a copy to avoid external mutation
	copyPB := *pb
	return &copyPB, nil
}

// Ensure PhoneBookMemDAO implements dao.PhoneBookDAO
var _ dao.PhoneBookDAO = (*PhoneBookMemDAO)(nil)
