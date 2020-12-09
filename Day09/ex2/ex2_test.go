package main

import "testing"

func TestBreakCypher(t *testing.T) {
	target := int64(62)
	c := Cypher{}
	invalid := c.checkCypher(5, "../input_test1.dat")
	code := c.breakCypher(invalid)

	if code != target {
		t.Errorf("Test failed. Got %d, exepcted %d",
			code, target)
	}
}
