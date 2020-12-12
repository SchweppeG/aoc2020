package main

import "testing"

func TestTravelDistance(t *testing.T) {
	target := 25
	s := Ship{}
	s.followInstructions("../input_test1.dat")
	result := s.getDistance()
	if target != result {
		t.Errorf("Got %d, expected %d", result, target)
	}

}
