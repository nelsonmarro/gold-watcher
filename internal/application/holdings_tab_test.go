package application

import (
	"slices"
	"testing"
)

func TestApp_GetHoldings(t *testing.T) {
	holdings, err := testApp.currentHoldings()
	if err != nil {
		t.Error("failed to get current holdings from db:", err)
	}

	if len(holdings) != 2 {
		t.Errorf("expected holdings slice to return 2 values, got:%v", len(holdings))
	}
}

func TestApp_GetHoldingsSlice(t *testing.T) {
	slice := testApp.getHoldingsSlice()

	if len(slice) != 3 {
		t.Errorf("expected to get a holdings slice with 3 rows, but got: %v", len(slice))
	}

	lable_row := slice[0]
	expected_lable_row := []interface{}{"ID", "Amount", "Price", "Date", "Options"}

	if !slices.Equal(lable_row, expected_lable_row) {
		t.Errorf("wrong values on table's lable row, expected: ['ID', 'Amount', 'Price', 'Date', 'Options'], but got: %s", lable_row)
	}
}
