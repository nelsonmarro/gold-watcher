package contracts

import "github.com/nelsonmarro/gold-watcher/internal/models"

type GoldService interface {
	GetPrices() (*models.Price, error)
}
