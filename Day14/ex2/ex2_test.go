package main

import "testing"

func TestGetSum(t *testing.T) {
	target := int64(208)
	d := Docking{}
	d.readProgramm("../input_test2.dat")
	res := d.getSum()

	if res != target {
		t.Errorf("Got %d, expected %d", res, target)
	}

}
