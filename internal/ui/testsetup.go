package ui

import (
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	os.Exit(m.Run())
}
