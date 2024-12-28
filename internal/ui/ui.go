package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/nelsonmarro/gold-watcher/internal/contracts"
)

type Config struct {
	App            fyne.App
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	MainWindow     fyne.Window
	PriceContainer *fyne.Container
}

func (app *Config) MakeUI(goldService contracts.GoldService) {
	// get the current price of gold
	openPrice, currentPrice, priceChange := app.getPriceText(goldService)

	// put price information into a container
	priceContent := container.NewGridWithColumns(3,
		openPrice, currentPrice, priceChange)

	app.PriceContainer = priceContent

	// get toolbar
	toolBar := app.getToolBar()

	// add container to the window
	finalContent := container.NewVBox(priceContent, toolBar)
	app.MainWindow.SetContent(finalContent)
}
