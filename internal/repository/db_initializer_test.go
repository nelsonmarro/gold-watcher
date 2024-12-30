package repository

import "testing"

func TestDBInitializer_Migrate(t *testing.T) {
	err := testDBInitalizer.Migrate()
	if err != nil {
		t.Errorf("Error migrating database: %v", err)
	}
}
