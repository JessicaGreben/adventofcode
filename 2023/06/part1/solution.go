package main

type race struct {
	raceDurationMs   int // in milliseconds
	recordDistanceMm int // in millimeters
}

func solution(races []race) int {
	product := 1
	for _, race := range races {
		product *= waysToWin(race)
	}

	return product
}

func waysToWin(r race) int {
	result := 0
	for hold := 0; hold <= r.raceDurationMs; hold++ {
		distance := (r.raceDurationMs - hold) * hold
		if distance > r.recordDistanceMm {
			result++
		}
	}
	return result
}
