package day1

import (
	"bufio"
	"os"
	"strconv"
)

func mostCalories() (int, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return -1, err
	}
	scanner := bufio.NewScanner(fd)
	maxCalories := -1
	currCalories := 0
	for scanner.Scan() {
		cals := scanner.Text()
		if cals == "" {
			maxCalories = max(maxCalories, currCalories)
			currCalories = 0
			continue
		}
		calsInt, err := strconv.Atoi(cals)
		if err != nil {
			return -1, err
		}
		currCalories += calsInt
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return maxCalories, nil
}

func top3Calories() (int, error) {
	fd, err := os.Open("input.txt")
	if err != nil {
		return -1, err
	}
	scanner := bufio.NewScanner(fd)
	top3 := make([]int, 3)
	currCalories := 0
	for scanner.Scan() {
		cals := scanner.Text()
		if cals == "" {
			top3 = insert(top3, currCalories)
			currCalories = 0
			continue
		}
		calsInt, err := strconv.Atoi(cals)
		if err != nil {
			return -1, err
		}
		currCalories += calsInt
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return top3[0] + top3[1] + top3[2], nil
}

// insert replaces the smallest value in arr with the newValue if
// its larger than the smallest value
func insert(arr []int, newValue int) []int {
	minVal, minIdx, maxVal := arr[0], 0, arr[0]
	for i, currVal := range arr {
		if currVal < minVal {
			minVal, minIdx = currVal, i
		}
		if currVal > maxVal {
			maxVal = currVal
		}
	}
	if newValue <= minVal {
		return arr
	}
	if newValue > minVal {
		arr[minIdx] = newValue
	}
	return arr
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
