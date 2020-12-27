package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func op(ix int, y string, o string) int {

	iy, err := strconv.Atoi(y)
	if err != nil {
		log.Fatal(err)
	}

	result := 0
	switch o {
	case "+":
		result = ix + iy
	case "*":
		result = ix * iy
	default:
		log.Fatal("Operator '", o, "' unknown")
	}

	return result

}

func resolveParethesis(input []string) []string {
	// return new string with all parenthesis resolved.
	output := make([]string, len(input))
	copy(output, input)

	run := true
	for run {
		ind1 := 0 // index of matching (
		ind2 := 0 // index of last )
		run = false
		for i := len(output) - 1; i >= 0; i-- {
			if output[i] == ")" {
				ind2 = i
				run = true
			}
			if output[i] == "(" {
				ind1 = i
				val := solveOps(output[ind1+1 : ind2])
				sval := strconv.Itoa(val)

				tmp := make([]string, 0)
				tmp = output[:ind1]
				tmp = append(tmp, sval)
				tmp = append(tmp, output[ind2+1:]...)
				output = make([]string, len(tmp))
				copy(output, tmp)
				i = 0
			}
		}
	}
	return output
}

func solveOps(line []string) int {
	result, err := strconv.Atoi(line[0])
	if err != nil {
		log.Fatal(err)
	}

	l := len(line)
	for i := 0; i <= l-2; i += 2 {
		result = op(result, line[i+2], line[i+1])
	}
	return result
}

func evaluateLine(line []string) int {
	// resolve parenthesis
	line = resolveParethesis(line)
	//evaluate
	result := solveOps(line)

	return result

}
func processInput(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := make([]string, 0)
		line := strings.Split(scanner.Text(), " ")

		for _, s := range line {
			if len(s) >= 2 {
				for _, c := range s {
					input = append(input, string(c))
				}
			} else {
				input = append(input, s)
			}
		}

		sum += evaluateLine(input)
	}
	return sum

}

func main() {
	v := processInput("../input1.dat")
	fmt.Println("The sum of all expressions is: ", v)
}
