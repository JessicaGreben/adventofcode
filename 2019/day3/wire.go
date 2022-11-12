package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertInputToPoints(input []string) ([]planePoint, error) {
	points := []planePoint{}
	currX, currY := 0, 0
	order := 0
	for _, path := range input {
		distance, err := strconv.Atoi(string(path[1:]))
		if err != nil {
			return points, err
		}
		switch direction := path[0]; direction {
		case 'R':
			for i := currX; i <= currX+distance; i++ {
				points = append(points, planePoint{Point{i, currY}, 0, order})
				order++
			}
			currX += distance
		case 'L':
			for i := currX; i >= currX-distance; i-- {
				points = append(points, planePoint{Point{i, currY}, 0, order})
				order++
			}
			currX -= distance
		case 'U':
			for i := currY; i <= currY+distance; i++ {
				points = append(points, planePoint{Point{currX, i}, 0, order})
				order++
			}
			currY += distance
		case 'D':
			for i := currY; i >= currY-distance; i-- {
				points = append(points, planePoint{Point{currX, i}, 0, order})
				order++
			}
			currY -= distance
		default:
			return points, fmt.Errorf("direction not supported: %s", string(direction))
		}
	}

	return points, nil
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
