package day6

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func part1() int {
	grid := parseInput("input.txt")
	return shortestPath(grid)
}

func part2() int {
	grid := parseInput("input.txt")
	return shortestPath2(grid)
}

func parseInput(file string) [][]byte {
	fd, err := os.Open(file)
	if err != nil {
		fmt.Println("err Open", err)
	}
	grid := [][]byte{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		row := []byte{}
		for i := range line {
			row = append(row, line[i])
		}

		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("scanner err", err)
	}
	return grid
}

type direction struct {
	rowDiff, colDiff int
}

func shortestPath(grid [][]byte) int {
	var start, end point
	for rowIdx := range grid {
		for colIdx := range grid[rowIdx] {
			curr := grid[rowIdx][colIdx]
			if curr == 'S' {
				start = point{rowIdx, colIdx, 0}
				grid[rowIdx][colIdx] = 'a'
			}
			if curr == 'E' {
				end = point{rowIdx, colIdx, -1}
				grid[rowIdx][colIdx] = 'z'
			}
		}
	}
	return bfs(start, end, grid)
}

func shortestPath2(grid [][]byte) int {
	var start, end point
	for rowIdx := range grid {
		for colIdx := range grid[rowIdx] {
			curr := grid[rowIdx][colIdx]
			if curr == 'E' {
				end = point{rowIdx, colIdx, -1}
				grid[rowIdx][colIdx] = 'z'
			}
		}
	}
	shortest := math.MaxInt32
	for rowIdx := range grid {
		curr := grid[rowIdx][0]
		if curr == 'S' {
			grid[rowIdx][0] = 'a'
		}
		start = point{rowIdx, 0, 0}
		shortest = min(shortest, bfs(start, end, grid))
	}
	return shortest
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type point struct {
	r, c, w int
}

func (p point) eq(p2 point) bool {
	return p.r == p2.r && p.c == p2.c
}

func bfs(start, end point, grid [][]byte) int {
	seen := map[point]bool{}

	distance := make([][]int, len(grid))
	for i := range distance {
		distance[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			distance[i][j] = -1
		}
	}
	distance[start.r][start.c] = 0

	h := MinHeap{}
	q := &h
	heap.Push(q, start)

	for q.Len() > 0 {
		curr := heap.Pop(q).(point)
		seen[point{r: curr.r, c: curr.c}] = true

		directions := []direction{
			{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		}
		for _, dir := range directions {
			nextRow, nextCol := curr.r+dir.rowDiff, curr.c+dir.colDiff
			if !inBounds(nextRow, nextCol, grid) {
				continue
			}
			nextPoint := point{r: nextRow, c: nextCol}
			if seen[nextPoint] {
				continue
			}
			currVal, nextVal := grid[curr.r][curr.c], grid[nextRow][nextCol]
			if nextVal > currVal+1 {
				continue
			}
			newDistance := distance[curr.r][curr.c] + 1
			nextDistance := distance[nextRow][nextCol]
			if nextDistance == -1 || newDistance < nextDistance {
				distance[nextRow][nextCol] = newDistance
				heap.Push(q, point{nextRow, nextCol, distance[nextRow][nextCol]})
			}
			if nextPoint.eq(end) {
				return distance[nextRow][nextCol]
			}
		}
	}
	return -1
}

func inBounds(r, c int, grid [][]byte) bool {
	return r >= 0 && r < len(grid) &&
		c >= 0 && c < len(grid[0])
}

type MinHeap []point

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].w < h[j].w }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(point))
}
func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}
