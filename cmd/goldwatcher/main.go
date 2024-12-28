package main

import (
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/nelsonmarro/gold-watcher/internal/contracts"
	client "github.com/nelsonmarro/gold-watcher/internal/http"
	"github.com/nelsonmarro/gold-watcher/internal/services"
	"github.com/nelsonmarro/gold-watcher/internal/ui"
)

var myApp ui.Config

func main() {
	// create a fyne application
	a := app.NewWithID("lilim.code.goldwatcher.preferences")
	myApp.App = a

	// create our loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to the database

	// create a database repository

	// dependency injection
	client := client.NewHttpClient(5 * time.Second)
	var goldService contracts.GoldService = services.NewGoldService(client)

	// create and size a fyne window
	myApp.MainWindow = a.NewWindow("Gold Watcher")
	myApp.MainWindow.Resize(fyne.NewSize(800, 500))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.CenterOnScreen()
	myApp.MainWindow.SetMaster()

	myApp.MakeUI(goldService)

	// show and run the application
	myApp.MainWindow.ShowAndRun()
}
