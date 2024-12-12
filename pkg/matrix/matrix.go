package matrix

import "fmt"

type Element struct {
	Row, Col int
}

var Directions = []Element{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func InBounds[V string | int](m [][]V, r, c int) bool {
	return r >= 0 && r < len(m) && c >= 0 && c < len(m[0])
}

func Print[V string | int](m [][]V) {
	fmt.Println()
	for r := range m {
		fmt.Println(m[r])
	}
	fmt.Println()
}
