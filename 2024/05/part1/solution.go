package main

import (
	"strconv"
	"strings"

	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	lines, err := fileinput.GetLines(file)
	if err != nil {
		return -1, err
	}
	g, paths, err := parseLines(lines)
	if err != nil {
		return -1, err
	}

	var result int64
	for _, path := range paths {
		var valid bool = true
		for i := range path {
			for j := i + 1; j < len(path); j++ {
				a, b := path[i], path[j]
				if _, ok := g[a+b]; !ok {
					valid = false
					continue
				}
			}
		}
		if valid {
			pathMidStr := path[len(path)/2]
			pathMidInt, err := strconv.Atoi(pathMidStr)
			if err != nil {
				return -1, nil
			}
			result += int64(pathMidInt)
		}
	}
	return result, nil
}

func parseLines(lines []string) (adjList map[string]bool, paths [][]string, err error) {
	adjList = map[string]bool{}
	paths = [][]string{}
	for _, line := range lines {
		switch {
		case strings.Contains(line, "|"):
			parts := strings.Split(line, "|")
			adjList[parts[0]+parts[1]] = true
		case strings.Contains(line, ","):
			parts := strings.Split(line, ",")
			paths = append(paths, parts)
		}
	}
	return adjList, paths, nil
}
