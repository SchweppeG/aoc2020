package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func (p *PassportScanner) validateField(key string, val string) bool {

	// byr (Birth Year) - four digits;
	// at least 1920 and at most 2002.
	if key == "byr" {
		year, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		if year >= 1920 && year <= 2002 {
			return true
		}
		return false
	}
	// iyr (Issue Year) - four digits;
	// at least 2010 and at most 2020.
	if key == "iyr" {
		year, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		if year >= 2010 && year <= 2020 {
			return true
		}
		return false
	}
	// eyr (Expiration Year) - four digits;
	// at least 2020 and at most 2030.
	if key == "eyr" {
		year, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		if year >= 2020 && year <= 2030 {
			return true
		}
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least
	//   150 and at most 193.
	// If in, the number must be at least
	//   59 and at most 76.
	if key == "hgt" {
		if strings.Contains(val, "cm") {
			val = strings.Replace(val, "cm", "", -1)
			h, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			if h >= 150 && h <= 193 {
				return true
			}
		} else if strings.Contains(val, "in") {
			val = strings.Replace(val, "in", "", -1)
			h, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			if h >= 59 && h <= 76 {
				return true
			}
		}
		return false

	}

	// hcl (Hair Color) - a # followed by exactly six characters
	// 0-9 or a-f.
	// char -> dec
	// 0 -> 48
	// 9 -> 57
	// a -> 97
	// f -> 102
	if key == "hcl" {
		for i, ic := range val {
			if i == 0 {
				if ic != 35 {
					return false
				}
			} else {
				if !((ic >= 48 && ic <= 57) ||
					(ic >= 97 && ic <= 102)) {
					return false
				}

			}
		}
		return true

	}
	// pid (Passport ID) - a nine-digit number,
	// including leading zeroes.
	if key == "pid" {
		if len(val) != 9 {
			return false
		}
		_, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return true
	}

	// ecl (Eye Color) - exactly one of:
	// amb blu brn gry grn hzl oth.
	if key == "ecl" {
		validEcl := []string{"amb", "blu", "brn", "gry",
			"grn", "hzl", "oth"}
		for _, vC := range validEcl {
			if vC == val {
				return true
			}
		}
		return false
	}
	return false
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

	for _, k := range req_keys {
		if p.m[k] == "" {
			return 0
		} else {
			isValid := p.validateField(k, p.m[k])
			if !isValid {
				return 0
			}
		}
	}
	return 1
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
	var line string
	for scanner.Scan() {
		line = scanner.Text()
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
	if line != "" {
		p.nValid += p.validatePassport()
	}

}

func main() {
	p := PassportScanner{}
	p.readPassports("../input1.dat")
	fmt.Printf("Found %d valid passports.", p.nValid)
}
