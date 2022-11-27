package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// https://adventofcode.com/2021/day/1
func main() {
	// printRolling3()
	printEach()
}

func printEach() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count, sum, prev, result int
	three := make([]int, 0, 3)
	for scanner.Scan() {
		count++
		num := scanner.Text()
		curr, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(curr)

		three = append(three, curr)
		if len(three) >= 3 {
			sum = three[0] + three[1] + three[2]
			if sum > prev && count >= 4 {
				result++
			}
			fmt.Println(three, sum)
			three = three[1:]
			prev = sum
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("res:", result)
}

func printRolling3() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var curr3 = make([]int, 0, 3)
	var prevSum, sum, result int
	var count int
	for scanner.Scan() {
		count++
		num := scanner.Text()
		curr, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		if len(curr3) < 3 {
			curr3 = append(curr3, curr)
		}
		if len(curr3) == 3 && count >= 3 {
			fmt.Println("curr3:", curr3)
			sum = curr3[0] + curr3[1] + curr3[2]
			if sum > prevSum {
				result++
			}
			prevSum = sum
			curr3 = curr3[1:]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", result)
}

func do() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var prev, result, count, currSum int
	begin := true

	for scanner.Scan() {
		count++
		num := scanner.Text()
		curr, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		if count <= 3 {
			currSum += curr
		}
		if currSum > prev && !begin {
			result++
		}
		if count == 3 {
			prev = currSum
			begin = false
			count = 0
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", result)
}
