package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Simulation struct {
	// cubes key will be xyz coordinates as string
	// the value will be the state active or inactive
	cubes    map[string]int
	tmpCubes map[string]int
	ncycle   int

	// simulation boundaries
	xmin, xmax int
	ymin, ymax int
	zmin, zmax int
}

func (s *Simulation) cToID(x, y, z int) string {
	// input coordinates, output string id
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	sz := strconv.Itoa(z)
	out := sx + "," + sy + "," + sz
	return out
}

func (s *Simulation) idToC(id string) (int, int, int) {
	// input cube id and output coordinates
	tmp := strings.Split(id, ",")
	x, err := strconv.Atoi(tmp[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(tmp[1])
	if err != nil {
		log.Fatal(err)
	}
	z, err := strconv.Atoi(tmp[2])
	if err != nil {
		log.Fatal(err)
	}

	return x, y, z
}

func (s *Simulation) updateBoundary(i int, min *int, max *int) {
	// update simulations boundaries
	if *min > i {
		*min = i
	}
	if *max < i {
		*max = i
	}

}

func (s *Simulation) readInitialState(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	s.cubes = make(map[string]int, 0)
	s.ncycle = 0

	// local coordinates
	lx := 0
	ly := 0
	lz := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		state := []rune(scanner.Text())
		// . - 46
		// # - 35
		_, _, _ = lx, ly, lz
		for _, c := range state {
			cid := s.cToID(lx, ly, lz)
			s.cubes[cid] = int(c)
			s.updateBoundary(lx, &s.xmin, &s.xmax)

			lx++
		}
		lx = 0
		s.updateBoundary(ly, &s.ymin, &s.ymax)
		ly++
	}
}

func (s *Simulation) addNewToMap(id string) {
	s.cubes[id] = 46

}

func (s *Simulation) applyToCubes(f func(x, y, z int)) {
	for z := s.zmin; z <= s.zmax; z++ {
		for y := s.ymin; y <= s.ymax; y++ {
			for x := s.xmin; x <= s.xmax; x++ {
				f(x, y, z)
			}
		}
	}
}

func (s *Simulation) getNeighbourState(x, y, z int) (int, int) {
	//fmt.Println("#+#+++#+#+#+#+#+#+#+#+#+#+#+#+#+#+#+#+#+#+#")
	isActive := 0
	isInactive := 0
	cId := s.cToID(x, y, z)
	//fmt.Println("Chekcing cube with id: " + cId)
	for i := z - 1; i <= z+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := x - 1; k <= x+1; k++ {
				id := s.cToID(k, j, i)
				if cId != id {
					if st, ok := s.cubes[id]; ok {
						switch st {
						case 46:
							isInactive++
						case 35:
							isActive++
							//						fmt.Println("Got Active cube id:",
							//							id,
							//							" increasing count to: ",
							//							isActive)

						}
					} else {
						isInactive++
					}
				}
			}

		}

	}
	//	fmt.Println("#+#+++#+#+#+#+#+#+#+#+#+#+#+#+#+#+#+#+#+#+#")
	return isActive, isInactive
}

func copyMap(dest map[string]int, src map[string]int) map[string]int {
	for key, value := range src {
		dest[key] = value
	}
	return dest
}

func (s *Simulation) cycle(steps int) {
	_ = steps

	for s.ncycle <= steps-1 {
		// first add padding, increase x,y,z +1 in each direction
		s.xmin--
		s.ymin--
		s.zmin--
		s.xmax++
		s.ymax++
		s.zmax++

		// add padding to the current state map
		s.applyToCubes(
			func(x, y, z int) {
				id := s.cToID(x, y, z)
				if _, ok := s.cubes[id]; !ok {
					s.addNewToMap(id)
				}
			})

		s.tmpCubes = make(map[string]int, 0)
		s.tmpCubes = copyMap(s.tmpCubes, s.cubes)
		// apply rules
		s.applyToCubes(
			func(x, y, z int) {
				nActive, nInactive := s.getNeighbourState(x, y, z)
				_ = nInactive
				id := s.cToID(x, y, z)
				myState := s.cubes[id]

				switch myState {
				case 46: // inactive (.)
					if nActive == 3 {
						s.tmpCubes[id] = 35
					}
				case 35: // active (#)
					if !(nActive >= 2 && nActive <= 3) {
						s.tmpCubes[id] = 46
					}
				}
			})

		s.cubes = make(map[string]int, 0)
		s.cubes = copyMap(s.cubes, s.tmpCubes)

		s.ncycle++

	}
}

func (s *Simulation) printState() {
	fmt.Printf("After %d cycle:\n", s.ncycle)
	for z := s.zmin; z <= s.zmax; z++ {
		fmt.Printf("z=%d\n", z)
		for y := s.ymin; y <= s.ymax; y++ {
			for x := s.xmin; x <= s.xmax; x++ {
				id := s.cToID(x, y, z)
				state := s.cubes[id]
				fmt.Printf("%v", string(state))
			}
			fmt.Printf("\n")
		}

	}
	fmt.Printf("\n\n")

}

func (s *Simulation) countActiveCubes() int {
	active := 0
	for _, value := range s.cubes {
		if value == 35 {
			active++
		}
	}
	return active
}

func main() {
	s := Simulation{}
	s.readInitialState("../input1.dat")
	s.cycle(6)
	fmt.Println(s.countActiveCubes())
}
