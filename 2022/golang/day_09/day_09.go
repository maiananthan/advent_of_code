package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {

	// parse_input("1.txt")
	// parse_input("2.txt")
	parse_input("input.txt")
}

func parse_input(file_name string) {
	var data []byte
	var err error

	data, err = os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	var ip_str string = string(data)
	var line string
	var l []string
	var route [][]string

	for _, line = range strings.Split(ip_str, "\n") {
		l = strings.Split(line, " ")
		route = append(route, l)
	}

	var position_visited int = p1(route)
	fmt.Printf("position visited (tail): %v\n", position_visited)

	var positionVisited int = p2(route)
	fmt.Printf("positionVisited (9): %v\n", positionVisited)

}

var moveCount int = 0

func p1(route [][]string) (pos_visited int) {
	return calcTravel(route, 2)
}

func p2(route [][]string) (posVisited int) {
	return calcTravel(route, 10)
}

func calcTravel(route [][]string, knots_og int) (disTravel int) {
	var posArr []point
	var posChange map[point]bool
	posChange = make(map[point]bool)
	var knots int = knots_og - 1
	// create position array
	for k := 0; k <= knots; k++ {
		posArr = append(posArr, point{0, 0})
	}
	// initial position
	posChange[posArr[knots]] = true

	// go through route one by one
	for _, direction := range route {
		var direction_count int
		direction_count, _ = strconv.Atoi(direction[1])

		// loop through number of direction moves
		for d1 := 0; d1 < direction_count; d1++ {

			// based on direction increment/decrement [0] position
			switch string(direction[0]) {
			case "L":
				posArr[0].x--
			case "R":
				posArr[0].x++
			case "U":
				posArr[0].y++
			case "D":
				posArr[0].y--

			}

			// adjust the following knots present based on previous knot
			for d2 := 0; d2 < knots; d2++ {
				posArr[d2+1] = adjCheck(posArr[d2+1], posArr[d2])
			}

			// map the position of last knot to true
			posChange[posArr[knots]] = true
		}
	}
	return len(posChange)

}

func adjCheck(tail point, head point) (newTail point) {
	newTail = tail

	// based on difference between head and tail, move tail
	switch (point{head.x - tail.x, head.y - tail.y}) {
	case point{-2, 1}, point{-1, 2}, point{0, 2}, point{1, 2}, point{2, 1}, point{2, 2}, point{-2, 2}:
		newTail.y++
	}
	switch (point{head.x - tail.x, head.y - tail.y}) {
	case point{1, 2}, point{2, 1}, point{2, 0}, point{2, -1}, point{1, -2}, point{2, 2}, point{2, -2}:
		newTail.x++
	}
	switch (point{head.x - tail.x, head.y - tail.y}) {
	case point{-2, -2}, point{2, -1}, point{1, -2}, point{0, -2}, point{-1, -2}, point{-2, -1}, point{2, -2}:
		newTail.y--
	}
	switch (point{head.x - tail.x, head.y - tail.y}) {
	case point{-2, -2}, point{-1, -2}, point{-2, -1}, point{-2, -0}, point{-2, 1}, point{-1, 2}, point{-2, 2}:
		newTail.x--
	}

	return newTail
}
