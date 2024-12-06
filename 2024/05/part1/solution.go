package main

import (
	"slices"
	"strings"

	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	lines, err := fileinput.GetLines(file)
	if err != nil {
		return -1, err
	}
	graph, paths, err := parseLines(lines)
	if err != nil {
		return -1, err
	}

	var result int64
	for _, path := range paths {
		if isValidPath(path, 0, graph) {
			result += int64(path[len(path)/2])
		}
	}
	return result, nil
}

func isValidPath(path []int64, pathIdx int, graph map[int64][]int64) bool {
	q := []int64{path[0]}

Outer:
	for len(q) > 0 {
		size := len(q)
		for range size {
			curr := q[0]
			q = q[1:]

			nextPathIdx := pathIdx + 1
			if nextPathIdx >= len(path) {
				return true
			}

			nextPath := path[nextPathIdx]
			for _, edge := range graph[curr] {
				if edge == nextPath {
					pathIdx++
					q = []int64{edge}
					goto Outer
				}

				// make sure that each downstream item in the path doesn't need to come
				// before the item getting added to the queue.
				correctOrder := true
				for i := nextPathIdx; i < len(path); i++ {
					downstreamPathValue := path[i]
					nextsPathEdges := graph[downstreamPathValue]
					if slices.Contains[[]int64](nextsPathEdges, edge) {
						correctOrder = false
					}
				}
				if correctOrder {
					q = append(q, edge)
				}
			}
		}
	}
	return false
}

func parseLines(lines []string) (adjList map[int64][]int64, paths [][]int64, err error) {
	adjList = map[int64][]int64{}
	paths = [][]int64{}
	for _, line := range lines {
		switch {
		case strings.Contains(line, "|"):
			parts, err := fileinput.ParseLineInt64(line, "|", 2)
			if err != nil {
				return adjList, paths, err
			}
			from, to := parts[0], parts[1]
			adjList[from] = append(adjList[from], to)
		case strings.Contains(line, ","):
			parts, err := fileinput.ParseLineInt64(line, ",", -1)
			if err != nil {
				return adjList, paths, err
			}
			paths = append(paths, parts)
		}
	}
	return adjList, paths, nil
}
