package handler

import (
	"encoding/json"
	"net/http"

	"evermos-api/internal/repository"
)

type StoreHandler struct {
	storeRepo repository.StoreRepository
}

func NewStoreHandler(sr repository.StoreRepository) *StoreHandler { return &StoreHandler{sr} }

func (h *StoreHandler) GetMyStore(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("user_id").(uint)
	store, err := h.storeRepo.FindByUserID(uid)
	if err != nil {
		http.Error(w, "toko tidak ditemukan", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(store)
}
