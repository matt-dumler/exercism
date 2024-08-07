package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var count int
    for _, n := range birdsPerDay {
        count += n
    }

    return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    startOfWeek := (week - 1) * 7
    endOfWeek := startOfWeek + 7
    birdsInWeek := birdsPerDay[startOfWeek:endOfWeek]

    return TotalBirdCount(birdsInWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for index := range birdsPerDay {
        if index % 2 == 0 {
            birdsPerDay[index]++
        }
    }

    return birdsPerDay
}
