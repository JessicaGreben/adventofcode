package main

import (
	"slices"
	"strconv"
	"strings"

	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	lines, err := fileinput.GetLines(file)
	if err != nil {
		return -1, err
	}
	edges, paths, err := parseLines(lines)
	if err != nil {
		return -1, err
	}

	var result int64
	for _, path := range paths {
		if !isValid(path, edges) {
			mid, ok := findValidPath(path, edges, map[string]bool{})
			if ok {
				result += mid
			}
		}
	}

	return result, nil
}

func findValidPath(path []string, g map[string]bool, seen map[string]bool) (int64, bool) {
	key := strings.Join(path, "")
	if _, ok := seen[key]; ok {
		return -1, false
	}
	seen[key] = true
	for i := range path {
		for j := i + 1; j < len(path); j++ {
			a, b := path[i], path[j]
			if _, ok := g[a+b]; !ok {
				cp := slices.Clone(path)
				cp[i], cp[j] = cp[j], cp[i]
				if x, ok := findValidPath(cp, g, seen); ok {
					return x, ok
				}
			}
		}
	}
	pathMidStr := path[len(path)/2]
	pathMidInt, err := strconv.Atoi(pathMidStr)
	if err != nil {
		return -1, false
	}
	return int64(pathMidInt), true
}

func isValid(path []string, edges map[string]bool) bool {
	for i := range path {
		for j := i + 1; j < len(path); j++ {
			a, b := path[i], path[j]
			if _, ok := edges[a+b]; !ok {
				return false
			}
		}
	}
	return true
}

func parseLines(lines []string) (edges map[string]bool, paths [][]string, err error) {
	edges = map[string]bool{}
	paths = [][]string{}
	for _, line := range lines {
		switch {
		case strings.Contains(line, "|"):
			parts := strings.Split(line, "|")
			edges[parts[0]+parts[1]] = true
		case strings.Contains(line, ","):
			parts := strings.Split(line, ",")
			paths = append(paths, parts)
		}
	}
	return edges, paths, nil
}
