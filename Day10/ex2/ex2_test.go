package main

import "testing"

func TestCountCombination1(t *testing.T) {
	target := 8
	a := Adapters{}
	a.readInput("../input_test1.dat")
	a.getJoltDiff()
	a.countCombination()

	if a.combinations != target {
		t.Errorf("Test failed. Expected %d, but got %d",
			target, a.combinations)
	}

}

func TestCountCombination2(t *testing.T) {
	target := 19208
	a := Adapters{}
	a.readInput("../input_test2.dat")
	a.getJoltDiff()
	a.countCombination()

	if a.combinations != target {
		t.Errorf("Test failed. Expected %d, but got %d",
			target, a.combinations)
	}

}
