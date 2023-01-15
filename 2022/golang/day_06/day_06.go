package main

import (
	"fmt"
	"os"
)

func main() {

	cal("1.txt", 4)
	cal("2.txt", 4)
	cal("3.txt", 4)
	cal("4.txt", 4)
	cal("5.txt", 4)
	cal("input.txt", 4)
	cal("1.txt", 14)
	cal("2.txt", 14)
	cal("3.txt", 14)
	cal("4.txt", 14)
	cal("5.txt", 14)
	cal("input.txt", 14)

}

func cal(file_name string, som int) {
	var data []byte
	var err error
	data, err = os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}
	var ip_str string
	ip_str = string(data)
	var index int
	var i1 string
	for index = 0; index < len(ip_str) && index+som <= len(ip_str); index++ {
		i1 = ""
		i1 = ip_str[index : index+som]
		var j, k int
		var flag int
		flag = 0
		for j = 0; j < len(i1)-1; j++ {
			for k = j + 1; k < len(i1); k++ {
				if i1[j] == i1[k] {
					// fmt.Print("j:", string(i1[j]), ", k:", string(i1[k]), "\n")
					flag = 1
					break
				}
			}
			if flag == 1 {
				break
			}
		}
		// fmt.Println()
		// fmt.Println("i1:", i1, "index:", index, ", flag:", flag)
		if flag == 0 {
			fmt.Println(file_name, "| match found:", index+som)
			break
		}

	}
}
