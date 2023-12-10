package input

import (
	"bufio"
	"os"
)

func ForEachLine(file string, fn func(string) (int, error)) (int, error) {
	fd, err := os.Open(file)
	if err != nil {
		return -1, err
	}
	scanner := bufio.NewScanner(fd)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		x, err := fn(line)
		if err != nil {
			return -1, err
		}
		sum += x
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return sum, nil
}

func ForEachLineInt64(file string, fn func(string) (int64, error)) (int64, error) {
	fd, err := os.Open(file)
	if err != nil {
		return -1, err
	}
	scanner := bufio.NewScanner(fd)
	var sum int64
	for scanner.Scan() {
		line := scanner.Text()
		x, err := fn(line)
		if err != nil {
			return -1, err
		}
		sum += x
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return sum, nil
}
