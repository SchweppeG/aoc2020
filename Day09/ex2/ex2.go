package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Cypher struct {
	data []int64
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
			c.data = append(c.data, val)
			i += 1
		} else {
			i := c.checkValid(buffer, val)
			if i != 0 {
				return i
			}
			buffer = buffer[1:]
			buffer = append(buffer, val)
			c.data = append(c.data, val)
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
		return target
	}
	return 0

}

func (c *Cypher) breakCypher(target int64) int64 {
	windowSize := 2
	notValid := true
	i := 0
	for notValid {
		sum := int64(0)
		for j := i; j <= i+windowSize-1; j += 1 {
			sum += c.data[j]
		}
		if sum == target {
			min := int64(9223372036854775807)
			max := int64(-9223372036854775808)

			for n := i; n <= i+windowSize-1; n += 1 {
				if c.data[n] > max {
					max = c.data[n]
				}
				if c.data[n] < min {
					min = c.data[n]
				}
			}
			notValid = false
			return min + max
		}
		i += 1
		if i >= len(c.data)-windowSize+1 {
			i = 0
			windowSize += 1
		}
		if windowSize >= len(c.data)-1 {
			log.Fatal("Window size too large. code not breakable")
		}

	}

	return int64(0)
}

func main() {
	c := Cypher{}
	invalid := c.checkCypher(25, "../input1.dat")
	code := c.breakCypher(invalid)

	fmt.Println("********************************")
	fmt.Printf("%d does not match the sum rule\n", invalid)
	fmt.Println("********************************")
	fmt.Println("********************************")
	fmt.Printf("%d is the encryption weakness\n", code)
	fmt.Println("********************************")
}
