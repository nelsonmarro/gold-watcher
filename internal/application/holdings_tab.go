package application

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/nelsonmarro/gold-watcher/internal/models"
)

func (app *Config) holdingsTab() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Holdings content goes here"),
	)
}

func (app *Config) getHoldingsTable() *widget.Table {
	return nil
}

func (app *Config) getHoldingSlice() [][]interface{} {
	var slice [][]interface{}

	return slice
}

func (app *Config) currentHoldings() ([]models.Holding, error) {
	holdings, err := app.HoldingRepository.GetAll()
	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}

	return holdings, nil
}
