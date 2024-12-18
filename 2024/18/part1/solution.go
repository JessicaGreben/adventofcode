package main

import (
	"fmt"
	"math"

	"github.com/jessicagreben/adventofcode/pkg/heap"
	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
	"github.com/jessicagreben/adventofcode/pkg/matrix"
)

func solution(file string, size, byteCount int) (int64, error) {
	input, err := fileinput.New(file)
	if err != nil {
		return -1, err
	}

	m := make([][]string, size)
	for r := range m {
		m[r] = make([]string, size)
		for c := range m[r] {
			m[r][c] = "."
		}
	}

	var count int = 0
	for line := range input.All() {
		if count == byteCount {
			break
		}
		lineInts, err := parseLine(line)
		if err != nil {
			return -1, err
		}
		x, y := lineInts[0], lineInts[1]
		m[y][x] = "#"
		count++
	}
	//defer matrix.Print(m)
	return dijkstras(m)
}

func parseLine(line string) ([]int64, error) {
	return fileinput.ParseLineInt64(line, ",", 2)
}

func dijkstras(m [][]string) (int64, error) {
	distances := make([][]int64, len(m))
	for i := range distances {
		distances[i] = make([]int64, len(m[0]))
		for j := range distances[i] {
			distances[i][j] = math.MaxInt64
		}
	}
	seen := map[nodeDistancePair]bool{}

	n := nodeDistancePair{matrix.Element{0, 0}, 0}
	h := heap.NewMinHeap[nodeDistancePair](func(x, y nodeDistancePair) bool { return x.distance < y.distance })
	h.Push(n)
	distances[0][0] = 0
	for h.Len() > 0 {
		curr, err := h.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		if _, ok := seen[curr]; ok {
			continue
		}
		seen[curr] = true

		if curr.distance > distances[curr.node.Row][curr.node.Col] {
			continue
		}
		distances[curr.node.Row][curr.node.Col] = curr.distance

		for _, dir := range matrix.Directions {
			nextRow, nextCol := curr.node.Row+dir.Row, curr.node.Col+dir.Col
			if !matrix.InBounds(m, nextRow, nextCol) {
				continue
			}
			if m[nextRow][nextCol] == "#" {
				continue
			}
			nextDistance := curr.distance + 1
			n := nodeDistancePair{matrix.Element{nextRow, nextCol}, nextDistance}
			h.Push(n)
		}
	}
	return distances[len(m)-1][len(m[0])-1], nil
}

type nodeDistancePair struct {
	node     matrix.Element
	distance int64
}
