package models

import "time"

type Item struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewItem(id uint, name string, price float64, created_at time.Time) Item {
	return Item{
		ID:        id,
		Name:      name,
		Price:     price,
		CreatedAt: created_at,
	}
}
