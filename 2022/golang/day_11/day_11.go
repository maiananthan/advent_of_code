package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if os.Args[1] == "1" {
		parseInput("1.txt")
	} else if os.Args[1] == "2" {
		parseInput("input.txt")
	}
}

type ops struct {
	monkey    int
	items     []int
	operator  rune
	value     int
	test      int
	testTrue  int
	testFalse int
	inspect   int
}

func parseInput(fileName string) {
	var data []byte
	var err error

	data, err = os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var ip_str string = string(data)

	// load data to struct
	var opsMonkey []string
	opsMonkey = append(opsMonkey, strings.Split(ip_str, "\n\n")...)

	var s_ops []ops
	for _, opsSMonkey := range opsMonkey {

		var monkey int
		var items []int
		var operator rune
		var value int
		var test int
		var testT int
		var testF int
		for _, line := range strings.Split(opsSMonkey, "\n") {
			inputData := strings.Split(string(line), " ")

			if string(inputData[0]) == "Monkey" {
				monkeyNum := inputData[1]
				monkey, _ = strconv.Atoi(string(monkeyNum[0]))
			} else if strings.Contains(line, "Starting items") {
				itemsList := line[strings.Index(line, ":")+1:]
				itemsListStr := strings.Split(itemsList, ",")
				for _, i := range itemsListStr {
					i = strings.ReplaceAll(i, " ", "")
					items1, err := strconv.Atoi(i)
					if err != nil {
						panic(err)
					}
					items = append(items, items1)
				}
			} else if strings.Contains(line, "Operation") {
				if string(line) == "  Operation: new = old * old" || string(line) == "  Operation: new = old + old" {
					fmt.Sscanf(line, "  Operation: new = old %c old", &operator)
				} else {
					fmt.Sscanf(line, "  Operation: new = old %c %d", &operator, &value)
				}
			} else if strings.Contains(line, "Test") {
				fmt.Sscanf(line, "  Test: divisible by %d", &test)
			} else if strings.Contains(line, "If true") {
				fmt.Sscanf(line, "    If true: throw to monkey %d", &testT)
			} else if strings.Contains(line, "If false") {
				fmt.Sscanf(line, "    If false: throw to monkey %d", &testF)
			}
		}
		s_ops = append(s_ops, ops{monkey: monkey, items: items, operator: operator, value: value, test: test, testTrue: testT, testFalse: testF})
	}
	fmt.Printf("monkey business: %v\n", p1(s_ops, 20, true))
	fmt.Printf("monkey business: %v\n", p1(s_ops, 10_000, false))
}

func p1(s_ops []ops, totalRound int, d3 bool) int {
	var index int = 0
	var item int
	var worry int
	var round int = 1
	var commonMultiple int = 1
	if !d3 {
		for i1 := 0; i1 < len(s_ops); i1++ {
			commonMultiple *= s_ops[i1].test
		}
	}

	for round = 1; round <= totalRound; round++ {
		for index = 0; index < len(s_ops); index++ {
			var itemCount int = len(s_ops[index].items)
			for i1 := 0; i1 < itemCount; i1++ {

				// get item from array
				_, item, s_ops[index] = popItem(s_ops[index])
				var ops ops = s_ops[index]
				// perform calculation
				worry = item
				s_ops[index].inspect++
				if ops.value == 0 {
					if ops.operator == '+' {
						worry = worry + worry
					} else {
						worry = worry * worry
					}
				} else if ops.value != 0 {
					if ops.operator == '+' {
						worry = worry + ops.value
					} else {
						worry = worry * ops.value
					}
				}
				if d3 {
					worry = worry / 3
				} else {
					// as worry is not getting divided by 3, worry levels will get bigger on consecutive rounds
					// to avoid that all dividers product is taken and modulus with worry level and the same is added
					worry = worry % commonMultiple
				}
				if worry%ops.test == 0 {
					s_ops = pushItem(s_ops, worry, ops.testTrue)
				} else {
					s_ops = pushItem(s_ops, worry, ops.testFalse)
				}
			}
		}
	}
	var inspect []int
	for _, i_ops := range s_ops {
		inspect = append(inspect, i_ops.inspect)
	}
	sort.Ints(inspect)
	return inspect[len(inspect)-2] * inspect[len(inspect)-1]
}

func pushItem(s_ops []ops, worry int, index int) []ops {
	s_ops[index].items = append(s_ops[index].items, worry)
	return s_ops
}

func popItem(_ops ops) (exists bool, item int, ret_ops ops) {
	if len(_ops.items) > 0 {
		item = _ops.items[0]
		_ops.items = _ops.items[1:len(_ops.items)]
		ret_ops = _ops
		return true, item, ret_ops
	} else {
		return false, 0, _ops
	}
}
