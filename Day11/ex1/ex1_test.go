package main

import "testing"

func TestSolve(t *testing.T) {
	target := 37
	s := Seating{}
	s.readLayer("../input_test1.dat")
	result := s.solve()

	if result != target {
		t.Errorf("Expected %d, but got %d.", target, result)
	}
}
