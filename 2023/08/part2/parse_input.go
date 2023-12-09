package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ForEachLine(file string) ([]string, map[string][]string, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(fd)

	path := []string{}
	nodes := map[string][]string{}
	if scanner.Scan() {
		firstLine := scanner.Text()
		path = strings.Split(firstLine, "")
	}
	if scanner.Scan() {
		secondLine := scanner.Text()
		if secondLine != "" {
			fmt.Println("2nd line not empty")
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		x := strings.Split(line, " = ")

		s := strings.FieldsFunc(x[1], func(r rune) bool {
			return r == '(' || r == ')' || r == ',' || r == ' '
		})

		if len(s) != 2 {
			fmt.Println("len is not 2", s)
		}
		nodes[x[0]] = s
		if err != nil {
			return nil, nil, err
		}

	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return path, nodes, nil
}
