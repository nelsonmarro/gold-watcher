package application

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

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
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)

	return toolBar
}

func (app *Config) addHoldingsDialog() dialog.Dialog {
	addAmountEntry := widget.NewEntry()
	purchaseDateEntry := widget.NewEntry()
	purchasePriceEntry := widget.NewEntry()

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
