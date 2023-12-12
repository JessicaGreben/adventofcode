package main

import (
	"fmt"

	"github.com/jessicagreben/adventofcode/pkg/input"
)

const (
	verticalPipe   = "|"
	horizontalPipe = "-"
	Lpipe          = "L"
	Jpipe          = "J"
	sevenPipe      = "7"
	Fpipe          = "F"
	ground         = "."
	startPipe      = "S"
)

var (
	north = []int{-1, 0}
	south = []int{1, 0}
	east  = []int{0, 1}
	west  = []int{0, -1}
)

var neighbors = map[string][][]int{
	verticalPipe:   [][]int{north, south},
	horizontalPipe: [][]int{east, west},
	Lpipe:          [][]int{north, east},
	Jpipe:          [][]int{north, west},
	sevenPipe:      [][]int{south, west},
	Fpipe:          [][]int{south, east},
}

type node struct {
	pipeType string
	position pos
	distance int64
}

type pos struct {
	r, c int
}

type graph struct {
	d           [][]*node
	s           *node
	maxDistance int64
	seen        map[pos]bool
	found       bool
}

func solution(file string, start string) (int64, error) {
	grid, err := input.ConvertToMatrix(file)
	if err != nil {
		return -1, err
	}
	g := generateGraph(grid)
	g.s.pipeType = start
	g.traverse(g.s)
	return g.maxDistance/2 + 1, nil
}

func (g *graph) traverse(curr *node) {
	g.seen[curr.position] = true

	for _, neighbor := range neighbors[curr.pipeType] {
		if g.found {
			return
		}
		r, c := curr.position.r+neighbor[0], curr.position.c+neighbor[1]
		if outOfBounds(g.d, r, c) {
			continue
		}

		nn := g.d[r][c]
		if nn.position.r != r || nn.position.c != c {
			fmt.Println("err pos doesn't match")
		}

		if nn == g.s && curr.distance > 1 {
			g.found = true
			return
		}
		if nn.pipeType == ground {
			continue
		}
		if _, ok := g.seen[nn.position]; ok {
			continue
		}

		g.seen[nn.position] = true
		nn.distance = max(nn.distance, curr.distance+1)
		g.maxDistance = max(g.maxDistance, nn.distance)

		g.traverse(nn)
		delete(g.seen, nn.position)
	}
	g.seen[curr.position] = true
}

func outOfBounds(g [][]*node, r, c int) bool {
	return r < 0 || r >= len(g) || c < 0 || c >= len(g[0])
}

func generateGraph(grid [][]string) *graph {
	gg := make([][]*node, len(grid))
	var startNode *node
	for rIdx := range grid {
		gg[rIdx] = make([]*node, len(grid[rIdx]))
		for cIdx := range grid[rIdx] {
			curr := grid[rIdx][cIdx]
			p := &node{curr, pos{rIdx, cIdx}, -1}
			// p := &node{pipeTypeEnum[curr], pos{rIdx, cIdx}, -1}
			gg[rIdx][cIdx] = p
			if curr == startPipe {
				p.distance = 0
				startNode = p
			}
		}
	}

	return &graph{
		d:           gg,
		s:           startNode,
		maxDistance: 0,
		seen:        map[pos]bool{},
	}
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
