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
	jolts                                        []int
	joltdiff                                     []int
	count1, count2, count3, combinations, rating int
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
		a.joltdiff = append(a.joltdiff, diff)
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
	a.joltdiff = append(a.joltdiff, 3)
}

func (a *Adapters) countCombination() {
	a.combinations = 1

	for i := 0; i < len(a.joltdiff); i += 1 {
		fmt.Println(a.joltdiff[i])
		ones := 0
		j := 0
		// count ones
		if a.joltdiff[i] == 1 {
			for j = i; j < len(a.joltdiff); j++ {
				if a.joltdiff[j] == 1 {
					ones += 1
				} else {
					break
				}
			}
			fmt.Println("number", i)
			sum := 0
			if i == 0 {
				// Block at start
				sum = 3 + a.joltdiff[j]
			} else {
				sum = a.joltdiff[i-1] + a.joltdiff[j]
			}

			switch ones {
			case 1:
				if sum == 6 {
					a.combinations *= 1
				} else if sum == 5 {
					a.combinations *= 2
				} else if sum == 4 {
					a.combinations *= 3
				}
			case 2:
				if sum == 6 {
					a.combinations *= 2
				} else if sum == 5 {
					a.combinations *= 3
				} else if sum == 4 {
					a.combinations *= 4
				}

			case 3:
				if sum == 6 {
					a.combinations *= 4
				} else if sum == 5 {
					a.combinations *= 5
				} else if sum == 4 {
					a.combinations *= 8
				}
			case 4:
				if sum == 6 {
					a.combinations *= 7
				} else if sum == 5 {
					a.combinations *= 10
				} else if sum == 4 {
					a.combinations *= 16
				}
			}
			i += ones
		}
	}
}

func main() {
	a := Adapters{}
	a.readInput("../input1.dat")
	a.getJoltDiff()
	a.countCombination()

	fmt.Println("###########################################")
	fmt.Printf("%d differences of 1 jolts.\n", a.count1)
	fmt.Printf("%d differences of 3 jolts.\n", a.count3)
	fmt.Printf("Adapter rating fo %d\n", a.rating)
	fmt.Printf("Product count 1 jolts and count 3 jolts: %d\n",
		a.count1*a.count3)
	fmt.Printf("Number of combinations: %d\n", a.combinations)
	fmt.Println("###########################################")
}
