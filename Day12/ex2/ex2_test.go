package main

import "testing"

func TestGetDistance(t *testing.T) {
	target := 286

	s := Ship{}
	s.followInstructions("../input_test1.dat")
	ret := s.getDistance()

	if target != ret {
		t.Errorf("Got %d, expected %d", ret, target)
	}

}
