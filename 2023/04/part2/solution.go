package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution(file string) (int, error) {
	cardToWinCount, err := ForEachLine(file, processLine)
	if err != nil {
		return -1, err
	}
	totalCards, err := totalCardCount(cardToWinCount)
	if err != nil {
		return -1, err
	}
	sumCards := 0
	for _, cardCount := range totalCards {
		sumCards += cardCount
	}
	return sumCards, nil
}

func ForEachLine(file string, fn func(string) (int, error)) (map[int]int, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fd)

	cardNum := 1
	cardToWinCount := map[int]int{}

	for scanner.Scan() {
		line := scanner.Text()
		x, err := fn(line)
		if err != nil {
			return nil, err
		}
		cardToWinCount[cardNum] = x
		cardNum++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cardToWinCount, nil
}

func totalCardCount(cardToWinCount map[int]int) (map[int]int, error) {
	totalCards := map[int]int{}
	for gameNum := 1; gameNum <= len(cardToWinCount); gameNum++ {
		totalCards[gameNum]++ // initialize the first card
		winCount := cardToWinCount[gameNum]
		nextGame := gameNum + 1
		for i := 0; i < winCount; i++ {
			for cardCount := totalCards[gameNum]; cardCount > 0; cardCount-- {
				totalCards[nextGame]++
			}
			nextGame++
		}
		fmt.Printf("gameNum=%d, winCount=%d, totalCards=%#v\n", gameNum, winCount, totalCards)
	}
	return totalCards, nil
}

// return the number of matches per card
func processLine(line string) (int, error) {
	nums := strings.Split(line, "|")
	winNums, gameNums := strings.Split(nums[0], " "), strings.Split(nums[1], " ")

	winNumsMap := map[string]bool{}
	for _, num := range winNums {
		winNumsMap[num] = true
	}

	var countMatchWinNums int = 0
	for _, n := range gameNums {
		if n == "" {
			continue
		}
		if _, ok := winNumsMap[n]; ok {
			countMatchWinNums++
		}
	}

	return countMatchWinNums, nil
}
