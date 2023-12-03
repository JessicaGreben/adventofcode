package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solution, err := readInput("", doPart1)
	if err != nil {
		fmt.Println("main err=", err)
		os.Exit(1)
	}
	fmt.Println("Part 1 solution:", solution)

	solution, err = readInput("", doPart2)
	if err != nil {
		fmt.Println("main err=", err)
		os.Exit(1)
	}
	fmt.Println("Part 2 solution:", solution)
}

func readInput(file string, fn func(string) (int, error)) (int, error) {
	if file == "" {
		file = "input.txt"
	}
	fd, err := os.Open(file)
	if err != nil {
		return -1, err
	}
	scanner := bufio.NewScanner(fd)
	var possibleGameSum int
	for scanner.Scan() {
		line := scanner.Text()
		num, err := fn(line)
		if err != nil {
			return -1, err
		}
		possibleGameSum += num
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return possibleGameSum, nil
}

const (
	red   = "red"
	green = "green"
	blue  = "blue"
)

var bagContents = map[string]int{
	red:   12,
	green: 13,
	blue:  14,
}

func doPart1(line string) (int, error) {
	s := strings.FieldsFunc(line, func(r rune) bool {
		return r == ':' || r == ';'
	})
	game := strings.Split(s[0], " ")
	gameNumber, err := strconv.Atoi(game[1])
	if err != nil {
		return 0, err
	}
	for i := 1; i < len(s); i++ {
		if !isValidCubes(s[i]) {
			return 0, nil
		}
	}

	return gameNumber, nil
}

func isValidCubes(cubes string) bool {
	sumColors := map[string]int{}

	s := strings.FieldsFunc(cubes, func(r rune) bool {
		return r == ' ' || r == ','
	})

	for i := 0; i <= len(s)-2; i += 2 {
		val, color := s[i], s[i+1]
		numVal, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		sumColors[color] += numVal
	}
	for color, v := range sumColors {
		if v > bagContents[color] {
			return false
		}
	}
	return true
}

func doPart2(line string) (int, error) {
	s := strings.Split(line, ":")
	hands := strings.Split(s[1], ";")
	minCubes := map[string]int{
		red:   0,
		blue:  0,
		green: 0,
	}

	for _, hand := range hands {
		h := strings.FieldsFunc(hand, func(r rune) bool {
			return r == ' ' || r == ','
		})
		for i := 0; i <= len(h)-2; i += 2 {
			val, color := h[i], h[i+1]
			numVal, err := strconv.Atoi(val)
			if err != nil {
				return -1, err
			}
			if numVal > minCubes[color] {
				minCubes[color] = numVal
			}
		}
	}

	for k, v := range minCubes {
		if v == 0 {
			minCubes[k] = 1
		}
	}
	return minCubes[red] * minCubes[blue] * minCubes[green], nil
}
