package main

import (
	"booking-api/internal/repository/postgres"
	"booking-api/internal/transport/http/handlers"
	"booking-api/internal/usecase/restaurant"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})

}
func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
func main() {
	dsn := "postgres://booking_user:booking_pass@localhost:5433/booking_db?sslmode=disable"

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Не удалось создать пул соединений:", err)
	}
	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("Не удалось подключится к базе:", err)
	}
	log.Println("Подключено к PostgreSQL!")

	restaurantRepo := postgres.NewRestaurantRepository(pool)
	createRestaurantUC := restaurant.NewCreateRestaurantUseCase(restaurantRepo)
	getAllRestaurantsUC := restaurant.NewGetAllRestaurantsUseCase(restaurantRepo)
	getRestaurantsByID := restaurant.NewGetRestaurantByIDUseCase(restaurantRepo)

	restaurantHandler := handlers.NewRestaurantHandler(createRestaurantUC, getAllRestaurantsUC, getRestaurantsByID)

	router := mux.NewRouter()

	router.HandleFunc("/ping", pingHandler).Methods("GET")
	router.HandleFunc("/status", statusHandler).Methods("GET")

	router.HandleFunc("/restaurants", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			restaurantHandler.Create(w, r)
		} else if r.Method == http.MethodGet {
			restaurantHandler.GetAll(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}).Methods("GET", "POST")

	router.HandleFunc("/restaurants/{id:[0-9]+}", restaurantHandler.GetRestaurantByID).Methods("GET")

	log.Println("Server started on localhost 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed:", err)
	}

}
