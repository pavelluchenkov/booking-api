package main

import (
	"booking-api/internal/repository/postgres"
	"booking-api/internal/transport/http/handlers"
	"booking-api/internal/usecase/restaurant"
	"booking-api/internal/usecase/table"
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
	updateRestaurant := restaurant.NewUpdateRestaurant(restaurantRepo)
	deleteRestaurant := restaurant.NewDeleteRestaurantUseCase(restaurantRepo)

	tablesRepo := postgres.NewTableRepository(pool)
	createTableUC := table.NewCreateTableUseCase(tablesRepo, restaurantRepo)
	getTableByTableIDUC := table.NewGetTableByTableID(tablesRepo)
	getTableByRestaurantIDUC := table.NewGetTableByRestaurantID(tablesRepo, restaurantRepo)

	restaurantHandler := handlers.NewRestaurantHandler(createRestaurantUC, getAllRestaurantsUC, getRestaurantsByID, updateRestaurant, deleteRestaurant)
	tablesHandler := handlers.NewTableHandler(createTableUC, getTableByTableIDUC, getTableByRestaurantIDUC)

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
	router.HandleFunc("/restaurants/{id:[0-9]+}", restaurantHandler.Update).Methods("PUT")
	router.HandleFunc("/restaurants/{id:[0-9]+}", restaurantHandler.Delete).Methods("DELETE")

	router.HandleFunc("/restaurants/{id:[0-9]+}/tables", tablesHandler.Create).Methods("POST")
	router.HandleFunc("/restaurants/{id:[0-9]+}/tables", tablesHandler.GetByRestaurantID).Methods("GET")
	router.HandleFunc("/tables/{id:[0-9]+}", tablesHandler.GetByID).Methods("GET")

	log.Println("Server started on localhost 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed:", err)
	}

}
