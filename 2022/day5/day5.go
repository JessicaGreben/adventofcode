package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func part1() string {
	cratemoves, err := movesInput()
	if err != nil {
		fmt.Println(err)
	}
	stacks := inputToStacks(inputStacks)
	return moveCratesPart1(cratemoves, stacks)
}

func part2() string {
	cratemoves, err := movesInput()
	if err != nil {
		fmt.Println(err)
	}
	stacks := inputToStacks(inputStacks)
	return moveCratesPart2(cratemoves, stacks)
}

func moveCratesPart1(moves []cratemove, stacks []*Stack) string {
	for _, move := range moves {
		sFrom := stacks[move.from-1]
		sTo := stacks[move.to-1]
		for i := move.count; i > 0; i-- {
			p := sFrom.Pop()
			if p != "" {
				sTo.Push(p)
			} else {
				fmt.Println("pop is ''")
			}
		}
	}

	var sb strings.Builder
	for _, s := range stacks {
		sb.WriteString(s.Peek())
	}
	return sb.String()
}

func moveCratesPart2(moves []cratemove, stacks []*Stack) string {
	for _, move := range moves {
		buffer := &Stack{}
		sFrom := stacks[move.from-1]

		for i := move.count; i > 0; i-- {
			p := sFrom.Pop()
			if p != "" {
				buffer.Push(p)
			} else {
				fmt.Println("pop is ''")
			}
		}

		sTo := stacks[move.to-1]
		for buffer.Len() > 0 {
			sTo.Push(buffer.Pop())
		}
	}

	var sb strings.Builder
	for _, s := range stacks {
		sb.WriteString(s.Peek())
	}
	return sb.String()
}

func inputToStacks(input [][]string) []*Stack {
	stacks := []*Stack{}
	for _, x := range input {
		s := &Stack{data: slices.Clone(x)}
		stacks = append(stacks, s)
	}
	return stacks
}

type cratemove struct {
	count int
	from  int
	to    int
}

func movesInput() ([]cratemove, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	out := []cratemove{}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		move := strings.Split(scanner.Text(), " ")
		if len(move) != 3 {
			fmt.Println("err len not 3", move)
		}
		count, err := strconv.Atoi(move[0])
		if err != nil {
			fmt.Println("strconv count", move[0])
		}
		to, err := strconv.Atoi(move[1])
		if err != nil {
			fmt.Println("strconv to", move[1])
		}
		from, err := strconv.Atoi(move[2])
		if err != nil {
			fmt.Println("strconv from", move[2])
		}
		cratemv := cratemove{
			count, to, from,
		}
		out = append(out, cratemv)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
