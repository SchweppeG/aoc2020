package main

import "testing"

func TestCountToTurn1(t *testing.T) {
	target := 436
	r := Recitation{}
	r.readInput("../input_test1.dat")
	r.countToTurn(2020)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}

}

func TestCountToTurn2(t *testing.T) {
	target := 1
	r := Recitation{}
	r.readInput("../input_test2.dat")
	r.countToTurn(2020)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}
}

func TestCountToTurn3(t *testing.T) {
	target := 10
	r := Recitation{}
	r.readInput("../input_test3.dat")
	r.countToTurn(2020)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}
}

func TestCountToTurn4(t *testing.T) {
	target := 27
	r := Recitation{}
	r.readInput("../input_test4.dat")
	r.countToTurn(2020)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}
}

func TestCountToTurn5(t *testing.T) {
	target := 78
	r := Recitation{}
	r.readInput("../input_test5.dat")
	r.countToTurn(2020)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}
}

func TestCountToTurn6(t *testing.T) {
	target := 438
	r := Recitation{}
	r.readInput("../input_test6.dat")
	r.countToTurn(2020)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}
}

func TestCountToTurn7(t *testing.T) {
	target := 1836
	r := Recitation{}
	r.readInput("../input_test7.dat")
	r.countToTurn(2020)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}
}
