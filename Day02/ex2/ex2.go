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

type PolicyChecker struct {
	ValidPasswords chan int
}

func (p *PolicyChecker) processInput(filename string) {
	p.ValidPasswords = make(chan int, 10)
	// Read input file and
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		s := strings.Split(scanner.Text(), " ")

		limits := strings.Split(s[0], "-")
		first, _ := strconv.Atoi(limits[0])
		second, _ := strconv.Atoi(limits[1])
		first -= 1
		second -= 1
		c := strings.Replace(s[1], ":", "", 1)
		pass := s[2]

		wg.Add(1)
		go p.checkPassword(first, second, c, pass)
	}
}

func (p *PolicyChecker) countValid() int {
	total := 0
	for v := range p.ValidPasswords {
		total += v
	}
	return total
}

func (p *PolicyChecker) checkPassword(first int,
	second int,
	c string,
	password string) {

	iC := c[0]
	iCFirst := password[first]
	iCSecond := password[second]

	charIsFirst := (iC == iCFirst)
	charIsSecond := (iC == iCSecond)

	if (charIsFirst && !charIsSecond) ||
		(!charIsFirst && charIsSecond) {
		p.ValidPasswords <- 1
	}
	wg.Done()
}

func main() {
	p := PolicyChecker{}
	p.processInput("../input1.dat")

	go func() {
		wg.Wait()
		close(p.ValidPasswords)
	}()

	nValidPasswords := p.countValid()

	fmt.Println("###############################")
	fmt.Printf("%d valid passwords found.\n", nValidPasswords)
	fmt.Println("###############################")
}
