package main

import "testing"

func TestGetJoltDiff1(t *testing.T) {
	goalCount1 := 7
	goalCount3 := 5
	goalRating := 22

	a := Adapters{}
	a.readInput("../input_test1.dat")
	a.getJoltDiff()

	if (goalCount1 != a.count1) ||
		(goalCount3 != a.count3) ||
		(goalRating != a.rating) {
		t.Errorf("Test Failed. Expected %d, %d, %d for count1,"+
			"count2 and rating, respectively,"+
			" but got %d,%d,%d.\n",
			goalCount1,
			goalCount3,
			goalRating,
			a.count1,
			a.count2,
			a.rating)
	}
}

func TestGetJoltDiff2(t *testing.T) {
	goalCount1 := 22
	goalCount3 := 10

	a := Adapters{}
	a.readInput("../input_test2.dat")
	a.getJoltDiff()

	if (goalCount1 != a.count1) ||
		(goalCount3 != a.count3) {
		t.Errorf("Test Failed. Expected %d, %d  for count1,"+
			"count2 and rating, respectively,"+
			" but got %d,%d.\n",
			goalCount1,
			goalCount3,
			a.count1,
			a.count2)
	}
}
