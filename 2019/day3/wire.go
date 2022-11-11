package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type lineType int

const (
	lineTypeHorizontal lineType = 1
	lineTypeVertical   lineType = 2
)

// Wire is a type of object that can be plotted on a coordinate plane.
type Wire struct {
	steps []Step
}

func NewWire(path []string) (*Wire, error) {
	steps, err := convertInputToSteps(path)
	if err != nil {
		return nil, err
	}
	return &Wire{
		steps: steps,
	}, nil
}

func (w *Wire) Steps() []Step {
	return w.steps
}

func readInput(filename string) ([]string, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	directions := []string{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		text := scanner.Text()
		if err != nil {
			return nil, err
		}
		directions = append(directions, text)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return directions, nil
}

func convertInputToSteps(input []string) ([]Step, error) {
	steps := []Step{}
	currX, currY := 0, 0
	currStep := 0
	for _, path := range input {
		distance, err := strconv.Atoi(string(path[1:]))
		if err != nil {
			return steps, err
		}
		switch direction := path[0]; direction {
		case 'R':
			for i := currX; i <= currX+distance; i++ {
				steps = append(steps, Step{lineTypeHorizontal, Point{i, currY}, currStep, 0})
				currStep++
			}
			currX += distance
		case 'L':
			for i := currX; i >= currX-distance; i-- {
				steps = append(steps, Step{lineTypeHorizontal, Point{i, currY}, currStep, 0})
				currStep++
			}
			currX -= distance
		case 'U':
			for i := currY; i <= currY+distance; i++ {
				steps = append(steps, Step{lineTypeVertical, Point{currX, i}, currStep, 0})
				currStep++
			}
			currY += distance
		case 'D':
			for i := currY; i >= currY-distance; i-- {
				steps = append(steps, Step{lineTypeVertical, Point{currX, i}, currStep, 0})
				currStep++
			}
			currY -= distance
		default:
			return steps, fmt.Errorf("direction not supported: %s", string(direction))
		}
	}

	return steps, nil
}
