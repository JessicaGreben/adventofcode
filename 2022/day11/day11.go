package day11

import (
	"fmt"
	"sort"

	"golang.org/x/exp/slices"
)

func part1() int {
	chasedMonkeys := chaseMonkeys(slices.Clone(inputMonkeys), 3, 20)
	return monkeyBusiness(chasedMonkeys)
}

func part2() int {
	chasedMonkeys := chaseMonkeys(slices.Clone(inputMonkeys), 1, 10000)
	return monkeyBusiness(chasedMonkeys)
}

func monkeyBusiness(input []monkey) int {
	sort.Slice(input, func(i, j int) bool {
		return input[i].inspectionCount > input[j].inspectionCount
	})
	return input[0].inspectionCount * input[1].inspectionCount
}

func chaseMonkeys(monkeys []monkey, divide, rounds int) []monkey {
	var round int
	for round < rounds {
		for i, monkey := range monkeys {
			throwItems := []throw{}
			for _, item := range monkey.items {
				monkeys[i].inspectionCount++
				nw := newWorryLevel(item, monkey.op, monkey.opInt, divide)
				t := throw{item: nw, to: monkey.testFalse}
				if t.item%monkey.test == 0 {
					t.to = monkey.testTrue
				}
				throwItems = append(throwItems, t)
			}
			monkeys[i].items = []int{}
			for _, throwitem := range throwItems {
				mto := throwitem.to
				monkeys[mto].items = append(monkeys[mto].items, throwitem.item)
			}
		}
		round++
	}
	return monkeys
}

func newWorryLevel(originalWorry int, op string, operand, divide int) int {
	var newWorry int
	if operand == -1 {
		operand = originalWorry
	}
	switch op {
	case "+":
		newWorry = originalWorry + operand
	case "*":
		newWorry = originalWorry * operand
	default:
		fmt.Println("not a valid op")
	}
	return newWorry / divide
}
