package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Adapters struct {
	jolts                          []int
	count1, count2, count3, rating int
}

func (a *Adapters) readInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	a.jolts = append(a.jolts, 0)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		a.jolts = append(a.jolts, v)

	}
	sort.Ints(a.jolts)
}

func (a *Adapters) getJoltDiff() {
	l := len(a.jolts) - 1
	for i := 1; i <= l; i += 1 {
		diff := a.jolts[i] - a.jolts[i-1]
		switch diff {
		case 0:
		case 1:
			a.count1 += 1
		case 2:
			a.count2 += 1
		case 3:
			a.count3 += 1
		default:
			log.Fatal("Invalid jolt diff: ", diff)
		}
	}
	a.count3 += 1
	a.rating = a.jolts[l] + 3
}

func main() {
	a := Adapters{}
	a.readInput("../input1.dat")
	a.getJoltDiff()
	fmt.Println("###########################################")
	fmt.Printf("%d differences of 1 jolts.\n", a.count1)
	fmt.Printf("%d differences of 3 jolts.\n", a.count3)
	fmt.Printf("Adapter rating fo %d\n", a.rating)
	fmt.Printf("Product count 1 jolts and count 3 jolts: %d\n",
		a.count1*a.count3)
	fmt.Println("###########################################")
}
