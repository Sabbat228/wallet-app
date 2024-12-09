package service

import (
	"fmt"
	"github.com/google/uuid"
	"wallet-app/repository"
)

type WalletService struct {
	repo repository.WalletRepositoryInterface
}

func NewWalletService(repo repository.WalletRepositoryInterface) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) ProcessTransaction(walletID uuid.UUID, operationType string, amount float64) error {
	// Получаем текущий баланс
	balance, err := s.repo.GetBalance(walletID)
	if err != nil {
		return fmt.Errorf("could not get balance: %w", err)
	}

	switch operationType {
	case "DEPOSIT":
		return s.repo.UpdateBalance(walletID, amount)
	case "WITHDRAW":
		if balance < amount {
			return fmt.Errorf("insufficient funds: current balance is %.2f, but tried to withdraw %.2f", balance, amount)
		}
		return s.repo.UpdateBalance(walletID, -amount)
	default:
		return fmt.Errorf("invalid operation type: %s", operationType)
	}
}

func (s *WalletService) GetBalance(walletID uuid.UUID) (float64, error) {
	return s.repo.GetBalance(walletID)
}
