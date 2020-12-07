package main

import "testing"

func TestColorCount(t *testing.T) {
	goal := 4
	r := Ruleset{}
	r.readRules("../input_test1.dat")
	result := r.countBagsColor("shiny gold")

	if result != goal {
		t.Errorf("Test failed. Expeceted value %d but got %d.",
			goal, result)
	}

}
