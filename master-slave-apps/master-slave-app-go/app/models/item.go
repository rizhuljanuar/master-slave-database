package models


type Item struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Quantity    int       `json:"qty"`
	Price       float64   `json:"price"`
	CreatedAt   CustomTime `json:"created_at"`
	UpdatedAt   *CustomTime `json:"updated_at,omitempty"`
}