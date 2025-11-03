package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"evermos-api/internal/usecase"
	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	usecase usecase.CategoryUsecase
}

func NewCategoryHandler(u usecase.CategoryUsecase) *CategoryHandler { return &CategoryHandler{u} }

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct{ Name string `json:"name"` }
	_ = json.NewDecoder(r.Body).Decode(&req)
	_ = h.usecase.Create(req.Name)
	json.NewEncoder(w).Encode(map[string]string{"message": "Kategori dibuat"})
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}
	name := r.URL.Query().Get("name")
	data, total, _ := h.usecase.GetAll(name, page, limit)
	json.NewEncoder(w).Encode(map[string]interface{}{"data": data, "total": total, "page": page, "limit": limit})
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var req struct{ Name string `json:"name"` }
	_ = json.NewDecoder(r.Body).Decode(&req)
	_ = h.usecase.Update(uint(id), req.Name)
	json.NewEncoder(w).Encode(map[string]string{"message": "Kategori diperbarui"})
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	_ = h.usecase.Delete(uint(id))
	json.NewEncoder(w).Encode(map[string]string{"message": "Kategori dihapus"})
}
