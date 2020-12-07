package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	color string
	quant int
}

type Rule struct {
	id    string
	rules []Bag
}

type Ruleset struct {
	ruleset []Rule
}

func (r *Ruleset) readRules(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newRule := Rule{}
		// split into rule  and the actual rules
		values := strings.Split(scanner.Text(), " bags contain ")
		newRule.id = strings.Replace(values[0], " ", "_", 1)

		bags := strings.Split(values[1], ", ")
		for _, b := range bags {
			bAttr := strings.Split(b, " ")
			bQuant, err := strconv.Atoi(bAttr[0])
			if err == nil {
				bColor := bAttr[1] + "_" + bAttr[2]
				newRule.rules = append(newRule.rules,
					Bag{bColor, bQuant})
			}
		}
		r.ruleset = append(r.ruleset, newRule)
	}
}

func (r *Ruleset) findBagsColor(color string) map[string]bool {
	color = strings.Replace(color, " ", "_", 1)
	bags := map[string]bool{}
	_ = bags
	for _, rule := range r.ruleset {
		for _, bag := range rule.rules {
			if bag.color == color {
				//				fmt.Println(rule.id)
				bags[rule.id] = true

				tmp := r.findBagsColor(rule.id)

				for k, _ := range tmp {
					bags[k] = true
				}
			}
		}
	}
	return bags
}

func (r *Ruleset) countBagsColor(color string) int {
	return len(r.findBagsColor(color))

}

func main() {
	r := Ruleset{}
	r.readRules("../input.dat")
	fmt.Printf("%d bag colors can contain 'shiny gold' bags.",
		r.countBagsColor("shiny gold"))
}
