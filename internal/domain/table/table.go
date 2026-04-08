package table

import "time"

type Table struct {
	ID           int64     `json:"id"`
	RestaurantID int64     `json:"restaurant_id"`
	Number       int       `json:"number"`
	Capacity     int       `json:"capacity"`
	CreatedAt    time.Time `json:"created_at"`
}
