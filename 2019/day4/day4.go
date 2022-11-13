package main

import (
	"strconv"
	"strings"
)

func day4part1() int {
	pwMin, pwMax := 134792, 675810
	var count int
	for i := pwMin; i <= pwMax; i++ {
		if isValidPasswordPart1(i) {
			count++
		}
	}
	return count
}

func day4part2() int {
	pwMin, pwMax := 134792, 675810
	var count int
	for i := pwMin; i <= pwMax; i++ {
		if isValidPasswordPart2(i) {
			count++
		}
	}
	return count
}

// a valid password has two adjacent digits that are the same
// and the digits from left to right are always increasing.
func isValidPasswordPart1(pw int) bool {
	intDigits := strings.Split(strconv.Itoa(pw), "")
	return twoAdjNums(intDigits) && isAsc(intDigits)
}

func twoAdjNums(pw []string) bool {
	for i := 1; i < len(pw); i++ {
		prev, curr := pw[i-1], pw[i]
		if prev == curr {
			return true
		}
	}
	return false
}

func isAsc(pw []string) bool {
	for i := 1; i < len(pw); i++ {
		prev, curr := pw[i-1], pw[i]
		if prev > curr {
			return false
		}
	}
	return true
}

// a valid password is defined as:
//   - has two adjacent matching digits that are not part of a larger group of matching digits
//   - digits from left to right are increasing
func isValidPasswordPart2(pw int) bool {
	intDigits := strings.Split(strconv.Itoa(pw), "")
	occurTwice := occurTwoAdjNums(intDigits)
	occurThree := occurThreeAdjNums(intDigits)
	return (len(occurTwice) >= 1 && len(occurTwice) > len(occurThree)) && isAsc(intDigits)
}

// occurTwoAdjNums counts how many times two adjacent values
// are matching
func occurTwoAdjNums(pw []string) map[string]bool {
	occur := map[string]bool{}
	for i := 1; i < len(pw); i++ {
		prev, curr := pw[i-1], pw[i]
		if prev == curr {
			occur[curr] = true
		}
	}
	return occur
}

// occurThreeAdjNums counts how many times three adjacent values
// are matching
func occurThreeAdjNums(pw []string) map[string]bool {
	occur := map[string]bool{}
	for i := 1; i < len(pw)-1; i++ {
		prev, curr, next := pw[i-1], pw[i], pw[i+1]
		if prev == curr && curr == next {
			occur[curr] = true
		}
	}
	return occur
}
