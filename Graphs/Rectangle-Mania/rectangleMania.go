// Given an array of coordinates, find out the total amount of rectangles could be formed
// limitation: all sides of rectangle are horizontal/perpendicular to cartesian axis

package rectanglemania

// Solution 1: iterate thru all coordinates, bucket up/right/down/left coords for each of them
// each iteration treat the current coordinate as bottom left corner of potential rectangle
// Big O: O(n^2) time | O(n^2) space (worst case scenario: all coordinates on the same horizontal/vertical line)
// as for recursive calls on the callstack, there's at most 4 calls at any given time - clockwise UP -> RIGHT -> DOWN -> LEFT
// func RectangleMania(coords [][2]int) int {
// 	coordsTable := getCoordsTable(coords)
// 	return getRectangleCount(coords, coordsTable)
// }

// func getCoordsTable(coords [][2]int) map[string]map[string][][2]int {
// 	coordsTable := make(map[string]map[string][][2]int)
// 	for _, coord1 := range coords {
// 		coord1Directions := map[string][][2]int{
// 			"UP":    make([][2]int, 0),
// 			"RIGHT": make([][2]int, 0),
// 			"DOWN":  make([][2]int, 0),
// 			"LEFT":  make([][2]int, 0),
// 		}
// 		for _, coord2 := range coords {
// 			coord2Direction := getCoordDirection(coord1, coord2)
// 			if _, ok := coord1Directions[coord2Direction]; ok {
// 				coord1Directions[coord2Direction] = append(coord1Directions[coord2Direction], coord2)
// 			}
// 		}
// 		coord1String := coordToString(coord1)
// 		coordsTable[coord1String] = coord1Directions
// 	}
// 	return coordsTable
// }

// func getCoordDirection(coord1, coord2 [2]int) string {
// 	x1, y1 := coord1[0], coord1[1]
// 	x2, y2 := coord2[0], coord2[1]
// 	if y1 == y2 {
// 		if x2 > x1 {
// 			return "RIGHT"
// 		} else if x2 < x1 {
// 			return "LEFT"
// 		}
// 	} else if x1 == x2 {
// 		if y2 > y1 {
// 			return "UP"
// 		} else if y2 < y1 {
// 			return "DOWN"
// 		}
// 	}
// 	return ""
// }

// func getRectangleCount(coords [][2]int, coordsTable map[string]map[string][][2]int) int {
// 	rectangleCount := 0
// 	for _, coord := range coords {
// 		rectangleCount += clockwiseCountRectangles(coord, coordsTable, "UP", coord)
// 	}
// 	return rectangleCount
// }

// func clockwiseCountRectangles(coord [2]int, coordsTable map[string]map[string][][2]int, direction string, origin [2]int) int {
// 	coordString := coordToString(coord)
// 	rectangleCount := 0
// 	if direction == "LEFT" {
// 		for _, coordToCheck := range coordsTable[coordString]["LEFT"] {
// 			if origin == coordToCheck {
// 				return 1
// 			}
// 		}
// 		return 0
// 	} else {
// 		nextDirection := getNextClockwiseDirection(direction)
// 		for _, nextCoord := range coordsTable[coordString][direction] {
// 			rectangleCount += clockwiseCountRectangles(nextCoord, coordsTable, nextDirection, origin)
// 		}
// 	}
// 	return rectangleCount
// }

// func getNextClockwiseDirection(direction string) string {
// 	if direction == "UP" {
// 		return "RIGHT"
// 	}
// 	if direction == "RIGHT" {
// 		return "DOWN"
// 	}
// 	if direction == "DOWN" {
// 		return "LEFT"
// 	}
// 	return ""
// }

// func coordToString(coord [2]int) string {
// 	x, y := coord[0], coord[1]
// 	return fmt.Sprintf("%d-%d", x, y)
// }

// ======================================================================
// Solution 2: a bit optimization on space complexity than solution 1
// Big O: O(n^2) time | O(n) space
// func RectangleMania(coords [][2]int) int {
// 	coordsTable := getCoordsTable(coords)
// 	return getRectangleCount(coords, coordsTable)
// }

// // O(n) time | O(2 * n) space
// func getCoordsTable(coords [][2]int) map[string]map[int][][2]int {
// 	coordsTable := map[string]map[int][][2]int{
// 		"x": make(map[int][][2]int),
// 		"y": make(map[int][][2]int),
// 	}
// 	for _, coord := range coords {
// 		x, y := coord[0], coord[1]
// 		if _, exists := coordsTable["x"][x]; !exists {
// 			coordsTable["x"][x] = [][2]int{}
// 		}
// 		coordsTable["x"][x] = append(coordsTable["x"][x], coord)
// 		if _, exists := coordsTable["y"][y]; !exists {
// 			coordsTable["y"][y] = [][2]int{}
// 		}
// 		coordsTable["y"][y] = append(coordsTable["y"][y], coord)
// 	}
// 	return coordsTable
// }

// func getRectangleCount(coords [][2]int, coordsTable map[string]map[int][][2]int) int {
// 	rectangleCount := 0

// 	for _, coord := range coords {
// 		lowerLeftY := coord[1]
// 		rectangleCount += clockwiseCountRectangles(coord, coordsTable, "UP", lowerLeftY)
// 	}
// 	return rectangleCount
// }

// func clockwiseCountRectangles(coord1 [2]int, coordsTable map[string]map[int][][2]int, direction string, lowerLeftY int) int {
// 	x1, y1 := coord1[0], coord1[1]
// 	rectangleCount := 0
// 	if direction == "DOWN" {
// 		relavantCords := coordsTable["x"][x1]
// 		for _, coord2 := range relavantCords {
// 			lowerRightY := coord2[1]
// 			if lowerRightY == lowerLeftY {
// 				return 1
// 			}
// 		}
// 		return 0
// 	} else {
// 		if direction == "UP" {
// 			relavantCords := coordsTable["x"][x1]
// 			for _, coord2 := range relavantCords {
// 				y2 := coord2[1]
// 				if isAbove := y2 > y1; isAbove {
// 					rectangleCount += clockwiseCountRectangles(coord2, coordsTable, "RIGHT", lowerLeftY)
// 				}
// 			}
// 		} else if direction == "RIGHT" {
// 			relavantCords := coordsTable["y"][y1]
// 			for _, coord2 := range relavantCords {
// 				x2 := coord2[0]
// 				if isRight := x2 > x1; isRight {
// 					rectangleCount += clockwiseCountRectangles(coord2, coordsTable, "DOWN", lowerLeftY)
// 				}
// 			}
// 		}
// 		return rectangleCount
// 	}
// }

// ======================================================================
// Solution 3: cleanest one, storing all coordinates in a set, double for loop where we have
// coordinate1 and coordinate2, if coord1 and coord2 are bottom-left / top-right corner of
// a potential rectangle, then we can proceed to find the upper-left & bottom-right coordinates
// if both corners are found, total count += 1
// Big O: O(n^2) time | O(n) space
func RectangleMania(coords [][2]int) int {
	coordsTable := getCoordsTable(coords)
	return getRectangleCount(coords, coordsTable)
}

func getCoordsTable(coords [][2]int) map[[2]int]struct{} {
	coordsTable := make(map[[2]int]struct{})
	var empty struct{}
	for _, coord := range coords {
		coordsTable[coord] = empty
	}
	return coordsTable
}

func getRectangleCount(coords [][2]int, coordsTable map[[2]int]struct{}) int {
	rectangleCount := 0
	for _, coord1 := range coords {
		x1, y1 := coord1[0], coord1[1]
		for _, coord2 := range coords {
			x2, y2 := coord2[0], coord2[1]
			if !isInUpperRight(coord1, coord2) {
				continue
			}
			_, upperLeftExists := coordsTable[[2]int{x1, y2}]
			_, bottomRightExists := coordsTable[[2]int{x2, y1}]
			if upperLeftExists && bottomRightExists {
				rectangleCount += 1
			}
		}
	}
	return rectangleCount
}

func isInUpperRight(coord1, coord2 [2]int) bool {
	x1, y1 := coord1[0], coord1[1]
	x2, y2 := coord2[0], coord2[1]
	return x2 > x1 && y2 > y1
}
