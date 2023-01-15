package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// parseInput("1.txt")
	parseInput("input.txt")
}

func parseInput(fileName string) {
	var data []byte
	var err error

	data, err = os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	var ip_str string = string(data)
	// var line string
	var instArr []string
	// for _, line = range strings.Split(ip_str, "\n") {
	// 	instArr = append(instArr, line)
	// }

	instArr = append(instArr, strings.Split(ip_str, "\n")...)

	fmt.Printf("signal strength: %v\n", p1(instArr))
	// fmt.Printf("p2: %v\n", p2(instArr))
	p2(instArr)
	fmt.Println()
}

func p1(instArr []string) (sigS int) {
	sigS = 0
	var inst string
	var l []string
	var cycles int = 0
	var RegX int = 1

	for _, inst = range instArr {
		l = strings.Split(inst, " ")
		{
			switch l[0] {
			case "noop":
				cycles += 1
				if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 || cycles == 180 || cycles == 220 {
					sigS = sigS + (cycles * RegX)
					fmt.Printf("sigS: %v\n", sigS)
				}
				RegX += 0
				continue
			case "addx":
				cycles += 1
				if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 || cycles == 180 || cycles == 220 {
					sigS = sigS + (cycles * RegX)
					fmt.Printf("sigS: %v\n", sigS)
				}
				RegX += 0
				cycles += 1
				if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 || cycles == 180 || cycles == 220 {
					sigS = sigS + (cycles * RegX)
					fmt.Printf("sigS: %v\n", sigS)
				}
				var addVal int
				addVal, _ = strconv.Atoi(string(l[1]))
				RegX += addVal
			}

		}
	}

	return sigS
}

func p2(instArr []string) {

	var regX, cycle int
	regX = 1
	cycle = 0

	for _, line := range instArr {
		l := strings.Split(line, " ")

		// for noop and addx first cycle, perform CRT draw
		cycle = drawCRT(regX, cycle)

		if l[0] == "addx" {
			// for addx perform CRT draw and add register X value
			cycle = drawCRT(regX, cycle)
			val, _ := strconv.Atoi(l[1])
			regX += val
		}
	}

}

func drawCRT(regX int, cycle int) int {

	// for every 40 cycles create new line
	if cycle%40 == 0 && cycle <= 240 {
		fmt.Println()
	}
	// if cycle aligns with register X (sprite), create lit(#) else dim(.)
	if regX-1 == cycle%40 || regX == cycle%40 || regX+1 == cycle%40 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	cycle++
	return cycle
}
