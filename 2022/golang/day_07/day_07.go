package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// parse_input("1.txt")
	parse_input("input.txt")
}

type node struct {
	name   string
	isFile bool
	size   int
	parent *node
	child  map[string]*node
}

var total_size int
var c int
var dir_size1 []int
var dir_size map[string]int

func parse_input(file string) {

	var data []byte
	var err error
	var parent_dir *node
	data, err = os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var in_str string = string(data)
	var line string
	var cur_dir *node

	// read line by line
	// split the line by space
	// if arr[0] is "$" and arr[1] is "cd"
	// 		if arr[2] is "/" 	-> create a node without parent, assign it to current dir and parent_dir
	// 		if arr[2] is ".." 	-> assign current dir parent to current dir
	// 		if arr[2] is else	-> assign current dir name to current dir
	// else if arr[0] is "dir" 	-> create a node with current dir as parent and add it to child
	// else if arr[0] is not "$"	-> create a node with current dir as parent and as file and no childs
	for _, line = range strings.Split(in_str, "\n") {
		l := strings.Split(line, " ")
		if string(l[0]) == "$" && string(l[1]) == "cd" {
			if string(l[2]) == "/" {
				cur_dir = &node{"/", false, 0, nil, make(map[string]*node)}
				parent_dir = cur_dir
			} else if string(l[2]) == ".." {
				cur_dir = cur_dir.parent
			} else {
				cur_dir = cur_dir.child[string(l[2])]
			}
		} else if string(l[0]) == "dir" {
			cur_dir.child[string(l[1])] = &node{string(l[1]), false, 0, cur_dir, make(map[string]*node)}
		} else if string(l[0]) != "$" {
			var file_size int
			file_size, _ = strconv.Atoi(string(l[0]))
			cur_dir.child[string(l[1])] = &node{string(l[1]), true, file_size, cur_dir, nil}
		}
	}

	// print_dir(0, parent_dir)

	size_calc(parent_dir)
	fmt.Printf("total_size: %v\n", total_size)
	// del_dir_size(parent_dir)
	// sort.Ints(dir_size)
	// fmt.Printf("del size: %v\n", dir_size[len(dir_size)-1])
	// var del_size int = deldir(parent_dir)
	// fmt.Printf("del_size: %v\n", del_size)

	dir_size = make(map[string]int)

	calc_dir_size(parent_dir)
	fmt.Printf("total_size: %v\n", total_size)

	// part 1
	var low_size_total int = 0

	fmt.Printf("len: %v\n", len(dir_size))
	for _, element := range dir_size {
		// fmt.Printf("%v: %v\n", key, element)
		// fmt.Printf("%T\n", element)

		if element < 100_000 {
			low_size_total += element
		}
	}
	// fmt.Printf("c: %v\n", c)

	// part 2
	del_dir_size(parent_dir)
	var total_space int = 70_000_000
	var req_space int = 30_000_000
	// var unused_space int = total_space - dir_size["/"]
	sort.Ints(dir_size1)
	var unused_space int = total_space - dir_size1[len(dir_size1)-1]

	fmt.Printf("unsed space: %v\n", unused_space)
	var small_space []int
	for _, element := range dir_size {
		if unused_space+element > req_space {
			small_space = append(small_space, element)
		}
	}
	sort.Ints(small_space)
	fmt.Printf("small space: %v\n", small_space[0])
}

// check if dir is file or not
//
//	-> if file, return size of file
//
// if it's not file, loop through the child of the dir
// and recursion to the child items to the same function
// also, add directory to map with name and size

func calc_dir_size(dir *node) (size int) {
	// fmt.Printf("dir.name: %v\n", dir.name)s
	if dir.isFile {
		return dir.size
	}
	for _, e := range dir.child {
		// c += 1
		size += calc_dir_size(e)
	}
	// if dir_size[dir.name] == nil {
	// 	dir_size[dir.name] = map[sttr]type
	// }
	// fmt.Printf("calc -> %v: %v\n", dir.name, size)
	dir_size[dir.name] = size
	return size
}

// func deldir(dir *node) (size int) {
// 	if dir.isFile {
// 		return dir.size
// 	}
// 	var s1 int = 0
// 	for _, e := range dir.child {
// 		fmt.Printf("-- %v: %v\n", e.name, e.size)
// 		s1 += e.size
// 		if !e.isFile {
// 			deldir(e)
// 		}
// 	}
// 	fmt.Printf("%v: %v\n", dir.name, s1)
// 	return s1
// }

func del_dir_size(dir *node) (del_size int) {
	del_size = 0
	var size int
	// var tot_ds int = 70_000_000
	// var up_ds int = 30_000_000
	// fmt.Printf("/ size: %v\n", ret_size(dir))

	for _, element := range dir.child {
		if !element.isFile {
			size = ret_size(element)
			// if  > up_up_ds
			// dir_size1 = append(dir_size1, size)
			// fmt.Printf("%v -> %v\n", element.name, size)
			del_dir_size(element)
		}
	}
	dir_size1 = append(dir_size1, size)
	return del_size
}

func size_calc(parent_dir *node) {
	for _, element := range parent_dir.child {
		if !element.isFile {
			var size int = ret_size(element)
			// fmt.Printf("%v -> size: %v\n", element.name, size)
			if size < 100_000 {
				total_size += size
			}
			size_calc(element)
		}
	}
}

func ret_size(dir *node) (dir_size int) {
	dir_size = 0
	for _, element := range dir.child {
		dir_size += element.size
		if !element.isFile {
			dir_size += ret_size(element)
		}
	}
	return dir_size
}

func print_dir(space int, parent_dir *node) {
	for _, element := range parent_dir.child {
		fmt.Printf("%s", strings.Repeat(" ", space))
		fmt.Printf("- %v (%v)\n", element.name, file_or_dir(element.isFile))
		if element.child != nil {
			print_dir(space+2, element)
		}
	}
}

func file_or_dir(isFile bool) (str string) {
	if isFile {
		return "file"
	} else {
		return "dir"
	}
}
