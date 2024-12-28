package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"github.com/nelsonmarro/gold-watcher/internal/contracts"
)

type Config struct {
	App        fyne.App
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	MainWindow fyne.Window
}

func (app *Config) MakeUI(goldService contracts.GoldService) {
	// get the current price of gold

	// put price information into a container

	// add container to the window
}
