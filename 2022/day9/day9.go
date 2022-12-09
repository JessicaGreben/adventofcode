package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() int {
	moves := parseInput()
	return countTailPoints(moves)
}

type move struct {
	direction string
	amount    int
}

func parseInput() []move {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("err Open", err)
	}
	out := []move{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) != 2 {
			fmt.Println("err line not lenght 2", line)
		}
		amount, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println("strconv err", err)
		}
		out = append(out, move{line[0], amount})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("scanner err", err)
	}
	return out
}

func countTailPoints(moves []move) int {
	tailPoints := map[point]bool{}

	update := func(p, h, t *point, amount int, moveHeadFn func(*point)) *point {
		for i := 0; i < amount; i++ {
			moveHeadFn(h)
			updateTail(p, h, t)
			tailPoints[point{t.r, t.c}] = true
			p = &point{h.r, h.c}
		}
		return p
	}

	head, tail := &point{}, &point{}
	prevHead := &point{head.r, head.c}
	for _, move := range moves {
		switch move.direction {
		case "R":
			prevHead = update(prevHead, head, tail, move.amount, func(h *point) { h.c++ })
		case "L":
			prevHead = update(prevHead, head, tail, move.amount, func(h *point) { h.c-- })
		case "U":
			prevHead = update(prevHead, head, tail, move.amount, func(h *point) { h.r++ })
		case "D":
			prevHead = update(prevHead, head, tail, move.amount, func(h *point) { h.r-- })
		}
	}

	return len(tailPoints)
}

func updateTail(prevHead, head, tail *point) {
	if tail.adjacent(head) {
		return
	}
	tail.set(prevHead.r, prevHead.c)
}

type point struct {
	r, c int
}

// two points are adjacent if either:
//   - they are at the same location or
//   - they are next to each other, i.e only 1 row/column away
func (p *point) adjacent(p2 *point) bool {
	diffR, diffC := abs(p.r-p2.r), abs(p.c-p2.c)
	return (diffR == 0 || diffR == 1) && (diffC == 0 || diffC == 1)
}

func (p *point) set(r, c int) {
	p.r, p.c = r, c
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
