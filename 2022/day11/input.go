package day11

/*
Monkey 0:
  Starting items: 92, 73, 86, 83, 65, 51, 55, 93
  Operation: new = old * 5
  Test: divisible by 11
    If true: throw to monkey 3
    If false: throw to monkey 4

Monkey 1:
  Starting items: 99, 67, 62, 61, 59, 98
  Operation: new = old * old
  Test: divisible by 2
    If true: throw to monkey 6
    If false: throw to monkey 7

Monkey 2:
  Starting items: 81, 89, 56, 61, 99
  Operation: new = old * 7
  Test: divisible by 5
    If true: throw to monkey 1
    If false: throw to monkey 5

Monkey 3:
  Starting items: 97, 74, 68
  Operation: new = old + 1
  Test: divisible by 17
    If true: throw to monkey 2
    If false: throw to monkey 5

Monkey 4:
  Starting items: 78, 73
  Operation: new = old + 3
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 5:
  Starting items: 50
  Operation: new = old + 5
  Test: divisible by 7
    If true: throw to monkey 1
    If false: throw to monkey 6

Monkey 6:
  Starting items: 95, 88, 53, 75
  Operation: new = old + 8
  Test: divisible by 3
    If true: throw to monkey 0
    If false: throw to monkey 7

Monkey 7:
  Starting items: 50, 77, 98, 85, 94, 56, 89
  Operation: new = old + 2
  Test: divisible by 13
    If true: throw to monkey 4
    If false: throw to monkey 0

*/

var inputMonkeys = []monkey{
	{
		items:     []int{92, 73, 86, 83, 65, 51, 55, 93},
		op:        "*",
		opInt:     5,
		test:      11,
		testTrue:  3,
		testFalse: 4,
	},
	{
		items:     []int{99, 67, 62, 61, 59, 98},
		op:        "*",
		opInt:     -1,
		test:      2,
		testTrue:  6,
		testFalse: 7,
	},
	{
		items:     []int{81, 89, 56, 61, 99},
		op:        "*",
		opInt:     7,
		test:      5,
		testTrue:  1,
		testFalse: 5,
	},
	{
		items:     []int{97, 74, 68},
		op:        "+",
		opInt:     1,
		test:      17,
		testTrue:  2,
		testFalse: 5,
	},
	{
		items:     []int{78, 73},
		op:        "+",
		opInt:     3,
		test:      19,
		testTrue:  2,
		testFalse: 3,
	},
	{
		items:     []int{50},
		op:        "+",
		opInt:     5,
		test:      7,
		testTrue:  1,
		testFalse: 6,
	},
	{
		items:     []int{95, 88, 53, 75},
		op:        "+",
		opInt:     8,
		test:      3,
		testTrue:  0,
		testFalse: 7,
	},
	{
		items:     []int{50, 77, 98, 85, 94, 56, 89},
		op:        "+",
		opInt:     2,
		test:      13,
		testTrue:  4,
		testFalse: 0,
	},
}
