package main

import (
	"math"
	"strconv"
	"strings"

	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

type vec struct {
	x, y int
}

type machine struct {
	a     *vec
	b     *vec
	prize *vec
}

func solution(file string) (int64, error) {
	lines, err := fileinput.GetLines(file)
	if err != nil {
		return -1, err
	}
	machines, err := parse(lines)
	if err != nil {
		return -1, err
	}
	var result int64
	for _, m := range machines {
		w := win(m)
		if w == math.MaxInt64 {
			continue
		}
		result += w
	}
	return result, nil
}

const (
	buttonAcost = 3
	buttonBcost = 1
)

func win(m *machine) int64 {
	var bestWin int64 = math.MaxInt64
	for aCount := range 100 {
		for bCount := range 100 {
			currX := aCount*m.a.x + bCount*m.b.x
			currY := aCount*m.a.y + bCount*m.b.y
			if currX == m.prize.x && currY == m.prize.y {
				bestWin = min(bestWin, int64(aCount*buttonAcost+bCount*buttonBcost))
			}
		}
	}
	return bestWin
}

func parse(lines []string) ([]*machine, error) {
	machines := []*machine{}
	var m *machine
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "A") {
			m = &machine{}
			a, err := parseLine(line, "+")
			if err != nil {
				return nil, err
			}
			m.a = a
		}
		if strings.Contains(line, "B") {
			b, err := parseLine(line, "+")
			if err != nil {
				return nil, err
			}
			m.b = b
		}
		if strings.Contains(line, "Prize") {
			p, err := parseLine(line, "=")
			if err != nil {
				return nil, err
			}
			m.prize = p
			machines = append(machines, m)
		}
	}
	return machines, nil
}

func parseLine(line string, delimeter string) (*vec, error) {
	split1 := strings.Split(line, ":")
	split2 := strings.Split(split1[1], ",")
	v := &vec{}
	for i, x := range split2 {
		split3 := strings.Split(x, delimeter)
		num, err := strconv.Atoi(split3[1])
		if err != nil {
			return nil, err
		}
		if i == 0 {
			v.x = num
		} else {
			v.y = num
		}
	}
	return v, nil
}
