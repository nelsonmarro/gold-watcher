package mocks

import (
	"net/http"
	"time"

	"github.com/nelsonmarro/gold-watcher/internal/models"
)

// fakeTransport simula el comportamiento de un transporte HTTP
type FakeTransport struct {
	Response *http.Response
	Err      error
}

func (f *FakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f.Response, f.Err
}

type MockGoldService struct {
	Price *models.Price
	Err   error
}

func NewMockGoldService() *MockGoldService {
	return &MockGoldService{
		Price: &models.Price{},
		Err:   nil,
	}
}

func (m *MockGoldService) GetPrices() (*models.Price, error) {
	m.Price.Change = -12.4525
	m.Price.Currency = "USD"
	m.Price.PreviousClose = 2634.6325
	m.Price.Price = 2622.18
	m.Price.Time = time.Now()

	return m.Price, m.Err
}
