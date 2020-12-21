package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Recitation struct {
	// map key = number; value = turn
	mem  map[int]int
	turn int

	// buffer before written to mem
	last_turn int
	last_num  int
}

func (r *Recitation) readInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	if err != nil {
		log.Fatal(err)
	}

	r.mem = make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		startnum := strings.Split(line, ",")
		l := len(startnum) - 1

		for i := 0; i <= l-1; i++ {
			ind, err := strconv.Atoi(startnum[i])
			if err != nil {
				log.Fatal(err)
			}
			r.mem[ind] = i + 1
		}

		r.last_turn = l + 1
		r.last_num, err = strconv.Atoi(startnum[l])
		if err != nil {
			log.Fatal(err)
		}

		r.turn = l + 2
	}

}

func (r *Recitation) countToTurn(turns int) {

	for r.turn = r.turn; r.turn <= turns; r.turn++ {
		if t, ok := r.mem[r.last_num]; ok {
			r.mem[r.last_num] = r.last_turn
			diffturn := r.turn - (t + 1)

			//fmt.Printf("Turn %d: %d was spoken %d and %d,"+
			//	" the diff is %d\n",
			//	r.turn, r.last_num, r.last_turn, t, diffturn)

			r.last_num = diffturn

		} else {
			r.mem[r.last_num] = r.last_turn

			//fmt.Printf("Turn %d: %d is new, so the %dth"+
			//	" number spoken is %d\n",
			//	r.turn, r.last_num, r.turn, 0)
			r.last_num = 0
		}
		r.last_turn = r.turn
	}

}

func main() {
	r := Recitation{}
	r.readInput("../input1.dat")
	r.countToTurn(30000000)
	fmt.Println("-----------------")
	fmt.Println(r.last_num)
	fmt.Println("-----------------")
}
