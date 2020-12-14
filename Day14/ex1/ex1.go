package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Docking struct {
	mem   map[int]int64
	value []rune
	mask  []rune
}

func (d *Docking) readProgramm(filename string) {
	d.mem = make(map[int]int64)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		v := strings.Split(line, " = ")
		if strings.Contains(v[0], "mask") {
			d.mask = []rune(v[1])
		} else {
			memv := strings.Replace(v[0], "mem[", "", -1)
			memv = strings.Replace(memv, "]", "", -1)

			memi, err := strconv.Atoi(memv)
			_ = memi // store index

			if err != nil {
				log.Fatal(err)
			}

			itmp, err := strconv.ParseInt(v[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			tmp := strconv.FormatInt(itmp, 2)
			l := len(tmp) + 1
			for x := 0; x <= 36-l; x++ {
				tmp = "0" + tmp
			}

			val := []rune(tmp)
			for i, c := range d.mask {
				switch c {
				case 49: // case 1
					val[i] = 49
				case 48: // case 0
					val[i] = 48
				}
			}
			result, err := strconv.ParseInt(string(val), 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			d.mem[memi] = result

		}
	}

}

func (d *Docking) getSum() int64 {
	sum := int64(0)
	for _, v := range d.mem {
		sum += v
	}
	return sum
}

func main() {
	d := Docking{}
	d.readProgramm("../input1.dat")
	fmt.Println(d.getSum())
}
