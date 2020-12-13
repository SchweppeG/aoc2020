package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Schedule struct {
	depart int
	lines  []int
}

func (s *Schedule) readSchedule(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			line = strings.Replace(line, ",x", "", -1)
			lines := strings.Split(line, ",")
			for _, l := range lines {
				v, err := strconv.Atoi(l)
				if err != nil {
					log.Fatal(err)
				}
				s.lines = append(s.lines, v)
			}
		} else {
			s.depart, err = strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

func (s *Schedule) earliestBus() (int, int) {
	eline := 4294967295
	etime := 4294967295
	for _, i := range s.lines {
		time := i - (s.depart % i)
		fmt.Println(time, s.depart, i)
		if time < etime {
			eline = i
			etime = time
		}
	}
	return eline, etime

}

func main() {
	s := Schedule{}
	s.readSchedule("../input1.dat")
	line, time := s.earliestBus()
	fmt.Println("ID multiplied by minutes: ", line*time)
}
