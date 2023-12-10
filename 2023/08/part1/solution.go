package main

func solution(file string) (int, error) {
	path, nodes, err := ForEachLine(file)
	if err != nil {
		return -1, err
	}
	s := countSteps(path, nodes)
	return s, nil
}

const (
	startNode = "AAA"
	endNode   = "ZZZ"

	right = "R"
	left  = "L"
)

func countSteps(path []string, nodes map[string][]string) int {
	currNode := startNode
	pathIdx := 0
	steps := 0
	for currNode != endNode {
		if pathIdx >= len(path) {
			pathIdx = 0
		}
		if path[pathIdx] == left {
			currNode = nodes[currNode][0]
		}
		if path[pathIdx] == right {
			currNode = nodes[currNode][1]
		}
		pathIdx++
		steps++
	}
	return steps
}
