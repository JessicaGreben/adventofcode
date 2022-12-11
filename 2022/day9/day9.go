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
	return countTailPointsPart1(moves)
}

func part2() int {
	moves := parseInput()
	return countTailPointsPart2(moves)
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

func countTailPointsPart1(moves []move) int {
	tailPoints := map[point]bool{}
	head, tail := &point{}, &point{}
	prevHead := &point{head.r, head.c}

	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			switch move.direction {
			case "R":
				head.c++
			case "L":
				head.c--
			case "U":
				head.r++
			case "D":
				head.r--
			}
			moveTail(prevHead, head, tail)
			tailPoints[point{tail.r, tail.c}] = true
			prevHead = &point{head.r, head.c}
		}
	}

	return len(tailPoints)
}

func countTailPointsPart2(moves []move) int {
	tailPoints := map[point]bool{}

	points := make([]*point, 10)
	for i := range points {
		points[i] = &point{}
	}

	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			for i, p := range points {
				if i == 0 {
					switch move.direction {
					case "R":
						points[i].c++
					case "L":
						points[i].c--
					case "U":
						points[i].r++
					case "D":
						points[i].r--
					}
				}
				if i == len(points)-1 {
					tailPoints[point{p.r, p.c}] = true
					continue
				}
				next := points[i+1]
				moveNext(p, next)
			}
		}
	}

	return len(tailPoints)
}

func moveTail(prevHead, head, tail *point) {
	if tail.adjacent(head) {
		return
	}
	tail.set(prevHead.r, prevHead.c)
}

func moveNext(head, tail *point) {
	if tail.adjacent(head) {
		return
	}
	tail.moveNextTo(head)
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

func (p *point) sameRowOrCol(p2 *point) bool {
	return p.r == p2.r || p.c == p2.c
}

func (p *point) moveNextTo(p2 *point) {
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
