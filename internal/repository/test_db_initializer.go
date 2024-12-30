package repository

import "database/sql"

type testDBInitializer struct {
	db *sql.DB
}

func NewTestDbInitializer() *testDBInitializer {
	return &testDBInitializer{}
}

func (dbi *testDBInitializer) Migrate() error {
	return nil
}
