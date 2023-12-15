package main

import (
	"bufio"
	"os"
	"strings"
)

func solution(file string) (int64, error) {
	return ForEachLineInt64(file)
}

func ForEachLineInt64(file string) (int64, error) {
	fd, err := os.Open(file)
	if err != nil {
		return -1, err
	}
	scanner := bufio.NewScanner(fd)
	var s = []string{}
	for scanner.Scan() {
		line := scanner.Text()
		s = strings.FieldsFunc(line, func(r rune) bool {
			return r == ',' || r == '\n'
		})
	}
	var sum int64
	for _, ss := range s {
		var curr int64
		for i := range ss {
			curr += int64(ss[i])
			curr *= 17
			curr %= 256
		}
		sum += curr
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return sum, nil
}
