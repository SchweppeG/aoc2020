package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	values []int
}

type Rule struct {
	low  int
	high int
}

type Field struct {
	name  string
	rules []Rule
}

type TicketCheck struct {
	fields      []Field
	myTicket    []int
	nearTickets []Ticket
}

func readFields(in string) Field {

	f := Field{}
	// read fields with rules
	tmp := strings.Split(in, ":")
	f.name = tmp[0]

	tmpr := strings.Split(tmp[1], "or")
	for i := range tmpr {
		var err error
		r := Rule{}
		vals := strings.Split(tmpr[i], "-")
		r.low, err = strconv.Atoi(vals[0])
		if err != nil {
			log.Fatal(err)
		}
		r.high, err = strconv.Atoi(vals[1])
		if err != nil {
			log.Fatal(err)
		}

		f.rules = append(f.rules, r)
	}

	return f

}

func readTicket(in string) []int {
	out := make([]int, 0)
	tmp := strings.Split(in, ",")
	for i := range tmp {
		v, err := strconv.Atoi(tmp[i])
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, v)
	}
	return out

}

func (t *TicketCheck) readNotes(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	section := 0 // 0 fields, 1 my ticket, 2 nearby tickets
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			section++
		}
		// remove pesky whitespace
		line = strings.Replace(line, " ", "", -1)
		switch section {
		case 0:
			f := readFields(line)
			t.fields = append(t.fields, f)

		case 1:
			// read my ticket
			if !(line == "yourticket:" || line == "") {
				t.myTicket = readTicket(line)
			}

		case 2:
			// read nearby ticket
			if !(line == "nearbytickets:" || line == "") {
				tic := Ticket{}
				tic.values = readTicket(line)
				t.nearTickets = append(t.nearTickets, tic)
			}
		}
	}
}

func (t *TicketCheck) scanTickets() int {
	invalid := make([]int, 0)

	// iterate all ticket
	for _, tics := range t.nearTickets {
		// iterate value in ticket
		for _, v := range tics.values {
			passedRule := false
			// check all fields and the accoring rules
			for _, f := range t.fields {
				for _, r := range f.rules {
					if v >= r.low && v <= r.high {
						passedRule = true
					}
				}
			}
			if passedRule == false {
				invalid = append(invalid, v)
			}
		}
	}

	sum := 0
	for _, j := range invalid {
		sum += j
	}
	return sum
}

func main() {
	t := TicketCheck{}
	t.readNotes("../input1.dat")
	errorRate := t.scanTickets()
	fmt.Println("-----------------")
	fmt.Println("The error rate is:", errorRate)
	fmt.Println("-----------------")
}
