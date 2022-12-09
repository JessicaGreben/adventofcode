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

	for _, move := range moves {
		switch move.direction {
		case "R":
			for i := 0; i < move.amount; i++ {
				head.c++
				updateTail(head, tail)
				tailPositions[position{tail.r, tail.c}] = true
			}
		case "L":
			for i := 0; i < move.amount; i++ {
				head.c--
				updateTail(head, tail)
				tailPositions[position{tail.r, tail.c}] = true
			}
		case "U":
			for i := 0; i < move.amount; i++ {
				head.r++
				updateTail(head, tail)
				tailPositions[position{tail.r, tail.c}] = true
			}
		case "D":
			for i := 0; i < move.amount; i++ {
				head.r--
				updateTail(head, tail)
				tailPositions[position{tail.r, tail.c}] = true
			}
		}
	}

	return len(tailPositions)
}

func updateTail(head, tail *position) {
	if tail.touching(head) {
		return
	}
	tail.moveNextTo(head)
}

type position struct {
	r, c int
}

// two positions are touching if either:
//   - the head and tail are at the same location or
//   - the head and tail are next to each other
func (p *position) touching(p2 *position) bool {
	return p.eq(p2) || p.isNextTo(p2)
}

func (p *position) eq(p2 *position) bool {
	return p.r == p2.r && p.c == p2.c
}

func (p *position) isNextTo(p2 *position) bool {
	return (p.r == p2.r && p.c == p2.c+1) || // right: row is the same, col +1
		(p.r == p2.r && p.c == p2.c-1) || // left: row is the same, col -1
		(p.r == p2.r+1 && p.c == p2.c) || // up: col is the same, row +1
		(p.r == p2.r-1 && p.c == p2.c) || // down: col is the same, row -1
		(p.r == p2.r+1 && p.c == p2.c+1) || // diagonal up/right: both differ by +1
		(p.r == p2.r-1 && p.c == p2.c-1) || // diagonal down/left: both differ by -1
		(p.r == p2.r+1 && p.c == p2.c-1) || // diagonal down/right: row+1, col-1
		(p.r == p2.r-1 && p.c == p2.c+1) // diagonal up/left: row-1, col+1
}

func (p *position) sameRowOrCol(p2 *position) bool {
	return p.r == p2.r || p.c == p2.c
}

func (p *position) moveNextTo(p2 *position) {
	if p.sameRowOrCol(p2) { // move rol or col
		switch {
		case p.r == p2.r:
			if p.c > p2.c+1 {
				p.c = p2.c + 1
			}
			if p.c < p2.c-1 {
				p.c = p2.c - 1
			}
		case p.c == p2.c:
			if p.r > p2.r+1 {
				p.r = p2.r + 1
			}
			if p.r < p2.r-1 {
				p.r = p2.r - 1
			}
		default:
			fmt.Println("err row or col should be the same", p, p2)
		}
	} else { // move diagonal
		switch {
		case p.c > p2.c && p.r > p2.r: // right/up
			p.r, p.c = p.r-1, p.c-1
		case p.c < p2.c && p.r < p2.r: // left/down
			p.r, p.c = p.r+1, p.c+1
		case p.c < p2.c && p.r > p2.r: // left/up
			p.r, p.c = p.r-1, p.c+1
		case p.c > p2.c && p.r < p2.r: // right/down
			p.r, p.c = p.r+1, p.c-1
		default:
			fmt.Println("err diagnoal invalid condition", p, p2)
		}
	}
}
