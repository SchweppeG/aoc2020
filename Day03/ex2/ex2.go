package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Vec2D struct{ x, y int }

func (v *Vec2D) add(v2 Vec2D) {
	v.x += v2.x
	v.y += v2.y
}

type Treemap struct {
	tiles []int // 0 space, 1 tree
	dir   Vec2D
	nRows int
	nCols int
}

func (t *Treemap) readMapData(filename string) {
	// read map data from file into 1D slice. All values will be stored
	// into a continuous manner
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	t.nCols = 0
	t.nRows = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		t.nCols = len(line)
		t.nRows += 1
		for i := 0; i <= t.nCols-1; i++ {
			tileType := int(0)
			if line[i] == 35 { // string(35) == #
				tileType = 1
			}
			t.tiles = append(t.tiles, tileType)
		}
	}
}

func (t *Treemap) xyToIndex(x int, y int) int {
	// receive coordinates in cartesian x,y and convert them
	// into the index of the according tile inside the storage
	// slice

	// check boundary conditions make maps invinite on x-component
	// 93, nCols 31
	if x > t.nCols-1 || x < 0 {
		x = x % t.nCols
	}

	index := t.nCols*y + x
	return index
}

func (t *Treemap) returnTileType(v Vec2D) int {
	// return the tile type. 1 == tree, 0 == space
	index := t.xyToIndex(v.x, v.y)
	return t.tiles[index]
}

type RoutePlaner struct {
	rMap Treemap
	cPos Vec2D
}

func (r *RoutePlaner) getNumberTrees(dir Vec2D) int {
	// follow the route trough the map and conut the number or trees
	treeCount := 0
	sPos := r.cPos
	for r.cPos.y < r.rMap.nRows-1 {
		r.cPos.add(dir)
		treeCount += r.rMap.returnTileType(r.cPos)
	}
	r.cPos = sPos // reset startpoint
	return treeCount
}

func main() {
	m := Treemap{}
	m.readMapData("../input1.dat")
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

	fmt.Println("-----------------------------------")
	fmt.Printf("Product of tree encounters is %d.\n", prod)
	fmt.Println("-----------------------------------")

}
