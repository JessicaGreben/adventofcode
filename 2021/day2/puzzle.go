package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ref: https://adventofcode.com/2021/day/2
func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var depth, horiz, aim int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		split := strings.Split(in, " ")
		if len(split) != 2 {
			log.Fatal("should be len 2, got", len(split))
		}
		cmd := split[0]
		amount, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cmd, amount)

		switch cmd {
		case "up":
			// depth -= amount
			aim -= amount
		case "forward":
			depth += aim * amount
			horiz += amount
		case "down":
			// depth += amount
			aim += amount
		default:
			log.Println("not valid command, got", cmd)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("final depth:", depth, "final horizontal position:", horiz)
	fmt.Println(depth * horiz)
}
