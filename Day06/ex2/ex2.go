package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Customs struct {
	sumOfCounts int
}

func stringToIntSlice(s string) []int {
	var sslice []int
	for i, _ := range s {
		sslice = append(sslice, int(s[i]))
	}
	return sslice
}

func intersect(aSlice, bSlice []int) []int {
	var out []int

	for _, a := range aSlice {
		for _, b := range bSlice {
			if a == b {
				out = append(out, a)
			}
		}
	}
	return out
}

func (c *Customs) CountAnswers(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// a:97 z 122
	answerBuffer := []int{}
	newGroup := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// new group
			c.sumOfCounts += len(answerBuffer)
			answerBuffer = []int{}
			newGroup = true
		} else {
			if newGroup == true {
				answerBuffer = stringToIntSlice(line)
				newGroup = false
			} else {
				answerBuffer = intersect(answerBuffer,
					stringToIntSlice(line))
			}

		}
	}
	c.sumOfCounts += len(answerBuffer)
}

func main() {
	c := Customs{}
	c.CountAnswers("../input1.dat")
	fmt.Printf("The sum of the counts is: %d", c.sumOfCounts)
}
