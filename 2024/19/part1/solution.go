package main

import (
	"strings"

	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
	"github.com/jessicagreben/adventofcode/pkg/trie"
)

func solution(file string, towelsInput string) (int64, error) {
	input, err := fileinput.New(file)
	if err != nil {
		return -1, err
	}

	towels := strings.Split(towelsInput, ", ")
	prefixTree := trie.New()
	for _, towel := range towels {
		prefixTree.Insert(towel)
	}
	var result int64
	for towelDesign := range input.All() {
		if isDesignPossible(towelDesign, 0, len(towelDesign), prefixTree) {
			result++
		}
	}
	return result, nil
}

func isDesignPossible(design string, startIdx, endIdx int, prefixTree trie.Trie) bool {
	if startIdx >= len(design) {
		return true
	}

	for size := 1; size <= endIdx-startIdx; size++ {
		if startIdx+size > endIdx {
			continue
		}
		curr := design[startIdx : startIdx+size]
		if prefixTree.Search(curr) {
			if ok := isDesignPossible(design, startIdx+size, endIdx, prefixTree); ok {
				return true
			}
		}
	}

	return false
}
