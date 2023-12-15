package main

import (
	"fmt"
	"math"
)

func main() {
	out, err := solution("../input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("solution:", out)
}

func solution(file string) (int64, error) {
	path, nodes, err := ForEachLine(file)
	if err != nil {
		return -1, err
	}
	starts := getStartNodes(nodes)
	m := countStepsWithMemo(path, nodes, starts)
	s := LCMSet(m)
	return s, nil
}

func getStartNodes(nodes map[string][]string) map[string]string {
	startNodes := map[string]string{}
	for n := range nodes {
		if n[2] == byte('A') {
			startNodes[n] = n
		}
	}
	return startNodes
}

func countStepsWithMemo(path []string, nodes map[string][]string, startNodes map[string]string) map[string]int {
	pathIdx := 0
	steps := 1

	memo := map[string]int{}

	for len(memo) < len(startNodes) {
		if pathIdx >= len(path) {
			pathIdx = 0
		}

		for n, currNode := range startNodes {
			if _, ok := memo[n]; ok {
				continue
			}
			var nextNode string
			if path[pathIdx] == "L" {
				nextNode = nodes[currNode][0]
			}
			if path[pathIdx] == "R" {
				nextNode = nodes[currNode][1]
			}
			startNodes[n] = nextNode
			if nextNode[2] == byte('Z') {
				memo[n] = steps
			}
		}
		pathIdx++
		steps++
	}
	return memo
}

func LCMSet(set map[string]int) int64 {
	var p int64
	for _, v := range set {
		if p == 0 {
			p = int64(v)
			continue
		}
		p = LCM(p, int64(v))
	}

	return p
}

func LCM(a, b int64) int64 {
	aFactors := primeFactoriaztion(a)
	bFactors := primeFactoriaztion(b)

	var lcm int64 = 1
	for p, aPow := range aFactors {
		bPow, ok := bFactors[p]
		if !ok {
			lcm *= int64(math.Pow(float64(p), float64(aPow)))
			continue
		}
		maxPow := aPow
		if bPow > aPow {
			maxPow = bPow
		}
		lcm *= int64(math.Pow(float64(p), float64(maxPow)))
	}
	for p, bPow := range bFactors {
		if _, ok := aFactors[p]; ok {
			// only process primes that haven't been seen already
			continue
		}
		lcm *= int64(math.Pow(float64(p), float64(bPow)))
	}
	return lcm
}

func primeFactoriaztion(x int64) map[int64]int64 {
	factors := map[int64]int64{}
	if x <= 1 {
		return factors
	}

	for _, prime := range primes {
		var pow int64
		r := x % prime
		for r == 0 {
			pow++
			x /= prime
			r = x % prime
		}
		if pow != 0 {
			factors[prime] = pow
		}
	}
	return factors
}

var primes = []int64{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89,
	97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191,
	193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293,
}
