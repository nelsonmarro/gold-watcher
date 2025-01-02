package application

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/nelsonmarro/gold-watcher/internal/helpers"
	"github.com/nelsonmarro/gold-watcher/internal/models"
)

func (app *Config) getToolBar() *widget.Toolbar {
	toolBar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			app.addHoldingsDialog()
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			app.refreshPriceContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			w := app.showPreferences()
			w.Resize(fyne.Size{Width: 300, Height: 200})
			w.Show()
		}),
	)

	return toolBar
}

func (app *Config) showPreferences() fyne.Window {
	win := app.App.NewWindow("Preferences")

	lbl := widget.NewLabel("Preferred Currency")
	curr := widget.NewSelect([]string{
		"USD",
		"CAD",
		"GBP",
	}, func(value string) {
		helpers.Currency = value
		app.App.Preferences().SetString("currency", value)
	})
	curr.Selected = helpers.Currency

	btn := widget.NewButton("Save", func() {
		win.Close()
		app.refreshPriceContent()
	})
	btn.Importance = widget.HighImportance

	win.SetContent(container.NewVBox(lbl, curr, btn))

	return win
}

func (app *Config) addHoldingsDialog() dialog.Dialog {
	purchaseDateEntry := widget.NewEntry()
	purchasePriceEntry := widget.NewEntry()
	addAmountEntry := widget.NewEntry()

	app.AddHoldingsPurchaseAmountEntry = addAmountEntry
	app.AddHoldingsPurchaseDateEntry = purchaseDateEntry
	app.AddHoldingsPurchasePriceEntry = purchasePriceEntry

	dateValidator := func(s string) error {
		if _, err := time.Parse("2006-01-02", s); err != nil {
			return err
		}
		return nil
	}
	purchaseDateEntry.Validator = dateValidator

	isIntValidator := func(s string) error {
		_, err := strconv.Atoi(s)
		if err != nil {
			return err
		}

		return nil
	}
	addAmountEntry.Validator = isIntValidator

	isFloatValidator := func(s string) error {
		_, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return err
		}

		return nil
	}
	purchasePriceEntry.Validator = isFloatValidator

	purchaseDateEntry.PlaceHolder = "YYYY-MM-DD"

	// create a Dialog
	addForm := dialog.NewForm("Add Gold", "Add", "Cancel", []*widget.FormItem{
		{Text: "Amount in toz", Widget: addAmountEntry},
		{Text: "Purchase Price", Widget: purchasePriceEntry},
		{Text: "Purchase Date", Widget: purchaseDateEntry},
	},
		func(valid bool) {
			if valid {
				amount, _ := strconv.Atoi(addAmountEntry.Text)
				date, _ := time.Parse("2006-01-02", purchaseDateEntry.Text)
				price, _ := strconv.ParseFloat(purchasePriceEntry.Text, 64)
				price *= 100

				_, err := app.HoldingRepository.Create(models.Holding{
					Amount:        amount,
					PurchaseDate:  date,
					PurchasePrice: price,
				})
				if err != nil {
					app.ErrorLog.Println(err)
				}
				app.refreshHoldingsTable()
			}
		}, app.MainWindow)

	// size and show dialog
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}
