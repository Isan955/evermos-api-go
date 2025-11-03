package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"evermos-api/internal/entity"
	"evermos-api/internal/middleware"
	"evermos-api/internal/usecase"

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	usecase usecase.TransactionUsecase
}

func NewTransactionHandler(u usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{usecase: u}
}

func (h *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var trx entity.Transaction
	if err := json.NewDecoder(r.Body).Decode(&trx); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.usecase.Create(userID, &trx); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(trx)
}

func (h *TransactionHandler) GetMyTransactions(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	trxs, err := h.usecase.GetMyTransactions(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(trxs)
}

func (h *TransactionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	trx, err := h.usecase.GetByID(uint(id), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(trx)
}

func (h *TransactionHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.usecase.Cancel(uint(id), userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "canceled"})
}
