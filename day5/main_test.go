package main

import "testing"

func TestSeatID(t *testing.T) {
	tables := []struct {
		input     string
		excpected int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
		{"BBBFFFBLRL", 0},
	}

	for _, table := range tables {
		seatID := SeatID(table.input)
		if seatID != table.excpected {
			t.Errorf("Seat id of %s was incorrect. got: %d, want: %d", table.input, seatID, table.excpected)
		}
	}
}
