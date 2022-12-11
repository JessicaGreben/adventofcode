package day11

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestPart1(t *testing.T) {
	if want, got := 120756, part1(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	if want, got := -1, part2(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

/*
Monkey 0:

	Starting items: 79, 98
	Operation: new = old * 19
	Test: divisible by 23
	  If true: throw to monkey 2
	  If false: throw to monkey 3

Monkey 1:

	Starting items: 54, 65, 75, 74
	Operation: new = old + 6
	Test: divisible by 19
	  If true: throw to monkey 2
	  If false: throw to monkey 0

Monkey 2:

	Starting items: 79, 60, 97
	Operation: new = old * old
	Test: divisible by 13
	  If true: throw to monkey 1
	  If false: throw to monkey 3

Monkey 3:

	Starting items: 74
	Operation: new = old + 3
	Test: divisible by 17
	  If true: throw to monkey 0
	  If false: throw to monkey 1
*/
func TestProgrampart1(t *testing.T) {
	testMonkeys := []monkey{
		{
			items:     []int{79, 98},
			op:        "*",
			opInt:     19,
			test:      23,
			testTrue:  2,
			testFalse: 3,
		},
		{
			items:     []int{54, 65, 75, 74},
			op:        "+",
			opInt:     6,
			test:      19,
			testTrue:  2,
			testFalse: 0,
		},
		{
			items:     []int{79, 60, 97},
			op:        "*",
			opInt:     -1,
			test:      13,
			testTrue:  1,
			testFalse: 3,
		},
		{
			items:     []int{74},
			op:        "+",
			opInt:     3,
			test:      17,
			testTrue:  0,
			testFalse: 1,
		},
	}

	chasedMonkeys := chaseMonkeys(slices.Clone(testMonkeys), 3, 20)
	if want, got := 10605, monkeyBusiness(chasedMonkeys); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
