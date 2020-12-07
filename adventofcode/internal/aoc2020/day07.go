package aoc2020

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

const MyBag = "shiny gold"

var bagRe = regexp.MustCompile("(.*) bags contain ([no0-9]+ .* bag[s,]?).+")
var bagExists = struct{}{}

type Bag struct {
	Type        string
	Contents    map[*Bag]int
	ContainedIn map[*Bag]struct{}
}

func Day07(reader *bufio.Scanner) (string, string) {
	bags := map[string]*Bag{}
	for reader.Scan() {
		l := reader.Text()
		_ = addBag(l, bags)
	}

	flattened := map[string]struct{}{}
	countContainedIn(bags[MyBag], flattened)
	part1 := len(flattened)
	part2 := countContents(bags[MyBag]) - 1
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func addBag(l string, bags map[string]*Bag) *Bag {
	bagString := bagRe.FindStringSubmatch(l)
	t := bagString[1]
	b := getOrCreateBag(t, bags)
	contents := strings.Split(bagString[2], ",")
	if contents[0] == "no other bags" {
		return b
	}

	for _, ll := range contents {
		ss := strings.SplitN(strings.TrimSpace(ll), " ", 2)
		amount := aoc.ToInt(ss[0])
		bt := ss[1]
		if amount == 1 {
			bt = bt[0 : len(bt)-4]
		} else {
			bt = bt[0 : len(bt)-5]
		}
		bag := getOrCreateBag(bt, bags)
		b.Contents[bag] = amount
		bag.ContainedIn[b] = bagExists
	}

	return b
}

func getOrCreateBag(t string, bags map[string]*Bag) *Bag {
	if _, ok := bags[t]; ok {
		return bags[t]
	}

	b := &Bag{Type: t, Contents: map[*Bag]int{}, ContainedIn: map[*Bag]struct{}{}}
	bags[t] = b
	return b
}

func countContainedIn(b *Bag, flattened map[string]struct{}) {
	for bb := range b.ContainedIn {
		flattened[bb.Type] = bagExists
		countContainedIn(bb, flattened)
	}
}

func countContents(b *Bag) int {
	s := 1
	for bb, count := range b.Contents {
		s += count * countContents(bb)
	}

	return s
}
