package main

import "testing"

func TestGetNumberTree(t *testing.T) {
	m := Treemap{}
	m.readMapData("../input_test1.dat")
	r := RoutePlaner{rMap: m, cPos: Vec2D{0, 0}}
	nT := r.getNumberTrees(Vec2D{3, 1})

	if nT != 7 {
		t.Errorf("Test failed. Got %d and expected %d",
			nT,
			7)
	}
}
