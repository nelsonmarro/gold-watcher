package repository

import "database/sql"

type DBInitializer interface {
	Migrate() error
}

type dbInitializer struct {
	db *sql.DB
}

func NewDbInitializer(db *sql.DB) *dbInitializer {
	return &dbInitializer{db: db}
}

func (dbi *dbInitializer) Migrate() error {
	_, err := dbi.db.Exec(`CREATE TABLE IF NOT EXISTS holdings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		purchase_date INTEGER NOT NULL,
		purchase_price INTEGER NOT NULL
	);`)
	return err
}
