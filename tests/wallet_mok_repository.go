package tests

import (
	"errors"
	"github.com/google/uuid"
	"wallet-app/repository"
)

type MockWalletRepository struct {
	balance map[uuid.UUID]float64
}

func NewMockWalletRepository() *MockWalletRepository {
	return &MockWalletRepository{
		balance: make(map[uuid.UUID]float64),
	}
}

func (m *MockWalletRepository) CreateWallet(walletID uuid.UUID) error {
	if _, exists := m.balance[walletID]; exists {
		return nil
	}
	m.balance[walletID] = 0.0
	return nil
}

func (m *MockWalletRepository) UpdateBalance(walletID uuid.UUID, amount float64) error {
	if _, exists := m.balance[walletID]; !exists {
		return errors.New("wallet not found")
	}
	m.balance[walletID] += amount
	return nil
}

func (m *MockWalletRepository) GetBalance(walletID uuid.UUID) (float64, error) {
	balance, exists := m.balance[walletID]
	if !exists {
		return 0, errors.New("wallet not found")
	}

	return balance, nil
}

var _ repository.WalletRepositoryInterface = (*MockWalletRepository)(nil)
