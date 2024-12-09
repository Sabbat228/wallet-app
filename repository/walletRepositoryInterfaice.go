package repository

import "github.com/google/uuid"

type WalletRepositoryInterface interface {
	CreateWallet(walletID uuid.UUID) error
	UpdateBalance(walletID uuid.UUID, amount float64) error
	GetBalance(walletID uuid.UUID) (float64, error)
}
