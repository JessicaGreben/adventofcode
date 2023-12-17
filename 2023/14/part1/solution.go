package main

import (
	"github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	m, err := input.ConvertToMatrix(file)
	if err != nil {
		return -1, err
	}

	tiltNorth(m)
	return count(m), nil
}

func tiltNorth(m [][]string) {
	for cIdx := range m[0] {
		startIdx, endIdx := 0, 1
		for startIdx < len(m) && endIdx < len(m) {
			startVal, endVal := m[startIdx][cIdx], m[endIdx][cIdx]
			if startVal == "#" && endVal == "#" {
				move(m, startIdx+1, endIdx-1, cIdx)
				startIdx, endIdx = endIdx, endIdx+1
				continue
			}
			if startVal == "#" && endIdx == len(m)-1 {
				move(m, startIdx+1, endIdx, cIdx)
				break
			}
			if startIdx == 0 && endVal == "#" {
				move(m, startIdx, endIdx-1, cIdx)
				startIdx, endIdx = endIdx, endIdx+1
				continue
			}
			if startIdx == 0 && endIdx == len(m)-1 {
				move(m, startIdx, endIdx, cIdx)
				break
			}

			endIdx++
		}
	}
}

func move(m [][]string, start, end, cIdx int) {
	i, j := start, end
	for i <= j {
		iVal, jVal := m[i][cIdx], m[j][cIdx]
		if iVal == "." && jVal == "O" {
			m[i][cIdx], m[j][cIdx] = m[j][cIdx], m[i][cIdx]
			i++
			j--
			continue

		}
		if iVal != "." {
			i++
		}
		if jVal != "O" {
			j--
		}
	}
}

func count(m [][]string) int64 {
	var result int64
	x := int64(len(m))
	for rIdx := range m {
		var countO int64
		for _, v := range m[rIdx] {
			if v == "O" {
				countO++
			}
		}
		result += x * countO
		x--
	}
	return result
}
