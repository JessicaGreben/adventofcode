package main

type matrix struct {
	data [][]string
}

func solution(file string) (int64, error) {
	puzzles := parseInput(file)
	var colSum, rowSum int64 = 0, 0
	for _, puzzle := range puzzles {
		h := getHorizontal(puzzle.data)
		if h > -1 {
			rowSum += h
			continue
		}
		colSum += getVertical(puzzle.data)
	}
	return (100 * rowSum) + colSum, nil
}

func getVertical(m [][]string) int64 {
	for colI, colJ := int64(0), int64(1); colJ < int64(len(m[0])); colI, colJ = colI+1, colJ+1 {
		i, j := colI, colJ
		for colInBounds(m, i) && colInBounds(m, j) && isMirrorCol(m, i, j) {
			i--
			j++
		}
		if colInBounds(m, i) && colInBounds(m, j) {
			continue
		}
		return colI + 1
	}
	return -1
}

func getHorizontal(m [][]string) int64 {
	for rowI, rowJ := int64(0), int64(1); rowJ < int64(len(m)); rowI, rowJ = rowI+1, rowJ+1 {
		i, j := rowI, rowJ
		for rowInBounds(m, i) && rowInBounds(m, j) && isMirrorRow(m, i, j) {
			i--
			j++
		}
		if rowInBounds(m, i) && rowInBounds(m, j) {
			continue
		}
		return rowI + 1
	}
	return -1
}

func rowInBounds(m [][]string, i int64) bool {
	return i >= 0 && i < int64(len(m))
}

func colInBounds(m [][]string, i int64) bool {
	return i >= 0 && i < int64(len(m[0]))
}

func isMirrorRow(m [][]string, i, j int64) bool {
	for c := 0; c < len(m[0]); c++ {
		rowI, rowJ := m[i][c], m[j][c]
		if rowI != rowJ {
			return false
		}
	}
	return true
}

func isMirrorCol(m [][]string, i, j int64) bool {
	for r := 0; r < len(m); r++ {
		colI, colJ := m[r][i], m[r][j]
		if colJ != colI {
			return false
		}
	}
	return true
}
