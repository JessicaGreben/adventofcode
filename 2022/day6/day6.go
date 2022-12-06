package day6

func part1() int {
	return findPacketMarker(dataInput)
}

func part2() int {
	return findMessage(dataInput)
}

func findPacketMarker(data string) int {
	f := map[string]int{}
	var i int
	for i < 4 {
		f[string(data[i])]++
		i++
	}
	if len(f) == 4 {
		return 0
	}
	for i < len(data) {
		f[string(data[i-4])]--
		if f[string(data[i-4])] == 0 {
			delete(f, string(data[i-4]))
		}
		f[string(data[i])]++
		if len(f) == 4 {
			return i + 1
		}
		i++
	}
	return -1
}

func findMessage(data string) int {
	f := map[string]int{}
	var i int
	for i < 14 {
		f[string(data[i])]++
		i++
	}
	if len(f) == 14 {
		return 0
	}
	for i < len(data) {
		f[string(data[i-14])]--
		if f[string(data[i-14])] == 0 {
			delete(f, string(data[i-14]))
		}
		f[string(data[i])]++
		if len(f) == 14 {
			return i + 1
		}
		i++
	}
	return -1
}
