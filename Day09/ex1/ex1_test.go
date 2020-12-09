package main

import "testing"

func TestCheckCypher(t *testing.T) {
	target := int64(127)
	c := Cypher{}
	res := c.checkCypher(5, "../input_test1.dat")

	if target != res {
		t.Errorf("Test failed. Got %d, but expected %d.",
			res, target)
	}

}
