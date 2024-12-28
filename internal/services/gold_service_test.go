package services

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	client "github.com/nelsonmarro/gold-watcher/internal/http"
	"github.com/nelsonmarro/gold-watcher/test/mocks"
)

func TestGold_GetPrices(t *testing.T) {
	jsonToReturn := `
	{"ts":1735359675320,"tsj":1735359665827,"date":"Dec 27th 2024, 11:21:05 pm NY","items":[{"curr":"USD","xauPrice":2622.18,"xagPrice":29.3625,"chgXau":-12.4525,"chgXag":-0.3695,"pcXau":-0.4726,"pcXag":-1.2428,"xauClose":2634.6325,"xagClose":29.73197}]}
	`
	fakeResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)), // Simula el cuerpo de la respuesta
		Header:     make(http.Header),
	}

	mockTransport := &mocks.FakeTransport{
		Response: fakeResponse,
		Err:      nil,
	}

	mockClient := client.NewHttpClientWithTransport(15*time.Second, mockTransport)
	goldService := NewGoldService(mockClient)

	data, err := goldService.GetPrices()
	if err != nil {
		t.Errorf("error while getting the data: %v", err)
	}

	if data.Price != 2622.18 {
		t.Errorf("expected price: 2622.18, got: %v", data.Price)
	}
}
