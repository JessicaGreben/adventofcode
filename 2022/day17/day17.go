package day17

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1() int {
	lines := readLines("input.txt")
	moves := strings.Split(lines[0], "")
	return fallingRocks(2022, moves)
}

type point struct {
	x, y int
}

var rockShapes = [][]point{
	horizontal, plus, l, vertical, square,
}

var horizontal = []point{
	{0, 0},
	{1, 0},
	{2, 0},
	{3, 0},
}
var vertical = []point{
	{0, 0},
	{0, 1},
	{0, 2},
	{0, 3},
}

var plus = []point{
	{0, 1},
	{1, 1},
	{2, 1},
	{1, 0},
	{1, 2},
}

var l = []point{
	{0, 0},
	{1, 0},
	{2, 0},
	{2, 1},
	{2, 2},
}

var square = []point{
	{0, 0},
	{0, 1},
	{1, 0},
	{1, 1},
}

func readLines(file string) []string {
	fd, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	lines := []string{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return lines
}

const (
	leftWall  = -1
	rightWall = 7
	floor     = -1
)

func fallingRocks(rounds int, moves []string) int {
	currRound := 1
	shapeIdx := 0
	moveIdx := 0
	start := point{2, 3}
	allCaveRocks := map[point]bool{}
	for currRound <= rounds {
		nextShape := rockShapes[shapeIdx]
		currRock := shapeToRock(start, nextShape)
		//fmt.Println(nextShape, " - ", currRock)

		// fmt.Println("start")
		// print(allCaveRocks, currRock, start)
		var falling bool = true
		for falling {
			nextMove := moves[moveIdx]
			//fmt.Println("before:", nextMove, currRock)
			moveSide(currRock, nextMove, allCaveRocks)
			// fmt.Println(nextMove)
			// print(allCaveRocks, currRock, start)
			falling = moveDown(currRock, allCaveRocks)
			// fmt.Println("down")
			// print(allCaveRocks, currRock, start)

			moveIdx++
			if moveIdx >= len(moves) {
				moveIdx = 0
			}
			//fmt.Println("after:", currRock)
		}
		for _, p := range currRock {
			allCaveRocks[p] = true
			if p.y+4 > start.y {
				start.y = p.y + 4
			}
		}
		shapeIdx++
		if shapeIdx >= len(rockShapes) {
			shapeIdx = 0
		}
		currRound++
	}
	return start.y - 3
}

func print(allRocks map[point]bool, currRock []point, start point) {
	largestY := start.y
	for _, p := range currRock {
		largestY = max(largestY, p.y)
	}
	grid := make([][]string, largestY+1)
	for i := range grid {
		row := make([]string, 0, 7)
		for i := 0; i < 7; i++ {
			row = append(row, ".")
		}
		grid[i] = row
	}

	for k := range allRocks {
		grid[k.y][k.x] = "#"
	}
	for _, k := range currRock {
		grid[k.y][k.x] = "@"
	}
	for i := len(grid) - 1; i >= 0; i-- {
		fmt.Println(grid[i])
	}
	fmt.Println([]string{"-", "-", "-", "-", "-", "-", "-"})
	fmt.Println()
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func shapeToRock(start point, shape []point) []point {
	rockPoints := make([]point, 0, len(shape))
	for _, p := range shape {
		rockPoints = append(rockPoints,
			point{p.x + start.x, p.y + start.y},
		)
	}
	return rockPoints
}

func moveSide(rock []point, move string, caveRocks map[point]bool) {
	xDiff := 1
	if move == "<" {
		xDiff = -1
	}
	if canMove(rock, xDiff, 0, caveRocks) {
		for i := range rock {
			rock[i].x += xDiff
		}
	}
}

func canMove(rock []point, xDiff, yDiff int, caveRocks map[point]bool) bool {
	for _, p := range rock {
		newX, newY := p.x+xDiff, p.y+yDiff
		if _, ok := caveRocks[point{newX, newY}]; ok {
			return false
		}
		if newX == leftWall || newX == rightWall || newY == floor {
			return false
		}
	}
	return true
}

func moveDown(rock []point, caveRocks map[point]bool) (falling bool) {
	if canMove(rock, 0, -1, caveRocks) {
		for i := range rock {
			rock[i].y--
		}
		return true
	}
	return false
}
