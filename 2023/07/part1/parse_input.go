package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ForEachLine(file string, fn func(string) (string, int64, error)) ([]*hand, error) {
	hands := []*hand{}

	fd, err := os.Open(file)
	if err != nil {
		return hands, err
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		hand, bid, err := fn(line)
		if err != nil {
			return hands, err
		}
		hands = append(hands, newHand(hand, bid))
	}
	if err := scanner.Err(); err != nil {
		return hands, err
	}
	return hands, nil
}

func processLine(line string) (hand string, bid int64, err error) {
	out := strings.Split(line, " ")
	if len(out) != 2 {
		return hand, bid, fmt.Errorf("expected len=2, got=%d", len(out))
	}
	hand = out[0]
	b, err := strconv.Atoi(out[1])
	if err != nil {
		return hand, bid, err
	}
	return hand, int64(b), nil
}
