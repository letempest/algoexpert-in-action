package greedy

// Solution 1: Brute Force approach
// Big O: O(n^2) time | O(1) space
// func ValidStartingCity(distances, fuel []int, mpg int) int {
// 	numberOfCities := len(distances)

// 	for startCityIdx := 0; startCityIdx < numberOfCities; startCityIdx++ {
// 		var milesRemaining int
// 		for currentCityUnwrappedIdx := startCityIdx; currentCityUnwrappedIdx < startCityIdx+numberOfCities; currentCityUnwrappedIdx++ {
// 			if milesRemaining < 0 {
// 				continue
// 			}
// 			currentCityIdx := currentCityUnwrappedIdx % numberOfCities

// 			fuelFromCurrentCity := fuel[currentCityIdx]
// 			distanceToNextCity := distances[currentCityIdx]
// 			milesRemaining += fuelFromCurrentCity*mpg - distanceToNextCity
// 		}
// 		if milesRemaining >= 0 {
// 			return startCityIdx
// 		}
// 	}

// 	return -1
// }

// Solution 2
// prerequsite: unique solution, only one valid starting city
// hence the valid starting city is the one we enter with fewest fuel that can
// travel the LEAST distance afterwards
// Big O: O(n) time | O(1) space
func ValidStartingCity(distances, fuel []int, mpg int) int {
	numberOfCities := len(distances)
	var milesRemaining int

	var indexOfStartingCityCandidate int
	var milesRemainingAtStartingCityCandidate int

	for cityIdx := 1; cityIdx < numberOfCities; cityIdx++ {
		distanceFromPreviousCity := distances[cityIdx-1]
		fuelFromPreviousCity := fuel[cityIdx-1]
		milesRemaining += fuelFromPreviousCity*mpg - distanceFromPreviousCity

		if milesRemaining < milesRemainingAtStartingCityCandidate {
			milesRemainingAtStartingCityCandidate = milesRemaining
			indexOfStartingCityCandidate = cityIdx
		}
	}

	return indexOfStartingCityCandidate
}
