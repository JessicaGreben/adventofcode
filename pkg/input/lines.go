package input

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	scanner *bufio.Scanner
}

func NewInput(filepath string) (*Input, error) {
	fd, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return &Input{
		scanner: bufio.NewScanner(fd),
	}, nil
}

func (i *Input) All() iter.Seq[string] {
	return func(yield func(string) bool) {
		defer func() {
			if err := i.scanner.Err(); err != nil {
				fmt.Println("error scanning input file", err)
			}
		}()
		for i.scanner.Scan() {
			line := i.scanner.Text()
			if !yield(line) {
				return
			}
		}
	}
}

// ParseLineInt64 splits the line by the delimeter and ensures there are partCount items after its split.
// Each part is converted to an Int64 and returned in a slice.
// The items in the slice are in order that they are in the original line.
// For example, if the line is this: "1 2 3" and delimeter is " ", then the output will be [1, 2, 3].
func ParseLineInt64(line, delimeter string, partCount int) ([]int64, error) {
	out := []int64{}
	lineParts := strings.Split(line, delimeter)
	if partCount != -1 && len(lineParts) != partCount {
		return out, fmt.Errorf("wrong number of parts, want=%d, got=%d, parts=%v", partCount, len(lineParts), lineParts)
	}

	for i := range partCount {
		part := lineParts[i]
		partInt, err := strconv.Atoi(part)
		if err != nil {
			return out, err
		}
		out = append(out, int64(partInt))
	}
	return out, nil
}

// GetLines returns a slice of strings where each string is one line in the file.
func GetLines(file string) ([]string, error) {
	lines := []string{}
	fd, err := os.Open(file)
	if err != nil {
		return lines, err
	}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

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
