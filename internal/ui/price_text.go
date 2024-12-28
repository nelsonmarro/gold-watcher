package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/nelsonmarro/gold-watcher/internal/contracts"
	"github.com/nelsonmarro/gold-watcher/internal/helpers"
)

func (app *Config) getPriceText(goldService contracts.GoldService) (*canvas.Text, *canvas.Text, *canvas.Text) {
	var open, current, change *canvas.Text

	gold, err := goldService.GetPrices()
	if err != nil {
		grey := color.NRGBA{R: 155, G: 155, B: 155, A: 255}
		open = canvas.NewText("Open: Unreachable", grey)
		current = canvas.NewText("Current: Unreachable", grey)
		change = canvas.NewText("Change: Unreachable", grey)

	} else {
		displayColor := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		if gold.Price < gold.PreviousClose {
			displayColor = color.NRGBA{R: 180, G: 0, B: 0, A: 255}
		}

		openTxt := fmt.Sprintf("Open: $%.4f %s", gold.PreviousClose, helpers.CURRENCY)
		currentTxt := fmt.Sprintf("Current: $%.4f %s", gold.Price, helpers.CURRENCY)
		changeTxt := fmt.Sprintf("Change: $%.4f %s", gold.Change, helpers.CURRENCY)

		open = canvas.NewText(openTxt, nil)
		current = canvas.NewText(currentTxt, displayColor)
		change = canvas.NewText(changeTxt, displayColor)
	}

	// Align the texts on the canvas
	open.Alignment = fyne.TextAlignLeading
	current.Alignment = fyne.TextAlignCenter
	change.Alignment = fyne.TextAlignTrailing

	return open, current, change
}