package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"wallet-app/service"
)

type WalletHandler struct {
	svc *service.WalletService
}

func NewWalletHandler(svc *service.WalletService) *WalletHandler {
	return &WalletHandler{svc: svc}
}

func (h *WalletHandler) HandleTransaction(w http.ResponseWriter, r *http.Request) {
	var request struct {
		WalletId      uuid.UUID `json:"walletId"`
		OperationType string    `json:"operationType"`
		Amount        float64   `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.svc.ProcessTransaction(request.WalletId, request.OperationType, request.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *WalletHandler) HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	walletIDStr := r.URL.Path[len("/api/v1/wallets/"):]
	walletID, err := uuid.Parse(walletIDStr)
	if err != nil {
		http.Error(w, "invalid wallet ID", http.StatusBadRequest)
		return
	}

	balance, err := h.svc.GetBalance(walletID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]float64{"balance": balance}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
