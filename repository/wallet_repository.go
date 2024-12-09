package repository

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"sync"
)

type WalletRepository struct {
	db *sql.DB
	mu sync.Mutex
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) CreateWallet(walletID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec("INSERT INTO wallets (id, balance) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING", walletID, 0)
	return err
}

func (r *WalletRepository) UpdateBalance(walletID uuid.UUID, amount float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec("UPDATE wallets SET balance = balance + $1 WHERE id = $2", amount, walletID)
	return err
}

func (r *WalletRepository) GetBalance(walletID uuid.UUID) (float64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var balance float64
	err := r.db.QueryRow("SELECT balance FROM wallets WHERE id = $1", walletID).Scan(&balance)
	return balance, err
}

var _ WalletRepositoryInterface = (*WalletRepository)(nil)
