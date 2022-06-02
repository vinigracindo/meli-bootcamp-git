package products

import "time"

type Product struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	Code      string    `json:"code"`
	Published bool      `json:"published"`
	CreatedAt time.Time `json:"created_at"`
}
