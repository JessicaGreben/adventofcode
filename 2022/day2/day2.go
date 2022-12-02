package day1

import (
	"bufio"
	"fmt"
	"os"
)

func part1() int {
	in, err := input()
	if err != nil {
		fmt.Println(err)
	}
	return totalScore(in)
}

func input() ([]string, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	out := []string{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func totalScore(games []string) int {
	// A,X for Rock, B,Y for Paper, and C,Z for Scissors
	// 1 for Rock, 2 for Paper, and 3 for Scissors
	possibleGames := map[string]int{
		// hand points + round points
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	totalScore := 0
	for _, round := range games {
		score, ok := possibleGames[round]
		if !ok {
			fmt.Println("err round not found", round)
		}
		totalScore += score
	}
	return totalScore
}

func part2() int {
	in, err := input()
	if err != nil {
		fmt.Println(err)
	}
	return totalScorePart2(in)
}

func totalScorePart2(games []string) int {
	// A for Rock, B for Paper, and C for Scissors
	// 1 for Rock, 2 for Paper, and 3 for Scissors
	// X lose, Y draw, Z win
	possibleGames := map[string]int{
		// hand points + round points
		"A X": 3 + 0,
		"A Y": 1 + 3,
		"A Z": 2 + 6,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 2 + 0,
		"C Y": 3 + 3,
		"C Z": 1 + 6,
	}

	totalScore := 0
	for _, round := range games {
		score, ok := possibleGames[round]
		if !ok {
			fmt.Println("err round not found", round)
		}
		totalScore += score
	}
	return totalScore
}
