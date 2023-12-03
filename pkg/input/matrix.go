package input

import (
	"bufio"
	"os"
	"strings"
)

func ConvertToMatrix(file string) ([][]string, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	m := [][]string{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		m = append(m, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return m, nil
}
