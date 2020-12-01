package main

import "testing"

func TestFindSum(t *testing.T) {
	e := ExpenseResport{}
	e.ReadInput("..//input_test1.dat")

	_, _, prod := e.FindSum()
	if prod != 514579 {
		t.Errorf("Test failed. Got %d and expected %d", prod, 514579)
	}
}
