package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/nelsonmarro/gold-watcher/internal/application"
	"github.com/nelsonmarro/gold-watcher/internal/helpers"
	client "github.com/nelsonmarro/gold-watcher/internal/http"
	"github.com/nelsonmarro/gold-watcher/internal/repository"
	"github.com/nelsonmarro/gold-watcher/internal/services"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var myApp application.Config

	// create a fyne application
	a := app.NewWithID("lilim.code.goldwatcher.preferences")
	myApp.App = a

	// create our loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to the database
	db, err := createSqlConn(&myApp)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// dependency injection
	setupDependencyInjection(&myApp, db)

	// create and size a fyne window
	myApp.MainWindow = a.NewWindow("Gold Watcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.CenterOnScreen()
	myApp.MainWindow.SetMaster()

	myApp.MakeUI()

	helpers.Currency = a.Preferences().StringWithFallback("currency", "CAD")

	// show and run the application
	myApp.MainWindow.ShowAndRun()
}

func createSqlConn(app *application.Config) (*sql.DB, error) {
	path := ""

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
		app.InfoLog.Println("DB in: ", path)
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func setupDependencyInjection(app *application.Config, db *sql.DB) {
	client := client.NewHttpClient("https://data-asg.goldprice.org/dbXRates/", 5*time.Second)
	var goldService services.GoldService = services.NewGoldService(client)

	app.GoldService = goldService

	setupDB(app, db)
}

func setupDB(app *application.Config, db *sql.DB) {
	app.HoldingRepository = repository.NewHoldingRepository(db)

	dbInitializer := repository.NewDbInitializer(db)
	err := dbInitializer.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Panic(err)
	}
}
