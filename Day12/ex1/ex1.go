package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Ship struct {
	N       int // y +1 cartesian
	E       int // x -1 cartesian
	facedir int // directoin in degree
	// 0 == N
	// 90 == W
	// 180 == S
	// 270 == E
}

func (s *Ship) followInstructions(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	s.N = 0
	s.E = 0
	s.facedir = 90

	scanner := bufio.NewScanner(file)
	// possible instructin
	// N S E W
	// L R F
	for scanner.Scan() {
		line := scanner.Text()
		action := string(line[0])
		val, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			log.Fatal(err)
		}
		s.processInput(action, val)
	}

}

func (s *Ship) processInput(action string, val int) {
	fmt.Printf("Starting at %d, %d facing %d\n", s.N, s.E, s.facedir)
	switch action {
	case "N":
		s.N += val
	case "E":
		s.E += val
	case "S":
		s.N -= val
	case "W":
		s.E -= val
	case "R":
		// cw
		s.turn(val)
	case "L":
		// ccw
		val *= -1
		s.turn(val)
	case "F":
		switch s.facedir {
		case 0:
			s.N += val
		case 90:
			s.E += val
		case 180:
			s.N -= val
		case 270:
			s.E -= val
		}
	}
	fmt.Printf("Went to %d, %d facing %d\n", s.N, s.E, s.facedir)
}

func (s *Ship) turn(val int) {
	s.facedir += val
	if val >= 360 {
		val = val % 360
	}
	if s.facedir < 0 {
		s.facedir = 360 + s.facedir
	}
	if s.facedir >= 360 {
		s.facedir = s.facedir % 360
	}

}

func (s *Ship) getDistance() int {
	x := s.N
	y := s.E
	if x < 0 {
		x *= -1
	}

	if y < 0 {
		y *= -1
	}
	return x + y

}

func main() {
	s := Ship{}
	s.followInstructions("../input1.dat")
	fmt.Printf("The Manhatten distance %d.\n", s.getDistance())
	fmt.Println(s)
}
