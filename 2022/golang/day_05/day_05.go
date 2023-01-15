package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// split crates stack and instructions
// parse crates stack into two dimensional array
// parse instructions and re-arrange crates stack

func main() {
	var data []byte
	var err error
	// data, err = os.ReadFile("input.sample.txt")
	data, err = os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	// split crates stack and instructions
	var stack string
	var inst string
	spl_str := strings.Split(string(data), "\n\n")
	stack = string(spl_str[0])
	inst = string(spl_str[1])

	// parse crates stack into 2d array

	inst = inst
	var arr2dlen int
	arr2dlen, err = strconv.Atoi(stack[len(stack)-2 : len(stack)-1])
	var arr2d = make([][]string, arr2dlen)
	var line string

	for _, line = range strings.Split(stack, "\n") {
		// log.Printf("line: \"%s\"", line)
		var index int
		var j int
		j = 0
		for index = 1; index < len(line); index += 4 {
			// log.Printf("index: \"%s\"", index)
			// log.Printf("line[%d]: \"%T\"", index, line[index])
			if string(line[index]) != " " {
				// log.Printf("line[%v]: %v", index, string(line[index]))
				b1, err := regexp.MatchString("[0-9]", string(line[index]))
				// log.Printf("b1: %v", b1)
				if err != nil {
					panic(err)
				}
				// arr1d = append(arr1d, string(line[index]))
				if !b1 {
					// arr2d[j] = append(arr2d[j], string(line[index]))
					arr2d[j] = append(arr2d[j], string(line[index]))

				}
			}
			j++
		}

		// line[2] 6 10
	}

	// reverse arr2d
	var arr2dr [][]string
	var p2arr [][]string
	for _, arr1d := range arr2d {
		fmt.Printf("\nlen: %v | ", len((arr1d)))
		var arr1dr []string
		for j := len(arr1d) - 1; j >= 0; j-- {
			// log.Printf("<-> arr[%v][%v]: %v", i, j, arr1d[j])
			arr1dr = append(arr1dr, arr1d[j])
			fmt.Print(arr1d[j], " ")
			// log.Printf("arr[%v][%v]: %v", i, j, arr1dr[j])
		}
		arr2dr = append(arr2dr, arr1dr)
		// for _, i1 := range arr1dr {
		// 	fmt.Print(" ", i1)
		// }
		// fmt.Println()
		// p2arr = append(p2arr, arr1dr)
	}

	for _, p2arr1d := range arr2d {
		// fmt.Printf("\nlen: %v | ", len((p2arr1d)))
		var p2arr1dr []string
		for j := len(p2arr1d) - 1; j >= 0; j-- {
			// log.Printf("<-> arr[%v][%v]: %v", i, j, arr1d[j])
			p2arr1dr = append(p2arr1dr, p2arr1d[j])
			// fmt.Print(p2arr1d[j], " ")
			// log.Printf("arr[%v][%v]: %v", i, j, arr1dr[j])
		}
		// arr2dr = append(arr2dr, arr1dr)
		// for _, i1 := range arr1dr {
		// 	fmt.Print(" ", i1)
		// }
		// fmt.Println()
		p2arr = append(p2arr, p2arr1dr)
	}

	// print the crates stack
	// for i, arr1d := range arr2d {
	// 	fmt.Print("\n", i, " | ")
	// 	for _, element := range arr1d {
	// 		// fmt.Printf("arr[%v][%v]: %v", i, j, element)
	// 		fmt.Printf("%v ", element)

	// 	}
	// }
	// for i, arr1d := range p2arr {
	// 	fmt.Print("\n r->", i, " | ")
	// 	for _, element := range arr1d {
	// 		fmt.Printf("%v ", element)

	// 	}
	// }
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		log.Printf("arr[%v][%v]: %v", i, j, arr2d[i][j])
	// 	}
	// }
	// load the stack to 2d arrays
	// var stackarr2d [][]string
	var last_line string
	last_line = stack[len(stack)-2 : len(stack)-1]

	// var dimen int
	// dimen = strconv.Atoi(last_line)
	// var stackarr2d [][]string

	// load into stack
	f1 := stack[:3]
	log.Printf("f1: \"%s\"", f1)
	if f1 != "  " {
		log.Printf("not null: f1: \"%s\"", f1)
	} else {
		log.Printf("null")
	}

	log.Printf("last_line: \"%s\"", last_line)
	// parse instructions
	// var index int
	var inst_line string
	var move, from, to int
	var arr2d_r []string
	var ans string
	// print(arr2dr)
	// log.Printf("len: %v %v", len(inst), strings.Count(inst, "\n"))
	// var i2 int
	// log.Printf("\"%v\"", inst)
	// var p2arr [][]string
	// p2arr = arr2dr
	for _, inst_line = range strings.Split(inst, "\n") {
		fmt.Sscanf(inst_line, "move %d from %d to %d", &move, &from, &to)

		// log.Printf("%v: move %v from %v to %v", i2, move, from, to)
		for i := 0; i < move; i++ {
			arr2d_r, ans = pop(arr2dr[from-1])
			arr2dr[from-1] = arr2d_r
			arr2d_r = push(arr2dr[to-1], ans)
			arr2dr[to-1] = arr2d_r
		}
		// print(arr2dr)
	}
	fmt.Print("part 1 -> ")
	for _, e1 := range arr2dr {
		fmt.Printf(e1[len(e1)-1])
	}
	fmt.Println()

	// print(arr2dr)
	log.Println(" - - - ")
	print(p2arr)
	log.Println(" - - - ")

	for _, inst_line = range strings.Split(inst, "\n") {
		var p2arr_ans []string
		fmt.Sscanf(inst_line, "move %d from %d to %d", &move, &from, &to)
		// log.Printf("move %v from %v to %v", move, from, to)

		p2arr_ans = p2arr[from-1][len(p2arr[from-1])-move : len(p2arr[from-1])]
		p2arr[from-1] = p2arr[from-1][:len(p2arr[from-1])-move]
		// print(p2arr)
		// log.Println(" - - - ")
		// for _, e1 := range p2arr_ans {
		// 	fmt.Print(e1, "-")
		// }
		// fmt.Println()
		p2arr[to-1] = append(p2arr[to-1], p2arr_ans...)
		// print(p2arr)
	}
	for _, e1 := range p2arr {
		fmt.Printf(e1[len(e1)-1])
	}
	// fmt.Println()
}

func pop(in_arr []string) (arr_r []string, ans string) {
	ans = in_arr[len(in_arr)-1]
	arr_r = in_arr[0 : len(in_arr)-1]
	return arr_r, ans
}

func push(in_arr []string, in string) (arr []string) {
	arr = append(in_arr, in)
	return arr
}

func print(in_arr [][]string) {
	for i1, arr1 := range in_arr {
		fmt.Print(" ", i1, "\"")
		for _, e := range arr1 {
			fmt.Print(" ", e)
		}
		fmt.Print("\" [", len(arr1), "]")
		fmt.Println(" ")
	}
}
