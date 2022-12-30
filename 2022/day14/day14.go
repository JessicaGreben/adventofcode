package day14

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func part1() int {
	pathPoints, minY := convertLines(readLines("input.txt"))
	return countSand(pathPoints, minY)
}

func part2() int {
	pathPoints, minY := convertLines(readLines("input.txt"))
	return countSandPart2(pathPoints, minY)
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

func convertLines(lines []string) (map[point]string, int) {
	points := map[point]string{}
	maxY := math.MinInt32
	for _, line := range lines {
		xypoints := strings.Split(line, " -> ")
		path := []point{}
		for _, p := range xypoints {
			xy := strings.Split(p, ",")
			x, err := strconv.Atoi(xy[0])
			if err != nil {
				fmt.Println("err strconv", err)
			}
			y, err := strconv.Atoi(xy[1])
			if err != nil {
				fmt.Println("err strconv", err)
			}
			path = append(path, point{x, y})
		}

		for i := 0; i < len(path)-1; i++ {
			curr, next := path[i], path[i+1]
			var start, end int

			// vertical line
			if curr.x == next.x {
				if curr.y < next.y {
					start, end = curr.y, next.y
				} else {
					start, end = next.y, curr.y
				}
				for i := start; i <= end; i++ {
					maxY = max(maxY, i)
					points[point{curr.x, i}] = "#"
				}
			}

			// horizontal line
			if curr.y == next.y {
				maxY = max(maxY, curr.y)
				if curr.x < next.x {
					start, end = curr.x, next.x
				} else {
					start, end = next.x, curr.x
				}
				for i := start; i <= end; i++ {
					points[point{i, curr.y}] = "#"
				}
			}
		}
	}
	//print(points)
	//fmt.Println("maxY:", maxY)
	return points, maxY
}

func print(points map[point]string) {
	minX, maxX := math.MaxInt32, math.MinInt32
	minY, maxY := math.MaxInt32, math.MinInt32
	for p := range points {
		minX, maxX = min(minX, p.x), max(maxX, p.x)
		minY, maxY = min(minY, p.y), max(maxY, p.y)
	}
	grid := make([][]string, maxY-minY+1)
	for i := range grid {
		row := make([]string, maxX-minX+1)
		for idx := range row {
			row[idx] = "."
		}
		grid[i] = row
	}
	for p, v := range points {
		grid[p.y-minY][p.x-minX] = v
	}
	for i := range grid {
		fmt.Println(grid[i])
	}
	fmt.Println()
}

func countSand(path map[point]string, maxY int) int {
	start := point{500, 0}
	countSand := 0
	for doSand(start, path, maxY) {
		countSand++
		//print(path)
	}
	return countSand
}

func doSand(s point, path map[point]string, maxY int) (rest bool) {
	curr := s
	for curr.y <= maxY {
		down, left, right := point{curr.x, curr.y + 1}, point{curr.x - 1, curr.y + 1}, point{curr.x + 1, curr.y + 1}
		_, downBlocked := path[down]
		_, leftBlocked := path[left]
		_, rightBlocked := path[right]

		if !downBlocked {
			curr = down
		} else if !leftBlocked {
			curr = left
		} else if !rightBlocked {
			curr = right
		} else {
			path[curr] = "o"
			return true
		}

	}

	return false
}
func countSandPart2(path map[point]string, maxY int) int {
	start := point{500, 0}
	countSand := 0
	for doSandPart2(start, path, maxY) {
		countSand++
	}
	//print(path)
	return countSand + 1
}

func doSandPart2(s point, path map[point]string, maxY int) (sourceNotBlocked bool) {
	curr := s
	for curr.y < maxY+2 {
		nextY := curr.y + 1
		if nextY == maxY+2 {
			path[curr] = "o"
			return true
		}
		down, left, right := point{curr.x, nextY}, point{curr.x - 1, nextY}, point{curr.x + 1, nextY}
		_, downBlocked := path[down]
		_, leftBlocked := path[left]
		_, rightBlocked := path[right]

		if !downBlocked {
			curr = down
		} else if !leftBlocked {
			curr = left
		} else if !rightBlocked {
			curr = right
		} else {
			path[curr] = "o"
			return curr != s
		}
	}

	path[curr] = "o"
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
