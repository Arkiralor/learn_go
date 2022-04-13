package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var num int64 = 600
	fmt.Println("Converting", num, "to binary.")
	fmt.Println("Binary:", convert_to_binary(num))
}

func convert_to_binary(num int64) int {
	num_str := strconv.FormatInt(num, 2)
	num_int, err := strconv.Atoi(num_str)
	if err != nil {
		log.Fatalf("\nError Code: %v\n", err.Error())
	}
	return num_int
}
