package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/jessicagreben/adventofcode/pkg/heap"
	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
	"github.com/jessicagreben/adventofcode/pkg/matrix"
)

func solution(file string, savings int) (int64, error) {
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

	if _, err := dijkstras(m, startRow, startCol, endRow, endCol); err != nil {
		return -1, err
	}
	return getCheatsWithSavings(savings, m), nil
}

func getCheatsWithSavings(savings int, m [][]string) int64 {
	var result int64
	re := regexp.MustCompile(`\d`)
	for r := range m {
		for c := range m[r] {
			curr := m[r][c]
			if re.MatchString(curr) {
				for _, dir := range matrix.Directions {
					nextRow, nextCol := r+dir.Row, c+dir.Col
					if !matrix.InBounds(m, nextRow, nextCol) {
						continue
					}
					if m[nextRow][nextCol] == "#" {
						nextnextRow, nextnextCol := nextRow+dir.Row, nextCol+dir.Col
						if !matrix.InBounds(m, nextnextRow, nextnextCol) {
							continue
						}
						nextnext := m[nextnextRow][nextnextCol]
						if re.MatchString(nextnext) {
							currInt, err := strconv.Atoi(curr)
							if err != nil {
								fmt.Println(err, curr)
							}
							nextnextInt, err := strconv.Atoi(nextnext)
							if err != nil {
								fmt.Println(err, curr)
							}
							if nextnextInt > currInt {
								continue
							}
							if abs(currInt-nextnextInt)-2 >= savings {
								result++
							}
						}
					}

				}
			}
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func markPath(m [][]string, startRow, startCol, endRow, endCol int, parents map[matrix.Element]matrix.Element) {
	curr := parents[matrix.Element{endRow, endCol}]
	m[endRow][endCol] = "0"
	step := 1
	for {
		m[curr.Row][curr.Col] = strconv.Itoa(step)
		step++
		if curr.Row == startRow && curr.Col == startCol {
			return
		}
		curr = parents[matrix.Element{curr.Row, curr.Col}]
	}
}

func dijkstras(m [][]string, startRow, startCol, endRow, endCol int) (int64, error) {
	distances := make([][]int64, len(m))
	for i := range distances {
		distances[i] = make([]int64, len(m[0]))
		for j := range distances[i] {
			distances[i][j] = math.MaxInt64
		}
	}
	seen := map[nodeDistancePair]bool{}
	parents := map[matrix.Element]matrix.Element{}

	n := nodeDistancePair{matrix.Element{startRow, startCol}, 0, matrix.Element{}}
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
		parents[curr.node] = curr.from

		for _, dir := range matrix.Directions {
			nextRow, nextCol := curr.node.Row+dir.Row, curr.node.Col+dir.Col
			if !matrix.InBounds(m, nextRow, nextCol) {
				continue
			}
			if m[nextRow][nextCol] == "#" {
				continue
			}
			nextDistance := curr.distance + 1
			n := nodeDistancePair{matrix.Element{nextRow, nextCol}, nextDistance, curr.node}
			h.Push(n)
		}
	}
	markPath(m, startRow, startCol, endRow, endCol, parents)
	return distances[len(m)-1][len(m[0])-1], nil
}

type nodeDistancePair struct {
	node     matrix.Element
	distance int64
	from     matrix.Element
}
