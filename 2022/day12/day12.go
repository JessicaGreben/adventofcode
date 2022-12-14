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

type path struct {
	r, c, d int
}

func bfs(start, end point, grid [][]byte) int {
	seen := map[point]bool{}

	distance := make([][]path, len(grid))
	for i := range distance {
		distance[i] = make([]path, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			distance[i][j] = path{-1, -1, -1}
		}
	}
	distance[start.r][start.c] = path{-1, -1, 0}

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
			nextPoint := point{r: nextRow, c: nextCol}
			if seen[nextPoint] {
				continue
			}
			currVal, nextVal := grid[curr.r][curr.c], grid[nextRow][nextCol]
			if nextVal > currVal+1 {
				continue
			}
			newDistance := distance[curr.r][curr.c].d + 1
			nextDistance := distance[nextRow][nextCol].d
			if nextDistance == -1 || newDistance < nextDistance {
				distance[nextRow][nextCol] = path{curr.r, curr.c, newDistance}
			}
			if nextPoint.eq(end) {
				// Rebuild path starting from the end
				for x := distance[nextRow][nextCol]; x.d != 0; x = distance[x.r][x.c] {
					grid[x.r][x.c] -= 32 // Upper Case
				}
				// Print grid similar to input file
				for _, r := range grid {
					for _, b := range r {
						fmt.Print(string(b))
					}
					fmt.Println()
				}
				return distance[nextRow][nextCol].d
			}
			heap.Push(q, point{nextRow, nextCol, distance[nextRow][nextCol].d})
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
