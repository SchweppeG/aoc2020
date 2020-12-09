package main

import "testing"

func TestHandheldSolve(t *testing.T) {
	goal := int64(8)
	h := Handheld{}
	h.readOperations("../input_test1.dat")
	h.solve()

	if goal != h.accumulator {
		t.Errorf("Test Failed. Expected %d, but got %d.",
			goal, h.accumulator)
	}

}
