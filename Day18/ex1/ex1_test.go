package main

import "testing"

func TestProcessInput1(t *testing.T) {
	target := 71
	v := processInput("../input_test1.dat")
	if v != target {
		t.Errorf("Got %d, expected %d.", v, target)
	}

}

func TestProcessInput2(t *testing.T) {
	target := 26335
	v := processInput("../input_test2.dat")
	if v != target {
		t.Errorf("Got %d, expected %d.", v, target)
	}
}
