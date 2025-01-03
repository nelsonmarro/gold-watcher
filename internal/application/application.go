package application

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/nelsonmarro/gold-watcher/internal/repository"
	"github.com/nelsonmarro/gold-watcher/internal/services"
)

type Config struct {
	App                            fyne.App
	InfoLog                        *log.Logger
	ErrorLog                       *log.Logger
	MainWindow                     fyne.Window
	PriceContainer                 *fyne.Container
	ToolBar                        *widget.Toolbar
	PriceChartContainer            *fyne.Container
	Holdings                       [][]interface{}
	HoldingsTable                  *widget.Table
	GoldService                    services.GoldService
	HoldingRepository              repository.Repository
	AddHoldingsPurchaseAmountEntry *widget.Entry
	AddHoldingsPurchaseDateEntry   *widget.Entry
	AddHoldingsPurchasePriceEntry  *widget.Entry
}

func (app *Config) MakeUI() {
	// get the current price of gold
	openPrice, currentPrice, priceChange := app.getPriceText()

	// put price information into a container
	priceContent := container.NewGridWithColumns(3,
		openPrice, currentPrice, priceChange)

	app.PriceContainer = priceContent

	// get toolbar
	toolBar := app.getToolBar()
	app.ToolBar = toolBar

	priceTabContent := app.pricesTab()
	holdingsTabContent := app.holdingsTab()

	// get app tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), holdingsTabContent),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// add container to the window
	finalContent := container.NewVBox(priceContent, toolBar, tabs)
	app.MainWindow.SetContent(finalContent)

	go func() {
		for range time.Tick(time.Second * 30) {
			app.refreshPriceContent()
		}
	}()
}

func (app *Config) refreshPriceContent() {
	open, current, change := app.getPriceText()
	app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	app.PriceChartContainer.Refresh()
}

func (app *Config) refreshHoldingsTable() {
	app.Holdings = app.getHoldingsSlice()
	app.HoldingsTable.Refresh()
}
