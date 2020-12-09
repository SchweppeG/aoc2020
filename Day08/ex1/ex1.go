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
	hist        []int
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
	fmt.Printf("Adding %d\n", ii)
	h.accumulator += ii
	h.pointer += 1
}

func (h *Handheld) jmp(i int) {
	fmt.Printf("Jump to %d\n", i)
	h.pointer += i
}

func (h *Handheld) nop() {
	h.pointer += 1
}

func (h *Handheld) run() {
	// infinite loop
	h.pointer = 0
	shouldBreak := false
	for {
		op := h.stack[h.pointer]

		switch op.inst {
		case "nop":
			fmt.Printf("nop pointer: %d\n", h.pointer)
			h.nop()
		case "jmp":
			fmt.Printf("jmp pointer: %d\n", h.pointer)
			h.jmp(op.val)
		case "acc":
			fmt.Printf("acc pointer: %d\n", h.pointer)
			h.acc(op.val)
		}

		for _, i := range h.hist {
			if h.pointer == i {
				shouldBreak = true
			}
		}
		if shouldBreak {
			break
		}
		h.hist = append(h.hist, h.pointer)
	}
}

func main() {
	h := Handheld{}
	h.readOperations("../input1.dat")
	h.run()
	fmt.Println("################################")
	fmt.Printf("Accumulator is: %d\n", h.accumulator)
	fmt.Println("################################")
}
