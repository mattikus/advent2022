package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	win  = 6
	draw = 3
	lose = 0
)

func part1(games [][]string) int {
	rps := map[string]map[string]int{
		"A": { // Rock
			"X": draw,
			"Y": win,
			"Z": lose,
		},
		"B": { // Paper
			"X": lose,
			"Y": draw,
			"Z": win,
		},
		"C": { // Scissors
			"X": win,
			"Y": lose,
			"Z": draw,
		},
	}
	playValue := map[string]int{"X": 1, "Y": 2, "Z": 3}

	var score int
	for _, game := range games {
		a, b := game[0], game[1]
		score += rps[a][b] + playValue[b]
	}

	return score
}

func part2(games [][]string) int {
	move := map[string]map[int]string{
		"A": { // Rock
			win:  "P",
			draw: "R",
			lose: "S",
		},
		"B": { // Paper
			win:  "S",
			draw: "P",
			lose: "R",
		},
		"C": { // Scissors
			win:  "R",
			draw: "S",
			lose: "P",
		},
	}
	winLoseDraw := map[string]int{"X": lose, "Y": draw, "Z": win}
	playValue := map[string]int{"R": 1, "P": 2, "S": 3}

	var score int
	for _, game := range games {
		outcome := winLoseDraw[game[1]]
		play := move[game[0]][outcome]
		score += outcome + playValue[play]
	}
	return score
}

func parseInput(reader *bufio.Reader) [][]string {
	var output [][]string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		items := strings.Split(string(line), " ")
		output = append(output, items)
	}
	return output
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	games := parseInput(reader)
	fmt.Printf("Day 2 Part 1: %v\n", part1(games))
	fmt.Printf("Day 2 Part 2: %v\n", part2(games))
}
