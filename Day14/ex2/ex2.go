package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Docking struct {
	mem   map[int64]int64
	value []rune
	mask  []rune
}

func (d *Docking) readProgramm(filename string) {
	d.mem = make(map[int64]int64)
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
			// get memory location
			memv := strings.Replace(v[0], "mem[", "", -1)
			memv = strings.Replace(memv, "]", "", -1)
			memi, err := strconv.ParseInt(memv, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			// val will be pushed to mem
			val, err := strconv.ParseInt(v[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			_ = val

			// apply mask
			xIndex := []int{}

			memMask := d.ItoMask(memi)
			for i, c := range d.mask {
				switch c {
				case 49: // case 1
					memMask[i] = 49
				case 88: // case X
					// floating action
					memMask[i] = 88
					xIndex = append(xIndex, i)
				case 48: // case 0 do nothing
				}
			}

			l := len(xIndex)

			numIter := int(math.Pow(2, float64(l)))
			tmpMask := make([]rune, len(memMask))

			for i := 0; i <= numIter-1; i++ {
				copy(tmpMask, memMask)
				bstr := strconv.FormatInt(int64(i), 2)

				nPadding := l - len(bstr) - 1
				for i := 0; i <= nPadding; i++ {
					bstr = "0" + bstr
				}

				br := []rune(bstr)

				for i, k := range xIndex {
					tmpMask[k] = br[i]
				}
				d.pushMem(tmpMask, val)
			}
		}
	}
}

func (d *Docking) ItoMask(i int64) []rune {
	tmp := strconv.FormatInt(i, 2)
	l := len(tmp) + 1
	for x := 0; x <= 36-l; x++ {
		tmp = "0" + tmp
	}
	return []rune(tmp)

}

func (d *Docking) pushMem(ind []rune, val int64) {
	index, err := strconv.ParseInt(string(ind), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	d.mem[index] = val
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
	fmt.Println("---------------------")
	fmt.Println(d.getSum())
	fmt.Println("---------------------")
}
