package handlers

import (
	"booking-api/internal/usecase/restaurant"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type CreateRestaurantRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone,omitempty"`
}

type RestaurantHandler struct {
	createUC            *restaurant.CreateRestaurantUseCase
	getAllUC            *restaurant.GetAllRestaurantsUseCase
	getRestaurantByIDUC *restaurant.GetRestaurantByIDUseCase
}

func NewRestaurantHandler(createUC *restaurant.CreateRestaurantUseCase, getAllUC *restaurant.GetAllRestaurantsUseCase, getRestaurantByIDUC *restaurant.GetRestaurantByIDUseCase) *RestaurantHandler {
	return &RestaurantHandler{
		createUC:            createUC,
		getAllUC:            getAllUC,
		getRestaurantByIDUC: getRestaurantByIDUC,
	}
}

func (h *RestaurantHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRestaurantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	rest, err := h.createUC.Execute(r.Context(), req.Name, req.Address, req.Phone)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	writeJSON(w, http.StatusCreated, rest)

}
func (h *RestaurantHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	restaurants, err := h.getAllUC.Execute(r.Context())
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	writeJSON(w, http.StatusOK, restaurants)
}
func (h *RestaurantHandler) GetRestaurantByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid id format"))
		return
	}
	restaurant, err := h.getRestaurantByIDUC.Execute(r.Context(), id)
	 if err != nil {
        if err.Error() == "restaurant not found" {
            writeError(w, http.StatusNotFound, err)
            return
        }
        writeError(w, http.StatusInternalServerError, err)
        return
    }

	writeJSON(w, http.StatusOK, restaurant)

}
func writeError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
