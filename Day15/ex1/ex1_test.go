package main

import "testing"

func TestCountToTurn(t *testing.T) {
	target := 436
	r := Recitation{}
	r.readInput("../input_test1.dat")
	r.countToTurn(2020)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}

}
