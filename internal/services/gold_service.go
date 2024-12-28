package services

import (
	"log"
	"time"

	"github.com/nelsonmarro/gold-watcher/internal/helpers"
	client "github.com/nelsonmarro/gold-watcher/internal/http"
	"github.com/nelsonmarro/gold-watcher/internal/models"
)

type GoldService struct {
	client *client.HttpClient
}

func NewGoldService(client *client.HttpClient) *GoldService {
	return &GoldService{client: client}
}

func (s *GoldService) GetPrices() (*models.Price, error) {
	data, err := s.client.Get(helpers.CURRENCY)
	if err != nil {
		log.Println("error contacting goldprice.org: ", err)
		return nil, err
	}

	gold, err := helpers.DeserializeJson[models.Gold](data)
	if err != nil {
		log.Println("error deserializing the response from goldprice.org: ", err)
		return nil, err
	}

	previous, current, change := gold.Prices[0].PreviousClose, gold.Prices[0].Price, gold.Prices[0].Change

	currentInfo := models.Price{
		Currency:      helpers.CURRENCY,
		Price:         current,
		Change:        change,
		PreviousClose: previous,
		Time:          time.Now(),
	}

	return &currentInfo, nil
}
