package application

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/nelsonmarro/gold-watcher/internal/models"
)

func (app *Config) holdingsTab() *fyne.Container {
	app.Holdings = app.getHoldingsSlice()
	app.HoldingsTable = app.getHoldingsTable()

	holdingsCtr := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, app.HoldingsTable),
	)

	return holdingsCtr
}

func (app *Config) getHoldingsTable() *widget.Table {
	t := widget.NewTable(
		func() (int, int) {
			return len(app.Holdings), len(app.Holdings[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == len(app.Holdings[0])-1 && i.Row != 0 {
				// last cell - put a button
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete?", "", func(deleted bool) {
						if deleted {
							id, _ := strconv.Atoi(app.Holdings[i.Row][0].(string))
							err := app.HoldingRepository.Delete(int64(id))
							if err != nil {
								app.ErrorLog.Println(err)
							}
						}
						app.refreshHoldingsTable()
					}, app.MainWindow)
				})
				w.Importance = widget.HighImportance

				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					w,
				}
			} else {
				// put in textual information
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(app.Holdings[i.Row][i.Col].(string)),
				}
			}
		})

	colWidths := []float32{50, 200, 200, 200, 110}
	for i := 0; i < len(colWidths); i++ {
		t.SetColumnWidth(i, colWidths[i])
	}

	return t
}

func (app *Config) getHoldingsSlice() [][]interface{} {
	var slice [][]interface{}

	holdings, err := app.currentHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	slice = append(slice, []interface{}{"ID", "Amount", "Price", "Date", "Options"})

	for _, x := range holdings {
		var currentRow []interface{}

		currentRow = append(currentRow, strconv.FormatInt(x.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("%v toz", x.Amount))
		currentRow = append(currentRow, fmt.Sprintf("$%2f", float32(x.PurchasePrice/100)))
		currentRow = append(currentRow, x.PurchaseDate.Format("2006-01-02"))
		currentRow = append(currentRow, widget.NewButton("Delete", func() {
		}))

		slice = append(slice, currentRow)
	}

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
