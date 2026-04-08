package handlers

import (
	"booking-api/internal/usecase/table"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CreateTableRequest struct {
	Number   int `json:"number"`
	Capacity int `json:"capacity"`
}

type TableHandler struct {
	createUC            *table.CreateTableUseCase
	getByIDUC           *table.GetTableByTableID
	getByRestaurantIDUC *table.GetTableByRestaurantID
}

func NewTableHandler(
	createUC *table.CreateTableUseCase,
	getByIDUC *table.GetTableByTableID,
	getByRestaurantIDUC *table.GetTableByRestaurantID,
) *TableHandler {
	return &TableHandler{
		createUC:            createUC,
		getByIDUC:           getByIDUC,
		getByRestaurantIDUC: getByRestaurantIDUC,
	}
}
func (h *TableHandler) Create(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	restIdStr := vars["id"]
	restID, err := strconv.ParseInt(restIdStr, 10, 64)
	if err != nil{
		writeError(w, http.StatusBadRequest, errors.New("invalid restaurant_id"))
		return
	}
	var req CreateTableRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err!=nil{
		writeError(w, http.StatusBadRequest, err)
		return
	}
	t, err := h.createUC.Execute(r.Context(),restID, req.Number, req.Capacity)
	if err != nil{
		writeError(w, http.StatusBadRequest, err)
		return
	}
	writeJSON(w, http.StatusCreated, t)
}
func (h *TableHandler) GetByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil{
		writeError(w, http.StatusBadRequest, errors.New("invalid id"))
		return 
	}
	t, err := h.getByIDUC.Execute(r.Context(), id)
	if err != nil{
		writeError(w, http.StatusBadRequest, err)
		return
	}
	writeJSON(w, http.StatusOK, t)
}
func (h *TableHandler) GetByRestaurantID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	restIdStr := vars["id"]
	restID, err := strconv.ParseInt(restIdStr, 10, 64)
	if err != nil{
		writeError(w, http.StatusBadRequest, errors.New("invalid id"))
		return 
	}
	t, err := h.getByRestaurantIDUC.Execute(r.Context(), restID)
	if err != nil{
		writeError(w, http.StatusBadRequest, err)
		return
	}
	writeJSON(w, http.StatusOK, t)
}