package repository

import (
	"testing"
	"time"

	"github.com/nelsonmarro/gold-watcher/internal/models"
)

func TestHoldingsRepository_Create(t *testing.T) {
	h := models.Holding{
		Amount:        1,
		PurchasePrice: 1000,
		PurchaseDate:  time.Now(),
	}

	result, err := testRepo.Create(h)
	if err != nil {
		t.Errorf("Error creating holding: %v", err)
	}

	if result.ID <= 0 {
		t.Errorf("Invalid ID sent back: %v", result.ID)
	}
}

func TestHoldingsRepository_GetAll(t *testing.T) {
	h := models.Holding{
		Amount:        1,
		PurchasePrice: 1000,
		PurchaseDate:  time.Now(),
	}

	_, err := testRepo.Create(h)
	if err != nil {
		t.Errorf("Error creating holding: %v", err)
	}

	holdings, err := testRepo.GetAll()
	if err != nil {
		t.Errorf("Error getting holdings: %v", err)
	}

	if len(holdings) == 0 {
		t.Errorf("No holdings returned")
	}
}

func TestHoldingsRepository_GetById(t *testing.T) {
	h := models.Holding{
		Amount:        1,
		PurchasePrice: 1000,
		PurchaseDate:  time.Now(),
	}

	result, err := testRepo.Create(h)
	if err != nil {
		t.Errorf("Error creating holding: %v", err)
	}

	holding, err := testRepo.GetByID(result.ID)
	if err != nil {
		t.Errorf("Error getting holding: %v", err)
	}

	if holding.ID != result.ID {
		t.Errorf("Invalid ID sent back: %v", holding.ID)
	}
}

func TestHoldingsRepository_GetById_NotFound(t *testing.T) {
	_, err := testRepo.GetByID(999)
	if err == nil {
		t.Errorf("get one returned a holding that does not exist")
	}
}

func TestHoldingsRepository_Update(t *testing.T) {
	updateHolding := models.Holding{
		Amount:        2,
		PurchasePrice: 3000,
		PurchaseDate:  time.Now(),
	}

	err := testRepo.Update(1, updateHolding)
	if err != nil {
		t.Errorf("Error updating holding: %v", err)
	}

	dbHolding, err := testRepo.GetByID(1)
	if err != nil {
		t.Errorf("Error getting holding: %v", err)
	}

	if dbHolding.Amount != updateHolding.Amount {
		t.Errorf("Amount not updated")
	}

	if dbHolding.PurchasePrice != updateHolding.PurchasePrice {
		t.Errorf("PurchasePrice not updated")
	}
}

func TestHoldingsRepository_Delete(t *testing.T) {
	h := models.Holding{
		Amount:        1,
		PurchasePrice: 1000,
		PurchaseDate:  time.Now(),
	}

	result, err := testRepo.Create(h)
	if err != nil {
		t.Errorf("Error creating holding: %v", err)
	}

	err = testRepo.Delete(result.ID)
	if err != nil {
		t.Errorf("Error deleting holding: %v", err)
	}

	_, err = testRepo.GetByID(result.ID)
	if err == nil {
		t.Errorf("Holding not deleted")
	}
}
