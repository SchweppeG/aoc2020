package main

import "testing"

func TestCountToTurn1(t *testing.T) {
	target := 175594
	r := Recitation{}
	r.readInput("../input_test1.dat")
	r.countToTurn(30000000)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}

}

func TestCountToTurn2(t *testing.T) {
	target := 2578
	r := Recitation{}
	r.readInput("../input_test2.dat")
	r.countToTurn(30000000)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}
}

func TestCountToTurn3(t *testing.T) {
	target := 3544142
	r := Recitation{}
	r.readInput("../input_test3.dat")
	r.countToTurn(30000000)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}
}

func TestCountToTurn4(t *testing.T) {
	target := 261214
	r := Recitation{}
	r.readInput("../input_test4.dat")
	r.countToTurn(30000000)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}
}

func TestCountToTurn5(t *testing.T) {
	target := 6895259
	r := Recitation{}
	r.readInput("../input_test5.dat")
	r.countToTurn(30000000)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}
}

func TestCountToTurn6(t *testing.T) {
	target := 18
	r := Recitation{}
	r.readInput("../input_test6.dat")
	r.countToTurn(30000000)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}
}

func TestCountToTurn7(t *testing.T) {
	target := 362
	r := Recitation{}
	r.readInput("../input_test7.dat")
	r.countToTurn(30000000)

	if r.last_num != target {
		t.Errorf("Got %d, expected %d.\n", r.last_num, target)
	}
}
