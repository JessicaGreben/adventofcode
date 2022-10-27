package input

import (
	"bufio"
	"os"
	"strconv"
)

func ReadIntsFromFile(filepath string) ([]int, error) {
	fd, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fd)
	input := []int{}
	for scanner.Scan() {
		valueInt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		input = append(input, valueInt)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return input, nil
}
