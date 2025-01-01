package models

import "time"

type Holding struct {
	ID            int64     `json:"id"`
	Amount        int       `json:"amount"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PurchasePrice float64   `json:"purchase_price"`
}
