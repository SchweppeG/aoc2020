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

func (s *Schedule) earliestBus() uint64 {
	var N uint64 = 1
	for _, k := range s.lines {
		N *= uint64(k)
	}

	time := uint64(s.offset[0])
	incr := uint64(s.lines[0])

nextTime:
	for i := 1; i < len(s.offset); i++ {
		for ctime := time; ctime < N; ctime += incr {
			bid := uint64(s.lines[i])
			off := uint64(s.offset[i])
			if (ctime+off)%bid == 0 {
				time = ctime
				incr *= bid
				continue nextTime
			}
		}
		return 0
	}

	return time
}

func main() {
	s := Schedule{}
	s.readSchedule("../input1.dat")

	fmt.Println("#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#")
	fmt.Printf("Depatrue time is %d\n", s.earliestBus())
	fmt.Println("#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#")
}
