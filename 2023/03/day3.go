package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func convertInputToMatrix(file string) ([][]string, error) {
	if file == "" {
		file = "input.txt"
	}
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	m := [][]string{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		m = append(m, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

const (
	period = "."
	digits = "0123456789"
)

type pos struct {
	r, c int
}

var seen = map[pos]bool{}

func doPart1(m [][]string) (int, error) {
	seen = map[pos]bool{}
	var sum int
	for r := range m {
		for c := range m[r] {
			if isSymbol(m[r][c]) {
				neighborsSum, _, err := getNeighboringNums(m, r, c)
				if err != nil {
					return -1, err
				}
				sum += neighborsSum
			}
		}
	}
	return sum, nil
}

func doPart2(m [][]string) (int, error) {
	seen = map[pos]bool{}
	var sum int
	for r := range m {
		for c := range m[r] {
			if m[r][c] == "*" {
				_, neighbors, err := getNeighboringNums(m, r, c)
				if err != nil {
					return -1, err
				}
				if len(neighbors) != 2 {
					continue
				}
				sum += (neighbors[0] * neighbors[1])
			}
		}
	}
	return sum, nil
}

func isSymbol(char string) bool {
	return char != period && !strings.Contains(digits, char)
}

func inMatrix(m [][]string, r, c int) bool {
	return r >= 0 && r < len(m[0]) && c >= 0 && c < len(m)
}

func getNeighboringNums(m [][]string, r, c int) (int, []int, error) {
	sum := 0
	result := []int{}
	neighbors := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, n := range neighbors {
		neighborR, neighborC := r+n[0], c+n[1]
		num, err := getNum(m, neighborR, neighborC)
		if err != nil {
			return -1, nil, err
		}
		if num == -1 {
			continue
		}
		sum += num
		result = append(result, num)
	}
	return sum, result, nil
}

func getNum(m [][]string, r, c int) (int, error) {
	if !inMatrix(m, r, c) {
		return -1, nil
	}
	if !strings.Contains(digits, m[r][c]) {
		return -1, nil
	}

	left := c - 1
	for inMatrix(m, r, left) && strings.Contains(digits, m[r][left]) {
		left--
	}
	right := c + 1
	for inMatrix(m, r, right) && strings.Contains(digits, m[r][right]) {
		right++
	}

	left++
	p := pos{r, left}
	if _, ok := seen[p]; ok {
		return -1, nil
	}
	seen[p] = true

	row := m[r]
	numStr := strings.Join(row[left:right], "")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return -1, err
	}

	return num, nil
}
