package application

import (
	"os"
	"testing"

	"fyne.io/fyne/v2/test"

	"github.com/nelsonmarro/gold-watcher/internal/repository"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("Test")
	testApp.HoldingRepository = repository.NewTestRepository()
	os.Exit(m.Run())
}
