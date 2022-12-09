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
	return countTailPositions(moves)
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

func countTailPositions(moves []move) int {
	tailPositions := map[position]bool{}

	head, tail := &position{}, &position{}
	prevHead := &position{head.r, head.c}
	for _, move := range moves {
		switch move.direction {
		case "R":
			for i := 0; i < move.amount; i++ {
				head.c++
				updateTail(prevHead, head, tail)
				prevHead = &position{head.r, head.c}
				tailPositions[position{tail.r, tail.c}] = true
			}
		case "L":
			for i := 0; i < move.amount; i++ {
				head.c--
				updateTail(prevHead, head, tail)
				prevHead = &position{head.r, head.c}
				tailPositions[position{tail.r, tail.c}] = true
			}
		case "U":
			for i := 0; i < move.amount; i++ {
				head.r++
				updateTail(prevHead, head, tail)
				prevHead = &position{head.r, head.c}
				tailPositions[position{tail.r, tail.c}] = true
			}
		case "D":
			for i := 0; i < move.amount; i++ {
				head.r--
				updateTail(prevHead, head, tail)
				prevHead = &position{head.r, head.c}
				tailPositions[position{tail.r, tail.c}] = true
			}
		}
	}

	return len(tailPositions)
}

func updateTail(prevHead, head, tail *position) {
	if tail.adjacent(head) {
		return
	}
	tail.set(prevHead)
}

type position struct {
	r, c int
}

// two positions are adjacent if either:
//   - they are at the same location or
//   - they are next to each other, i.e only 1 row/column away
func (p *position) adjacent(p2 *position) bool {
	diffR, diffC := abs(p.r-p2.r), abs(p.c-p2.c)
	return (diffR == 0 || diffR == 1) && (diffC == 0 || diffC == 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p *position) set(p2 *position) {
	p.r = p2.r
	p.c = p2.c
}
