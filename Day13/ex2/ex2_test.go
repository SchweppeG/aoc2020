package main

import "testing"

func TestEarliestBus1(t *testing.T) {
	target := int(1068781)
	s := Schedule{}
	s.readSchedule("../input_test1.dat")
	result := s.earliestBus(1068000)

	if target != result {
		t.Errorf("Got %d, expected %d\n", result, target)
	}
}

func TestEarliestBus2(t *testing.T) {
	target := int(3417)
	s := Schedule{}
	s.readSchedule("../input_test2.dat")
	result := s.earliestBus(3000)

	if target != result {
		t.Errorf("Got %d, expected %d\n", result, target)
	}
}

func TestEarliestBus3(t *testing.T) {
	target := int(754018)
	s := Schedule{}
	s.readSchedule("../input_test3.dat")
	result := s.earliestBus(754000)

	if target != result {
		t.Errorf("Got %d, expected %d\n", result, target)
	}
}

func TestEarliestBus4(t *testing.T) {
	target := int(779210)
	s := Schedule{}
	s.readSchedule("../input_test4.dat")
	result := s.earliestBus(779000)

	if target != result {
		t.Errorf("Got %d, expected %d\n", result, target)
	}
}

func TestEarliestBus5(t *testing.T) {
	target := int(1261476)
	s := Schedule{}
	s.readSchedule("../input_test5.dat")
	result := s.earliestBus(1261000)

	if target != result {
		t.Errorf("Got %d, expected %d\n", result, target)
	}
}

func TestEarliestBus6(t *testing.T) {
	target := int(1202161486)
	s := Schedule{}
	s.readSchedule("../input_test6.dat")
	result := s.earliestBus(1202161000)

	if target != result {
		t.Errorf("Got %d, expected %d\n", result, target)
	}
}
