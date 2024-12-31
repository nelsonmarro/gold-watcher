package application

import (
	"fmt"
	"image/png"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/nelsonmarro/gold-watcher/internal/helpers"
	"github.com/nelsonmarro/gold-watcher/internal/resources"
)

func (app *Config) pricesTab() *fyne.Container {
	chart := app.getChart()
	chartContainer := container.NewVBox(chart)
	app.PriceChartContainer = chartContainer

	return chartContainer
}

func (app *Config) getChart() *canvas.Image {
	apiUrl := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(helpers.CURRENCY))
	var img *canvas.Image

	err := app.downloadFile(apiUrl, "gold_chart.png")
	if err != nil {
		// use bundle image
		img = canvas.NewImageFromResource(resources.ResourceUnreachablePng)
	} else {
		img = canvas.NewImageFromFile("gold_chart.png")
	}

	img.SetMinSize(fyne.NewSize(800, 500))
	img.FillMode = canvas.ImageFillStretch

	return img
}

func (app *Config) downloadFile(URL, fileName string) error {
	img, err := app.GoldService.GetGoldChartImage(URL)
	if err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("./%s", fileName))
	if err != nil {
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}
