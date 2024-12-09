package tests

import (
	"github.com/google/uuid"
	"testing"
	"wallet-app/service"
)

func TestProcessTransaction(t *testing.T) {
	walletID := uuid.New()
	mockRepo := NewMockWalletRepository()

	svc := service.NewWalletService(mockRepo)

	err := mockRepo.CreateWallet(walletID)
	if err != nil {
		t.Fatalf("expected no error when creating wallet, got %v", err)
	}

	err = svc.ProcessTransaction(walletID, "DEPOSIT", 50.0)
	if err != nil {
		t.Fatalf("expected no error on DEPOSIT, got %v", err)
	}

	balance, _ := mockRepo.GetBalance(walletID)
	if balance != 50.0 {
		t.Errorf("expected balance to be 50.0, got %v", balance)
	}

	err = svc.ProcessTransaction(walletID, "WITHDRAW", 30.0)
	if err != nil {
		t.Fatalf("expected no error on WITHDRAW, got %v", err)
	}

	balance, _ = mockRepo.GetBalance(walletID)
	if balance != 20.0 {
		t.Errorf("expected balance to be 20.0, got %v", balance)
	}

	err = svc.ProcessTransaction(walletID, "WITHDRAW", 50.0)
	if err == nil {
		t.Errorf("expected error for insufficient funds, got none")
	}

	nonexistentWalletID := uuid.New()
	err = svc.ProcessTransaction(nonexistentWalletID, "DEPOSIT", 50.0)
	if err == nil {
		t.Errorf("expected error for nonexistent wallet, got none")
	}
}
