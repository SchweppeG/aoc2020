package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Half struct {
	low  int
	high int
}

func (h *Half) walk(s string) {
	if h.high != h.low {
		loUp := (h.high-h.low)/2 + h.low
		if s == "L" {
			h.high = loUp
		} else if s == "U" {
			h.low = loUp + 1
		}
	}
}

type BoardingPass struct {
	highestSeatId int
	rows          Half
	cols          Half
	seatIDs       []int
}

func (b *BoardingPass) CheckAll(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	rowsOriginal := b.rows
	colsOriginal := b.cols

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		_, _, s := b.Check(scanner.Text())
		b.seatIDs = append(b.seatIDs, s)
		if s > b.highestSeatId {
			b.highestSeatId = s
		}
		// reset values
		b.rows = rowsOriginal
		b.cols = colsOriginal
	}

}

func (b *BoardingPass) Check(s string) (int,
	int, int) {
	for _, c := range s {
		cc := string(c)
		if cc == "B" {
			b.rows.walk("U")
		} else if cc == "F" {
			b.rows.walk("L")
		} else if cc == "R" {
			b.cols.walk("U")
		} else if cc == "L" {
			b.cols.walk(cc)
		}
	}

	return b.rows.low, b.cols.low, b.rows.low*8 + b.cols.low
}

func main() {
	bp := BoardingPass{rows: Half{0, 127}, cols: Half{0, 7}}
	bp.CheckAll("../input1.dat")

	fmt.Printf("Highest Seat id is: %d\n", bp.highestSeatId)
	sort.Ints(bp.seatIDs)
	for i := 1; i <= len(bp.seatIDs)-1; i++ {
		a := bp.seatIDs[i]
		b := bp.seatIDs[i-1]
		diff := a - b
		if diff == 2 {
			fmt.Printf("Seat %d is free has neighbor %d, %d",
				a-1, a, b)
		}

	}
}
