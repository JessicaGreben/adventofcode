package day1

import (
	"strconv"
)

func Part1(moduleMasses []string) (int, error) {
	totalFuel, err := moduleFuelRequirements(moduleMasses)
	if err != nil {
		return -1, err
	}
	return totalFuel, nil
}

func moduleFuelRequirements(moduleMasses []string) (int, error) {
	fuelRequired := 0
	for _, mass := range moduleMasses {
		mass, err := strconv.Atoi(mass)
		if err != nil {
			return -1, err
		}
		fuelRequired += calculateModuleFuel(mass)
	}
	return fuelRequired, nil
}

func Part2(moduleMasses []string) (int, error) {
	totalFuel, err := totalFuelRequirements(moduleMasses)
	if err != nil {
		return -1, err
	}
	return totalFuel, nil
}

func totalFuelRequirements(moduleMasses []string) (int, error) {
	fuelRequired := 0
	for _, mass := range moduleMasses {
		mass, err := strconv.Atoi(mass)
		if err != nil {
			return -1, err
		}
		moduleFuel := calculateModuleFuel(mass)
		fuelRequired += moduleFuel
		fuelRequired += calculateFuelFuel(moduleFuel)

	}
	return fuelRequired, nil
}

func calculateModuleFuel(mass int) int {
	return (mass / 3) - 2
}

func calculateFuelFuel(fuel int) int {
	var totalFuel int
	currFuel := calculateModuleFuel(fuel)
	for currFuel > 0 {
		totalFuel += currFuel
		currFuel = calculateModuleFuel(currFuel)
	}
	return totalFuel
}
