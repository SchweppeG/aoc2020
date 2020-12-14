package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type Schedule struct {
	lines  []int
	offset []int
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
			lines := strings.Split(line, ",")
			i := 0
			for _, l := range lines {
				if string(l) != "x" {
					v, err := strconv.Atoi(l)
					if err != nil {
						log.Fatal(err)
					}
					s.lines = append(s.lines, v)
					s.offset = append(s.offset, i)
				}
				i += 1
			}
		}
	}
}

func (s *Schedule) checkTime(ct uint64, d chan uint64) {
	for i := range s.lines {
		if (ct+uint64(s.offset[i]))%uint64(s.lines[i]) !=
			0 {
			return
		}
	}
	//log.Fatal("Found", ct)
	//fmt.Println("--", ct)
	d <- ct
	close(d)
}

func (s *Schedule) earliestBus(starttime uint64) uint64 {

	depart := make(chan uint64, 1)
	//timestamps := make(chan uint64)
	time := starttime
	dt := uint64(s.lines[0])
	// shift time to be mod of dt
	if time%dt != 0 {
		time = time - (time % dt)
	}
	var ret uint64
	running := true

	//go s.makeTimestamps(timestamps, time, dt)
	for running {
		select {
		case x, _ := <-depart:
			ret = x
			running = false
			break
		default:
			go s.checkTime(time, depart)
		}
		time += dt
	}

	return ret
}

func main() {
	s := Schedule{}
	s.readSchedule("../input1.dat")
	fmt.Println(s.lines)
	fmt.Println(s.offset)
	fmt.Println("#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#")
	fmt.Printf("Depatrue time is %d\n", s.earliestBus(100000000000000))
	fmt.Println("#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#")
	//line, time := s.earliestBus()
	//fmt.Println("ID multiplied by minutes: ", line*time)
}
