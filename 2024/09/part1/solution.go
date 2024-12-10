package main

import (
	fileinput "github.com/jessicagreben/adventofcode/pkg/input"
)

func solution(file string) (int64, error) {
	input, err := fileinput.New(file)
	if err != nil {
		return -1, err
	}

	var result int64
	for line := range input.All() {
		diskMap, err := parseLine(line)
		if err != nil {
			return -1, nil
		}
		diskMapIDs := convertToIDs(diskMap)
		compact(diskMapIDs)
		for i, x := range diskMapIDs {
			if x == -1 {
				return result, nil
			}
			result += int64(i) * int64(x)
		}

	}
	return result, nil
}

const freeSpace = -1

func compact(diskMapIDs []int) {
	var freeSpaceIdx int
	moveIdx := len(diskMapIDs) - 1

	for freeSpaceIdx < moveIdx {
		if diskMapIDs[freeSpaceIdx] != freeSpace {
			freeSpaceIdx++
			continue
		}
		if diskMapIDs[moveIdx] == freeSpace {
			moveIdx--
			continue
		}
		diskMapIDs[freeSpaceIdx], diskMapIDs[moveIdx] = diskMapIDs[moveIdx], diskMapIDs[freeSpaceIdx]
		freeSpaceIdx++
		moveIdx--
	}
}

func convertToIDs(diskMap []int64) []int {
	result := []int{}
	for i, x := range diskMap {
		if i%2 == 0 { // file
			for range x {
				id := i / 2
				result = append(result, id)
			}

		} else { // free space
			for range x {
				result = append(result, freeSpace)
			}
		}

	}
	return result
}

func parseLine(line string) ([]int64, error) {
	return fileinput.ParseLineInt64(line, "", -1)
}
