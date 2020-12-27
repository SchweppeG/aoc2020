package main

import "testing"

func TestProcessInput(t *testing.T) {
	target := 693942
	v := processInput("../input_test2.dat")

	if v != target {
		t.Errorf("Got %d, expected %d.", v, target)
	}

}
