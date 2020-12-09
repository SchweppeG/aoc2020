package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	inst string
	val  int
}

type Handheld struct {
	accumulator int64
	stack       []Operation
	pointer     int
}

func (h *Handheld) readOperations(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		op := strings.Split(scanner.Text(), " ")
		opI, err := strconv.Atoi(op[1])
		if err != nil {
			log.Fatal(err)
		}
		h.stack = append(h.stack, Operation{op[0], opI})
	}
}

func (h *Handheld) acc(i int) {
	ii := int64(i)
	h.accumulator += ii
	h.pointer += 1
}

func (h *Handheld) jmp(i int) {
	h.pointer += i
}

func (h *Handheld) nop() {
	h.pointer += 1
}

func (h *Handheld) run(stack []Operation) int {

	var hist []int
	h.pointer = 0
	h.accumulator = 0
	addr := len(stack) - 1
	shouldBreak := false
	for {
		op := stack[h.pointer]

		if h.pointer == addr {
			shouldBreak = true
		}

		switch op.inst {
		case "nop":
			h.nop()
		case "jmp":
			h.jmp(op.val)
		case "acc":
			h.acc(op.val)
		}

		for _, i := range hist {
			if h.pointer == i {
				return -1
			}
		}

		if shouldBreak {
			return 0
		}
		hist = append(hist, h.pointer)
	}
}

func (h *Handheld) solve() {
	// loop something solve until run returns 0
	i := 0
	ret := -1
	for ret == -1 {
		fmt.Println("Tial: ", i)
		if h.stack[i].inst == "nop" {
			h.stack[i].inst = "jmp"
			ret = h.run(h.stack)
			h.stack[i].inst = "nop"
			i += 1
		} else if h.stack[i].inst == "jmp" {
			h.stack[i].inst = "nop"
			ret = h.run(h.stack)
			h.stack[i].inst = "jmp"
			i += 1
		} else {
			i += 1
		}
	}
}

func main() {
	h := Handheld{}
	h.readOperations("../input1.dat")
	h.solve()
	fmt.Println("################################")
	fmt.Printf("Accumulator is: %d\n", h.accumulator)
	fmt.Println("################################")
}
