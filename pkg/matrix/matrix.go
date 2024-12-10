package matrix

func InBounds(m [][]string, r, c int) bool {
	return r >= 0 && r < len(m) && c >= 0 && c < len(m[0])
}
