package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() int {
	in, err := input("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	return countNoBeacons(in, 2000000)
}

func input(file string) ([]sensor, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	out := []sensor{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		sensorX, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println("err strconv", err)
		}
		sensorY, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println("err strconv", err)
		}
		beaconX, err := strconv.Atoi(line[2])
		if err != nil {
			fmt.Println("err strconv", err)
		}
		beaconY, err := strconv.Atoi(line[3])
		if err != nil {
			fmt.Println("err strconv", err)
		}
		out = append(out, sensor{point{sensorX, sensorY}, point{beaconX, beaconY}})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func countNoBeacons(sensors []sensor, desiredY int) int {
	yRowXvalues := map[int]bool{}
	yRowBeaconLocations := map[int]bool{}

	for _, sensor := range sensors {
		d := manhattanDistance(sensor.position, sensor.beacon)
		x, y := sensor.position.x, sensor.position.y
		if desiredY > y+d || desiredY < y-d {
			continue
		}
		if sensor.beacon.y == desiredY {
			yRowBeaconLocations[sensor.beacon.x] = true
		}
		if y == desiredY {
			yRowBeaconLocations[x] = true
		}
		yOffset := 0
		for xOffset := d; xOffset >= 0; xOffset-- {
			yUp, yDown := y+yOffset, y-yOffset
			yOffset++
			if yUp != desiredY && yDown != desiredY {
				continue
			}
			xLeft, xRight := x-xOffset, x+xOffset
			for x := xLeft; x <= xRight; x++ {
				yRowXvalues[x] = true
			}
		}
	}

	return len(yRowXvalues) - len(yRowBeaconLocations)
}

type sensor struct {
	position point
	beacon   point
}

type point struct {
	x, y int
}

// |x1 - x2| + |y1 - y2|
func manhattanDistance(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
