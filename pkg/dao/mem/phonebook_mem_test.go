package mem

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/yourusername/truecaller-lite/pkg/dao/daoerrors"
	"github.com/yourusername/truecaller-lite/pkg/models"
)

func TestPhoneBookMemDAO_HappyPath(t *testing.T) {
	dao := NewPhoneBookMemDAO()
	ctx := context.Background()
	pb := &models.PhoneBook{PhoneNumber: "919876543210", Contacts: []models.Contact{{PhoneNumber: "919123456789", Name: "Bob"}}}
	if err := dao.CreateOrUpdatePhoneBook(ctx, pb); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := dao.GetPhoneBookByUserPhoneNumber(ctx, "919876543210")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.GetContacts()) != 1 || got.GetContacts()[0].GetName() != "Bob" {
		t.Errorf("expected contact Bob, got %+v", got.GetContacts())
	}
}

func TestPhoneBookMemDAO_NotFound(t *testing.T) {
	dao := NewPhoneBookMemDAO()
	ctx := context.Background()
	_, err := dao.GetPhoneBookByUserPhoneNumber(ctx, "919999999999")
	if err == nil || !errors.Is(err, daoerrors.ErrPhoneBookNotFound) {
		t.Errorf("expected phone book not found error, got %v", err)
	}
}

func TestPhoneBookMemDAO_ValidationError(t *testing.T) {
	dao := NewPhoneBookMemDAO()
	ctx := context.Background()
	pb := &models.PhoneBook{PhoneNumber: "123", Contacts: nil}
	if err := dao.CreateOrUpdatePhoneBook(ctx, pb); err == nil {
		t.Error("expected validation error, got nil")
	}
}

func TestPhoneBookMemDAO_ContextCanceled(t *testing.T) {
	dao := NewPhoneBookMemDAO()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	pb := &models.PhoneBook{PhoneNumber: "919876543210", Contacts: nil}
	if err := dao.CreateOrUpdatePhoneBook(ctx, pb); err != context.Canceled {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}

func TestPhoneBookMemDAO_ConcurrentAccess(t *testing.T) {
	dao := NewPhoneBookMemDAO()
	ctx := context.Background()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			num := 919000000000 + int64(i)
			pb := &models.PhoneBook{PhoneNumber: string(rune(num)), Contacts: []models.Contact{{PhoneNumber: "919123456789", Name: "Bob"}}}
			_ = dao.CreateOrUpdatePhoneBook(ctx, pb)
			_, _ = dao.GetPhoneBookByUserPhoneNumber(ctx, pb.GetPhoneNumber())
		}(i)
	}
	wg.Wait()
}
