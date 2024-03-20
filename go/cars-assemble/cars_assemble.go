package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	perHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(perHour / float64(60))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// cost of 1 is 10,000
	// cost of a group of 10 is 95,000
	if carsCount >= 10 {
		groupsOfTen := carsCount / 10
		remaining := carsCount % 10
		return uint((groupsOfTen * 95000) + (remaining * 10000))
	}
	return uint(carsCount * 10000)
}
