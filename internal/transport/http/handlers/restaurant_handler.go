package handlers

import (
	"booking-api/internal/usecase/restaurant"
	"encoding/json"
	"net/http"
)

type CreateRestaurantRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone,omitempty"`
}

type RestaurantHandler struct {
	createUC *restaurant.CreateRestaurantUseCase
	getAllUC *restaurant.GetAllRestaurantsUseCase
}

func NewRestaurantHandler(createUC *restaurant.CreateRestaurantUseCase, getAllUC *restaurant.GetAllRestaurantsUseCase) *RestaurantHandler {
	return &RestaurantHandler{
		createUC: createUC,
		getAllUC: getAllUC,
	}
}

func (h *RestaurantHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRestaurantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		writeError(w, http.StatusBadRequest, err)
		return
	}
	rest, err := h.createUC.Execute(r.Context(), req.Name, req.Address, req.Phone)
	if err != nil{
		writeError(w, http.StatusBadRequest, err)
		return
	}
	writeJSON(w, http.StatusCreated, rest)

}
func (h *RestaurantHandler) GetAll(w http.ResponseWriter, r *http.Request){
	restaurants, err := h.getAllUC.Execute(r.Context())
	if err != nil{
		writeError(w, http.StatusBadRequest, err)
		return
	}
	writeJSON(w, http.StatusOK, restaurants)
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
