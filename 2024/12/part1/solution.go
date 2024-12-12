package main

import (
	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
	"github.com/jessicagreben/adventofcode/pkg/matrix"
)

type region struct {
	area      int64
	perimeter int64
}

func solution(file string) (int64, error) {
	m, err := fileinput.ConvertToMatrix(file)
	if err != nil {
		return -1, err
	}

	mInt := convert(m)
	regionInfo := getRegionInfo(mInt)
	var result int64
	for _, info := range regionInfo {
		result += info.area * info.perimeter
	}

	return result, nil
}

func getRegionInfo(m [][]int) map[int]region {
	info := map[int]region{}
	for r := range m {
		for c := range m[r] {
			curr := m[r][c]
			if _, ok := info[curr]; !ok {
				info[curr] = region{}
			}
			currInfo := info[curr]
			currInfo.area++
			perimeterCount := 0
			for _, direction := range matrix.Directions {
				nextRow, nextCol := r+direction.Row, c+direction.Col
				if !matrix.InBounds(m, nextRow, nextCol) {
					perimeterCount++
					continue
				}
				if m[nextRow][nextCol] != m[r][c] {
					perimeterCount++
				}
			}
			currInfo.perimeter += int64(perimeterCount)
			info[curr] = currInfo
		}
	}
	return info
}

func convert(m [][]string) [][]int {
	converted := make([][]int, len(m))
	for r := range m {
		converted[r] = make([]int, len(m[r]))
	}

	count := 0
	for r := range m {
		for c := range m[r] {
			if converted[r][c] == 0 {
				count++
				helper(m, converted, r, c, count, map[matrix.Element]bool{})
			}
		}
	}

	return converted
}

func helper(m [][]string, mInt [][]int, r, c, count int, seen map[matrix.Element]bool) {
	if _, ok := seen[matrix.Element{r, c}]; ok {
		return
	}
	seen[matrix.Element{r, c}] = true
	for _, direction := range matrix.Directions {
		nextRow, nextCol := r+direction.Row, c+direction.Col
		if !matrix.InBounds(m, nextRow, nextCol) {
			continue
		}

		if m[r][c] == m[nextRow][nextCol] {
			mInt[r][c] = count
			helper(m, mInt, nextRow, nextCol, count, seen)
		}
	}
	mInt[r][c] = count
}

func parseLine(line string) ([]int64, error) {
	return fileinput.ParseLineInt64(line, " ", 2)
}
