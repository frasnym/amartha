package main

type Coordinate struct {
	Y int
	X int
}

type Input [][]int

func isPathAvailable(input Input) bool {
	targetPosition := Coordinate{
		Y: len(input) - 1,
		X: len(input[0]) - 1,
	}

	right := Coordinate{Y: 0, X: 1}
	left := Coordinate{Y: 0, X: -1}
	down := Coordinate{Y: 1, X: 0}
	up := Coordinate{Y: -1, X: 0}
	directionToCheck := []Coordinate{right, down, left, up}

	startPosition := Coordinate{0, 0}
	pathToCheck := []Coordinate{startPosition}

	for len(pathToCheck) > 0 {
		pathPosition := pathToCheck[len(pathToCheck)-1]

		// Remove checked path from path
		pathToCheck = pathToCheck[:len(pathToCheck)-1]

		// Set to visited, don't check again
		input[pathPosition.Y][pathPosition.X] = -1

		for _, direction := range directionToCheck {
			// fmt.Printf("direction: %+v\n", direction)
			pathPlusDirection := Coordinate{
				Y: pathPosition.Y + direction.Y,
				X: pathPosition.X + direction.X,
			}
			// fmt.Printf("pathPlusDirection: %+v\n", pathPlusDirection)

			isReachTarget := pathPlusDirection.X == targetPosition.X && pathPlusDirection.Y == targetPosition.Y
			if isReachTarget {
				return true
			}

			isLessThanMin := pathPlusDirection.X < 0 || pathPlusDirection.Y < 0
			// fmt.Println("isLessThanMin", isLessThanMin)
			isMoreThanMax := pathPlusDirection.X > targetPosition.X || pathPlusDirection.Y > targetPosition.Y
			// fmt.Println("isMoreThanMax", isMoreThanMax)
			isOutOfBound := isLessThanMin || isMoreThanMax
			// fmt.Println("isOutOfBound", isOutOfBound)
			if !isOutOfBound {
				isGoodPath := input[pathPlusDirection.Y][pathPlusDirection.X] == 0
				// fmt.Println("isGoodPath", isGoodPath)
				if isGoodPath {
					pathToCheck = append(pathToCheck, pathPlusDirection)
				}
			}
			// fmt.Println("===")
		}
	}

	return false
}
