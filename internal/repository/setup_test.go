package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var (
	testRepo         HoldingRepository
	testDBInitalizer DBInitializer
)

func TestMain(m *testing.M) {
	_ = os.Remove("./testdata/sql.db")
	path := "./testdata/sql.db"

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Println(err)
	}

	testDBInitalizer = NewDbInitializer(db)
	testRepo = NewHoldingRepository(db)
	os.Exit(m.Run())
}
