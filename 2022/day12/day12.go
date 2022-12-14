package day6

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func part1() int {
	grid := parseInput("input.txt")
	return shortestPath(grid)
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

type point struct {
	r, c, w int
}

func (p point) eq(p2 point) bool {
	return p.r == p2.r && p.c == p2.c
}

func showPathOnGrid(end point, distance [][]point, grid [][]byte) {
	// Rebuild point starting from the end
	for p := distance[end.r][end.c]; p.w != 0; p = distance[p.r][p.c] {
		grid[p.r][p.c] -= 32 // Upper Case
	}
	// Print grid similar to input file
	for _, r := range grid {
		for _, b := range r {
			fmt.Print(string(b))
		}
		fmt.Println()
	}
}

func bfs(start, end point, grid [][]byte) int {
	seen := map[point]bool{}

	distance := make([][]point, len(grid))
	for i := range distance {
		distance[i] = make([]point, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			distance[i][j] = point{-1, -1, -1}
		}
	}
	distance[start.r][start.c] = point{-1, -1, 0}

	h := MinHeap{}
	q := &h
	heap.Push(q, start)

	for q.Len() > 0 {
		curr := q.Pop().(point)
		seen[point{r: curr.r, c: curr.c}] = true

		directions := []direction{
			{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		}
		for _, dir := range directions {
			nextRow, nextCol := curr.r+dir.rowDiff, curr.c+dir.colDiff
			if !inBounds(nextRow, nextCol, grid) {
				continue
			}
			currVal, nextVal := grid[curr.r][curr.c], grid[nextRow][nextCol]
			if nextVal > currVal+1 {
				continue
			}
			newDistance := distance[curr.r][curr.c].w + 1
			nextDistance := distance[nextRow][nextCol].w
			if nextDistance == -1 || newDistance < nextDistance {
				distance[nextRow][nextCol] = point{curr.r, curr.c, newDistance}
				heap.Push(q, point{nextRow, nextCol, distance[nextRow][nextCol].w})
			}
		}
	}
	showPathOnGrid(end, distance, grid)
	return distance[end.r][end.c].w
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
