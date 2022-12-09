package day9

import (
	"testing"
)

func TestPart1(t *testing.T) {
	if want, got := 5878, part1(); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestCountTailPositions(t *testing.T) {
	moves := []move{
		{"R", 4},
		{"U", 4},
		{"L", 3},
		{"D", 1},
		{"R", 4},
		{"D", 1},
		{"L", 5},
		{"R", 2},
	}
	if want, got := 13, countTailPositions(moves); want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestDiagonalMove(t *testing.T) {
	testCases := []struct {
		name             string
		prev, head, tail *position
		row, col         int
		same             bool
	}{
		{
			/*
				  012345
				5 ..H... p2 (r=5,c=2)
				4 ....T. p  (r=4,c=4)
				3 ......
				1 ......
				0 s.....
			*/
			"1",
			&position{5, 3}, &position{5, 2}, &position{4, 4},
			5, 3,
			false,
		},
		{
			/*
				  012345
				5 ...H.. p2 (r=5,c=3)
				4 .T.... p  (r=4,c=1)
				3 ......
				1 ......
				0 s.....
			*/
			"2",
			&position{5, 2}, &position{5, 3}, &position{4, 1},
			5, 2,
			false,
		},
		{
			/*
				  012345
				5 ......
				4 .T.... p  (r=4,c=1)
				3 ...H.. p2 (r=3,c=3)
				1 ......
				0 s.....
			*/
			"3",
			&position{3, 2}, &position{3, 3}, &position{4, 1},
			3, 2,
			false,
		},
		{
			/*
				  012345
				5 ......
				4 ....T. p  (r=4,c=4)
				3 ..H... p2 (r=3,c=2)
				1 ......
				0 s.....
			*/
			"4",
			&position{3, 3}, &position{3, 2}, &position{4, 4},
			3, 3,
			false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.same, tc.tail.adjacent(tc.head); want != got {
				t.Errorf("same want: %v, got: %v", want, got)
			}
			updateTail(tc.prev, tc.head, tc.tail)
			if want, got := tc.row, tc.tail.r; want != got {
				t.Errorf("row want: %v, got: %v", want, got)
			}
			if want, got := tc.col, tc.tail.c; want != got {
				t.Errorf("col want: %v, got: %v", want, got)
			}
			if want, got := tc.same, tc.tail.adjacent(tc.head); want == got {
				t.Errorf("same want: %v, got: %v", want, got)
			}
		})
	}
}
