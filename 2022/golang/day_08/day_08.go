package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// parse_input("1.txt")
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
	var l string
	var tree_arr2d [][]int
	var tree_arr1d []int
	for _, line = range strings.Split(ip_str, "\n") {
		tree_arr1d = nil
		for _, l = range strings.Split(line, "") {
			var i_l int
			i_l, _ = strconv.Atoi(l)
			tree_arr1d = append(tree_arr1d, i_l)
		}
		tree_arr2d = append(tree_arr2d, tree_arr1d)
	}

	// print_arr(tree_arr2d)
	var count int = p1(tree_arr2d)
	fmt.Printf("count: %v\n", count)

	var scenic_score int = p2(tree_arr2d)
	fmt.Printf("scenic score: %v\n", scenic_score)
}

func p1(arr2d [][]int) (count int) {
	count = 0
	for i, arr1d := range arr2d {
		for j, _ := range arr1d {
			// create variable to check if visible is found on either left/right/top/bottom
			l_f := 0
			r_f := 0
			t_f := 0
			b_f := 0

			// outermost layer
			if i == 0 || i == len(arr1d)-1 || j == 0 || j == len(arr1d)-1 {
				count += 1
			} else {
				// interior layer
				// if any side there is having height is more, then make the side to 1
				// left
				for i1 := 0; i1 < j; i1++ {
					if !(arr1d[j] > arr1d[i1]) {
						l_f = 1
						break
					}

				}
				// right
				if l_f == 1 {
					for i1 := j + 1; i1 < len(arr1d); i1++ {
						if !(arr1d[j] > arr1d[i1]) {
							r_f = 1
							break
						}
					}
				}
				// top
				if l_f == 1 && r_f == 1 {
					for i1 := 0; i1 < i; i1++ {
						if !(arr2d[i][j] > arr2d[i1][j]) {
							t_f = 1
							break
						}
					}
				}

				// bottom
				if l_f == 1 && r_f == 1 && t_f == 1 {
					for i1 := i + 1; i1 < len(arr1d); i1++ {
						if !(arr2d[i][j] > arr2d[i1][j]) {
							b_f = 1
							break
						}
					}
				}
				// if any side is visible, then increase visible count
				if t_f == 0 || b_f == 0 || r_f == 0 || l_f == 0 {
					count += 1
				}
			}
		}
	}

	return count
}

func p2(arr2d [][]int) (scenic_score int) {
	scenic_score = 0
	for i, arr1d := range arr2d {
		for j, _ := range arr1d {
			// create variable to check if visible is found on either left/right/top/bottom
			l_f := 0
			r_f := 0
			t_f := 0
			b_f := 0

			if !(i == 0 || i == len(arr1d)-1 || j == 0 || j == len(arr1d)-1) {
				// interior layer
				// if any side there is having height is more, then make the side to 1
				// left
				for i1 := 0; i1 < j; i1++ {
					if !(arr1d[j] > arr1d[i1]) {
						l_f = 1
						break
					}

				}
				// right
				if l_f == 1 {
					for i1 := j + 1; i1 < len(arr1d); i1++ {
						if !(arr1d[j] > arr1d[i1]) {
							r_f = 1
							break
						}
					}
				}
				// top
				if l_f == 1 && r_f == 1 {
					for i1 := 0; i1 < i; i1++ {
						if !(arr2d[i][j] > arr2d[i1][j]) {
							t_f = 1
							break
						}
					}
				}

				// bottom
				if l_f == 1 && r_f == 1 && t_f == 1 {
					for i1 := i + 1; i1 < len(arr1d); i1++ {
						if !(arr2d[i][j] > arr2d[i1][j]) {
							b_f = 1
							break
						}
					}
				}
				lc := 0
				rc := 0
				tc := 0
				bc := 0

				// check scenic_score for the visible trees alone
				if t_f == 0 || b_f == 0 || r_f == 0 || l_f == 0 {
					// check if height is less and add count
					// if height is same or more, add count and break
					// left check
					for i1 := j - 1; i1 >= 0; i1-- {
						if arr1d[j] > arr1d[i1] {
							lc += 1
						} else if arr1d[j] <= arr1d[i1] {
							lc += 1
							break
						} else {
							break
						}
					}

					// right check
					for i1 := j + 1; i1 < len(arr1d); i1++ {
						if arr1d[j] > arr1d[i1] {
							rc += 1
						} else if arr1d[j] <= arr1d[i1] {
							rc += 1
							break
						} else {
							break
						}
					}

					// top check
					for i1 := i - 1; i1 >= 0; i1-- {
						if arr2d[i][j] > arr2d[i1][j] {
							tc += 1
						} else if arr2d[i][j] <= arr2d[i1][j] {
							tc += 1
							break
						} else {
							break
						}
					}

					// bottom check
					for i1 := i + 1; i1 < len(arr1d); i1++ {
						if arr2d[i][j] > arr2d[i1][j] {
							bc += 1
						} else if arr2d[i][j] <= arr2d[i1][j] {
							bc += 1
							break
						} else {
							break
						}
					}

				}

				if (lc * rc * tc * bc) > scenic_score {
					scenic_score = (lc * rc * tc * bc)
				}
			}

		}
	}

	return scenic_score
}

func print_arr(arr2d [][]int) {
	for _, i := range arr2d {
		for _, j := range i {
			fmt.Printf("%v ", j)
		}
		fmt.Printf("\n")
	}
}
