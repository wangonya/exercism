package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, count := range birdsPerDay {
		sum += count
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsInWeek := birdsPerDay[(week-1)*7 : ((week-1)*7)+7]
	return TotalBirdCount(birdsInWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, count := range birdsPerDay {
		if i%2 == 0 || i == 0 {
			birdsPerDay[i] = count + 1
		}
	}
	return birdsPerDay
}
