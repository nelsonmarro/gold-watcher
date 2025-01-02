package services

import (
	"bytes"
	"image"
	"log"
	"time"

	"github.com/nelsonmarro/gold-watcher/internal/helpers"
	client "github.com/nelsonmarro/gold-watcher/internal/http"
	"github.com/nelsonmarro/gold-watcher/internal/models"
)

type GoldService interface {
	GetPrices() (*models.Price, error)
	GetGoldChartImage(URL string) (image.Image, error)
}

type goldService struct {
	client *client.HttpClient
}

func NewGoldService(client *client.HttpClient) *goldService {
	return &goldService{client: client}
}

func (s *goldService) GetPrices() (*models.Price, error) {
	data, err := s.client.Get(helpers.Currency, true)
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
		Currency:      helpers.Currency,
		Price:         current,
		Change:        change,
		PreviousClose: previous,
		Time:          time.Now(),
	}

	return &currentInfo, nil
}

func (s *goldService) GetGoldChartImage(URL string) (image.Image, error) {
	response, err := s.client.Get(URL, false)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(response))
	if err != nil {
		return nil, err
	}

	return img, nil
}
