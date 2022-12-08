package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() int {
	grid := inputGrid()
	return findVisible(grid)
}

type tree struct {
	height  int
	visible bool
}

func inputGrid() [][]tree {
	fd, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("err Open", err)
	}
	out := [][]tree{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		heights := strings.Split(scanner.Text(), "")
		row := make([]tree, 0, len(heights))
		for _, h := range heights {
			height, err := strconv.Atoi(h)
			if err != nil {
				fmt.Println("err strconv", h)
			}
			t := tree{
				height: height,
			}
			row = append(row, t)
		}
		out = append(out, row)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("scanner err", err)
	}
	return out
}

func findVisible(grid [][]tree) int {
	var totalVis int
	for rowIdx := range grid {
		for colIdx := range grid[rowIdx] {
			if isVisible(rowIdx, colIdx, grid) {
				totalVis++
				grid[rowIdx][colIdx].visible = true
			}
		}
	}
	return totalVis
}

func isVisible(r, c int, grid [][]tree) bool {
	if isEdge(r, c, grid) {
		return true
	}
	return rowShorter(r, c, grid) || colShorter(r, c, grid)
}

func isEdge(r, c int, grid [][]tree) bool {
	return r == 0 || r == len(grid)-1 || c == 0 || c == len(grid[r])-1
}

func rowShorter(r, c int, grid [][]tree) bool {
	curr := grid[r][c].height

	lShorter := true
	for i := 0; i < c; i++ {
		left := grid[r][i]
		if left.height >= curr {
			lShorter = false
			break
		}
	}
	rShorter := true
	for i := c + 1; i < len(grid[r]); i++ {
		right := grid[r][i]
		if right.height >= curr {
			rShorter = false
			break
		}
	}

	return lShorter || rShorter
}

func colShorter(r, c int, grid [][]tree) bool {
	curr := grid[r][c].height

	uShorter := true
	for i := 0; i < r; i++ {
		up := grid[i][c]
		if up.height >= curr {
			uShorter = false
			break
		}
	}
	dShorter := true
	for i := r + 1; i < len(grid); i++ {
		down := grid[i][c]
		if down.height >= curr {
			dShorter = false
			break
		}
	}

	return uShorter || dShorter
}
