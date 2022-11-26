package day2

import (
	"errors"
	"fmt"

	"golang.org/x/exp/slices"

	"github.com/jessicagreben/adventofcode/2019/pkg/intcode"
)

var (
	errInvalidLen = errors.New("invalid length")
	errNounRange  = errors.New("noun out of [0,99] range")
	errVerbRange  = errors.New("verb out of [0,99] range")
)

func Part1(input []int) (int, error) {
	const (
		noun = 12
		verb = 2
	)
	if err := restore(noun, verb, input); err != nil {
		return -1, err
	}
	p := intcode.NewProgram(input)
	if err := p.Run(-1); err != nil {
		return -1, err
	}
	return p.Instructions()[0], nil
}

func restore(noun, verb int, input []int) error {
	const minInputLen = 3
	if len(input) < minInputLen {
		return fmt.Errorf("%w: minimum length is %d, got %d", errInvalidLen, minInputLen, len(input))
	}
	if noun < 0 || noun > 99 {
		return fmt.Errorf("%w: %d", errNounRange, noun)
	}
	if verb < 0 || verb > 99 {
		return fmt.Errorf("%w: %d", errVerbRange, verb)
	}
	input[1] = noun
	input[2] = verb
	return nil
}

func Part2(input []int) (int, int, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			cp := slices.Clone(input)
			restore(noun, verb, cp)
			p := intcode.NewProgram(cp)
			if err := p.Run(-1); err != nil {
				return -1, -1, err
			}
			if p.Instructions()[0] == 19690720 {
				return noun, verb, nil
			}
		}
	}
	return -1, -1, nil
}
