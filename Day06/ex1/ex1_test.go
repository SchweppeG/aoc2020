package main

import "testing"

func TestInput1(t *testing.T) {
	result := 6
	c := Customs{}
	c.CountAnswers("../input_test1.dat")

	if c.sumOfCounts != result {
		t.Errorf("Test Failed. Expected %d and got %d",
			result, c.sumOfCounts)
	}
}

func TestInput2(t *testing.T) {
	result := 11
	c := Customs{}
	c.CountAnswers("../input_test2.dat")

	if c.sumOfCounts != result {
		t.Errorf("Test Failed. Expected %d and got %d",
			result, c.sumOfCounts)
	}
}
