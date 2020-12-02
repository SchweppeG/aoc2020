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
		low, _ := strconv.Atoi(limits[0])
		high, _ := strconv.Atoi(limits[1])
		c := strings.Replace(s[1], ":", "", 1)
		pass := s[2]

		wg.Add(1)
		go p.checkPassword(low, high, c, pass)
	}
}

func (p *PolicyChecker) countValid() int {
	total := 0
	for v := range p.ValidPasswords {
		total += v
	}
	return total
}

func (p *PolicyChecker) checkPassword(low int,
	high int,
	c string,
	password string) {

	count := strings.Count(password, c)

	if count >= low && count <= high {
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
