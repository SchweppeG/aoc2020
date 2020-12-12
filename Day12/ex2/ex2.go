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

	wN int // waypoint North component
	wE int // waypoint East component
}

func (s *Ship) followInstructions(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	s.N = 0
	s.E = 0
	s.facedir = 90

	s.wN = 1
	s.wE = 10

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
	fmt.Printf("Starting at %d, %d\n", s.N, s.E)
	fmt.Println(action, val)
	switch action {
	case "N":
		s.wN += val
	case "E":
		s.wE += val
	case "S":
		s.wN -= val
	case "W":
		s.wE -= val
	case "R":
		// cw
		n := val / 90
		for i := 0; i < n; i++ {
			s.turn(val)
		}
		fmt.Printf(" to %d, %d\n", s.wN, s.wE)
	case "L":
		// ccw
		n := val / 90
		val *= -1
		for i := 0; i < n; i++ {
			s.turn(val)
		}
	case "F":
		s.N += val * s.wN
		s.E += val * s.wE
	}
	fmt.Printf("Went to %d, %d\n.", s.N, s.E)
}

func (s *Ship) turn(val int) {
	// rotate WAYPOINT
	tmp := s.wN
	s.wN = s.wE
	s.wE = tmp
	if val > 0 {
		// ROTATE CW
		s.wN *= -1
	} else {
		// ROTATE CCW
		s.wE *= -1
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
