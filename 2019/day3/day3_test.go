package day3

import (
	"testing"
)

func TestProgramPart1AoC(t *testing.T) {
	wirePath1, err := readInput("day3wire1.txt")
	if err != nil {
		t.Error(err)
	}
	wire1, err := NewWire(wirePath1)
	if err != nil {
		t.Error(err)
	}
	wirePath2, err := readInput("day3wire2.txt")
	if err != nil {
		t.Error(err)
	}
	wire2, err := NewWire(wirePath2)
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
	/*            y
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
	gotSteps, err := convertInputToSteps(wirePath)
	if err != nil {
		t.Error(err)
	}
	//t.Logf("%#v\n", gotSteps)
	t.Run("Test ConvertInputToSteps", func(t *testing.T) {
		wantSteps := []Step{
			{1, Point{X: 0, Y: 0}, 0, 0},
			{1, Point{X: 1, Y: 0}, 1, 0},
			{1, Point{X: 2, Y: 0}, 2, 0},
			{1, Point{X: 3, Y: 0}, 3, 0},
			{2, Point{X: 3, Y: 0}, 4, 0},
			{2, Point{X: 3, Y: 1}, 5, 0},
			{2, Point{X: 3, Y: 2}, 6, 0},
			{1, Point{X: 3, Y: 2}, 7, 0},
			{1, Point{X: 2, Y: 2}, 8, 0},
			{1, Point{X: 1, Y: 2}, 9, 0},
			{1, Point{X: 0, Y: 2}, 10, 0},
			{1, Point{X: -1, Y: 2}, 11, 0},
			{1, Point{X: -2, Y: 2}, 12, 0},
			{2, Point{X: -2, Y: 2}, 13, 0},
			{2, Point{X: -2, Y: 1}, 14, 0},
			{2, Point{X: -2, Y: 0}, 15, 0},
			{2, Point{X: -2, Y: -1}, 16, 0},
			{2, Point{X: -2, Y: -2}, 17, 0},
			{1, Point{X: -2, Y: -2}, 18, 0},
			{1, Point{X: -1, Y: -2}, 19, 0},
			{1, Point{X: 0, Y: -2}, 20, 0},
			{1, Point{X: 1, Y: -2}, 21, 0},
		}
		if want, got := len(wantSteps), len(gotSteps); want != got {
			t.Fatalf("want: %d, got: %d", want, got)
		}
		for i, got := range gotSteps {
			want := wantSteps[i]
			if want.lineDirection != got.lineDirection {
				t.Errorf("i=%d line direction want:%d, got:%d", i, want.lineDirection, got.lineDirection)
			}
			if want.point.X != got.point.X {
				t.Errorf("i=%d x want:%d, got:%d", i, want.point.X, got.point.X)
			}
			if want.point.Y != got.point.Y {
				t.Errorf("i=%d y want:%d, got:%d", i, want.point.Y, got.point.Y)
			}
			if want.Value != got.Value {
				t.Errorf("i=%d value want:%d, got:%d", i, want.Value, got.Value)
			}
		}
	})

	p := NewPlane(Point{0, 0})

	t.Run("Test NewPlane Plot", func(t *testing.T) {
		wantPoints := map[Point][]Step{
			{X: 0, Y: 0}:   []Step{{1, Point{X: 0, Y: 0}, 0, 1}},
			{X: 1, Y: 0}:   []Step{{1, Point{X: 1, Y: 0}, 1, 1}},
			{X: 2, Y: 0}:   []Step{{1, Point{X: 2, Y: 0}, 2, 1}},
			{X: 3, Y: 0}:   []Step{{1, Point{X: 3, Y: 0}, 3, 1}, {2, Point{X: 3, Y: 0}, 4, 1}},
			{X: 3, Y: 1}:   []Step{{2, Point{X: 3, Y: 1}, 5, 1}},
			{X: 3, Y: 2}:   []Step{{2, Point{X: 3, Y: 2}, 6, 1}, {1, Point{X: 3, Y: 2}, 7, 1}},
			{X: 2, Y: 2}:   []Step{{1, Point{X: 2, Y: 2}, 8, 1}},
			{X: 1, Y: 2}:   []Step{{1, Point{X: 1, Y: 2}, 9, 1}},
			{X: 0, Y: 2}:   []Step{{1, Point{X: 0, Y: 2}, 10, 1}},
			{X: -1, Y: 2}:  []Step{{1, Point{X: -1, Y: 2}, 11, 1}},
			{X: -2, Y: 2}:  []Step{{1, Point{X: -2, Y: 2}, 12, 1}, {2, Point{X: -2, Y: 2}, 13, 1}},
			{X: -2, Y: 1}:  []Step{{2, Point{X: -2, Y: 1}, 14, 1}},
			{X: -2, Y: 0}:  []Step{{2, Point{X: -2, Y: 0}, 15, 1}},
			{X: -2, Y: -1}: []Step{{2, Point{X: -2, Y: -1}, 16, 1}},
			{X: -2, Y: -2}: []Step{{2, Point{X: -2, Y: -2}, 17, 1}, {1, Point{X: -2, Y: -2}, 18, 1}},
			{X: -1, Y: -2}: []Step{{1, Point{X: -1, Y: -2}, 19, 1}},
			{X: 0, Y: -2}:  []Step{{1, Point{X: 0, Y: -2}, 20, 1}},
			{X: 1, Y: -2}:  []Step{{1, Point{X: 1, Y: -2}, 21, 1}},
		}
		wire2, err := NewWire(wirePath)
		if err != nil {
			t.Fatal(err)
		}
		p.Plot(wire2)

		want, got := Point{0, 0}, p.centralPoint
		if want.X != got.X || want.Y != got.Y {
			t.Errorf("central point: want: %v, got: %v", want, got)
		}
		//t.Logf("%#v\n", p.points)
		if want, got := len(wantPoints), len(p.points); want != got {
			t.Fatalf("points: want: %d, got: %d", want, got)
		}

		for gotPoint, gotSteps := range p.points {
			wantSteps, ok := wantPoints[gotPoint]
			if !ok {
				t.Errorf("key not found: %v", gotPoint)
			}

			if want, got := len(wantSteps), len(gotSteps); want != got {
				t.Fatalf("wrong len: want: %d, got: %d", want, got)
			}

			for i, gotStep := range gotSteps {
				wantStep := wantSteps[i]
				if want, got := wantStep.lineDirection, gotStep.lineDirection; want != got {
					t.Errorf("line direction: want: %v, got: %v", want, got)
				}
				if want, got := wantStep.objectID, gotStep.objectID; want != got {
					t.Errorf("objectID: want: %v, got: %v", want, got)
				}
				if want, got := wantStep.Value, gotStep.Value; want != got {
					t.Errorf("Value: want: %v, got: %v", want, got)
				}
				if want, got := wantStep.point.X, gotStep.point.X; want != got {
					t.Errorf("x: want: %v, got: %v", want, got)
				}
				if want, got := wantStep.point.Y, gotStep.point.Y; want != got {
					t.Errorf("y: want: %v, got: %v", want, got)
				}
			}

		}
	})

	/*  A coordinate plane with 2 "wires" labeled "1" and "2". A "B" denotes where both wires are.
	              y
		          ^
			      |
		       111|111
			  2B222  1
	      <---21--o111----> x
		      2B2222
			   1111
			      |
				  v
	*/
	wirePath2 := []string{"U1", "L4", "D2", "R5"}
	//t.Logf("%#v\n", gotSteps2)

	t.Run("Plot second object", func(t *testing.T) {
		wire2, err := NewWire(wirePath2)
		if err != nil {
			t.Fatal(err)
		}

		p.Plot(wire2)

		if want, got := 2, p.objectCount; want != got {
			t.Errorf("object count: want: %d, got: %d", want, got)
		}
		if want, got := 3, p.closestIntersection; want != got {
			t.Errorf("closest intersection: want: %d, got: %d", want, got)
		}
	})
}
