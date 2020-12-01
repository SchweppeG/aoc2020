package main

import "testing"

func TestFindSum(t *testing.T) {
	e := ExpenseResport{}

	e.ReadInput("..//input_test1.dat")
	_, _, prod := e.FindSum()
	if prod != 241861950 {
		t.Errorf("Test failed. Got %d expected %d", prod, 241861950)
	}
}
