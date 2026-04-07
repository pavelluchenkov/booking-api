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
}

func NewRestaurantHandler(createUC *restaurant.CreateRestaurantUseCase) *RestaurantHandler {
	return &RestaurantHandler{createUC: createUC}
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
