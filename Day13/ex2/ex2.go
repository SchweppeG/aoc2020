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
	d <- ct
	close(d)
}

func GCD(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a

}

func LCM(a int, b int, values ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(values); i++ {
		result = LCM(result, values[i])
	}
	return result
}

func (s *Schedule) earliestBus(starttime uint64) int {

	var tmp []int
	for i := range s.lines {
		tmp = append(tmp, s.lines[i]-s.offset[i])
	}
	fmt.Println(tmp)
	fmt.Println(s.lines[0])
	fmt.Println(s.offset[1])
	fmt.Println(s.offset[2:])
	ret := LCM(s.lines[0], s.offset[1], s.offset[2:]...)

	return ret
}

func main() {
	s := Schedule{}
	s.readSchedule("../input_test1.dat")
	fmt.Println(s.lines)
	fmt.Println(s.offset)
	fmt.Println("#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#")
	fmt.Printf("Depatrue time is %d\n", s.earliestBus(100000))
	//fmt.Printf("Depatrue time is %d\n", s.earliestBus(100000000000000))
	fmt.Println("#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#-#")
	//line, time := s.earliestBus()
	//fmt.Println("ID multiplied by minutes: ", line*time)
}
