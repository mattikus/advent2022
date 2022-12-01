package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

// part1 returns the highest calorie count contained by an elf.
func part1(items []int) int {
	var most int
	for _, i := range items {
		if i > most {
			most = i
		}
	}
	return most
}

// part2 return the combined calories from the top 3 elves.
func part2(items []int) int {
	sort.Slice(items, func(i, j int) bool {
		return items[i] > items[j]
	})
	return items[0] + items[1] + items[2]
}

// processInput takes the data from a reader and returns a slice of ints with each int corresponding
// to the total calories of each elf. The number of entries in the slice should match the number of
// elves.
func processInput(reader *bufio.Reader) []int {
	var elf int
	items := []int{0}
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		// double newline separates elves.
		if len(line) == 0 {
			items = append(items, 0)
			elf++
			continue
		}

		cal, err := strconv.Atoi(string(line))
		if err != nil {
			panic(err)
		}
		items[elf] += cal
	}

	return items
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	items := processInput(reader)

	fmt.Printf("Day 1 Part 1: %d\n", part1(items))
	fmt.Printf("Day 1 Part 2: %d\n", part2(items))
}
