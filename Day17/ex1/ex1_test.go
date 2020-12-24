package main

import "testing"

func TestCountActive(t *testing.T) {
	target := 112
	s := Simulation{}
	s.readInitialState("../input_test1.dat")
	s.cycle(6)
	res := s.countActiveCubes()

	if target != res {
		t.Errorf("Got %d, expected %d.", res, target)
	}

}
