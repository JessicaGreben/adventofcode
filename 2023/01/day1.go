package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const nums = "0123456789"

func main() {
	solution, err := readInput(doPart1)
	if err != nil {
		fmt.Println("main err=", err)
		os.Exit(1)
	}
	fmt.Println("Day 1 Part 1 solution:", solution)

	solution, err = readInput(doPart2)
	if err != nil {
		fmt.Println("main err=", err)
		os.Exit(1)
	}
	fmt.Println("Day 1 Part 2 solution:", solution)
}

func readInput(fn func(string) (int, error)) (int, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return -1, err
	}
	scanner := bufio.NewScanner(fd)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		num, err := fn(line)
		if err != nil {
			fmt.Println("err=", err)
		}
		sum += num
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return sum, nil
}

func doPart1(line string) (int, error) {
	var digit1, digit10 string

	i, j := 0, len(line)-1
	for i <= j && i < len(line) && j >= 0 {
		if digit10 == "" && strings.Contains(nums, string(line[i])) {
			digit10 = string(line[i])
		}
		if digit1 == "" && strings.Contains(nums, string(line[j])) {
			digit1 = string(line[j])
		}
		if digit10 != "" && digit1 != "" {
			tensPlace, err := strconv.Atoi(digit10)
			if err != nil {
				fmt.Println("tens:", digit10)
				return -1, err
			}
			onesPlace, err := strconv.Atoi(digit1)
			if err != nil {
				fmt.Println("one:", digit1)
				return -1, err
			}
			num := (10 * tensPlace) + onesPlace
			return num, nil
		}
		if digit10 == "" {
			i++
		}
		if digit1 == "" {
			j--
		}
	}
	return -1, errors.New("got to end of loop, should not have")
}

var strNums = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func doPart2(line string) (int, error) {
	indexes := []int{}
	idxToNums := map[int]string{}
	for k := range strNums {
		firstIdx := strings.Index(line, k)
		lastIndex := strings.LastIndex(line, k)
		if firstIdx != -1 && firstIdx == lastIndex {
			indexes = append(indexes, firstIdx)
			idxToNums[firstIdx] = k
		}

		if firstIdx != -1 && lastIndex != -1 && firstIdx != lastIndex {
			indexes = append(indexes, firstIdx)
			idxToNums[firstIdx] = k
			indexes = append(indexes, lastIndex)
			idxToNums[lastIndex] = k
		}

	}
	if len(indexes) < 1 {
		return -1, errors.New("no digits found")
	}

	slices.Sort(indexes)
	tensPlaceIdx := indexes[0]
	onesPlaceIdx := indexes[len(indexes)-1]

	tensPlaceStrNum := idxToNums[tensPlaceIdx]
	onesPlaceStrNum := idxToNums[onesPlaceIdx]

	tensPlaceNum := strNums[tensPlaceStrNum]
	onesPlaceNum := strNums[onesPlaceStrNum]

	return (10 * tensPlaceNum) + onesPlaceNum, nil
}
