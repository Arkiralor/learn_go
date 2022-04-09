package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func NumberFromCon() int {
	reader := bufio.NewReader(os.Stdin)
	input_str, inp_err := reader.ReadString('\n')
	if inp_err != nil {
		log.Fatalf("Error Code: %v", inp_err)
	}
	input_str = RemoveEOLMarkers(input_str)
	input_int, err := strconv.Atoi(input_str)

	if err != nil {
		log.Fatalf("Error Code: %v", err)
	}

	return input_int
}

func ListFromString(s string) []string {
	s = RemoveEOLMarkers(s)
	s_2 := strings.Split(s, " ")
	return s_2

	// return s_2
}

func RemoveEOLMarkers(s string) string {
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)

	return s
}

func main() {
	// var arr []int
	// fmt.Println("Enter an array: ")
	// for i := 0; i < 10; i++ {
	// 	input_int := NumberFromCon()
	// 	arr = append(arr, input_int)
	// }

	// var arr [10]int
	// fmt.Println("Enter an array: ")
	// for i := 0; i < 10; i++ {
	// 	arr[i] = NumberFromCon()
	// }

	// fmt.Println(arr)
	arr := "Oh what a travesty it is to be in love!"
	fmt.Println(arr)
	new_var := ListFromString(arr)
	fmt.Println(new_var)

}
