package main

type race struct {
	raceDurationMs   int64 // in milliseconds
	recordDistanceMm int64 // in millimeters
}

func solution(races []race) int64 {
	var product int64 = 1
	for _, race := range races {
		product *= waysToWin(race)
	}

	return product
}

func waysToWin(r race) int64 {
	var result int64
	for hold := int64(0); hold <= r.raceDurationMs; hold++ {
		distance := (r.raceDurationMs - hold) * hold
		if distance > r.recordDistanceMm {
			result++
		}
	}
	return result
}
