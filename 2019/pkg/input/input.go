package input

import (
	"bufio"
	"os"
)

func ReadFromFile(filepath string) ([]string, error) {
	fd, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fd)
	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return input, nil
}
