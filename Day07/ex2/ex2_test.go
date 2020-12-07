package main

import "testing"

func TestCountRequiredBags(t *testing.T) {
	goal := 32
	r := Ruleset{}
	r.readRules("../input_test1.dat")
	result := r.countRequiredBags("shiny gold")

	if goal != result {
		t.Errorf("Test Failed. Expected %d, but got %d",
			goal, result)
	}

	goal = 126
	r = Ruleset{}
	r.readRules("../input_test2.dat")
	result = r.countRequiredBags("shiny gold")

	if goal != result {
		t.Errorf("Test Failed. Expected %d, but got %d",
			goal, result)
	}
}
