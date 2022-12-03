package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type rucksack [2]string

func (r rucksack) both() string {
	return r[0] + r[1]
}

func (r rucksack) findCommon() rune {
	for _, item := range strings.Split(r[0], "") {
		if strings.Contains(r[1], item) {
			return []rune(item)[0]
		}
	}
	return ' '
}

func newRucksack(line string) rucksack {
	left := line[0 : len(line)/2]
	right := line[len(line)/2:]
	return rucksack{left, right}
}

type rucksacks []rucksack

// group groups a set of rucksacks into groups of 3 rucksacks
func (rs rucksacks) group() []rucksacks {
	var out []rucksacks
	for i := 0; i < len(rs); i += 3 {
		out = append(out, rucksacks{rs[i], rs[i+1], rs[i+2]})
	}
	return out
}

func (rs rucksacks) findCommon() rune {
	for _, item := range strings.Split(rs[0].both(), "") {
		if strings.Contains(rs[1].both(), item) && strings.Contains(rs[2].both(), item) {
			return []rune(item)[0]
		}
	}
	return ' '
}

// priority converts a given run to a priority value where a-z = 1-26 and A-Z = 27-52. In ASCII, 'A'
// is a lower value than 'a', so we do some conversion to get lowercase 'a' correct, then bump the
// uppercase values back in a positive range.
func priority(r rune) int {
	out := int(r) - 96
	if out < 1 {
		out += 58
	}
	return out
}

func part1(rs rucksacks) int {
	var out int
	for _, r := range rs {
		out += priority(r.findCommon())
	}
	return out
}

func part2(rs rucksacks) int {
	var out int
	for _, g := range rs.group() {
		out += priority(g.findCommon())
	}
	return out
}

func parseInput(reader *bufio.Reader) []rucksack {
	var output []rucksack
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		output = append(output, newRucksack(string(line)))
	}
	return output
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	rucksacks := parseInput(reader)
	fmt.Printf("Day 3 Part 1: %+v\n", part1(rucksacks))
	fmt.Printf("Day 3 Part 2: %+v\n", part2(rucksacks))
}
