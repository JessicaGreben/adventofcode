package main

import (
	"strings"

	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string, towelsInput string) (int64, error) {
	input, err := fileinput.New(file)
	if err != nil {
		return -1, err
	}

	towels := map[string]bool{}
	maxSize := 0
	for _, towel := range strings.Split(towelsInput, ", ") {
		towels[towel] = true
		maxSize = max(maxSize, len(towel))
	}
	var result int64
	for towelDesign := range input.All() {
		result += countDesignWays(towelDesign, 0, len(towelDesign), maxSize, towels, map[seenKey]int64{})
	}
	return result, nil
}

type seenKey struct {
	startIdx, endIdx int
}

func countDesignWays(design string, startIdx, endIdx, maxSize int, towels map[string]bool, seen map[seenKey]int64) int64 {
	if v, ok := seen[seenKey{startIdx, endIdx}]; ok {
		return v
	}
	if startIdx >= len(design) {
		return 1
	}

	end := maxSize
	if endIdx-startIdx < maxSize {
		end = endIdx - startIdx
	}
	var total int64
	for size := 1; size <= end; size++ {
		if startIdx+size > endIdx {
			break
		}
		curr := design[startIdx : startIdx+size]
		if _, ok := towels[curr]; ok {
			ways := countDesignWays(design, startIdx+size, endIdx, maxSize, towels, seen)
			seen[seenKey{startIdx + size, endIdx}] = ways
			total += ways
		}
	}

	return total
}
