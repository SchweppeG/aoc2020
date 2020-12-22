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
	pos   []int
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

func remove(sl []Ticket, ind int) []Ticket {
	return append(sl[:ind], sl[ind+1:]...)
}

func (t *TicketCheck) scanTickets() int {
	invalid := make([]int, 0)

	failedTickets := make([]int, 0)
	// iterate all ticket
	for i, tics := range t.nearTickets {
		ticketFailed := false
		// iterate value in ticket
		for _, v := range tics.values {
			passedRule := false
			// check all fields and the accoring rules
			for _, f := range t.fields { // field has position arg
				for _, r := range f.rules {
					if v >= r.low && v <= r.high {
						passedRule = true
					}
				}
			}
			if passedRule == false {
				invalid = append(invalid, v)
				ticketFailed = true
				break
			}
		}
		if ticketFailed {
			failedTickets = append(failedTickets, i)
		}
	}

	// remove invalid tickets
	for j := len(failedTickets) - 1; j >= 0; j-- {
		ind := failedTickets[j]
		t.nearTickets = remove(t.nearTickets, ind)
	}

	sum := 0
	for _, j := range invalid {
		sum += j
	}
	return sum
}

func removeInt(sl []int, ind int) []int {
	return append(sl[:ind], sl[ind+1:]...)
}

func intInSlice(i int, sl []int) bool {
	for _, b := range sl {
		if b == i {
			return true
		}
	}
	return false
}

func (t *TicketCheck) identFields() {

	// iterate all ticket
	for _, tics := range t.nearTickets {
		// iterate value in ticket
		for n, v := range tics.values {
			// check all fields and the accoring rules
			for j, f := range t.fields { // field has position arg
				for _, r := range f.rules {
					if v >= r.low && v <= r.high {
						t.fields[j].pos = append(t.fields[j].pos, n)

					}
				}
			}
		}
	}

	npos := len(t.nearTickets)

	for i, f := range t.fields {
		pp := make([]int, 0)
		for _, p := range f.pos {
			count := 0
			for _, p2 := range f.pos {
				if p == p2 {
					count++
				}
			}
			if count == npos {
				if !intInSlice(p, pp) {
					pp = append(pp, p)
				}
			}
		}
		t.fields[i].pos = pp
	}

	runReduce := true
	for runReduce {
		// reduce position slices
		for _, f := range t.fields {
			if len(f.pos) == 1 {
				p := f.pos[0]
				for j, f2 := range t.fields {
					if len(f2.pos) > 1 {
						for i := range f2.pos {
							if f2.pos[i] == p {
								t.fields[j].pos = removeInt(
									t.fields[j].pos, i)
							}
						}
					}
				}
			}
		}

		maxlen := 1
		for _, f := range t.fields {
			l := len(f.pos)
			if l > maxlen {
				maxlen = 2
			}
		}
		if maxlen == 1 {
			runReduce = false

		}
	}
}

func main() {
	t := TicketCheck{}
	t.readNotes("../input1.dat")
	errorRate := t.scanTickets()
	t.identFields()
	fmt.Println("-----------------")
	fmt.Println("The error rate is:", errorRate)
	fmt.Println("-----------------")

	prod := 1
	for _, f := range t.fields {
		if strings.Contains(f.name, "departure") {
			prod *= t.myTicket[f.pos[0]]
		}
		fmt.Printf("%v has position %d\n", f.name, f.pos[0])
	}

	fmt.Println(".......................")
	fmt.Printf("Product of depature fields of my ticket is: %d\n",
		prod)
	fmt.Println(".......................")

}
