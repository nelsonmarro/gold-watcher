package models

import "time"

type Holding struct {
	ID            int64     `json:"id"`
	Amount        float64   `json:"amount"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PurchasePrice int       `json:"purchase_price"`
}
