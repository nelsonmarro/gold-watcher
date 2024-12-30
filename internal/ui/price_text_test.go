package ui

import (
	"testing"

	"github.com/nelsonmarro/gold-watcher/test/mocks"
)

func TestApp_getPriceText(t *testing.T) {
	goldService := mocks.NewMockGoldService()
	testApp.GoldService = goldService

	open, current, change := testApp.getPriceText()

	if open.Text != "Open: $2634.6325 USD" {
		t.Errorf("Expected Open: $2634.6325 USD, got %s", open.Text)
	}

	if current.Text != "Current: $2622.1800 USD" {
		t.Errorf("Expected Current: $2622.1800 USD, got %s", current.Text)
	}

	if change.Text != "Change: $-12.4525 USD" {
		t.Errorf("Expected Change: $-12.4525 USD, got %s", change.Text)
	}
}
