package day3

import (
	"testing"
)

func TestProgramPart1AoC(t *testing.T) {
	wirePath1, err := readInput("day3wire1.txt")
	if err != nil {
		t.Error(err)
	}
	wire1, err := convertInputToPoints(wirePath1)
	if err != nil {
		t.Error(err)
	}
	wirePath2, err := readInput("day3wire2.txt")
	if err != nil {
		t.Error(err)
	}
	wire2, err := convertInputToPoints(wirePath2)
	if err != nil {
		t.Error(err)
	}
	p := NewPlane(Point{0, 0})
	p.Plot(wire1)
	p.Plot(wire2)
	if want, got := 557, p.closestIntersection; want != got {
		t.Fatalf("closest intersection want: %d, got: %d", want, got)
	}
}

func TestProgramPart1(t *testing.T) {
	/*
			   ^
			   |
			xxx|xxx
			x  |  x
		<----x--oxxx----> x
			x  |
			xxxx
			   |
			   v
	*/
	wirePath := []string{"R3", "U2", "L5", "D4", "R3"}
	gotPoints, err := convertInputToPoints(wirePath)
	if err != nil {
		t.Error(err)
	}
	//t.Logf("%#v\n", gotPoints)
	t.Run("Test convertInputToPoints", func(t *testing.T) {
		wantPoints := []planePoint{
			{Point{X: 0, Y: 0}, 0, 0},
			{Point{X: 1, Y: 0}, 0, 1},
			{Point{X: 2, Y: 0}, 0, 2},
			{Point{X: 3, Y: 0}, 0, 3},
			{Point{X: 3, Y: 0}, 0, 4},
			{Point{X: 3, Y: 1}, 0, 5},
			{Point{X: 3, Y: 2}, 0, 6},
			{Point{X: 3, Y: 2}, 0, 7},
			{Point{X: 2, Y: 2}, 0, 8},
			{Point{X: 1, Y: 2}, 0, 9},
			{Point{X: 0, Y: 2}, 0, 10},
			{Point{X: -1, Y: 2}, 0, 11},
			{Point{X: -2, Y: 2}, 0, 12},
			{Point{X: -2, Y: 2}, 0, 13},
			{Point{X: -2, Y: 1}, 0, 14},
			{Point{X: -2, Y: 0}, 0, 15},
			{Point{X: -2, Y: -1}, 0, 16},
			{Point{X: -2, Y: -2}, 0, 17},
			{Point{X: -2, Y: -2}, 0, 18},
			{Point{X: -1, Y: -2}, 0, 19},
			{Point{X: 0, Y: -2}, 0, 20},
			{Point{X: 1, Y: -2}, 0, 21},
		}
		if want, got := len(wantPoints), len(gotPoints); want != got {
			t.Fatalf("want: %d, got: %d", want, got)
		}
		for i, got := range gotPoints {
			want := wantPoints[i]
			if want.x() != got.x() {
				t.Errorf("i=%d x want:%d, got:%d", i, want.x(), got.x())
			}
			if want.y() != got.y() {
				t.Errorf("i=%d y want:%d, got:%d", i, want.y(), got.y())
			}
			if want.getOrder() != got.getOrder() {
				t.Errorf("i=%d order want:%d, got:%d", i, want.getOrder(), got.getOrder())
			}
		}
	})

	p := NewPlane(Point{0, 0})

	t.Run("Test NewPlane Plot", func(t *testing.T) {
		p.Plot(gotPoints)

		want, got := Point{0, 0}, p.centralPoint
		if want.x() != got.x() || want.y() != got.y() {
			t.Errorf("central point: want: %v, got: %v", want, got)
		}

		wantPoints := map[Point][]planePoint{
			{X: 0, Y: 0}:   {{Point{X: 0, Y: 0}, 1, 0}},
			{X: 1, Y: 0}:   {{Point{X: 1, Y: 0}, 1, 1}},
			{X: 2, Y: 0}:   {{Point{X: 2, Y: 0}, 1, 2}},
			{X: 3, Y: 0}:   {{Point{X: 3, Y: 0}, 1, 3}, {Point{X: 3, Y: 0}, 1, 4}},
			{X: 3, Y: 1}:   {{Point{X: 3, Y: 1}, 1, 5}},
			{X: 3, Y: 2}:   {{Point{X: 3, Y: 2}, 1, 6}, {Point{X: 3, Y: 2}, 1, 7}},
			{X: 2, Y: 2}:   {{Point{X: 2, Y: 2}, 1, 8}},
			{X: 1, Y: 2}:   {{Point{X: 1, Y: 2}, 1, 9}},
			{X: 0, Y: 2}:   {{Point{X: 0, Y: 2}, 1, 10}},
			{X: -1, Y: 2}:  {{Point{X: -1, Y: 2}, 1, 11}},
			{X: -2, Y: 2}:  {{Point{X: -2, Y: 2}, 1, 12}, {Point{X: -2, Y: 2}, 1, 13}},
			{X: -2, Y: 1}:  {{Point{X: -2, Y: 1}, 1, 14}},
			{X: -2, Y: 0}:  {{Point{X: -2, Y: 0}, 1, 15}},
			{X: -2, Y: -1}: {{Point{X: -2, Y: -1}, 1, 16}},
			{X: -2, Y: -2}: {{Point{X: -2, Y: -2}, 1, 17}, {Point{X: -2, Y: -2}, 1, 18}},
			{X: -1, Y: -2}: {{Point{X: -1, Y: -2}, 1, 19}},
			{X: 0, Y: -2}:  {{Point{X: 0, Y: -2}, 1, 20}},
			{X: 1, Y: -2}:  {{Point{X: 1, Y: -2}, 1, 21}},
		}
		//t.Logf("%#v\n", p.points)
		if want, got := len(wantPoints), len(p.points); want != got {
			t.Fatalf("points: want: %d, got: %d", want, got)
		}

		for gotPoint, gotPoints := range p.points {
			wantPoints, ok := wantPoints[gotPoint]
			if !ok {
				t.Errorf("key not found: %v", gotPoint)
			}

			if want, got := len(wantPoints), len(gotPoints); want != got {
				t.Fatalf("wrong len: want: %d, got: %d", want, got)
			}

			for i, gotPoint := range gotPoints {
				wantPoint := wantPoints[i]
				if want, got := wantPoint.getObjectID(), gotPoint.getObjectID(); want != got {
					t.Errorf("objectID: want: %v, got: %v", want, got)
				}
				if want, got := wantPoint.getOrder(), gotPoint.getOrder(); want != got {
					t.Errorf("order: want: %v, got: %v", want, got)
				}
				if want, got := wantPoint.x(), gotPoint.x(); want != got {
					t.Errorf("x: want: %v, got: %v", want, got)
				}
				if want, got := wantPoint.y(), gotPoint.y(); want != got {
					t.Errorf("y: want: %v, got: %v", want, got)
				}
			}
		}
	})

	/*  A coordinate plane with 2 "wires" labeled "1" and "2". A "B" denotes where both wires are.
		   ^
		   |
		111|111
		2B222  1
	<---21--o111---->
		2B2222
		1111
		   |
		   v
	*/
	wirePath2 := []string{"U1", "L4", "D2", "R5"}
	//t.Logf("%#v\n", gotPoints2)

	t.Run("Plot second object", func(t *testing.T) {
		wire2, err := convertInputToPoints(wirePath2)
		if err != nil {
			t.Fatal(err)
		}
		p.Plot(wire2)

		if want, got := 2, p.objectIDCounter; want != got {
			t.Errorf("object count: want: %d, got: %d", want, got)
		}
		if want, got := 3, p.closestIntersection; want != got {
			t.Errorf("closest intersection: want: %d, got: %d", want, got)
		}
	})
}
