package main

import "testing"

func TestEarliestBus(t *testing.T) {
	target := 295
	s := Schedule{}
	s.readSchedule("../input_test1.dat")
	line, time := s.earliestBus()
	res := line * time
	if res != target {
		t.Errorf("Got %d, but expected %d.", res, target)
	}
}
