package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"wallet-app/config"
	"wallet-app/handler"
	"wallet-app/repository"
	"wallet-app/service"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("postgres", cfg.DBConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	migrate(db)

	repo := repository.NewWalletRepository(db)

	repo.CreateWallet(uuid.New())

	svc := service.NewWalletService(repo)

	wh := handler.NewWalletHandler(svc)

	http.HandleFunc("/api/v1/wallet", wh.HandleTransaction)
	http.HandleFunc("/api/v1/wallets/", wh.HandleGetBalance)

	log.Println("Starting server on :8081...")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

func migrate(db *sql.DB) {

	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'wallets')").Scan(&exists)
	if err != nil {
		log.Fatalf("Error checking if wallets table exists: %v", err)
	}

	if !exists {
		createTableQuery := `
			CREATE TABLE wallets (
			id UUID PRIMARY KEY,
			balance DECIMAL(10, 2) NOT NULL DEFAULT 0
		)`

		if _, err := db.Exec(createTableQuery); err != nil {
			log.Fatalf("Error creating wallets table: %v", err)
		}

		log.Println("Migration completed: wallets table created.")
	} else {
		log.Println("Migration skipped: wallets table already exists.")
	}
}
