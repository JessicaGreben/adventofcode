package main

import (
	"fmt"
	"math"

	"github.com/jessicagreben/adventofcode/pkg/heap"
	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
	"github.com/jessicagreben/adventofcode/pkg/matrix"
)

func solution(file string) (int64, error) {
	m, err := fileinput.ConvertToMatrix(file)
	if err != nil {
		return -1, err
	}
	var startRow, startCol int
	var endRow, endCol int
	for r := range m {
		for c := range m[r] {
			if m[r][c] == "S" {
				startRow, startCol = r, c
			}
			if m[r][c] == "E" {
				endRow, endCol = r, c
			}
		}
	}

	//return dijkstras(m, startRow, startCol, endRow, endCol)
	return shortestPath(m, startRow, startCol, endRow, endCol)
}

type nodeDistancePair struct {
	node      matrix.Element
	distance  int64
	direction string
}

type edge struct {
	row, col int
	weight   int64
}

const (
	north = "north"
	east  = "east"
	west  = "west"
	south = "south"
)

var directions = map[string]map[string]edge{
	north: map[string]edge{
		north: {-1, 0, 1},
		east:  {0, 1, 1001},
		west:  {0, -1, 1001},
	},
	east: map[string]edge{
		east:  {0, 1, 1},
		north: {-1, 0, 1001},
		south: {1, 0, 1001},
	},
	south: map[string]edge{
		south: {1, 0, 1},
		east:  {0, 1, 1001},
		west:  {0, -1, 1001},
	},
	west: map[string]edge{
		west:  {0, -1, 1},
		south: {1, 0, 1001},
		north: {-1, 0, 1001},
	},
}

func shortestPath(m [][]string, row, col, endRow, endCol int) (int64, error) {
	distances := make([][]int64, len(m))
	for i := range distances {
		distances[i] = make([]int64, len(m[0]))
		for j := range distances[i] {
			distances[i][j] = math.MaxInt64
		}
	}
	distances[row][col] = 0
	q := []nodeDistancePair{nodeDistancePair{matrix.Element{row, col}, 0, east}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.distance > distances[curr.node.Row][curr.node.Col] {
			continue
		}
		distances[curr.node.Row][curr.node.Col] = curr.distance

		for dir, edge := range directions[curr.direction] {
			nextRow, nextCol := curr.node.Row+edge.row, curr.node.Col+edge.col
			if m[nextRow][nextCol] == "#" {
				continue
			}
			nextDistance := curr.distance + edge.weight
			q = append(q, nodeDistancePair{matrix.Element{nextRow, nextCol}, nextDistance, dir})
		}
	}
	return distances[endRow][endCol], nil
}

func dijkstras(m [][]string, row, col, endRow, endCol int) (int64, error) {
	distances := make([][]int64, len(m))
	for i := range distances {
		distances[i] = make([]int64, len(m[0]))
		for j := range distances[i] {
			distances[i][j] = math.MaxInt64
		}
	}
	n := nodeDistancePair{matrix.Element{row, col}, 0, east}
	h := heap.NewMinHeap[nodeDistancePair](func(x, y nodeDistancePair) bool { return x.distance < y.distance })
	h.Push(n)
	distances[row][col] = 0
	for h.Len() > 0 {
		curr, err := h.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}

		if curr.distance > distances[curr.node.Row][curr.node.Col] {
			continue
		}
		distances[curr.node.Row][curr.node.Col] = curr.distance

		for dir, edge := range directions[curr.direction] {
			nextRow, nextCol := curr.node.Row+edge.row, curr.node.Col+edge.col
			if m[nextRow][nextCol] == "#" {
				continue
			}
			nextDistance := curr.distance + edge.weight
			n := nodeDistancePair{matrix.Element{nextRow, nextCol}, nextDistance, dir}
			h.Push(n)
		}
	}
	return distances[endRow][endCol], nil
}
