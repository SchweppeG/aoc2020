package main

import "testing"

func TestHandheldRun(t *testing.T) {
	goal := int64(5)
	h := Handheld{}

	h.readOperations("../input_test1.dat")
	h.run()

	if h.accumulator != goal {
		t.Errorf("Test failed. Expected %d, but got %d",
			goal, h.accumulator)

	}

}
