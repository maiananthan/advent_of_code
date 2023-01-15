package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// read the file into variable
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println("data" , data);
	var input_str string
	// fmt.Println("data len: ", len(data))
	// convert byte data into string
	input_str = string(data[:])
	// fmt.Println("input_str '" , input_str, "'")

	// fmt.Println("input_str:", strings.TrimSpace(input_str))

	var arr2d [][]int
	for _, group := range strings.Split(input_str, "\n\n") {
		// fmt.Println("group: ", group)
		row_arr := []int{}
		for _, line := range strings.Split(group, "\n") {
			// fmt.Println("line: ", line, "len:", len(line))
			if len(line) != 0 {
				row_arr = append(row_arr, toInt(line))
			}

		}
		arr2d = append(arr2d, row_arr)
	}

	var totals []int
	var items []int
	for _, items = range arr2d {
		// fmt.Println("items: ", items)
		var tot int
		var item int
		for _, item = range items {
			tot = tot + item
		}
		totals = append(totals, tot)
	}
	// fmt.Println("totals:", totals)
	// log.Println("totals:", totals)
	fmt.Println("max calories", maxint(totals...))

	// sort.Ints(totals)

	// fmt.Println("totals:", totals)
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	// fmt.Println("totals:", totals)
	// fmt.Println("totals:", totals)
	var topThree int
	var i int
	for _, i = range totals[0:3] {
		// fmt.Println("i", i)
		topThree += i
	}
	log.Println("topThree", topThree)
	// totals := []int{}
	/*
	   for _, items := range ans {
	       totals = append(totals, )
	   }
	*/
}

func maxint(nums ...int) int {
	var max int
	var i int
	max = nums[0]
	for _, i = range nums {
		if i > max {
			max = i
		}
	}

	return max
}
func toInt(arg interface{}) int {
	var val int
	switch arg.(type) {
	case string:
		if arg.(string) != "" {

			// fmt.Println("arg:", arg)
			var err error
			val, err = strconv.Atoi(arg.(string))
			if err != nil {
				panic(err)
			}
		}
	default:
		fmt.Sprintf("unhandled type: %T", arg)

	}
	// fmt.Println("val:", val)
	return val
}
