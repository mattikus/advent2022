package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type rng struct {
	start, end int
}

func (a rng) contains(b rng) bool {
	return (a.start <= b.start && a.end >= b.end)
}

func (a rng) overlaps(b rng) bool {
	return (a.start <= b.end && a.end >= b.start)
}

func newRng(s string) rng {
	bounds := strings.Split(s, "-")
	// assuming valid input to avoid too much boilerplate.
	start, _ := strconv.Atoi(bounds[0])
	end, _ := strconv.Atoi(bounds[1])
	return rng{start, end}
}

type pair [2]rng

func parseInput(reader *bufio.Reader) []pair {
	var output []pair
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		ranges := strings.Split(string(line), ",")
		p := pair{newRng(ranges[0]), newRng(ranges[1])}
		output = append(output, p)
	}
	return output
}

func part1(pairs []pair) int {
	var out int
	for _, p := range pairs {
		if p[0].contains(p[1]) || p[1].contains(p[0]) {
			out++
		}
	}
	return out
}

func part2(pairs []pair) int {
	var out int
	for _, p := range pairs {
		if p[0].overlaps(p[1]) || p[1].overlaps(p[0]) {
			out++
		}
	}
	return out
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	pairs := parseInput(reader)
	fmt.Printf("Day 4 Part 1 = %+v\n", part1(pairs))
	fmt.Printf("Day 4 Part 2 = %+v\n", part2(pairs))
}
