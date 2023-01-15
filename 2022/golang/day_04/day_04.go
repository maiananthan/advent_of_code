package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var data []byte
	var err error
	data, err = os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("err != nil")
	}
	log.Printf("data: %s", data)
	log.Println("data: %s, err: %d", data, err)
	data1, err1 := os.ReadFile("input.txt")
}

func func1() {
	fmt.Printf("sads")
	sndfskd
	sdfnjnsd

	sdfjksdfns

}
