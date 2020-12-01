package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// struct
// - read input file into slices of int
//  - compare async and print result

type ExpenseResport struct {
	expeneses []int
	name      string
}

func (e *ExpenseResport) ReadInput(filename string) {
	e.expeneses = make([]int, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//e.expeneses = make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		exp, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		e.expeneses = append(e.expeneses, exp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Read file", filename)
}

func (e *ExpenseResport) FindSum() (int, int, int) {
	fmt.Println("Finding sum")
	length := len(e.expeneses) - 1

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			a := e.expeneses[i]
			b := e.expeneses[j]
			if a+b == 2020 {
				fmt.Println("The sum of ", a, " and ", b, " is 2020")
				fmt.Println("The product is:", a*b)
				return a, b, a * b
			}
		}
	}
	return 0, 0, 0
}

func main() {
	e := ExpenseResport{}

	e.ReadInput("..//input_test1.dat")
	_, _, prod := e.FindSum()
	if prod != 514579 {
		log.Fatal("Test failed")
	}

	fmt.Println("#############################")
	fmt.Println("Test passed")
	fmt.Println("#############################")

	e.ReadInput("../input1.dat")
	_, _, _ = e.FindSum()

	fmt.Println("Done")
}
