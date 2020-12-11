package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Seating struct {
	// L = 76
	// . = 46
	// # = 35
	layout                 []int
	nRow, nCol             int
	empty, occupied, space int
}

func (s *Seating) readLayer(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	s.empty = 76    // corresponds L
	s.space = 46    // corresponds .
	s.occupied = 35 // corresponds #

	scanner := bufio.NewScanner(file)
	s.nRow = 0
	s.nCol = 0
	for scanner.Scan() {
		s.nCol = len(scanner.Text())
		for _, c := range scanner.Text() {
			s.layout = append(s.layout, int(c))
		}
		s.nRow++
	}

}

func (s *Seating) XYtoI(x int, y int) int {
	i := y*s.nCol + x
	return i
}

func (s *Seating) ItoXY(i int) (int, int) {
	var x, y int
	y = i / s.nCol
	x = i % s.nCol

	return x, y
}

func (s *Seating) printLayout() {
	for y := 0; y <= s.nRow-1; y++ {
		for x := 0; x <= s.nCol-1; x++ {
			out := string(s.layout[s.XYtoI(x, y)])
			fmt.Printf("%v", out)
		}
		fmt.Printf("\n")

	}

}
func (s *Seating) lookDirection(posx, posy, dirx, diry int) int {
	inRange := true

	for inRange {
		posx += dirx
		posy += diry
		if posx >= 0 && posx <= s.nCol-1 &&
			posy >= 0 && posy <= s.nRow-1 {
			val := s.layout[s.XYtoI(posx, posy)]
			if val == s.occupied {
				return 1
			} else if val == s.empty {
				return 0
			}
		} else {
			inRange = false
		}
	}
	return 0
}

func (s *Seating) applyRules() int {
	newLayout := make([]int, len(s.layout))
	totalSum := 0

	for i := 0; i <= len(s.layout)-1; i++ {
		newLayout[i] = s.layout[i]
		if s.layout[i] != s.space {
			occupyCount := 0
			x, y := s.ItoXY(i)

			// [0,1] up
			occupyCount += s.lookDirection(x, y, 0, 1)
			// [0,-1] down
			occupyCount += s.lookDirection(x, y, 0, -1)
			// [1,0] right
			occupyCount += s.lookDirection(x, y, 1, 0)
			// [-1,0] left
			occupyCount += s.lookDirection(x, y, -1, 0)
			// [-1,-1] upledt
			occupyCount += s.lookDirection(x, y, -1, 1)
			// [1,1] upright
			occupyCount += s.lookDirection(x, y, 1, 1)
			// [1,-1] downright
			occupyCount += s.lookDirection(x, y, 1, -1)
			// [-1,1] downleft
			occupyCount += s.lookDirection(x, y, -1, -1)

			if occupyCount == 0 {
				newLayout[i] = s.occupied
			} else if occupyCount >= 5 {
				newLayout[i] = s.empty
			}
		}
		totalSum += newLayout[i] - s.layout[i]
	}
	s.layout = newLayout
	return totalSum
}

func (s *Seating) solve() int {
	run := true
	count := 0
	for run {
		count++
		diff := s.applyRules()
		if diff == 0 {
			run = false
		}
	}
	// conut seats
	seatoccupied := 0
	for i := 0; i <= len(s.layout)-1; i++ {
		if s.layout[i] == s.occupied {
			seatoccupied++
		}
	}
	return seatoccupied
}

func main() {
	s := Seating{}
	s.readLayer("../input1.dat")
	fmt.Println("------------------------")
	fmt.Printf("%d seats are taken.\n", s.solve())
	fmt.Println("------------------------")
}
