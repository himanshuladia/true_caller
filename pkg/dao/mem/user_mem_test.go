package mem

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/yourusername/truecaller-lite/pkg/dao/daoerrors"
	"github.com/yourusername/truecaller-lite/pkg/models"
)

func TestUserMemDAO_HappyPath(t *testing.T) {
	dao := NewUserMemDAO()
	ctx := context.Background()
	user := &models.User{PhoneNumber: "919876543210", Name: "Alice"}
	if err := dao.CreateOrUpdateUser(ctx, user); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := dao.GetUserByPhoneNumber(ctx, "919876543210")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.GetName() != "Alice" {
		t.Errorf("expected name Alice, got %s", got.GetName())
	}
}

func TestUserMemDAO_NotFound(t *testing.T) {
	dao := NewUserMemDAO()
	ctx := context.Background()
	_, err := dao.GetUserByPhoneNumber(ctx, "919999999999")
	if err == nil || !errors.Is(err, daoerrors.ErrUserNotFound) {
		t.Errorf("expected user not found error, got %v", err)
	}
}

func TestUserMemDAO_ValidationError(t *testing.T) {
	dao := NewUserMemDAO()
	ctx := context.Background()
	user := &models.User{PhoneNumber: "123", Name: ""}
	if err := dao.CreateOrUpdateUser(ctx, user); err == nil {
		t.Error("expected validation error, got nil")
	}
}

func TestUserMemDAO_ContextCanceled(t *testing.T) {
	dao := NewUserMemDAO()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	user := &models.User{PhoneNumber: "919876543210", Name: "Alice"}
	if err := dao.CreateOrUpdateUser(ctx, user); err != context.Canceled {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}

func TestUserMemDAO_ConcurrentAccess(t *testing.T) {
	dao := NewUserMemDAO()
	ctx := context.Background()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			num := 919000000000 + int64(i)
			u := &models.User{PhoneNumber: string(rune(num)), Name: "User"}
			_ = dao.CreateOrUpdateUser(ctx, u)
			_, _ = dao.GetUserByPhoneNumber(ctx, u.GetPhoneNumber())
		}(i)
	}
	wg.Wait()
}

func TestUserMemDAO_UpdateSpamStatus(t *testing.T) {
	dao := NewUserMemDAO()
	ctx := context.Background()
	user := &models.User{PhoneNumber: "919876543210", Name: "Alice"}
	_ = dao.CreateOrUpdateUser(ctx, user)
	if err := dao.UpdateSpamStatus(ctx, "919876543210", true); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, _ := dao.GetUserByPhoneNumber(ctx, "919876543210")
	if !got.GetIsSpam() {
		t.Error("expected IsSpam true, got false")
	}
}
