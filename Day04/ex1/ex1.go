package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type PassportScanner struct {
	m      map[string]string
	nValid int
}

func (p *PassportScanner) resetMap() {
	p.m = map[string]string{"byr": "",
		"iyr": "",
		"eyr": "",
		"hgt": "",
		"hcl": "",
		"ecl": "",
		"pid": "",
		"cid": "",
	}
}

func (p *PassportScanner) validatePassport() int {
	// validate passports
	// required fields
	// byr (Birth Year)
	// iyr (Issue Year)
	// eyr (Expiration Year)
	// hgt (Height)
	// hcl (Hair Color)
	// ecl (Eye Color)
	// pid (Passport ID)
	//
	// optional field
	// cid (Country ID)
	req_keys := []string{"byr", "iyr", "eyr", "hgt",
		"hcl", "ecl", "pid"}

	isValid := 1 // 1 is valid 0 is not valid
	for _, k := range req_keys {
		if p.m[k] == "" {
			isValid = 0
		}
	}

	return isValid

}

func (p *PassportScanner) readPassports(filename string) {
	// read input data fill password field map
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	p.nValid = 0
	p.resetMap()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			p.nValid += p.validatePassport()
			p.resetMap()
		} else {
			attr := strings.Split(line, " ")
			for _, field := range attr {
				data := strings.Split(field, ":")
				p.m[data[0]] = data[1]
			}
		}
	}

}

func main() {
	p := PassportScanner{}
	p.readPassports("../input1.dat")
	fmt.Printf("Found %d valid passports.", p.nValid)
}
