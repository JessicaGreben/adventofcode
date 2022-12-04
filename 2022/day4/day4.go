package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() int {
	pairs, err := input()
	if err != nil {
		fmt.Println(err)
	}
	result := 0
	for _, pair := range pairs {
		if pair.fulloverlap() {
			result++
		}
	}
	return result
}

func part2() int {
	pairs, err := input()
	if err != nil {
		fmt.Println(err)
	}
	result := 0
	for _, pair := range pairs {
		if pair.partialoverlap() {
			result++
		}
	}
	return result
}

type pair struct {
	a []int
	b []int
}

func (p pair) fulloverlap() bool {
	return (p.a[0] >= p.b[0] && p.a[1] <= p.b[1]) || // a within b
		(p.b[0] >= p.a[0] && p.b[1] <= p.a[1]) // b within a
}

func (p pair) partialoverlap() bool {
	return (p.a[0] >= p.b[0] && p.a[0] <= p.b[1]) || // a start overlap
		(p.a[1] >= p.b[0] && p.a[1] <= p.b[1]) || // a end overlap
		(p.b[0] >= p.a[0] && p.b[0] <= p.a[1]) || // b start overlap
		(p.b[1] >= p.a[0] && p.b[1] <= p.a[1]) // b end overlap
}

func input() ([]pair, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	out := []pair{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")
		if len(pairs) != 2 {
			fmt.Println("err len not 2", pairs)
		}
		a := strings.Split(pairs[0], "-")
		if len(pairs) != 2 {
			fmt.Println("err len not 2 a", a)
		}
		a1, err := strconv.Atoi(a[0])
		if err != nil {
			return nil, err
		}
		a2, err := strconv.Atoi(a[1])
		if err != nil {
			return nil, err
		}
		b := strings.Split(pairs[1], "-")
		if len(pairs) != 2 {
			fmt.Println("err len not 2 b", b)
		}
		b1, err := strconv.Atoi(b[0])
		if err != nil {
			return nil, err
		}
		b2, err := strconv.Atoi(b[1])
		if err != nil {
			return nil, err
		}
		p := pair{
			a: []int{a1, a2},
			b: []int{b1, b2},
		}
		out = append(out, p)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
