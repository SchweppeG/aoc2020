package main

import "testing"

func TestValidPassports(t *testing.T) {
	result := 2
	p := PassportScanner{}
	p.readPassports("../input_test1.dat")

	if p.nValid != result {
		t.Errorf("Test failed. Got %d and expected %d.",
			p.nValid, result)
	}
}
