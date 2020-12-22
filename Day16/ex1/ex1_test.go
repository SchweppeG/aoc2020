package main

import "testing"

func TestScanTickets(t *testing.T) {
	target := 71
	tic := TicketCheck{}
	tic.readNotes("../input_test1.dat")
	errorRate := tic.scanTickets()

	if errorRate != target {
		t.Errorf("Got %d, expecte %d.\n", errorRate, target)
	}
}
