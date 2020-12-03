package main

import (
	"testing"
)

func TestGetNumberTree(t *testing.T) {
	m := Treemap{}
	m.readMapData("../input_test1.dat")
	r := RoutePlaner{rMap: m, cPos: Vec2D{0, 0}}
	prod := 1
	nT := r.getNumberTrees(Vec2D{1, 1})
	prod *= nT

	nT = r.getNumberTrees(Vec2D{3, 1})
	prod *= nT

	nT = r.getNumberTrees(Vec2D{5, 1})
	prod *= nT

	nT = r.getNumberTrees(Vec2D{7, 1})
	prod *= nT

	nT = r.getNumberTrees(Vec2D{1, 2})
	prod *= nT

	if prod != 336 {
		t.Errorf("Test failed. Got %d and expected %d",
			prod,
			336)
	}
}
