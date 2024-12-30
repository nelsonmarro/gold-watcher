package repository

import (
	"time"

	"github.com/nelsonmarro/gold-watcher/internal/models"
)

type TestRepository struct{}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (r *TestRepository) Create(holding models.Holding) (*models.Holding, error) {
	return &holding, nil
}

func (r *TestRepository) GetAll() ([]models.Holding, error) {
	holdings := make([]models.Holding, 0)
	return holdings, nil
}

func (r *TestRepository) GetByID(id int64) (*models.Holding, error) {
	h := models.Holding{
		Amount:        1,
		PurchasePrice: 1000,
		PurchaseDate:  time.Now(),
	}
	return &h, nil
}

func (r *TestRepository) Update(id int64, holding models.Holding) error {
	return nil
}

func (r *TestRepository) Delete(id int64) error {
	return nil
}
