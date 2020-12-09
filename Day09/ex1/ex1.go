package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Cypher struct {
}

func (c *Cypher) checkCypher(preamble int, filename string) int64 {
	buffer := make([]int64, preamble)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if i <= preamble-1 {
			buffer[i] = val
			i += 1
		} else {
			i := c.checkValid(buffer, val)
			if i != 0 {
				return i
			}
			buffer = buffer[1:]
			buffer = append(buffer, val)
		}
	}
	return 0
}

func (c *Cypher) checkValid(buffer []int64, target int64) int64 {
	isValid := false
	for i := 0; i <= len(buffer)-1; i += 1 {
		for j := i + 1; j <= len(buffer)-1; j += 1 {
			if buffer[i]+buffer[j] == target {
				isValid = true
			}
		}
	}

	if !isValid {
		fmt.Println("********************************")
		fmt.Printf("%d does not match the sum rule\n", target)
		fmt.Println("********************************")
		return target
	}
	return 0

}

func main() {
	c := Cypher{}
	c.checkCypher(25, "../input1.dat")
}
