package day1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	upper = "ABCDEFGHIJKLMNOPQRSTUVWYZ"
	lower = "abcdefghijklmnopqrstuvwxyz"
)

func part1() int {
	in, err := input1()
	if err != nil {
		fmt.Println(err)
	}
	return sumPriority(in)
}

func part2() int {
	in, err := input2()
	if err != nil {
		fmt.Println(err)
	}
	return sumPriority2(in)
}

func sumPriority2(groups [][]string) int {
	sum := 0
	for _, group := range groups {
		if len(group) != 3 {
			fmt.Println("err not len 3", group)
		}
		f := map[rune]int{}
		for _, char := range group[0] {
			f[char]++
		}
		f2 := map[rune]int{}
		for _, char := range group[1] {
			_, ok := f[char]
			if ok {
				f2[char]++
			}
		}
		for _, char := range group[2] {
			_, ok := f2[char]
			if ok {
				if strings.Contains(upper, string(char)) {
					sum += int(byte(char)) - 64 + 26
				} else {
					sum += int(byte(char)) - 96
				}
				break
			}
		}

	}
	return sum
}

func sumPriority(sacks []string) int {
	sum := 0
	for _, sack := range sacks {
		n := len(sack) / 2
		first, second := sack[:n], sack[n:]
		f := map[rune]int{}
		for _, char := range first {
			f[char]++
		}
		for _, char := range second {
			if _, ok := f[char]; ok {
				if strings.Contains(upper, string(char)) {
					sum += int(byte(char)) - 64 + 26
				} else {
					sum += int(byte(char)) - 96
				}
				break
			}
		}
	}
	return sum
}

func input1() ([]string, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	out := []string{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func input2() ([][]string, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	out := []string{}
	res := [][]string{}
	scanner := bufio.NewScanner(fd)
	count := 1
	for scanner.Scan() {
		out = append(out, scanner.Text())
		count++
		if count == 4 {
			res = append(res, out)
			out = []string{}
			count = 1
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
