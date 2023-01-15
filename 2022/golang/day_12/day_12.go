// 2023-01-04 14:07:22.457865339 +0530

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "1" {
		parseInput("1.txt")
	} else {
		parseInput("input.txt")
	}
}

func parseInput(fileName string) {
	var data []byte
	var err error
	data, err = os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var ip_str string = string(data)

	heightMap := make([][]rune, 0)
	var start, end point

	for j, line := range strings.Split(ip_str, "\n") {
		var hM1d []rune
		for i, l := range strings.Split(line, "") {
			elevation := []rune(l)[0]
			if elevation == 'S' {
				elevation = 'a'
				// get the starting point
				start = point{i, j}
			} else if elevation == 'E' {
				elevation = 'z'
				// get the ending point
				end = point{i, j}
			}
			hM1d = append(hM1d, elevation)
		}
		heightMap = append(heightMap, hM1d)
		hM1d = nil
	}

	fmt.Printf("minimum distance: %v\n", p1(heightMap, start, end))

	var a_arr []point
	for i, arr1d := range heightMap {
		// fmt.Printf("arr1d: %v\n", arr1d)
		for j, elev := range arr1d {
			if elev == 'a' {
				// fmt.Printf("elev: %v | %v | %v\n", elev, j, i)
				a_arr = append(a_arr, point{j, i})
			}
		}
	}

	// fmt.Printf("len: %v\n", len(a_arr))
	// fmt.Printf("arr: %v\n", a_arr)
	var shortArr []int
	for _, arr := range a_arr {
		dis := p1(heightMap, arr, end)
		if dis != 0 {
			shortArr = append(shortArr, dis)

		}
	}

	sort.Ints(shortArr)
	fmt.Printf("shortest of 'a': %v\n", shortArr[0])
}

func p1(heightmap [][]rune, start point, end point) (dis int) {

	visited := make(map[point]bool)
	toVisit := []point{start}
	distanceFromStart := map[point]int{start: 0}
	for {
		// if toVisit is nil; break
		if len(toVisit) == 0 {
			break
		}
		// assign currentPoint to toVisit[0]
		currentPoint := toVisit[0]
		// assign visited map to true for currentPoint
		visited[currentPoint] = true

		// remove currentPoint from toVisit
		toVisit = toVisit[1:]

		// if currentPoint is end (point{}) break with distance covered
		if currentPoint == end {
			dis = distanceFromStart[end]
			break
		}

		// loop through left, right, up, down
		for _, near := range [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
			j, i := near[1], near[0]
			// take nextPoint
			nextPoint := point{currentPoint.x + j, currentPoint.y + i}

			// if not visited already && nextPoint is within limits && diff bt current and next is one
			if !visited[nextPoint] && nextPoint.x >= 0 && nextPoint.y >= 0 &&
				nextPoint.x < len(heightmap[0]) && nextPoint.y < len(heightmap) &&
				(heightmap[nextPoint.y][nextPoint.x]-heightmap[currentPoint.y][currentPoint.x] <= 1) {
				// if nextPoint doesn't have distance append to toVisit and increment the distance
				if distanceFromStart[nextPoint] == 0 {
					toVisit = append(toVisit, nextPoint)
					distanceFromStart[nextPoint] = distanceFromStart[currentPoint] + 1
				}

				// if distance is more than one, increment by 1
				if distanceFromStart[nextPoint] >= distanceFromStart[currentPoint]+1 {
					distanceFromStart[nextPoint] = distanceFromStart[currentPoint] + 1
				}
			}
		}

		// after travelling to heightMap, sort the toVisit array
		sort.Slice(toVisit, func(i, j int) bool {
			return distanceFromStart[toVisit[i]] < distanceFromStart[toVisit[j]]
		})

	}
	// return the distance travelled by the start point till end point
	return dis
}
