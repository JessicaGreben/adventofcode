package day1

func Part1(moduleMasses []int) int {
	return moduleFuelRequirements(moduleMasses)
}

func moduleFuelRequirements(moduleMasses []int) int {
	fuelRequired := 0
	for _, mass := range moduleMasses {
		fuelRequired += calculateModuleFuel(mass)
	}
	return fuelRequired
}

func Part2(moduleMasses []int) int {
	return totalFuelRequirements(moduleMasses)
}

func totalFuelRequirements(moduleMasses []int) int {
	fuelRequired := 0
	for _, mass := range moduleMasses {
		moduleFuel := calculateModuleFuel(mass)
		fuelRequired += moduleFuel
		fuelRequired += calculateFuelFuel(moduleFuel)

	}
	return fuelRequired
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
