package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"evermos-api/internal/entity"
	"evermos-api/internal/usecase"
	"github.com/gorilla/mux"
)

type AddressHandler struct {
	usecase usecase.AddressUsecase
}

func NewAddressHandler(u usecase.AddressUsecase) *AddressHandler { return &AddressHandler{u} }

func (h *AddressHandler) Create(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("user_id").(uint)
	var req entity.Address
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	if err := h.usecase.Create(uid, req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "alamat berhasil ditambahkan"})
}

func (h *AddressHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("user_id").(uint)
	adds, _ := h.usecase.GetAll(uid)
	json.NewEncoder(w).Encode(adds)
}

func (h *AddressHandler) Update(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("user_id").(uint)
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	var req entity.Address
	_ = json.NewDecoder(r.Body).Decode(&req)
	if err := h.usecase.Update(uid, uint(id), req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Alamat diperbarui"})
}

func (h *AddressHandler) Delete(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("user_id").(uint)
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	if err := h.usecase.Delete(uid, uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Alamat dihapus"})
}
