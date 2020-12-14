package main

import "testing"

func TestGetSum(t *testing.T) {
	target := int64(165)
	d := Docking{}
	d.readProgramm("../input_test1.dat")
	result := d.getSum()

	if target != result {
		t.Errorf("Got %d, expected %d.", result, target)
	}

}
