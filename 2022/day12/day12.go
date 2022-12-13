package day6

import (
	"bufio"
	"fmt"
	"math"
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

var bestSeen int

func shortestPath(grid [][]byte) int {
	var start, end point
	for rowIdx := range grid {
		for colIdx := range grid[rowIdx] {
			curr := grid[rowIdx][colIdx]
			if curr == 'S' {
				start = point{rowIdx, colIdx}
				grid[rowIdx][colIdx] = 'a'
			}
			if curr == 'E' {
				end = point{rowIdx, colIdx}
				grid[rowIdx][colIdx] = 'z'
			}
		}
	}
	// bestSeen = 0
	// dfs(rowIdx, colIdx, 'S', grid, 0)
	// return bestSeen
	return bfs(start, end, grid)
}

type point struct {
	r, c int
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
	q := []point{start}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		seen[curr] = true
		fmt.Println("pop", curr)

		directions := []direction{
			{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		}
		for _, dir := range directions {
			nextRow, nextCol := curr.r+dir.rowDiff, curr.c+dir.colDiff
			nextPoint := point{nextRow, nextCol}
			if seen[nextPoint] {
				continue
			}
			if !inBounds(nextRow, nextCol, grid) {
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
			}
			if nextPoint.eq(end) {
				return distance[nextRow][nextCol]
			}

			q = append(q, point{nextRow, nextCol})
		}
	}
	return -1
}

func dfs(r, c int, prev byte, grid [][]byte, count int) int {
	if !inBounds(r, c, grid) {
		return math.MaxInt32
	}
	if count > bestSeen {
		return math.MaxInt32
	}
	curr := grid[r][c]
	if curr == 'V' {
		return math.MaxInt32
	}
	if prev+1 < curr {
		return math.MaxInt32
	}
	if curr == 'E' {
		return count
	}
	directions := []direction{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}

	grid[r][c] = 'V'
	for _, dir := range directions {
		nextRow, nextCol := r+dir.rowDiff, c+dir.colDiff
		bestSeen = min(bestSeen, dfs(nextRow, nextCol, curr, grid, count+1))
	}
	grid[r][c] = curr
	return bestSeen
}

func inBounds(r, c int, grid [][]byte) bool {
	return r >= 0 && r < len(grid) &&
		c >= 0 && c < len(grid[0])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
