package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func notContains(intslice []int, val int) bool {
	for _, i := range intslice {
		if i == val {
			return false
		}
	}
	return true
}

type Customs struct {
	sumOfCounts int
}

func (c *Customs) CountAnswers(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// a:97 z 122
	answerBuffer := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// new group
			c.sumOfCounts += len(answerBuffer)
			answerBuffer = []int{}

		} else {
			for _, c := range line {
				i := int(c)
				if 97 <= i && i <= 122 {
					if notContains(answerBuffer, i) {
						answerBuffer = append(answerBuffer, i)
					}
				}
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
