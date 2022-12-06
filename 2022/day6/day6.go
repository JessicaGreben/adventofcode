package day6

func part1() int {
	return findMarker(dataInput, 4)
}

func part2() int {
	return findMarker(dataInput, 14)
}

// find the first substring in the input data of size n with all unique
// characters then return the first index after the end of the substring
func findMarker(data string, n int) int {
	f := map[string]int{}
	var i int
	for i < n {
		f[string(data[i])]++
		i++
	}
	if len(f) == n {
		return 0
	}
	for i < len(data) {
		f[string(data[i-n])]--
		if f[string(data[i-n])] == 0 {
			delete(f, string(data[i-n]))
		}
		f[string(data[i])]++
		if len(f) == n {
			return i + 1
		}
		i++
	}
	return -1
}
