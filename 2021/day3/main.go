package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ref: https://adventofcode.com/2021/day/3
func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count0 := make([]int, 12)
	count1 := make([]int, 12)
	var total int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		fmt.Println(in)
		total++
		for i, x := range in {
			switch string(x) {
			case "0":
				count0[i]++
			case "1":
				count1[i]++
			default:
				log.Fatalf("not valid, expected 0 or 1, got:", x)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	gam, eps := gamma(count0, count1)
	y := strings.Join(gam, "")
	gamma, err := strconv.ParseInt(y, 2, 64)
	if err != nil {
		fmt.Println("err:", err)
	}
	epsilon, err := strconv.ParseInt(strings.Join(eps, ""), 2, 64)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println(gamma * epsilon)
}

func gamma(zeros, ones []int) ([]string, []string) {
	var i int
	var countOnes, countZeros int
	var gam = make([]string, 0, len(ones))
	var eps = make([]string, 0, len(ones))
	for i < len(zeros) && i < len(ones) {
		countOnes = ones[i]
		countZeros = zeros[i]
		if countOnes > countZeros {
			gam = append(gam, "1")
			eps = append(eps, "0")
		} else if countZeros > countOnes {
			gam = append(gam, "0")
			eps = append(eps, "1")
		} else {
			fmt.Println("equal", i)
		}
		i++
	}
	return gam, eps
}

func decimalToBinary2(dec int) []int {
	result := make([]int, 32)
	var k int = dec
	for i := 31; i >= 0; i-- {
		if k%2 != 0 {
			result[i] = 1
		} else {
			result[i] = 0
		}
		k = k >> 1
	}
	return result
}

func decimalToBinary3(dec int) []int {
	result := make([]int, 32)
	var k int = dec
	for i := 31; i >= 0; i-- {
		if k%2 != 0 {
			result[i] = 1
		} else {
			result[i] = 0
		}
		k /= 2
	}
	return result
}
